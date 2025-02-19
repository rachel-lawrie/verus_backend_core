package utils

import (
	"time"

	"github.com/golang-jwt/jwt"
)

var JwtKey = []byte("secret_key")

type Claims struct {
	UserID string `json:"cockpit_user_id"`
	jwt.StandardClaims
}

func GenerateJWT(userID string) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &Claims{
		UserID: userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(JwtKey)
}
