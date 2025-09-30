package main

import (
	"log"
	"net/http"
	"os"

	"book-api/handlers"
	"book-api/middleware"
	"book-api/storage"
)

func main() {
	// Initialize storage
	bookStorage := storage.NewMemoryStorage()
	tokenStorage := storage.NewTokenStorage()
	
	// Initialize handlers
	bookHandler := handlers.NewBookHandler(bookStorage)
	authHandler, err := handlers.NewAuthHandler(tokenStorage, "config.yaml")
	if err != nil {
		log.Fatalf("Failed to initialize auth handler: %v", err)
	}
	
	// Setup routes
	mux := http.NewServeMux()
	
	// Auth routes (no authentication required)
	mux.HandleFunc("/api/login", authHandler.Login)
	mux.HandleFunc("/api/logout", authHandler.Logout)
	
	// Protected book routes (authentication required)
	authMiddleware := middleware.AuthMiddleware(tokenStorage)
	mux.Handle("/api/books", authMiddleware(http.HandlerFunc(bookHandler.HandleBooks)))
	mux.Handle("/api/books/", authMiddleware(http.HandlerFunc(bookHandler.HandleBookByID)))
	
	// Add CORS middleware
	handler := corsMiddleware(mux)
	
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	
	log.Printf("Server starting on port %s", port)
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
