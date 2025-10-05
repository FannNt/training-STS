package storage

import (
	"errors"

	"book-api/models"

	"gorm.io/gorm"
)

// PostgresStorage implements BookStorage using PostgreSQL
type PostgresStorage struct {
	db *gorm.DB
}

// NewPostgresStorage creates a new PostgreSQL storage instance
func NewPostgresStorage(db *gorm.DB) *PostgresStorage {
	return &PostgresStorage{db: db}
}

// Create adds a new book to storage
func (s *PostgresStorage) Create(book *models.Book) error {
	return s.db.Create(book).Error
}

// GetByID retrieves a book by its ID
func (s *PostgresStorage) GetByID(id int) (*models.Book, error) {
	var book models.Book
	err := s.db.First(&book, id).Error
	
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("book not found")
		}
		return nil, err
	}
	
	return &book, nil
}

// GetAll retrieves all books
func (s *PostgresStorage) GetAll() ([]*models.Book, error) {
	var books []*models.Book
	err := s.db.Order("id asc").Find(&books).Error
	
	if err != nil {
		return nil, err
	}
	
	return books, nil
}

// Update modifies an existing book
func (s *PostgresStorage) Update(id int, updatedBook *models.Book) error {
	var book models.Book
	
	// Check if book exists
	if err := s.db.First(&book, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("book not found")
		}
		return err
	}
	
	// Update fields
	book.Title = updatedBook.Title
	book.Author = updatedBook.Author
	book.ISBN = updatedBook.ISBN
	book.PublishedAt = updatedBook.PublishedAt
	
	return s.db.Save(&book).Error
}

// Delete removes a book from storage
func (s *PostgresStorage) Delete(id int) error {
	result := s.db.Delete(&models.Book{}, id)
	
	if result.Error != nil {
		return result.Error
	}
	
	if result.RowsAffected == 0 {
		return errors.New("book not found")
	}
	
	return nil
}
