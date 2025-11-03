package utils

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

func CreateToken(userID uint64) (string, error) {
	signingKey := []byte("mySecretKey")

	claims := &jwt.StandardClaims{
		ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
		Issuer:    "myApp",
		Subject:   string(rune(userID)),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(signingKey)
}
