package models

import (
	"time"
)

// Book represents a book entity
// @Description Book object with all details
type Book struct {
	ID          int       `json:"id" gorm:"primaryKey;autoIncrement" example:"1"`
	Title       string    `json:"title" gorm:"not null" example:"The Go Programming Language"`
	Author      string    `json:"author" gorm:"not null" example:"Alan Donovan"`
	ISBN        string    `json:"isbn" gorm:"uniqueIndex;not null" example:"978-0134190440"`
	PublishedAt time.Time `json:"published_at" example:"2015-10-26T00:00:00Z"`
	CreatedAt   time.Time `json:"created_at" gorm:"autoCreateTime" example:"2024-01-15T10:30:00Z"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"autoUpdateTime" example:"2024-01-15T10:30:00Z"`
}

// CreateBookRequest represents the request payload for creating a book
// @Description Request body for creating a new book
type CreateBookRequest struct {
	Title       string `json:"title" example:"Clean Code"`
	Author      string `json:"author" example:"Robert C. Martin"`
	ISBN        string `json:"isbn" example:"978-0132350884"`
	PublishedAt string `json:"published_at" example:"2008-08-01"` // Format: "2006-01-02"
}

// UpdateBookRequest represents the request payload for updating a book
// @Description Request body for updating a book (all fields optional)
type UpdateBookRequest struct {
	Title       *string `json:"title,omitempty" example:"Clean Code: A Handbook of Agile Software Craftsmanship"`
	Author      *string `json:"author,omitempty" example:"Robert C. Martin"`
	ISBN        *string `json:"isbn,omitempty" example:"978-0132350884"`
	PublishedAt *string `json:"published_at,omitempty" example:"2008-08-01"` // Format: "2006-01-02"
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

// ErrorResponse represents an error response
// @Description Error response
type ErrorResponse struct {
	Error string `json:"error" example:"Invalid request"`
}

// MessageResponse represents a success message response
// @Description Success message response
type MessageResponse struct {
	Message string `json:"message" example:"Operation successful"`
}
