package security

type PasswordHasherRepository interface {
	HashPassword(string) (string, error)
	CheckPasswordHash(string, string) bool
}
