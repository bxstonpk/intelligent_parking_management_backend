package db

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

type Config struct {
	User       string
	Password   string
	Host       string
	Port       int
	Name       string
	DisableTLS bool
}

func Open(cfg Config) (*sqlx.DB, error) {
	sslmode := "require"
	if cfg.DisableTLS {
		sslmode = "disable"
	}
	dataSoruce := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s", cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.Name, sslmode)
	return sqlx.Open("postgres", dataSoruce)
}
