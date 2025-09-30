# Book RESTful API

A simple RESTful API for managing books built with native Go (no external frameworks).

## Features

- **CRUD Operations**: Create, Read, Update, Delete books
- **Authentication**: Token-based authentication with login/logout
- **Protected Endpoints**: All book operations require authentication
- **RESTful Design**: Follows REST conventions
- **Clean Architecture**: Separated concerns with models, handlers, storage, middleware, and utilities
- **Thread-Safe**: Concurrent-safe in-memory storage
- **JSON API**: All requests and responses in JSON format
- **CORS Support**: Cross-origin requests enabled
- **Web Interface**: Simple HTML frontend for testing

## API Endpoints

### Authentication Endpoints (Public)

| Method | Endpoint | Description |
|--------|----------|-------------|
| POST | `/api/login` | Login and get authentication token |
| POST | `/api/logout` | Logout and invalidate token |

### Book Endpoints (Protected - Requires Authentication)

| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/api/books` | Get all books |
| POST | `/api/books` | Create a new book |
| GET | `/api/books/{id}` | Get a specific book |
| PUT | `/api/books/{id}` | Update a specific book |
| DELETE | `/api/books/{id}` | Delete a specific book |

**Note**: All book endpoints require an `Authorization` header with a valid token.

## Project Structure

```
task3-RESTful_API/
├── main.go              # Entry point and server setup
├── config.yaml          # User credentials configuration
├── models/
│   ├── book.go          # Book model and data structures
│   └── auth.go          # Authentication models
├── handlers/
│   ├── book.go          # HTTP handlers for CRUD operations
│   └── auth.go          # Authentication handlers (login/logout)
├── middleware/
│   └── auth.go          # Authentication middleware
├── storage/
│   ├── memory.go        # In-memory book storage
│   └── token.go         # Token storage and management
├── utils/
│   └── response.go      # JSON response utilities
├── fe/
│   └── index.html       # Web frontend for testing
├── go.mod               # Go module file
└── README.md            # This file
```

## Running the API

1. Navigate to the project directory:
   ```bash
   cd task3-RESTful_API
   ```

2. Install dependencies:
   ```bash
   go mod tidy
   ```

3. Run the server:
   ```bash
   go run main.go
   ```

4. The server will start on port 8080 (or the PORT environment variable if set)

5. Open the web interface:
   - Open `fe/index.html` in your browser
   - Or visit `http://localhost:8080` if serving the frontend

## Test Credentials

The following users are configured in `config.yaml`:

| Username | Password |
|----------|----------|
| admin | admin123 |
| user1 | password1 |
| testuser | test123 |

## API Usage Examples

### Login
```bash
curl -X POST http://localhost:8080/api/login \
  -H "Content-Type: application/json" \
  -d '{
    "username": "admin",
    "password": "admin123"
  }'
```

**Response:**
```json
{
  "token": "a1b2c3d4e5f6...",
  "message": "Login successful"
}
```

### Create a Book (Authenticated)
```bash
TOKEN="your_token_here"

curl -X POST http://localhost:8080/api/books \
  -H "Content-Type: application/json" \
  -H "Authorization: $TOKEN" \
  -d '{
    "title": "The Go Programming Language",
    "author": "Alan Donovan",
    "isbn": "978-0134190440",
    "published_at": "2015-11-16"
  }'
```

### Get All Books (Authenticated)
```bash
curl http://localhost:8080/api/books \
  -H "Authorization: $TOKEN"
```

### Get a Specific Book (Authenticated)
```bash
curl http://localhost:8080/api/books/1 \
  -H "Authorization: $TOKEN"
```

### Update a Book (Authenticated)
```bash
curl -X PUT http://localhost:8080/api/books/1 \
  -H "Content-Type: application/json" \
  -H "Authorization: $TOKEN" \
  -d '{
    "title": "Updated Title"
  }'
```

### Delete a Book (Authenticated)
```bash
curl -X DELETE http://localhost:8080/api/books/1 \
  -H "Authorization: $TOKEN"
```

### Logout
```bash
curl -X POST http://localhost:8080/api/logout \
  -H "Authorization: $TOKEN"
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

## Authentication Flow

1. **Login**: Send username and password to `/api/login`
2. **Receive Token**: Server returns a unique authentication token
3. **Use Token**: Include token in `Authorization` header for all book operations
4. **Logout**: Send token to `/api/logout` to invalidate it

## Security Notes

⚠️ **This is a demonstration project. For production use, consider:**

- Using HTTPS/TLS for encrypted communication
- Implementing password hashing (bcrypt, argon2)
- Using JWT tokens with expiration
- Adding rate limiting for login attempts
- Storing tokens in a database or Redis
- Implementing refresh tokens
- Adding role-based access control (RBAC)

## Notes

- This API uses in-memory storage, so data and tokens will be lost when the server restarts
- Date format for `published_at` should be YYYY-MM-DD
- The API includes CORS headers for cross-origin requests
- All responses are in JSON format
- Tokens are simple random hex strings (32 characters)
- User credentials are stored in plain text in `config.yaml` (not production-ready)
