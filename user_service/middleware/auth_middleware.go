package middleware

import (
	"context"
	"net/http"
	"os"
	"strings"
	"user_services/security"

	"github.com/golang-jwt/jwt/v5"
)

type authMiddleware struct {
}

func NewAuthMiddleware() authMiddleware {
	return authMiddleware{}
}

type contextKey string

const UserIDKey contextKey = "userID"

func (m *authMiddleware) JWTAuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("Authorization")
		if tokenString == "" {
			http.Error(w, "Missing token", http.StatusUnauthorized)
			return
		}

		tokenString = strings.TrimPrefix(tokenString, "Bearer ")
		token, err := security.NewBcryptHasher(os.Getenv("SECRET_KEY")).ValidateToken(tokenString)
		if err != nil || !token.Valid {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			http.Error(w, "Invalid token claims", http.StatusUnauthorized)
			return
		}

		userID, ok := claims["user_id"].(string)
		if !ok {
			http.Error(w, "Invalid user ID in token", http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), UserIDKey, userID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
