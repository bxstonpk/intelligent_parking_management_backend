package service

import (
	"user_services/errs"
	"user_services/security"
)

type curityService struct {
	SecretKey string
}

// NewSecurityService creates a new instance of SecurityService
func NewSecurityService(secretKey string) curityService {
	return curityService{
		SecretKey: secretKey,
	}
}

func (s curityService) CheckToken(token string) (bool, error) {
	// Here you would typically make an HTTP request to the security service
	// to validate the token. For this example, we'll just simulate a check.
	if token == "" {
		return false, nil // Invalid token
	}

	_, err := security.NewBcryptHasher(s.SecretKey).ValidateToken(token)
	if err != nil {
		return false, errs.NewBadRequestError("invalid token")
	}

	return true, nil
}
