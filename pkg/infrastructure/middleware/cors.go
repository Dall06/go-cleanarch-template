package middleware

import (
	"net/http"
)

type corsMiddleware struct {
	methods    string
	corsDomain string
}

func NewCORSMiddleware(m string, d string) *corsMiddleware{
	return &corsMiddleware{
		methods: m,
		corsDomain: d,
	}
}

func (cors *corsMiddleware) SetCORSParameters(methods string, domain string) {
	cors.methods = methods
	cors.corsDomain = domain
}

func (cors *corsMiddleware) EnableCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, req *http.Request) {
			w.Header().Set("Access-Control-Allow-Origin", cors.corsDomain)
			w.Header().Set("Access-Control-Allow-Credentials", "true")
			w.Header().Set("Access-Control-Allow-Methods", cors.methods)
			w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
			next.ServeHTTP(w, req)
		})
}
