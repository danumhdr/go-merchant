package implementation

import (
	"go-merchant/model"
	"go-merchant/repository/implementation"
	"go-merchant/service"
	"log"
	"strconv"
	"strings"
)

type AuthServe struct {
}

func NewAuthServe() service.AuthService {
	return &AuthServe{}
}

func (authserve *AuthServe) CheckCredential(username string, password string) (model.User, error) {
	user, credentialError := implementation.NewAuthRepo().CheckPassword(username, password)
	if credentialError != nil {
		log.Println(credentialError.Error())
		return user, credentialError
	}
	log.Println(user)
	log.Println(strconv.Itoa(user.Id))
	tokenString, err := implementation.NewAuthRepo().GenerateJWT(user.UserName, strconv.Itoa(user.Id))
	if err != nil {
		log.Println(err.Error())
		return user, err
	}
	user.Token = tokenString
	return user, nil
}

func (authserve *AuthServe) ValidateToken(signedToken string) (userid string, err error) {
	valueToken := strings.Split(signedToken, " ")
	id, err := implementation.NewAuthRepo().ValidateToken(valueToken[1])
	if err != nil {
		log.Println(err.Error())
		return "", err
	}
	return id, nil
}
