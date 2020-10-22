package auth

import "github.com/dgrijalva/jwt-go"

var jwtKey = []byte("my_secret_key") //TODO: make envvar

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}



