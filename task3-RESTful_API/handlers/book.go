package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
	"time"

	"book-api/models"
	"book-api/storage"
	"book-api/utils"
)

// BookHandler handles HTTP requests for book operations
type BookHandler struct {
	storage storage.BookStorage
}

// NewBookHandler creates a new book handler
func NewBookHandler(storage storage.BookStorage) *BookHandler {
	return &BookHandler{storage: storage}
}

// HandleBooks handles requests to /api/books
func (h *BookHandler) HandleBooks(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.getAllBooks(w, r)
	case http.MethodPost:
		h.createBook(w, r)
	default:
		utils.WriteErrorResponse(w, http.StatusMethodNotAllowed, "Method not allowed")
	}
}

// HandleBookByID handles requests to /api/books/{id}
func (h *BookHandler) HandleBookByID(w http.ResponseWriter, r *http.Request) {
	// Extract ID from URL path
	path := strings.TrimPrefix(r.URL.Path, "/api/books/")
	if path == "" {
		utils.WriteErrorResponse(w, http.StatusBadRequest, "Book ID is required")
		return
	}
	
	id, err := strconv.Atoi(path)
	if err != nil {
		utils.WriteErrorResponse(w, http.StatusBadRequest, "Invalid book ID")
		return
	}
	
	switch r.Method {
	case http.MethodGet:
		h.getBookByID(w, r, id)
	case http.MethodPut:
		h.updateBook(w, r, id)
	case http.MethodDelete:
		h.deleteBook(w, r, id)
	default:
		utils.WriteErrorResponse(w, http.StatusMethodNotAllowed, "Method not allowed")
	}
}

// getAllBooks retrieves all books
func (h *BookHandler) getAllBooks(w http.ResponseWriter, r *http.Request) {
	books, err := h.storage.GetAll()
	if err != nil {
		utils.WriteErrorResponse(w, http.StatusInternalServerError, "Failed to retrieve books")
		return
	}
	
	utils.WriteJSONResponse(w, http.StatusOK, map[string]interface{}{
		"books": books,
		"count": len(books),
	})
}

// getBookByID retrieves a specific book
func (h *BookHandler) getBookByID(w http.ResponseWriter, r *http.Request, id int) {
	book, err := h.storage.GetByID(id)
	if err != nil {
		utils.WriteErrorResponse(w, http.StatusNotFound, "Book not found")
		return
	}
	
	utils.WriteJSONResponse(w, http.StatusOK, book)
}

// createBook creates a new book
func (h *BookHandler) createBook(w http.ResponseWriter, r *http.Request) {
	var req models.CreateBookRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.WriteErrorResponse(w, http.StatusBadRequest, "Invalid JSON payload")
		return
	}
	
	if err := req.Validate(); err != nil {
		utils.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}
	
	// Parse published date
	publishedAt, err := time.Parse("2006-01-02", req.PublishedAt)
	if err != nil {
		utils.WriteErrorResponse(w, http.StatusBadRequest, "Invalid published_at format. Use YYYY-MM-DD")
		return
	}
	
	book := &models.Book{
		Title:       req.Title,
		Author:      req.Author,
		ISBN:        req.ISBN,
		PublishedAt: publishedAt,
	}
	
	if err := h.storage.Create(book); err != nil {
		utils.WriteErrorResponse(w, http.StatusInternalServerError, "Failed to create book")
		return
	}
	
	utils.WriteJSONResponse(w, http.StatusCreated, book)
}

// updateBook updates an existing book
func (h *BookHandler) updateBook(w http.ResponseWriter, r *http.Request, id int) {
	// Check if book exists
	existingBook, err := h.storage.GetByID(id)
	if err != nil {
		utils.WriteErrorResponse(w, http.StatusNotFound, "Book not found")
		return
	}
	
	var req models.UpdateBookRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.WriteErrorResponse(w, http.StatusBadRequest, "Invalid JSON payload")
		return
	}
	
	// Update fields if provided
	updatedBook := *existingBook
	if req.Title != nil {
		updatedBook.Title = *req.Title
	}
	if req.Author != nil {
		updatedBook.Author = *req.Author
	}
	if req.ISBN != nil {
		updatedBook.ISBN = *req.ISBN
	}
	if req.PublishedAt != nil {
		publishedAt, err := time.Parse("2006-01-02", *req.PublishedAt)
		if err != nil {
			utils.WriteErrorResponse(w, http.StatusBadRequest, "Invalid published_at format. Use YYYY-MM-DD")
			return
		}
		updatedBook.PublishedAt = publishedAt
	}
	
	if err := h.storage.Update(id, &updatedBook); err != nil {
		utils.WriteErrorResponse(w, http.StatusInternalServerError, "Failed to update book")
		return
	}
	
	// Get updated book
	book, _ := h.storage.GetByID(id)
	utils.WriteJSONResponse(w, http.StatusOK, book)
}

// deleteBook deletes a book
func (h *BookHandler) deleteBook(w http.ResponseWriter, r *http.Request, id int) {
	if err := h.storage.Delete(id); err != nil {
		utils.WriteErrorResponse(w, http.StatusNotFound, "Book not found")
		return
	}
	
	utils.WriteJSONResponse(w, http.StatusOK, map[string]string{
		"message": "Book deleted successfully",
	})
}
