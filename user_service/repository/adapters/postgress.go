package repository

import "github.com/jmoiron/sqlx"

type postgresUserRepository struct {
	db *sqlx.DB
}

func NewPostgresUserRepository(db *sqlx.DB) postgresUserRepository {
	return postgresUserRepository{db: db}
}
