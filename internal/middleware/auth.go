package middleware

import (
	"context"
	jwt "github.com/dgrijalva/jwt-go"
	"net/http"
	"os"
	"profiler/internal/message"
	"profiler/internal/responser"
	"profiler/internal/user"
	"strings"
)

var JwtAuthentication = func(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// List of non-authentication URLs
		noAuth := []string{"/api/users/new", "/api/users/login"}
		request := r.URL.Path

		// Handling for no authentication URLs
		for _, value := range noAuth {
			if value == request {
				next.ServeHTTP(w, r)
				return
			}
		}

		response := make(map[string]interface{})
		tokenRequest := r.Header.Get("Authorization")

		// Check if request is missing token
		if tokenRequest == "" {
			response = message.Message(false, "NG00001", "Missing auth token")
			w.WriteHeader(http.StatusForbidden)
			w.Header().Add("Content-Type", "application/json")
			responser.Respond(w, response)
			return
		}

		// Check if the token is in the correct format
		split := strings.Split(tokenRequest, " ")
		if len(split) != 2 {
			response = message.Message(false, "NG00001", "Invalid/Malformed auth token")
			w.WriteHeader(http.StatusForbidden)
			w.Header().Add("Content-Type", "application/json")
			responser.Respond(w, response)
			return
		}

		tokenPart := split[1]
		tk := &user.Token{}

		token, err := jwt.ParseWithClaims(tokenPart, tk, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("token_password")), nil
		})

		// Malformed token
		if err != nil {
			response = message.Message(false, "NG0001", "Malformed authentication token")
			w.WriteHeader(http.StatusForbidden)
			w.Header().Add("Content-Type", "application/json")
			responser.Respond(w, response)
			return
		}

		// Token is invalid
		if !token.Valid {
			response = message.Message(false, "NG0001", "Token is not valid.")
			w.WriteHeader(http.StatusForbidden)
			w.Header().Add("Content-Type", "application/json")
			responser.Respond(w, response)
			return
		}

		// Token is valid, everything is ok
		ctx := context.WithValue(r.Context(), "user", tk.UserId)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}
