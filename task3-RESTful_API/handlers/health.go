package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"
)

// HealthCheckHandler handles health check requests
type HealthCheckHandler struct {
	db *sql.DB
}

// NewHealthCheckHandler creates a new health check handler
func NewHealthCheckHandler(db *sql.DB) *HealthCheckHandler {
	return &HealthCheckHandler{db: db}
}

// Check performs health check
func (h *HealthCheckHandler) Check(w http.ResponseWriter, r *http.Request) {
	response := map[string]string{
		"status":   "healthy",
		"database": "connected",
	}

	// Ping database
	if err := h.db.Ping(); err != nil {
		response["status"] = "unhealthy"
		response["database"] = "error"
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusServiceUnavailable)
		json.NewEncoder(w).Encode(response)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
