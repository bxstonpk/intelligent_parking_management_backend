package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
	"user_services/handler"
	"user_services/logs"
	"user_services/middleware"
	repository "user_services/repository"
	"user_services/service"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func main() {
	// Load env variables
	loadEnv()

	// Connect to database
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", dbHost, dbPort, dbUser, dbPassword, dbName)

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second, // Slow SQL threshold
			LogLevel:      logger.Info, // Log level
			Colorful:      true,        // Enable color
		},
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		panic(err)
	}

	// Migrate database
	err = db.AutoMigrate(
		&repository.Users{},
	)
	if err != nil {
		logs.Error("Failed to migrate database: " + err.Error())
		return
	}

	// Create User Repository
	userRepository := repository.NewPostgresUserRepository(db)
	userService := service.NewUserService(userRepository)
	userHandler := handler.NewUserHandler(userService)

	// Create Router
	router := mux.NewRouter()

	// Public routes
	router.HandleFunc("/login", userHandler.LoginUserHandler).Methods("POST")
	router.HandleFunc("/register", userHandler.RegisterUserHandler).Methods("POST")
	router.HandleFunc("/checktoken", userHandler.CheckTokenHandler).Methods("POST")

	// Protected routes
	protected := router.PathPrefix("/secure").Subrouter()

	// Use middleware for protected routes
	authMiddleware := middleware.NewAuthMiddleware()
	protected.Use(authMiddleware.JWTAuthMiddleware)

	// Define protected routes
	protected.HandleFunc("/getuser/{userId:[0-9]+}", userHandler.GetUserHandler).Methods("GET")
	protected.HandleFunc("/updateuser/{userId:[0-9]+}/info", userHandler.UpdateUserInfoHandler).Methods("POST")
	protected.HandleFunc("/updateuser/{userId:[0-9]+}/password", userHandler.UpdateUserPasswordHandler).Methods("POST")
	protected.HandleFunc("/updateuser/{userId:[0-9]+}/email", userHandler.UpdateUserEmailHandler).Methods("POST")
	protected.HandleFunc("/updateuser/{userId:[0-9]+}/username", userHandler.UpdateUserUsernameHandler).Methods("POST")
	protected.HandleFunc("/updateuser/{userId:[0-9]+}/profile", userHandler.UpdateUserProfileHandler).Methods("POST")
	protected.HandleFunc("/deleteuser/{userId:[0-9]+}", userHandler.DeleteUserHandler).Methods("DELETE")

	// Start server
	logs.Info("User service starting on port " + os.Getenv("API_PORT"))
	http.ListenAndServe(os.Getenv("API_PORT"), router)
}

// ENV
func loadEnv() {
	env_err := godotenv.Load()
	if env_err != nil {
		panic("Error loading .env file")
	}
}

func GetSecretKey() string {
	loadEnv()

	return os.Getenv("SECRET_KEY")
}
