package model

import "github.com/golang-jwt/jwt"

type User struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	UserName string `json:"user_name"`
	Token    string `json:"token"`
}

type SignInRequest struct {
	UserName string `json:"user_name"`
	Password string `json:"password"`
}

type JWTUserClaim struct {
	Username string `json:"username"`
	UserId   string `json:"Id"`
	jwt.StandardClaims
}
