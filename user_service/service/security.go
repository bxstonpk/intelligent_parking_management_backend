package service

type SecurityService interface {
	CheckToken(token string) (bool, error)
}
