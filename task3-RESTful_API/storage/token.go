package storage

import (
	"crypto/rand"
	"encoding/hex"
	"sync"
)

// TokenStorage manages active authentication tokens
type TokenStorage struct {
	tokens map[string]string // token -> username
	mutex  sync.RWMutex
}

// NewTokenStorage creates a new token storage instance
func NewTokenStorage() *TokenStorage {
	return &TokenStorage{
		tokens: make(map[string]string),
	}
}

// GenerateToken generates a simple random token
func (s *TokenStorage) GenerateToken() (string, error) {
	bytes := make([]byte, 16)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}

// StoreToken stores a token for a username
func (s *TokenStorage) StoreToken(token, username string) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.tokens[token] = username
}

// ValidateToken checks if a token is valid and returns the username
func (s *TokenStorage) ValidateToken(token string) (string, bool) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	username, exists := s.tokens[token]
	return username, exists
}

// RemoveToken removes a token from storage
func (s *TokenStorage) RemoveToken(token string) bool {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	_, exists := s.tokens[token]
	if exists {
		delete(s.tokens, token)
	}
	return exists
}
