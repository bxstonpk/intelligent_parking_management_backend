package middleware

import (
	"context"
	"net/http"
	"strings"
	"user_services/security"

	"github.com/golang-jwt/jwt/v5"
)

type authMiddleware struct {
	secretKey string
}

func NewAuthMiddleware(secretKey string) authMiddleware {
	return authMiddleware{secretKey: secretKey}
}

type contextKey string

const userIDKey contextKey = contextKey("userID")

func (m *authMiddleware) JWTAuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("Authorization")
		if tokenString == "" {
			http.Error(w, "Missing token", http.StatusUnauthorized)
			return
		}

		tokenString = strings.TrimPrefix(tokenString, "Bearer ")
		token, err := security.NewBcryptHasher(m.secretKey).ValidateToken(tokenString)
		if err != nil || !token.Valid {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}
		claims, _ := token.Claims.(jwt.MapClaims)
		userID, _ := claims["user_id"].(string)

		ctx := context.WithValue(r.Context(), userIDKey, userID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
