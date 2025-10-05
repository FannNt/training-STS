package main

import (
	"log"
	"net/http"
	"os"

	"book-api/database"
	"book-api/handlers"
	"book-api/middleware"
	"book-api/storage"
	
	httpSwagger "github.com/swaggo/http-swagger"
	_ "book-api/docs" // Import generated docs
)

// @title Book API
// @version 1.0
// @description A RESTful API for managing books with authentication
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.email support@bookapi.com

// @license.name MIT
// @license.url https://opensource.org/licenses/MIT

// @host localhost:8080
// @BasePath /

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Type "Bearer" followed by a space and JWT token.

func main() {
	// Load database configuration
	dbConfig := database.LoadConfigFromEnv()
	
	// Connect to database
	db, err := database.Connect(dbConfig)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	
	// Run migrations
	if err := database.RunMigrations(db); err != nil {
		log.Fatalf("Failed to run migrations: %v", err)
	}
	
	// Seed initial users
	if err := database.SeedUsers(db); err != nil {
		log.Fatalf("Failed to seed users: %v", err)
	}
	
	// Initialize storage
	bookStorage := storage.NewPostgresStorage(db)
	sessionStorage := storage.NewSessionStorage(db)
	
	// Initialize handlers
	bookHandler := handlers.NewBookHandler(bookStorage)
	authHandler := handlers.NewAuthHandler(sessionStorage, db)
	
	// Get SQL DB for health check
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("Failed to get SQL DB: %v", err)
	}
	
	// Setup routes
	mux := http.NewServeMux()
	
	// Swagger documentation
	mux.HandleFunc("/api/docs/", httpSwagger.WrapHandler)
	
	// Health check endpoint
	healthHandler := handlers.NewHealthCheckHandler(sqlDB)
	mux.HandleFunc("/health", healthHandler.Check)
	
	// Auth routes (no authentication required)
	mux.HandleFunc("/api/login", authHandler.Login)
	mux.HandleFunc("/api/logout", authHandler.Logout)
	
	// Protected book routes (authentication required)
	authMiddleware := middleware.AuthMiddleware(sessionStorage)
	mux.Handle("/api/books", authMiddleware(http.HandlerFunc(bookHandler.HandleBooks)))
	mux.Handle("/api/books/", authMiddleware(http.HandlerFunc(bookHandler.HandleBookByID)))
	
	// Add CORS middleware
	handler := corsMiddleware(mux)
	
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	
	log.Printf("Server starting on port %s", port)
	log.Printf("Swagger docs available at http://localhost:%s/api/docs/", port)
	log.Fatal(http.ListenAndServe(":"+port, handler))
}

// CORS middleware
func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}
		
		next.ServeHTTP(w, r)
	})
}
