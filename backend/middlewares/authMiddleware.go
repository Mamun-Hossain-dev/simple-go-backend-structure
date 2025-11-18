package middlewares

import (
	"net/http"
	"strings"

	"ecommerce/config"
	"ecommerce/utils"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Read Authorization Header
		auth := r.Header.Get("Authorization")
		if auth == "" {
			http.Error(w, "Missing Authorization Header", http.StatusUnauthorized)
			return
		}

		// expecting: Bearer token
		parts := strings.Split(auth, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			http.Error(w, "Invalid Authorization Header", http.StatusUnauthorized)
			return
		}

		tokenString := parts[1]
		secret := []byte(config.LoadConfig().Jwt_secret)

		// verify token
		claims, err := utils.VerifyJWT(tokenString, secret)
		if err != nil {
			http.Error(w, "Invalid Token: "+err.Error(), http.StatusUnauthorized)
			return
		}

		// store claims so handler can use it
		r.Header.Set("x-user-id", claims.Sub)
		r.Header.Set("x-user-role", claims.Role)
		r.Header.Set("x-user-name", claims.Name)

		next.ServeHTTP(w, r)
	})
}
