package handlers

import (
	"encoding/json"
	"net/http"
	"os"

	"book-api/models"
	"book-api/storage"
	"book-api/utils"

	"gopkg.in/yaml.v3"
)

// AuthHandler handles authentication requests
type AuthHandler struct {
	tokenStorage *storage.TokenStorage
	config       *models.Config
}

// NewAuthHandler creates a new auth handler
func NewAuthHandler(tokenStorage *storage.TokenStorage, configPath string) (*AuthHandler, error) {
	config, err := loadConfig(configPath)
	if err != nil {
		return nil, err
	}

	return &AuthHandler{
		tokenStorage: tokenStorage,
		config:       config,
	}, nil
}

// loadConfig loads the configuration from YAML file
func loadConfig(path string) (*models.Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var config models.Config
	if err := yaml.Unmarshal(data, &config); err != nil {
		return nil, err
	}

	return &config, nil
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

	// Check credentials
	if !h.validateCredentials(req.Username, req.Password) {
		utils.WriteErrorResponse(w, http.StatusUnauthorized, "Invalid username or password")
		return
	}

	// Generate token
	token, err := h.tokenStorage.GenerateToken()
	if err != nil {
		utils.WriteErrorResponse(w, http.StatusInternalServerError, "Failed to generate token")
		return
	}

	// Store token
	h.tokenStorage.StoreToken(token, req.Username)

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

	// Remove token
	if !h.tokenStorage.RemoveToken(token) {
		utils.WriteErrorResponse(w, http.StatusUnauthorized, "Invalid token")
		return
	}

	utils.WriteJSONResponse(w, http.StatusOK, map[string]string{
		"message": "Logout successful",
	})
}

// validateCredentials checks if the username and password match
func (h *AuthHandler) validateCredentials(username, password string) bool {
	for _, user := range h.config.Users {
		if user.Username == username && user.Password == password {
			return true
		}
	}
	return false
}
