package main

import (
	_ "fmt"
	"net/http"
	"os"
	"user_services/handler"
	repository "user_services/repository"
	"user_services/service"

	"github.com/gorilla/mux"
	_ "github.com/gorilla/mux"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
)

func main() {
	dsn := "postgres://projectEngineer:S@uEngin33r@localhost:5432/userDB?sslmode=disable"
	db, err := sqlx.Open("pgx", dsn)
	if err != nil {
		panic(err)
	}

	// Load env variables
	env_err := godotenv.Load()
	if env_err != nil {
		panic("Error loading .env file")
	}

	// Keys
	SecretKey := os.Getenv("SECRET_KEY")

	func GetSecretKey() string {
		return SecretKey
	}

	// Create User Repository
	userRepository := repository.NewPostgresUserRepository(db)
	userService := service.NewUserService(userRepository)
	userHandler := handler.NewUserHandler(userService)

	// Create Router
	router := mux.NewRouter()

	router.HandleFunc("/login", userHandler.LoginUserHandler).Methods("POST")
	router.HandleFunc("/getuser/{userId:[0-9]+}", userHandler.GetUserHandler).Methods("GET")
	router.HandleFunc("/register", userHandler.RegisterUserHandler).Methods("POST")

	http.ListenAndServe(":8081", router)
}
