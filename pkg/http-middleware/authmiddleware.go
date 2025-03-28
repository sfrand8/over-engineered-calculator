package http_middleware

import (
	"context"
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
		ctx = ContextWithUserID(ctx, token)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

const userIDKey string = "userID"

func ContextWithUserID(ctx context.Context, userID string) context.Context {
	return context.WithValue(ctx, userIDKey, userID)
}

func GetUserIDFromContext(ctx context.Context) string {
	if userID, ok := ctx.Value(userIDKey).(string); ok {
		return userID
	}
	return ""
}
