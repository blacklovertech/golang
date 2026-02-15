# API Gateway ⚡

**Phase 4: Production Ready** | **Duration:** 5-7 days | **Status:** ⬜ Not Started

## Description

Reverse proxy and API gateway for microservices with load balancing and circuit breaker.

## Features

- Route requests to multiple backend services
- Load balancing (round-robin, least connections)
- Rate limiting per endpoint
- Authentication and authorization
- Request/response logging
- Health checks for backends
- Circuit breaker pattern
- Request transformation
- CORS handling

## Learning Outcomes

- Reverse proxy implementation
- Load balancing algorithms
- Circuit breaker pattern
- Middleware chains
- Service discovery
- Health check mechanisms

## Tech Stack

- `net/http/httputil` for reverse proxy
- Redis for rate limiting
- Consul or etcd for service discovery (optional)

## Quick Start

```bash
cd 10-api-gateway
go mod init github.com/blacklovertech/10-api-gateway
go get github.com/gin-gonic/gin
go run cmd/gateway/main.go
```
