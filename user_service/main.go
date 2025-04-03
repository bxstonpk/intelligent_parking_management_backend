package main

import (
	"fmt"
	_ "fmt"
	"net/http"
	"os"
	"user_services/handler"
	"user_services/middleware"
	repository "user_services/repository"
	"user_services/service"

	"github.com/gorilla/mux"
	_ "github.com/gorilla/mux"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// Load env variables
	env_err := godotenv.Load()
	if env_err != nil {
		panic("Error loading .env file")
	}

	// Connect to database
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", dbHost, dbPort, dbUser, dbPassword, dbName)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	println("Connected to database")

	// Keys
	SecretKey := os.Getenv("SECRET_KEY")

	// Create User Repository
	userRepository := repository.NewPostgresUserRepository(db)
	userService := service.NewUserService(userRepository)
	userHandler := handler.NewUserHandler(userService, SecretKey)

	// Create Router
	router := mux.NewRouter()

	// Public routes
	router.HandleFunc("/login", userHandler.LoginUserHandler).Methods("POST")
	router.HandleFunc("/register", userHandler.RegisterUserHandler).Methods("POST")

	// Protected routes
	protected := router.PathPrefix("/secure").Subrouter()

	// Use middleware for protected routes
	authMiddleware := middleware.NewAuthMiddleware(SecretKey)
	protected.Use(authMiddleware.JWTAuthMiddleware)
	protected.HandleFunc("/getuser/{userId:[0-9]+}", userHandler.GetUserHandler).Methods("GET")

	http.ListenAndServe(":8081", router)
}
