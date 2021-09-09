package config

import (
	"github.com/Dall06/go-cleanarch-template/src/domain"
	"github.com/golang-jwt/jwt"
)

// SECRET PASSWORD FOR THE TOKENS
var JwtKey = []byte(SecretPassword)

type Clamis struct {
	UserAccount domain.UserAccount `json:"userAccount"`
	jwt.StandardClaims
}
