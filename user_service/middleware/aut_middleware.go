package middleware

import (
	"context"
	"net/http"
	"strings"
	"user_services/security"

	"github.com/golang-jwt/jwt/v5"
)

type authMiddleware struct {
	authService security.PasswordHasherRepository
	secretKey   string
}

func NewAuthMiddleware(authService security.PasswordHasherRepository, secretKey string) authMiddleware {
	return authMiddleware{authService: authService, secretKey: secretKey}
}

func (m *authMiddleware) JWTAuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("Authorization")
		if tokenString == "" {
			http.Error(w, "Missing token", http.StatusUnauthorized)
			return
		}

		tokenString = strings.TrimPrefix(tokenString, "Bearer ")
		token, err := m.authService.ValidateToken(tokenString)
		if err != nil || !token.Valid {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		claims, _ := token.Claims.(jwt.MapClaims)
		userID := claims["user_id"].(string)

		ctx := context.WithValue(r.Context(), "userID", userID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
