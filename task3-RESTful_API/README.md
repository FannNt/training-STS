# Book RESTful API

A RESTful API for managing books built with Go, PostgreSQL, and GORM.

## Features

- **CRUD Operations**: Create, Read, Update, Delete books
- **Authentication**: Token-based authentication with bcrypt password hashing
- **Protected Endpoints**: All book operations require authentication
- **PostgreSQL Database**: Persistent storage with GORM ORM
- **Database Migrations**: Automated schema management with golang-migrate
- **RESTful Design**: Follows REST conventions
- **Clean Architecture**: Separated concerns with models, handlers, storage, middleware, and utilities
- **Session Management**: Database-backed sessions with 24-hour expiration
- **Health Check**: Endpoint to monitor API and database status
- **JSON API**: All requests and responses in JSON format
- **CORS Support**: Cross-origin requests enabled
- **Web Interface**: Simple HTML frontend for testing

## API Endpoints

### Health Check

| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/health` | Check API and database health |

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
├── database/
│   ├── db.go            # Database connection
│   ├── migrate.go       # Migration runner
│   └── seed.go          # Database seeding
├── migrations/          # SQL migration files
│   ├── 000001_create_books_table.up.sql
│   ├── 000001_create_books_table.down.sql
│   ├── 000002_create_users_table.up.sql
│   ├── 000002_create_users_table.down.sql
│   ├── 000003_create_sessions_table.up.sql
│   └── 000003_create_sessions_table.down.sql
├── models/
│   ├── book.go          # Book model with GORM tags
│   └── auth.go          # User and Session models with bcrypt
├── handlers/
│   ├── book.go          # HTTP handlers for CRUD operations
│   ├── auth.go          # Authentication handlers
│   └── health.go        # Health check handler
├── middleware/
│   └── auth.go          # Authentication middleware
├── storage/
│   ├── memory.go        # In-memory storage (legacy)
│   ├── postgres.go      # PostgreSQL book storage
│   ├── session.go       # PostgreSQL session storage
│   └── token.go         # Token storage (legacy)
├── utils/
│   └── response.go      # JSON response utilities
├── fe/
│   └── index.html       # Web frontend for testing
├── docs/
│   ├── docs.go          # Swagger documentation
│   └── swagger.json     # OpenAPI specification
├── .env.example         # Environment variables template
├── go.mod               # Go module file
└── README.md            # This file
```

## Quick Start

### 1. Setup PostgreSQL Database

Install and start PostgreSQL on your system:

```bash
# Ubuntu/Debian
sudo apt-get install postgresql postgresql-contrib
sudo systemctl start postgresql

# macOS
brew install postgresql
brew services start postgresql

# Create database
psql -U postgres -c "CREATE DATABASE bookapi;"
```

### 2. Set Environment Variables

```bash
# Copy environment template
cp .env.example .env

# Edit .env with your PostgreSQL credentials
# DB_HOST=localhost
# DB_PORT=5432
# DB_USER=postgres
# DB_PASSWORD=your_password
# DB_NAME=bookapi
# DB_SSLMODE=disable
```

### 3. Run the Application

```bash
# Install dependencies
go mod tidy

# Run the server
go run main.go
```

The application will:
- Connect to PostgreSQL
- Run database migrations automatically
- Seed initial users (admin, user1, testuser)
- Start server on port 8080

### 4. Access the Application

- **Web Interface**: Open `fe/index.html` in your browser
- **Login**: Use `admin` / `admin123`

## Test Credentials

The following users are automatically seeded in the database:

| Username | Password |
|----------|----------|
| admin | admin123 |
| user1 | password1 |
| testuser | test123 |

**Note**: Passwords are hashed using bcrypt before storage.

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

## Database Schema

### Books Table
- Stores book information with unique ISBN constraint
- Indexes on ISBN, title, and author for fast queries

### Users Table
- Stores user credentials with bcrypt hashed passwords
- Unique username constraint

### Sessions Table
- Stores authentication tokens with 24-hour expiration
- Automatic cleanup of expired sessions

## Notes

- **Persistent Storage**: All data is stored in PostgreSQL
- **Password Security**: Passwords are hashed using bcrypt
- **Session Expiration**: Tokens expire after 24 hours
- **Date Format**: `published_at` should be YYYY-MM-DD
- **CORS**: Enabled for cross-origin requests
- **Response Format**: All responses are in JSON
- **Token Format**: 32-character hex strings

## Production Considerations

- Use environment variables for sensitive configuration
- Enable SSL/TLS for database connections
- Set up database backups
- Configure connection pooling
- Implement rate limiting
- Add monitoring and logging
- Use a secrets manager for credentials
