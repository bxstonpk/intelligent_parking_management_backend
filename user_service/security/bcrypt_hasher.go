package security

import (
	"os"
	"time"
	"user_services/errs"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

// SecretKey is the key used to sign the JWT tokens
var secretKey []byte

type bcryptHasher struct {
	SecretKey string
}

func NewBcryptHasher(secret string) bcryptHasher {
	return bcryptHasher{SecretKey: string(secretKey)}
}

func (b bcryptHasher) HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func (b bcryptHasher) CheckPasswordHash(password string, passwordHash string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(password)); err != nil {
		return false
	}
	return true
}

func (s bcryptHasher) GenerateJWT(userID string) (string, error) {
	// Ensure the secret key is set
	secretKey = []byte(os.Getenv("SECRET_KEY"))
	if secretKey == nil {
		return "", errs.NewUnexpectedError()
	}

	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secretKey)
}

func (s bcryptHasher) ValidateToken(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Ensure the signing method is HS256
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errs.NewUnexpectedError()
		}
		return secretKey, nil
	})
}
