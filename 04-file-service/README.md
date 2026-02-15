# File Upload/Download Service 📤

**Phase 2: Intermediate** | **Duration:** 3-4 days | **Status:** ⬜ Not Started

## Description

API service for uploading and downloading files with JWT authentication, rate limiting, and S3-compatible storage support.

## Features

- File upload with size limits
- Store files locally or S3-compatible storage
- Generate time-limited download links
- JWT authentication
- File metadata tracking in database
- Rate limiting

## Learning Outcomes

- Multipart file uploads
- JWT authentication middleware
- File storage strategies
- Presigned URLs
- Rate limiting middleware
- Database relationships

## Tech Stack

- Gin/Echo framework
- PostgreSQL
- JWT (`golang-jwt/jwt`)
- MinIO or AWS S3 SDK (optional)

## Project Structure

```
04-file-service/
├── cmd/api/main.go
├── internal/
│   ├── handlers/
│   ├── middleware/
│   └── models/
├── uploads/
├── docker-compose.yml
└── README.md
```

## Quick Start

```bash
cd 04-file-service
go mod init github.com/blacklovertech/04-file-service
go get github.com/gin-gonic/gin gorm.io/gorm github.com/golang-jwt/jwt/v5
go run cmd/api/main.go
```

## API Endpoints

```
POST   /upload          - Upload file
GET    /files/:id       - Download file
GET    /files           - List user's files
DELETE /files/:id       - Delete file
POST   /auth/login      - User login
POST   /auth/register   - User registration
```
