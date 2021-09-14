package middleware

import (
	"net/http"
	"time"

	"github.com/Dall06/go-cleanarch-template/config"
	"github.com/Dall06/go-cleanarch-template/pkg/internal/internal_user"
	"github.com/golang-jwt/jwt"
)

var jwtKey = []byte(config.SecretPassword)

type claims struct {
	UserAccount internal_user.UserAccount `json:"userAccount"`
	jwt.StandardClaims
}

type JWTHandler struct{}

func NewJWTHandler() *JWTHandler {
	return &JWTHandler{}
}

func (j *JWTHandler) SetTokenCookie(w http.ResponseWriter, user internal_user.UserAccount) error {
	expirationTime := time.Now().Add(time.Minute * 5)
	// Setting up token
	claims := &claims{
		UserAccount: user,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	// if error when creating token retunr Internal Server Error
	if err != nil {
		return err
	}

	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   tokenString,
		Expires: expirationTime,
	})

	return nil
}

func (*JWTHandler) RefreshTokenCookie(r *http.Request, user internal_user.UserAccount) (string, time.Time, error) {
	expirationTime := time.Now().Add(time.Minute * 5)
	//  get token cookie value
	cookie, err := r.Cookie("token")
	if err != nil {
		return "", time.Time{}, err
	}
	tokenStr := cookie.Value
	// setting claims up and generate token
	claims := &claims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token, err := jwt.ParseWithClaims(tokenStr, claims,
		func(t *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		},
	)

	if err != nil {
		return "", time.Time{}, err
	}
	// if token is not valid respond authorization and then re flow continues
	if !token.Valid {
		return "", time.Time{}, http.ErrAbortHandler
		// return
	}
	// Generate new token
	newToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := newToken.SignedString(jwtKey)
	// if error exists when creating new token send error
	if err != nil {
		return "", time.Time{}, err
	}
	// Set new Cookie, then finishes
	return tokenString, expirationTime, nil
}

func (*JWTHandler) CleanTokenCookie(w http.ResponseWriter) {
	// Cleans your t=cookie in the 'token' field
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   "",
		Expires: time.Now().Add(-time.Hour),
	})
}

func (*JWTHandler) ValidateTokenCookie(r *http.Request) (bool, error) {
	// Get token from cookies
	cookie, err := r.Cookie("token")
	if err == http.ErrNoCookie || err != nil {
		return false, err
	}
	tokenStr := cookie.Value
	// Generate empty claims and parse token with them
	claims := &claims{}
	tkn, err := jwt.ParseWithClaims(tokenStr, claims,
		func(t *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		},
	)
	// if error exists breaks
	if err != nil {
		return false, err
	}
	// if token is not valid breaks not authorized
	if !tkn.Valid {
		return false, nil
	}
	return true, nil
}