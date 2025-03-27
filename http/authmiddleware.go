package http

import (
	"OverEngineeredCalculator/helpers"
	"github.com/google/uuid"
	"net/http"
	"strings"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		// Remove the Bearer prefix and validate the token as a UUID
		// In a real-world scenario, we could validate the token with a JWT library something else depending on the auth method
		token := strings.TrimPrefix(authHeader, "Bearer ")
		if _, err := uuid.Parse(token); err != nil {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		// Store user ID in context (for future use)
		ctx := r.Context()
		ctx = helpers.ContextWithUserID(ctx, token)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
