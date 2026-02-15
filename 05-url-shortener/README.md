# URL Shortener with Redis 🔗

**Phase 2: Intermediate** | **Duration:** 3-4 days | **Status:** ⬜ Not Started

## Description

URL shortening service with analytics, caching, and rate limiting using Redis.

## Features

- Shorten long URLs to unique codes
- Redirect from short URL to original
- Click tracking and analytics
- Redis caching for fast lookups
- Rate limiting per IP
- Custom short codes (optional)

## Learning Outcomes

- Redis integration
- Caching strategies
- Analytics and metrics
- Hash generation (short codes)
- Database + cache patterns
- Rate limiting with Redis

## Tech Stack

- Gin/Echo framework
- PostgreSQL (URL storage)
- Redis (caching + rate limiting)
- `go-redis/redis`

## Quick Start

```bash
cd 05-url-shortener
go mod init github.com/blacklovertech/05-url-shortener
go get github.com/gin-gonic/gin gorm.io/gorm github.com/redis/go-redis/v9
docker-compose up -d
go run cmd/api/main.go
```

## API Endpoints

```
POST   /shorten         - Create short URL
GET    /:code           - Redirect to original URL
GET    /stats/:code     - Get URL analytics
GET    /urls            - List user's URLs
DELETE /:code           - Delete short URL
```
