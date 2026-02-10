# REST API - Notes/Todo App 📝

**Phase 1: Foundations** | **Duration:** 3-4 days | **Status:** ⬜ Not Started

## Description

A RESTful API for managing notes with PostgreSQL database integration, input validation, and comprehensive error handling.

## Features

- ✅ CRUD operations (Create, Read, Update, Delete)
- ✅ PostgreSQL database integration
- ✅ JSON request/response handling
- ✅ Input validation
- ✅ Basic error handling
- ✅ Environment variables configuration

## Learning Outcomes

- [x] HTTP handlers and routing
- [x] Gin framework basics
- [x] PostgreSQL with `database/sql` or `pgx`
- [x] Environment variables configuration
- [x] RESTful API design
- [x] Middleware basics

## Tech Stack

- **Framework**: Gin (https://gin-gonic.com)
- **Database**: PostgreSQL 15
- **ORM**: GORM
- **Config**: godotenv

## API Endpoints

```
GET    /notes          - List all notes
GET    /notes/:id      - Get single note by ID
POST   /notes          - Create new note
PUT    /notes/:id      - Update existing note
DELETE /notes/:id      - Delete note
```

## Project Structure

```
02-rest-api-notes/
├── cmd/
│   └── api/
│       └── main.go              # Application entry point
├── internal/
│   ├── handlers/                # HTTP handlers
│   │   └── notes.go
│   ├── models/                  # Database models
│   │   └── note.go
│   └── database/                # Database connection
│       └── postgres.go
├── docker-compose.yml           # PostgreSQL setup
├── Dockerfile                   # App containerization
├── Makefile                     # Build automation
├── .env.example                 # Environment template
├── .gitignore
├── go.mod                       # Dependencies
└── README.md                    # This file
```

## Quick Start

### Prerequisites

- Go 1.22+
- PostgreSQL 15+ (or use Docker)
- Make (optional but recommended)

### Setup

1. **Clone and navigate:**

   ```bash
   cd 02-rest-api-notes
   ```

2. **Install dependencies:**

   ```bash
   go mod download
   ```

3. **Setup environment:**

   ```bash
   cp .env.example .env
   # Edit .env with your database credentials
   ```

4. **Start PostgreSQL (with Docker):**

   ```bash
   docker-compose up -d
   ```

5. **Run the application:**

   ```bash
   go run cmd/api/main.go
   # OR
   make run
   ```

6. **Test the API:**
   ```bash
   curl http://localhost:8080/notes
   ```

## Environment Variables

```env
DATABASE_URL=postgres://postgres:postgres@localhost:5432/notes?sslmode=disable
PORT=8080
```

## Example Requests

### Create Note

```bash
curl -X POST http://localhost:8080/notes \
  -H "Content-Type: application/json" \
  -d '{"title":"My First Note","content":"This is a test note"}'
```

### Get All Notes

```bash
curl http://localhost:8080/notes
```

### Get Single Note

```bash
curl http://localhost:8080/notes/1
```

### Update Note

```bash
curl -X PUT http://localhost:8080/notes/1 \
  -H "Content-Type: application/json" \
  -d '{"title":"Updated Title","content":"Updated content"}'
```

### Delete Note

```bash
curl -X DELETE http://localhost:8080/notes/1
```

## Development

### Run with hot reload (using air)

```bash
# Install air first
go install github.com/cosmtrek/air@latest

# Run
air
```

### Build

```bash
make build
# OR
go build -o bin/api cmd/api/main.go
```

### Run tests

```bash
make test
# OR
go test -v ./...
```

## Docker Commands

```bash
# Start services
docker-compose up -d

# Stop services
docker-compose down

# View logs
docker-compose logs -f

# Rebuild and start
docker-compose up -d --build
```

## Next Steps

1. Add authentication (JWT)
2. Implement pagination
3. Add search functionality
4. Add tags/categories
5. Implement rate limiting
6. Add unit tests
7. Add Swagger documentation

## Resources

- [Gin Documentation](https://gin-gonic.com/docs/)
- [GORM Documentation](https://gorm.io/docs/)
- [PostgreSQL Documentation](https://www.postgresql.org/docs/)
