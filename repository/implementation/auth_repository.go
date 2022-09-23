package implementation

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"go-merchant/config"
	"go-merchant/database"
	"go-merchant/model"
	"go-merchant/repository"
	"log"
	"time"

	"github.com/golang-jwt/jwt"
)

type AuthRepo struct {
}

func NewAuthRepo() repository.AuthRepository {
	return &AuthRepo{}
}

func (authrepo *AuthRepo) CheckPassword(username string, password string) (model.User, error) {
	var user model.User
	log.Println()
	hash := md5.Sum([]byte(password))
	record := database.DB.Table("users").Where("user_name = ? and password = ?", username, hex.EncodeToString(hash[:])).First(&user)
	if record.Error != nil {
		log.Println(record.Error)
		return user, record.Error
	}
	return user, nil
}

func (authrepo *AuthRepo) GenerateJWT(username string, userid string) (tokenString string, err error) {
	expirationTime := time.Now().Add(1 * time.Hour)
	claims := &model.JWTUserClaim{
		UserId:   userid,
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err = token.SignedString(config.JwtKey)
	return
}

func (authrepo *AuthRepo) ValidateToken(signedToken string) (userId string, err error) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&model.JWTUserClaim{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(config.JwtKey), nil
		},
	)
	if err != nil {
		return
	}
	claims, ok := token.Claims.(*model.JWTUserClaim)
	if !ok {
		err = errors.New("couldn't parse claims")
		return
	}
	if claims.ExpiresAt < time.Now().Local().Unix() {
		err = errors.New("token expired")
		return
	}
	return claims.UserId, nil
}
