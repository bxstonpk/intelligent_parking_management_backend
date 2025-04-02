package security

import (
	"github.com/golang-jwt/jwt/v5"
)

type Password struct {
	userID       string `json:"user_id"`
	Username     string `json:"username"`
	Password     string `json:"password"`
	PasswordHash string `json:"password_hash"`
}

type PasswordHasherRepository interface {
	HashPassword(Password) (*Password, error)
	CheckPasswordHash(Password) bool
	GenerateJWT(userID string) (string, error)
	ValidateToken(tokenString string) (*jwt.Token, error)
}
