package handlers

import (
	"encoding/json"
	"net/http"

	"book-api/models"
	"book-api/storage"
	"book-api/utils"

	"gorm.io/gorm"
)

// AuthHandler handles authentication requests
type AuthHandler struct {
	sessionStorage *storage.SessionStorage
	db             *gorm.DB
}

// NewAuthHandler creates a new auth handler
func NewAuthHandler(sessionStorage *storage.SessionStorage, db *gorm.DB) *AuthHandler {
	return &AuthHandler{
		sessionStorage: sessionStorage,
		db:             db,
	}
}

// Login handles user login
func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		utils.WriteErrorResponse(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	var req models.LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.WriteErrorResponse(w, http.StatusBadRequest, "Invalid JSON payload")
		return
	}

	if err := req.Validate(); err != nil {
		utils.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	// Check credentials from database
	if !h.validateCredentials(req.Username, req.Password) {
		utils.WriteErrorResponse(w, http.StatusUnauthorized, "Invalid username or password")
		return
	}

	// Generate token
	token, err := h.sessionStorage.GenerateToken()
	if err != nil {
		utils.WriteErrorResponse(w, http.StatusInternalServerError, "Failed to generate token")
		return
	}

	// Store token in database
	if err := h.sessionStorage.StoreToken(token, req.Username); err != nil {
		utils.WriteErrorResponse(w, http.StatusInternalServerError, "Failed to store session")
		return
	}

	utils.WriteJSONResponse(w, http.StatusOK, models.LoginResponse{
		Token:   token,
		Message: "Login successful",
	})
}

// Logout handles user logout
func (h *AuthHandler) Logout(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		utils.WriteErrorResponse(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	// Get token from Authorization header
	token := r.Header.Get("Authorization")
	if token == "" {
		utils.WriteErrorResponse(w, http.StatusBadRequest, "Authorization token is required")
		return
	}

	// Remove "Bearer " prefix if present
	if len(token) > 7 && token[:7] == "Bearer " {
		token = token[7:]
	}

	// Remove token from database
	if !h.sessionStorage.RemoveToken(token) {
		utils.WriteErrorResponse(w, http.StatusUnauthorized, "Invalid token")
		return
	}

	utils.WriteJSONResponse(w, http.StatusOK, map[string]string{
		"message": "Logout successful",
	})
}

// validateCredentials checks if the username and password match using bcrypt
func (h *AuthHandler) validateCredentials(username, password string) bool {
	var user models.User
	err := h.db.Where("username = ?", username).First(&user).Error
	
	if err != nil {
		return false
	}
	
	return user.CheckPassword(password)
}
