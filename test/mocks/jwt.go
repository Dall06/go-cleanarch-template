package mocks

import (
	"time"

	"github.com/Dall06/go-cleanarch-template/config"
	"github.com/Dall06/go-cleanarch-template/src/domain"
	"github.com/golang-jwt/jwt"
)

func GenerateJWT(u *domain.UserAccount) (string, error){
	expirationTime := time.Now().Add(time.Minute * 5)
	// Setting up token
	claims := &config.Clamis{
		UserAccount: *u,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(config.JwtKey)
	// if error when creating token retunr Internal Server Error
	if err != nil {
		return "", err
	}

	return tokenString, err
}