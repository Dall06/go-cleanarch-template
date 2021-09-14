package middleware

import (
	"log"
	"net/http"
)

// Define our struct
type authenticationMiddleware struct {
	tokenUsers map[string]string
}

func NewAuthenticationMiddleWre() *authenticationMiddleware {
	return &authenticationMiddleware{}
}

// Initialize it somewhere
func (amw *authenticationMiddleware) Populate() {
	amw.tokenUsers = make(map[string]string)
	amw.tokenUsers["00000000"] = "userDall"
}

// Middleware function, which will be called for each request
func (amw *authenticationMiddleware) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("X-Session-Token")

		if user, found := amw.tokenUsers[token]; found {
			log.Printf("Authenticated user %s\n", user)
			next.ServeHTTP(w, r)
			return
		}
		// Write an error and stop the handler chain
		http.Error(w, "Forbidden", http.StatusForbidden)
	})
}
