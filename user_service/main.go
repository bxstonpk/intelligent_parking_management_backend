package main

import (
	_ "fmt"
	"net/http"
	"user_services/handler"
	repository "user_services/repository"
	"user_services/service"

	"github.com/gorilla/mux"
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

	// Create User Repository
	userRepository := repository.NewPostgresUserRepository(db)
	userService := service.NewUserService(userRepository)
	userHandler := handler.NewUserHandler(userService)

	// Create Router
	router := mux.NewRouter()

	router.HandleFunc("/login", userHandler.LoginUserHandler).Methods("POST")
	router.HandleFunc("/getuser/{userId:[0-9]+}", userHandler.GetUserHandler).Methods("GET")

	http.ListenAndServe(":8081", router)
}
