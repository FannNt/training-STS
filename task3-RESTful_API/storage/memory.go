package storage

import (
	"errors"
	"sync"
	"time"

	"book-api/models"
)

// BookStorage defines the interface for book storage operations
type BookStorage interface {
	Create(book *models.Book) error
	GetByID(id int) (*models.Book, error)
	GetAll() ([]*models.Book, error)
	Update(id int, book *models.Book) error
	Delete(id int) error
}

// MemoryStorage implements BookStorage using in-memory storage
type MemoryStorage struct {
	books  map[int]*models.Book
	nextID int
	mutex  sync.RWMutex
}

// NewMemoryStorage creates a new memory storage instance
func NewMemoryStorage() *MemoryStorage {
	return &MemoryStorage{
		books:  make(map[int]*models.Book),
		nextID: 1,
	}
}

// Create adds a new book to storage
func (s *MemoryStorage) Create(book *models.Book) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	
	book.ID = s.nextID
	book.CreatedAt = time.Now()
	book.UpdatedAt = time.Now()
	
	s.books[book.ID] = book
	s.nextID++
	
	return nil
}

// GetByID retrieves a book by its ID
func (s *MemoryStorage) GetByID(id int) (*models.Book, error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	
	book, exists := s.books[id]
	if !exists {
		return nil, errors.New("book not found")
	}
	
	return book, nil
}

// GetAll retrieves all books
func (s *MemoryStorage) GetAll() ([]*models.Book, error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	
	books := make([]*models.Book, 0, len(s.books))
	for _, book := range s.books {
		books = append(books, book)
	}
	
	return books, nil
}

// Update modifies an existing book
func (s *MemoryStorage) Update(id int, updatedBook *models.Book) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	
	book, exists := s.books[id]
	if !exists {
		return errors.New("book not found")
	}
	
	// Update fields
	book.Title = updatedBook.Title
	book.Author = updatedBook.Author
	book.ISBN = updatedBook.ISBN
	book.PublishedAt = updatedBook.PublishedAt
	book.UpdatedAt = time.Now()
	
	return nil
}

// Delete removes a book from storage
func (s *MemoryStorage) Delete(id int) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	
	_, exists := s.books[id]
	if !exists {
		return errors.New("book not found")
	}
	
	delete(s.books, id)
	return nil
}
