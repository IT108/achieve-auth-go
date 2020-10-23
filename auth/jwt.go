package auth

import (
	. "github.com/IT108/achieve-models-go/auth"
	"github.com/dgrijalva/jwt-go"
	"time"
)

var jwtKey = []byte("my_secret_key") //TODO: make envvar

func GenerateToken(username string) (ok bool, token AuthToken, err string) {
	result := AuthToken{
		TokenString:    "",
		ExpirationTime: time.Time{},
	}

	expirationTime := time.Now().Add(5 * time.Minute)
	result.ExpirationTime = expirationTime
	// Create the JWT claims, which includes the username and expiry time
	claims := &Claims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			// In JWT, the expiry time is expressed as unix milliseconds
			ExpiresAt: expirationTime.Unix(),
		},
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, signErr := jwtToken.SignedString(jwtKey)

	if signErr != nil {
		return false, result, "Internal server error"
	}
	result.TokenString = tokenString

	return true, result, ""
}
