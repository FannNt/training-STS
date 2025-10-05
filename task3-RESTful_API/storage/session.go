package storage

import (
	"crypto/rand"
	"encoding/hex"
	"time"

	"book-api/models"

	"gorm.io/gorm"
)

// SessionStorage manages authentication sessions in database
type SessionStorage struct {
	db *gorm.DB
}

// NewSessionStorage creates a new session storage instance
func NewSessionStorage(db *gorm.DB) *SessionStorage {
	return &SessionStorage{db: db}
}

// GenerateToken generates a simple random token
func (s *SessionStorage) GenerateToken() (string, error) {
	bytes := make([]byte, 16)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}

// StoreToken stores a token for a username
func (s *SessionStorage) StoreToken(token, username string) error {
	session := models.Session{
		Token:     token,
		Username:  username,
		ExpiresAt: time.Now().Add(24 * time.Hour), // Token expires in 24 hours
	}
	
	return s.db.Create(&session).Error
}

// ValidateToken checks if a token is valid and returns the username
func (s *SessionStorage) ValidateToken(token string) (string, bool) {
	var session models.Session
	
	err := s.db.Where("token = ? AND expires_at > ?", token, time.Now()).First(&session).Error
	
	if err != nil {
		return "", false
	}
	
	return session.Username, true
}

// RemoveToken removes a token from storage
func (s *SessionStorage) RemoveToken(token string) bool {
	result := s.db.Where("token = ?", token).Delete(&models.Session{})
	return result.RowsAffected > 0
}
