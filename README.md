# Go Learning Journey 🚀

A comprehensive roadmap for learning Go (Golang) from scratch to building production-ready systems. This repository tracks my progress from PHP/Laravel/React/Next.js/FastAPI/Flask background to mastering Go for system engineering.

## 📋 Table of Contents

- [Learning Goals](#learning-goals)
- [Why Go?](#why-go)
- [Timeline Overview](#timeline-overview)
- [Phase 1: Foundations](#phase-1-foundations-week-1-2)
- [Phase 2: Intermediate](#phase-2-intermediate-week-3-4)
- [Phase 3: Advanced Projects](#phase-3-advanced-projects-week-5-8)
- [Phase 4: Production Ready](#phase-4-production-ready-week-9-12)
- [Additional Projects](#additional-project-ideas)
- [Resources](#resources)
- [Progress Tracker](#progress-tracker)

---

## 🎯 Learning Goals

- Master Go fundamentals and idioms
- Build production-ready REST APIs
- Work with Docker and Kubernetes programmatically
- Create system daemons and CLI tools
- Implement WebSocket-based real-time applications
- Deploy containerized Go applications

## 🤔 Why Go?

Coming from PHP/Laravel/FastAPI/Flask, Go offers:

- **Performance:** Compiled language, 10-100x faster than interpreted languages
- **Concurrency:** Built-in goroutines for handling thousands of concurrent operations
- **Single Binary:** No dependency hell, easy deployment
- **System Programming:** Perfect for CLI tools, daemons, infrastructure
- **Cloud Native:** Docker, Kubernetes, and most DevOps tools are written in Go
- **Simple Syntax:** Easier learning curve compared to Rust
- **Strong Ecosystem:** Excellent libraries for web APIs, databases, networking

## 📅 Timeline Overview

| Phase | Duration | Focus Area | Projects |
|-------|----------|------------|----------|
| **Phase 1** | Week 1-2 | Foundations | 3 projects |
| **Phase 2** | Week 3-4 | Intermediate + Docker | 3 projects |
| **Phase 3** | Week 5-8 | Advanced (Main Ideas) | 3 major projects |
| **Phase 4** | Week 9-12 | Production Systems | 3 projects |

**Total Duration:** 12 weeks (3 months)

---

## 🏁 Phase 1: Foundations (Week 1-2)

**Goal:** Learn Go syntax, basic data structures, file I/O, and simple HTTP servers

### Project 1: CLI Task Manager
**Duration:** 2-3 days  
**Status:** ⬜ Not Started | 🟡 In Progress | ✅ Completed

**Description:**  
A command-line task management application with persistence.

**Features:**
- Add, list, complete, and delete tasks
- Save tasks to JSON file
- Filter tasks by status (pending/completed)
- Mark tasks with priority levels

**Learning Outcomes:**
- [ ] Go syntax and basic types
- [ ] Structs and methods
- [ ] File I/O operations
- [ ] JSON marshaling/unmarshaling
- [ ] Command-line argument parsing
- [ ] Error handling

**Tech Stack:**
- Pure Go (no frameworks)
- `encoding/json` for persistence
- `flag` package for CLI arguments

**Folder:** `01-cli-task-manager/`

---

### Project 2: REST API - Notes/Todo App
**Duration:** 3-4 days  
**Status:** ⬜ Not Started | 🟡 In Progress | ✅ Completed

**Description:**  
A RESTful API for managing notes with PostgreSQL database.

**Features:**
- CRUD operations (Create, Read, Update, Delete)
- PostgreSQL database integration
- JSON request/response handling
- Input validation
- Basic error handling

**Learning Outcomes:**
- [ ] HTTP handlers and routing
- [ ] Gin or Echo framework basics
- [ ] PostgreSQL with `database/sql` or `pgx`
- [ ] Environment variables configuration
- [ ] RESTful API design
- [ ] Middleware basics

**Tech Stack:**
- Gin or Echo framework
- PostgreSQL
- `godotenv` for env variables

**Endpoints:**
```
GET    /notes          - List all notes
GET    /notes/:id      - Get single note
POST   /notes          - Create note
PUT    /notes/:id      - Update note
DELETE /notes/:id      - Delete note
```

**Folder:** `02-rest-api-notes/`

---

### Project 3: System Info CLI
**Duration:** 2-3 days  
**Status:** ⬜ Not Started | 🟡 In Progress | ✅ Completed

**Description:**  
Cross-platform CLI tool to display system information.

**Features:**
- Display CPU, memory, disk usage
- Show running processes
- Network interface information
- Cross-platform support (Windows/Linux)
- Formatted output (tables or JSON)

**Learning Outcomes:**
- [ ] Working with system calls
- [ ] Third-party packages (`gopsutil`)
- [ ] Cross-platform development
- [ ] Formatted console output
- [ ] Goroutines for concurrent data fetching

**Tech Stack:**
- `github.com/shirou/gopsutil` for system stats
- `github.com/olekukonko/tablewriter` for formatting

**Folder:** `03-system-info-cli/`

---

## 🔧 Phase 2: Intermediate (Week 3-4)

**Goal:** HTTP middleware, authentication, file handling, Redis, Docker SDK

### Project 4: File Upload/Download Service
**Duration:** 3-4 days  
**Status:** ⬜ Not Started | 🟡 In Progress | ✅ Completed

**Description:**  
API service for uploading and downloading files with authentication.

**Features:**
- File upload with size limits
- Store files locally or S3-compatible storage
- Generate time-limited download links
- JWT authentication
- File metadata tracking in database
- Rate limiting

**Learning Outcomes:**
- [ ] Multipart file uploads
- [ ] JWT authentication middleware
- [ ] File storage strategies
- [ ] Presigned URLs (if using S3)
- [ ] Rate limiting middleware
- [ ] Database relationships

**Tech Stack:**
- Gin/Echo framework
- PostgreSQL
- JWT (`golang-jwt/jwt`)
- MinIO or AWS S3 SDK (optional)

**Folder:** `04-file-service/`

---

### Project 5: URL Shortener with Redis
**Duration:** 3-4 days  
**Status:** ⬜ Not Started | 🟡 In Progress | ✅ Completed

**Description:**  
URL shortening service with analytics and caching.

**Features:**
- Shorten long URLs to unique codes
- Redirect from short URL to original
- Click tracking and analytics
- Redis caching for fast lookups
- Rate limiting per IP
- Custom short codes (optional)

**Learning Outcomes:**
- [ ] Redis integration
- [ ] Caching strategies
- [ ] Analytics and metrics
- [ ] Hash generation (short codes)
- [ ] Database + cache patterns
- [ ] Rate limiting with Redis

**Tech Stack:**
- Gin/Echo framework
- PostgreSQL (URL storage)
- Redis (caching + rate limiting)
- `go-redis/redis`

**Folder:** `05-url-shortener/`

---

### Project 6: Simple Docker Container Manager
**Duration:** 4-5 days  
**Status:** ⬜ Not Started | 🟡 In Progress | ✅ Completed

**Description:**  
CLI tool to manage Docker containers.

**Features:**
- List all containers with status
- Start, stop, restart containers
- View container logs (real-time)
- Show container resource stats (CPU, memory)
- Remove containers
- Pull images

**Learning Outcomes:**
- [ ] Docker SDK for Go
- [ ] Working with Docker API
- [ ] Streaming logs
- [ ] Context and cancellation
- [ ] Real-time stats monitoring
- [ ] CLI design patterns

**Tech Stack:**
- `github.com/docker/docker` (Docker SDK)
- `cobra` for CLI (optional)

**Folder:** `06-docker-manager-cli/`

---

## 🚀 Phase 3: Advanced Projects (Week 5-8)

**Goal:** Build the three main project ideas with production-quality code

### Project 7: WebSocket Chat/Notification System ⭐
**Duration:** 7-10 days  
**Status:** ⬜ Not Started | 🟡 In Progress | ✅ Completed

**Description:**  
Real-time WebSocket-based communication system with React frontend.

**Features:**
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

**Learning Outcomes:**
- [ ] WebSocket protocol implementation
- [ ] Gorilla WebSocket library
- [ ] Connection pooling and management
- [ ] Redis pub/sub for multi-server scaling
- [ ] Concurrent connection handling
- [ ] Message broadcasting patterns
- [ ] Frontend WebSocket integration

**Tech Stack:**
- Gin framework + `gorilla/websocket`
- PostgreSQL (message history)
- Redis (pub/sub, presence)
- React/Next.js frontend
- JWT authentication

**Architecture:**
```
Client (React) <--WebSocket--> Go Server <--> Redis Pub/Sub
                                    ↓
                              PostgreSQL
```

**Folder:** `07-websocket-chat/`

---

### Project 8: Windows Service Monitor Daemon ⭐
**Duration:** 7-10 days  
**Status:** ⬜ Not Started | 🟡 In Progress | ✅ Completed

**Description:**  
System daemon to monitor and auto-restart Windows services with web dashboard.

**Features:**
- Background daemon running as Windows service
- Monitor configured Windows services
- Auto-restart failed services
- Configurable health check intervals
- REST API for configuration
- Email/Slack notifications on failures
- Service status dashboard (React)
- Logs of all service actions
- Start/stop monitoring via API

**Learning Outcomes:**
- [ ] Windows service creation in Go
- [ ] Windows API calls (`golang.org/x/sys/windows`)
- [ ] Service management and monitoring
- [ ] Background task scheduling
- [ ] Email/Slack integration
- [ ] System daemon patterns
- [ ] Alert mechanisms

**Tech Stack:**
- `golang.org/x/sys/windows` for Windows APIs
- `kardianos/service` for cross-platform service management
- Gin for REST API
- PostgreSQL for logs
- React dashboard

**Folder:** `08-windows-service-monitor/`

---

### Project 9: Docker Reseller Control Panel with Kubernetes ⭐⭐
**Duration:** 14-21 days  
**Status:** ⬜ Not Started | 🟡 In Progress | ✅ Completed

**Description:**  
Multi-tenant container management platform (like a hosting reseller panel).

**Features:**
- Multi-tenant architecture (admin + clients)
- Provision Docker containers per client
- Kubernetes integration for orchestration
- Resource quotas (CPU, memory, storage per client)
- Container templates (WordPress, Node.js, databases)
- Real-time resource usage monitoring
- Billing/usage tracking system
- Client portal to manage their containers
- Admin dashboard for oversight
- Container logs access
- Auto-scaling based on load
- Backup and restore functionality

**Learning Outcomes:**
- [ ] Docker SDK advanced usage
- [ ] Kubernetes client-go library
- [ ] Multi-tenancy architecture
- [ ] Resource isolation and quotas
- [ ] Billing and metering systems
- [ ] Complex authorization (RBAC)
- [ ] Container orchestration
- [ ] Microservices patterns

**Tech Stack:**
- Gin framework
- Docker SDK
- Kubernetes client-go
- PostgreSQL (users, billing, metadata)
- Redis (caching, sessions)
- Prometheus (metrics)
- React admin + client dashboards

**Architecture:**
```
React Admin ----\
                 \
React Client --> Go API --> Docker SDK --> Docker Engine
                     |
                     ├--> Kubernetes API --> K8s Cluster
                     |
                     └--> PostgreSQL + Redis
```

**Folder:** `09-docker-reseller-panel/`

---

## 🏭 Phase 4: Production Ready (Week 9-12)

**Goal:** Build production-grade infrastructure tools

### Project 10: API Gateway
**Duration:** 5-7 days  
**Status:** ⬜ Not Started | 🟡 In Progress | ✅ Completed

**Description:**  
Reverse proxy and API gateway for microservices.

**Features:**
- Route requests to multiple backend services
- Load balancing (round-robin, least connections)
- Rate limiting per endpoint
- Authentication and authorization
- Request/response logging
- Health checks for backends
- Circuit breaker pattern
- Request transformation
- CORS handling

**Learning Outcomes:**
- [ ] Reverse proxy implementation
- [ ] Load balancing algorithms
- [ ] Circuit breaker pattern
- [ ] Middleware chains
- [ ] Service discovery
- [ ] Health check mechanisms

**Tech Stack:**
- `net/http/httputil` for reverse proxy
- Redis for rate limiting
- Consul or etcd for service discovery (optional)

**Folder:** `10-api-gateway/`

---

### Project 11: Log Aggregator/Processor
**Duration:** 6-8 days  
**Status:** ⬜ Not Started | 🟡 In Progress | ✅ Completed

**Description:**  
Centralized log collection and analysis system.

**Features:**
- Collect logs from multiple sources (files, syslog, HTTP)
- Parse and structure log data
- Store in Elasticsearch or PostgreSQL
- Search and filter API
- Real-time log streaming
- Alerting on patterns/errors
- Web dashboard for analysis
- Log retention policies

**Learning Outcomes:**
- [ ] File tailing and watching
- [ ] Log parsing and structuring
- [ ] Elasticsearch integration
- [ ] Stream processing
- [ ] Pattern matching
- [ ] Time-series data handling

**Tech Stack:**
- `fsnotify` for file watching
- Elasticsearch or PostgreSQL
- Gin for API
- WebSocket for real-time streaming

**Folder:** `11-log-aggregator/`

---

### Project 12: Metrics Collection Agent
**Duration:** 5-7 days  
**Status:** ⬜ Not Started | 🟡 In Progress | ✅ Completed

**Description:**  
Lightweight monitoring agent for system metrics.

**Features:**
- Collect CPU, memory, disk, network metrics
- Export to Prometheus format
- Run as system service (systemd/Windows service)
- Plugin system for custom metrics
- Configurable collection intervals
- Low resource footprint
- Multiple export formats (Prometheus, JSON, InfluxDB)

**Learning Outcomes:**
- [ ] System metrics collection
- [ ] Prometheus exporter format
- [ ] Time-series data
- [ ] Running as system service
- [ ] Plugin architecture
- [ ] Efficient data collection

**Tech Stack:**
- `gopsutil` for metrics
- `prometheus/client_golang`
- `kardianos/service`

**Folder:** `12-metrics-agent/`

---

## 💡 Additional Project Ideas

For future exploration after completing the main timeline:

### Infrastructure & DevOps
- **Backup Automation Tool** - Scheduled backups with compression, encryption, and rotation
- **SSH Tunnel Manager** - Manage SSH tunnels and port forwarding
- **Certificate Manager** - Auto-renew SSL certificates (Let's Encrypt)
- **DNS Server** - Custom DNS resolver
- **Load Balancer** - TCP/HTTP load balancer with health checks

### Data & Processing
- **Job Queue System** - Background job processor with workers (like Sidekiq)
- **Database Migration Tool** - Version control for database schemas
- **ETL Pipeline** - Extract, transform, load data from multiple sources
- **Data Sync Service** - Sync data between different databases/APIs

### Networking & Security
- **VPN Server** - WireGuard-based VPN
- **Firewall Manager** - iptables/nftables management API
- **Network Scanner** - Port scanner and service detection
- **Rate Limiter Service** - Distributed rate limiting

### Monitoring & Observability
- **Health Check Monitor** - Ping endpoints and alert on failures
- **Uptime Monitor** - Website/service uptime tracking
- **Distributed Tracing** - Request tracing across microservices
- **Alerting System** - Flexible alerting with multiple channels

---

## 📚 Resources

### Official Documentation
- [Go Official Website](https://go.dev)
- [Go Tour](https://go.dev/tour) - Interactive tutorial
- [Go by Example](https://gobyexample.com) - Annotated examples
- [Effective Go](https://go.dev/doc/effective_go) - Best practices

### Books
- **Free:**
  - [The Go Programming Language Specification](https://go.dev/ref/spec)
  - [Go 101](https://go101.org) - Free online book
- **Paid:**
  - "The Go Programming Language" by Donovan & Kernighan
  - "Learning Go" by Jon Bodner
  - "Concurrency in Go" by Katherine Cox-Buday

### Video Courses
- [freeCodeCamp - Learn Go Programming](https://www.youtube.com/watch?v=YS4e4q9oBaU)
- [Tech With Tim - Go Tutorial](https://www.youtube.com/playlist?list=PLzMcBGfZo4-mtY_SE3HuzQJzuj4VlUG0q)

### Framework Documentation
- [Gin Web Framework](https://gin-gonic.com/docs/)
- [Echo Framework](https://echo.labstack.com/guide/)
- [Fiber Framework](https://docs.gofiber.io/)

### Libraries & Tools
- [Docker SDK for Go](https://docs.docker.com/engine/api/sdk/)
- [Kubernetes client-go](https://github.com/kubernetes/client-go)
- [GORM](https://gorm.io/) - ORM library
- [sqlx](https://github.com/jmoiron/sqlx) - SQL extensions
- [Gorilla WebSocket](https://github.com/gorilla/websocket)
- [go-redis](https://redis.uptrace.dev/)
- [Cobra](https://github.com/spf13/cobra) - CLI framework
- [Viper](https://github.com/spf13/viper) - Configuration

### Communities
- [r/golang](https://reddit.com/r/golang)
- [Gophers Slack](https://gophers.slack.com)
- [Go Forum](https://forum.golangbridge.org/)

### Practice Platforms
- [Exercism - Go Track](https://exercism.org/tracks/go)
- [LeetCode](https://leetcode.com/) - Filter by Go
- [HackerRank - Go](https://www.hackerrank.com/domains/tutorials/10-days-of-go)

---

## 📊 Progress Tracker

### Overall Progress
- **Started:** [YYYY-MM-DD]
- **Current Phase:** Phase 1
- **Completed Projects:** 0/12
- **Completion:** 0%

### Phase Completion
- [ ] Phase 1: Foundations (0/3 projects)
- [ ] Phase 2: Intermediate (0/3 projects)
- [ ] Phase 3: Advanced (0/3 projects)
- [ ] Phase 4: Production (0/3 projects)

### Skills Checklist

#### Core Go Concepts
- [ ] Variables, types, and constants
- [ ] Functions and methods
- [ ] Structs and interfaces
- [ ] Pointers
- [ ] Arrays and slices
- [ ] Maps
- [ ] Error handling
- [ ] Defer, panic, recover

#### Concurrency
- [ ] Goroutines
- [ ] Channels
- [ ] Select statement
- [ ] Mutex and sync primitives
- [ ] Context package
- [ ] Worker pools

#### Web Development
- [ ] HTTP handlers
- [ ] Routing
- [ ] Middleware
- [ ] JSON encoding/decoding
- [ ] Template rendering
- [ ] WebSockets
- [ ] Authentication (JWT)
- [ ] CORS handling

#### Database
- [ ] database/sql basics
- [ ] PostgreSQL integration
- [ ] Redis operations
- [ ] Connection pooling
- [ ] Transactions
- [ ] ORM usage (optional)

#### Testing
- [ ] Unit testing
- [ ] Table-driven tests
- [ ] Mocking
- [ ] Benchmarking
- [ ] Integration testing

#### DevOps & Deployment
- [ ] Docker containerization
- [ ] Docker SDK usage
- [ ] Kubernetes basics
- [ ] Environment configuration
- [ ] Logging
- [ ] Metrics/monitoring
- [ ] CI/CD pipeline

#### System Programming
- [ ] File I/O
- [ ] OS signals
- [ ] System calls
- [ ] Process management
- [ ] Service/daemon creation

---

## 🗂️ Repository Structure

```
go-learning-journey/
│
├── README.md                          # This file
├── .gitignore
├── docs/
│   ├── notes/                         # Learning notes
│   ├── cheatsheets/                   # Quick reference guides
│   └── resources.md                   # Additional resources
│
├── 01-cli-task-manager/
│   ├── README.md
│   ├── main.go
│   ├── task/
│   └── ...
│
├── 02-rest-api-notes/
│   ├── README.md
│   ├── main.go
│   ├── handlers/
│   ├── models/
│   ├── db/
│   └── ...
│
├── 03-system-info-cli/
├── 04-file-service/
├── 05-url-shortener/
├── 06-docker-manager-cli/
├── 07-websocket-chat/
├── 08-windows-service-monitor/
├── 09-docker-reseller-panel/
├── 10-api-gateway/
├── 11-log-aggregator/
└── 12-metrics-agent/
```

---

## 🎯 Success Criteria

By the end of this learning journey, I will be able to:

✅ Build production-ready REST APIs in Go  
✅ Work confidently with Docker and Kubernetes APIs  
✅ Create system daemons and CLI tools  
✅ Implement real-time features with WebSockets  
✅ Handle concurrent operations with goroutines  
✅ Deploy containerized Go applications  
✅ Debug and optimize Go applications  
✅ Write idiomatic Go code following best practices  

---

## 📝 Weekly Reflection Template

### Week X Reflection (YYYY-MM-DD to YYYY-MM-DD)

**Projects Completed:**
- 

**Key Learnings:**
- 

**Challenges Faced:**
- 

**Solutions Found:**
- 

**Next Week Goals:**
- 

**Time Spent:** X hours


---

## 🤝 Contributing to This Repo

This is a personal learning repository, but if you're following along:
- Feel free to fork and adapt to your needs
- Share improvements or suggestions via issues
- Document your own journey!

---

## 📜 License

MIT License - Feel free to use this roadmap for your own learning journey.

---

## 🔗 Connect

- **GitHub:** [your-username]
- **LinkedIn:** [your-profile]
- **Twitter:** [your-handle]
- **Blog:** [your-blog] (optional)

---

**Last Updated:** YYYY-MM-DD  
**Current Status:** Just getting started! 🚀

---

## 📌 Quick Start

1. **Clone this repo:**
   ```bash
   git clone https://github.com/your-username/go-learning-journey.git
   cd go-learning-journey
   ```

2. **Install Go:**
   - Download from [go.dev/dl](https://go.dev/dl/)
   - Verify: `go version`

3. **Start with Project 1:**
   ```bash
   mkdir 01-cli-task-manager
   cd 01-cli-task-manager
   go mod init cli-task-manager
   ```

4. **Update progress as you go!**

---

**Let's build something awesome! 🎉**
