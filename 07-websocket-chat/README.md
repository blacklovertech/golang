# WebSocket Chat/Notification System 💬

**Phase 3: Advanced** | **Duration:** 7-10 days | **Status:** ⬜ Not Started

## Description

Real-time WebSocket communication system with React frontend, room support, and Redis pub/sub for scaling.

## Features

- WebSocket server with room/channel support
- User authentication and presence detection
- Real-time messaging with typing indicators
- Message persistence in PostgreSQL
- Redis pub/sub for horizontal scaling
- Message history and search
- File/image sharing
- Online/offline status
- Read receipts
- React/Next.js frontend client

## Learning Outcomes

- WebSocket protocol implementation
- Gorilla WebSocket library
- Connection pooling and management
- Redis pub/sub for multi-server scaling
- Concurrent connection handling
- Message broadcasting patterns
- Frontend WebSocket integration

## Tech Stack

- Gin framework + `gorilla/websocket`
- PostgreSQL (message history)
- Redis (pub/sub, presence)
- React/Next.js frontend
- JWT authentication

## Quick Start

```bash
cd 07-websocket-chat
# Backend
cd backend
go mod init github.com/blacklovertech/07-websocket-chat
go get github.com/gin-gonic/gin github.com/gorilla/websocket
go run cmd/server/main.go

# Frontend (in another terminal)
cd frontend
npm install
npm run dev
```

## Architecture

```
Client (React) <--WebSocket--> Go Server <--> Redis Pub/Sub
                                   ↓
                             PostgreSQL
```
