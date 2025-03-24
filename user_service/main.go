package main

import (
	"fmt"
	repository "user_services/repository"

	_ "github.com/gorilla/mux"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
)

func main() {
	dsn := "postgres://projectEngineer:S@uEngin33r@localhost:5432/userDB?sslmode=disable"
	db, err := sqlx.Open("pgx", dsn)
	if err != nil {
		panic(err)
	}

	// Create Router
	/* r := mux.NewRouter() */

	// Create User Repository
	userRepository := repository.NewPostgresUserRepository(db)

	_ = userRepository

	/* // Create Route Handler for get users
	r.HandleFunc("/user/{id}", func(w http.ResponseWriter, r *http.Request) { */

	user, err := userRepository.LoginUser("sconklin0", "$2a$04$BLkZ7c2Ft6DoxpGU4ZyncezlG8ma/SrX5DPoAe8M2HRrRrUP2Wpq2")
	if err != nil {
		panic(err)
	}

	fmt.Println(user)
}
