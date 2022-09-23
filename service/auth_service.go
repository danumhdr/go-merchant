package service

import "go-merchant/model"

type AuthService interface {
	CheckCredential(username string, password string) (model.User, error)
	ValidateToken(signedToken string) (userid string, err error)
}
