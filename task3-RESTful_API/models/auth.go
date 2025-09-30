package models

// User represents a user from config
type User struct {
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

// Config represents the application configuration
type Config struct {
	Users []User `yaml:"users"`
}

// LoginRequest represents the login request payload
type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// LoginResponse represents the login response
type LoginResponse struct {
	Token   string `json:"token"`
	Message string `json:"message"`
}

// LogoutRequest represents the logout request payload
type LogoutRequest struct {
	Token string `json:"token"`
}

// Validate validates the login request
func (r *LoginRequest) Validate() error {
	if r.Username == "" {
		return NewValidationError("username is required")
	}
	if r.Password == "" {
		return NewValidationError("password is required")
	}
	return nil
}
