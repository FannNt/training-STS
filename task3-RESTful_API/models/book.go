package models

import (
	"time"
)

// Book represents a book entity
type Book struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Author      string    `json:"author"`
	ISBN        string    `json:"isbn"`
	PublishedAt time.Time `json:"published_at"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// CreateBookRequest represents the request payload for creating a book
type CreateBookRequest struct {
	Title       string `json:"title"`
	Author      string `json:"author"`
	ISBN        string `json:"isbn"`
	PublishedAt string `json:"published_at"` // Format: "2006-01-02"
}

// UpdateBookRequest represents the request payload for updating a book
type UpdateBookRequest struct {
	Title       *string `json:"title,omitempty"`
	Author      *string `json:"author,omitempty"`
	ISBN        *string `json:"isbn,omitempty"`
	PublishedAt *string `json:"published_at,omitempty"` // Format: "2006-01-02"
}

// Validate validates the create book request
func (r *CreateBookRequest) Validate() error {
	if r.Title == "" {
		return NewValidationError("title is required")
	}
	if r.Author == "" {
		return NewValidationError("author is required")
	}
	if r.ISBN == "" {
		return NewValidationError("isbn is required")
	}
	return nil
}

// ValidationError represents a validation error
type ValidationError struct {
	Message string
}

func (e *ValidationError) Error() string {
	return e.Message
}

func NewValidationError(message string) *ValidationError {
	return &ValidationError{Message: message}
}
