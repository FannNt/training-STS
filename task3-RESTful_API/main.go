package main

import (
	"log"
	"net/http"
	"os"

	"book-api/handlers"
	"book-api/storage"
)

func main() {
	// Initialize storage
	bookStorage := storage.NewMemoryStorage()
	
	// Initialize handlers
	bookHandler := handlers.NewBookHandler(bookStorage)
	
	// Setup routes
	mux := http.NewServeMux()
	
	// Book routes
	mux.HandleFunc("/api/books", bookHandler.HandleBooks)
	mux.HandleFunc("/api/books/", bookHandler.HandleBookByID)
	
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
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}
		
		next.ServeHTTP(w, r)
	})
}
