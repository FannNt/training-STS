# Book RESTful API

A simple RESTful API for managing books built with native Go (no external frameworks).

## Features

- **CRUD Operations**: Create, Read, Update, Delete books
- **RESTful Design**: Follows REST conventions
- **Clean Architecture**: Separated concerns with models, handlers, storage, and utilities
- **Thread-Safe**: Concurrent-safe in-memory storage
- **JSON API**: All requests and responses in JSON format
- **CORS Support**: Cross-origin requests enabled

## API Endpoints

| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/api/books` | Get all books |
| POST | `/api/books` | Create a new book |
| GET | `/api/books/{id}` | Get a specific book |
| PUT | `/api/books/{id}` | Update a specific book |
| DELETE | `/api/books/{id}` | Delete a specific book |

## Project Structure

```
task3-RESTful_API/
├── main.go              # Entry point and server setup
├── models/
│   └── book.go          # Book model and data structures
├── handlers/
│   └── book.go          # HTTP handlers for CRUD operations
├── storage/
│   └── memory.go        # In-memory storage implementation
├── utils/
│   └── response.go      # JSON response utilities
├── go.mod               # Go module file
└── README.md            # This file
```

## Running the API

1. Navigate to the project directory:
   ```bash
   cd task3-RESTful_API
   ```

2. Run the server:
   ```bash
   go run main.go
   ```

3. The server will start on port 8080 (or the PORT environment variable if set)

## API Usage Examples

### Create a Book
```bash
curl -X POST http://localhost:8080/api/books \
  -H "Content-Type: application/json" \
  -d '{
    "title": "The Go Programming Language",
    "author": "Alan Donovan",
    "isbn": "978-0134190440",
    "published_at": "2015-11-16"
  }'
```

### Get All Books
```bash
curl http://localhost:8080/api/books
```

### Get a Specific Book
```bash
curl http://localhost:8080/api/books/1
```

### Update a Book
```bash
curl -X PUT http://localhost:8080/api/books/1 \
  -H "Content-Type: application/json" \
  -d '{
    "title": "Updated Title"
  }'
```

### Delete a Book
```bash
curl -X DELETE http://localhost:8080/api/books/1
```

## Book Model

```json
{
  "id": 1,
  "title": "Book Title",
  "author": "Author Name",
  "isbn": "978-1234567890",
  "published_at": "2023-01-01T00:00:00Z",
  "created_at": "2023-12-01T10:00:00Z",
  "updated_at": "2023-12-01T10:00:00Z"
}
```

## Notes

- This API uses in-memory storage, so data will be lost when the server restarts
- Date format for `published_at` should be YYYY-MM-DD
- The API includes CORS headers for cross-origin requests
- All responses are in JSON format
