package repository

import "go-merchant/model"

type AuthRepository interface {
	CheckPassword(username string, password string) (model.User, error)
	GenerateJWT(username string, userid string) (tokenString string, err error)
	ValidateToken(signedToken string) (userId string, err error)
}
