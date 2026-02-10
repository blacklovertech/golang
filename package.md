# Go Packages Reference Guide
**Essential packages for System Engineering, APIs, DevOps, and Backend Development**

---

## 🚀 Web Frameworks & Routers

### **Gin** - Fast HTTP framework (most popular)
```go
import "github.com/gin-gonic/gin"
```
- Best for: REST APIs, middleware, JSON handling
- Like: Express.js for Node
- Speed: Fastest in benchmarks

### **Echo** - High performance, minimalist
```go
import "github.com/labstack/echo/v4"
```
- Best for: REST APIs, middleware
- Like: Express.js
- Great docs, easy to learn

### **Fiber** - Express-inspired web framework
```go
import "github.com/gofiber/fiber/v2"
```
- Best for: If you love Express.js syntax
- Built on Fasthttp (very fast)
- Similar API to Express

### **Chi** - Lightweight, idiomatic router
```go
import "github.com/go-chi/chi/v5"
```
- Best for: Composable routing, standard library compatible
- Minimalist approach

### **Gorilla Mux** - Powerful URL router
```go
import "github.com/gorilla/mux"
```
- Best for: Complex routing patterns
- Older but stable

---

## 🗄️ Database Drivers & ORMs

### **GORM** - Full-featured ORM (like Laravel Eloquent)
```go
import "gorm.io/gorm"
import "gorm.io/driver/postgres"
import "gorm.io/driver/mysql"
import "gorm.io/driver/sqlite"
```
- Best for: Quick CRUD, migrations, relationships
- Like: Laravel Eloquent, Django ORM
- Auto migrations, associations

### **sqlx** - Extensions to standard database/sql
```go
import "github.com/jmoiron/sqlx"
```
- Best for: Raw SQL with struct mapping
- More control than ORM
- Get/Select directly into structs

### **pgx** - PostgreSQL driver (fastest)
```go
import "github.com/jackc/pgx/v5"
```
- Best for: PostgreSQL-specific features
- Better performance than database/sql
- Connection pooling built-in

### **go-redis** - Redis client
```go
import "github.com/redis/go-redis/v9"
```
- Best for: Caching, pub/sub, queues
- Full Redis features
- Context support

### **mongo-driver** - Official MongoDB driver
```go
import "go.mongodb.org/mongo-driver/mongo"
```
- Best for: MongoDB operations
- Official driver

### **ent** - Entity framework (modern ORM)
```go
import "entgo.io/ent"
```
- Best for: Type-safe schema-based ORM
- Code generation
- Graph traversal

---

## 🔐 Authentication & Security

### **jwt-go (golang-jwt)** - JSON Web Tokens
```go
import "github.com/golang-jwt/jwt/v5"
```
- Best for: API authentication, token generation
- Standard JWT implementation

### **bcrypt** - Password hashing (built-in)
```go
import "golang.org/x/crypto/bcrypt"
```
- Best for: Secure password storage
- Standard library extension

### **oauth2** - OAuth2 client
```go
import "golang.org/x/oauth2"
```
- Best for: OAuth2 authentication (Google, GitHub, etc.)

### **casbin** - Authorization library
```go
import "github.com/casbin/casbin/v2"
```
- Best for: RBAC, ABAC, ACL
- Complex permission systems

### **argon2** - Password hashing (more secure than bcrypt)
```go
import "golang.org/x/crypto/argon2"
```
- Best for: Modern password hashing

---

## 🐳 Docker & Kubernetes

### **docker/docker** - Official Docker SDK
```go
import "github.com/docker/docker/client"
import "github.com/docker/docker/api/types"
```
- Best for: Managing Docker containers programmatically
- Full Docker API access

### **client-go** - Kubernetes client
```go
import "k8s.io/client-go/kubernetes"
```
- Best for: Kubernetes automation
- Official Kubernetes Go client

### **docker-compose** - Docker Compose in Go
```go
import "github.com/docker/compose/v2"
```
- Best for: Managing docker-compose from Go

---

## 🌐 WebSockets & Real-time

### **gorilla/websocket** - WebSocket implementation
```go
import "github.com/gorilla/websocket"
```
- Best for: WebSocket servers
- Most popular, battle-tested

### **melody** - WebSocket framework
```go
import "github.com/olahol/melody"
```
- Best for: Easy WebSocket with broadcasting
- Built on gorilla/websocket

### **centrifugo** - Real-time messaging
```go
import "github.com/centrifugal/centrifuge-go"
```
- Best for: Scalable real-time apps
- Pub/sub, presence

---

## 📝 Configuration & Environment

### **viper** - Configuration management
```go
import "github.com/spf13/viper"
```
- Best for: Config files (JSON, YAML, ENV)
- Like: dotenv but more powerful
- Hot reloading

### **godotenv** - .env file loader
```go
import "github.com/joho/godotenv"
```
- Best for: Simple .env loading
- Like: dotenv in Node.js

### **envconfig** - Environment to struct
```go
import "github.com/kelseyhightower/envconfig"
```
- Best for: Mapping env vars to structs

---

## 🛠️ CLI Tools

### **cobra** - CLI framework
```go
import "github.com/spf13/cobra"
```
- Best for: Building complex CLIs
- Used by: Docker, Kubernetes, GitHub CLI
- Subcommands, flags, auto-docs

### **urfave/cli** - Simple CLI builder
```go
import "github.com/urfave/cli/v2"
```
- Best for: Simple CLI apps
- Less verbose than Cobra

### **survey** - Interactive CLI prompts
```go
import "github.com/AlecAivazis/survey/v2"
```
- Best for: User input, selections
- Like: inquirer.js

### **color** - Terminal colors
```go
import "github.com/fatih/color"
```
- Best for: Colored output
- Easy formatting

### **progressbar** - Progress bars
```go
import "github.com/schollz/progressbar/v3"
```
- Best for: Download/upload progress

---

## 📊 Logging

### **logrus** - Structured logger
```go
import "github.com/sirupsen/logrus"
```
- Best for: Structured JSON logging
- Most popular

### **zap** - Fast structured logger (Uber)
```go
import "go.uber.org/zap"
```
- Best for: High-performance logging
- Fastest logger

### **zerolog** - Zero allocation logger
```go
import "github.com/rs/zerolog"
```
- Best for: Minimal allocations
- JSON logging

---

## ✅ Validation

### **validator** - Struct validation
```go
import "github.com/go-playground/validator/v10"
```
- Best for: Request validation
- Tag-based validation
- Used by Gin

### **govalidator** - String validators
```go
import "github.com/asaskevich/govalidator"
```
- Best for: Email, URL, IP validation

---

## 🔄 Background Jobs & Queues

### **asynq** - Distributed task queue
```go
import "github.com/hibiken/asynq"
```
- Best for: Background jobs (like Celery)
- Redis-backed
- Cron jobs, retries

### **machinery** - Asynchronous task queue
```go
import "github.com/RichardKnop/machinery/v1"
```
- Best for: Distributed task processing
- Multiple brokers (Redis, RabbitMQ)

### **watermill** - Event-driven architecture
```go
import "github.com/ThreeDotsLabs/watermill"
```
- Best for: Message streaming, pub/sub
- Multiple backends

---

## 🧪 Testing & Mocking

### **testify** - Testing toolkit
```go
import "github.com/stretchr/testify/assert"
import "github.com/stretchr/testify/mock"
```
- Best for: Assertions, mocking
- Most popular testing library

### **gomock** - Mock generation
```go
import "github.com/golang/mock/gomock"
```
- Best for: Interface mocking
- Code generation

### **httptest** - HTTP testing (built-in)
```go
import "net/http/httptest"
```
- Best for: Testing HTTP handlers
- Standard library

---

## 📅 Date & Time

### **now** - Time toolkit
```go
import "github.com/jinzhu/now"
```
- Best for: Beginning/end of day, week, month
- Useful helpers

### **carbon** - DateTime library
```go
import "github.com/golang-module/carbon/v2"
```
- Best for: Date manipulation
- Like: Moment.js

---

## 🔧 Utilities

### **cast** - Type conversions
```go
import "github.com/spf13/cast"
```
- Best for: Safe type casting
- toString, toInt, etc.

### **copier** - Struct copying
```go
import "github.com/jinzhu/copier"
```
- Best for: Deep copying structs

### **lo** - Lodash-style utilities
```go
import "github.com/samber/lo"
```
- Best for: Map, Filter, Reduce on slices
- Like: lodash for Go

### **uuid** - UUID generation
```go
import "github.com/google/uuid"
```
- Best for: Generating UUIDs

### **hashids** - Short unique IDs
```go
import "github.com/speps/go-hashids/v2"
```
- Best for: URL-safe short IDs

---

## 📨 Email

### **gomail** - Email sending
```go
import "gopkg.in/gomail.v2"
```
- Best for: SMTP email sending
- Attachments, HTML emails

### **email** - Email package
```go
import "github.com/jordan-wright/email"
```
- Best for: Simple email sending

---

## 🌍 HTTP Client

### **resty** - HTTP client
```go
import "github.com/go-resty/resty/v2"
```
- Best for: REST API calls
- Like: axios for Go
- Retry, middleware

### **req** - HTTP client
```go
import "github.com/imroc/req/v3"
```
- Best for: Advanced HTTP requests
- Auto-retry, debugging

---

## 📁 File & Path

### **afero** - Filesystem abstraction
```go
import "github.com/spf13/afero"
```
- Best for: Testable file operations
- Mock filesystems

### **filepath** - Path manipulation (built-in)
```go
import "path/filepath"
```
- Best for: Cross-platform paths

---

## 🗜️ Compression & Archive

### **compress** - Compression (built-in)
```go
import "compress/gzip"
import "archive/zip"
import "archive/tar"
```
- Best for: Gzip, zip, tar operations

---

## 🖼️ Image Processing

### **imaging** - Image manipulation
```go
import "github.com/disintegration/imaging"
```
- Best for: Resize, crop, rotate images

### **bimg** - Fast image processing (libvips)
```go
import "github.com/h2non/bimg"
```
- Best for: High-performance image ops
- Requires libvips

---

## 📄 Excel & CSV

### **excelize** - Excel files
```go
import "github.com/xuri/excelize/v2"
```
- Best for: Reading/writing Excel files
- XLSX format

### **csv** - CSV handling (built-in)
```go
import "encoding/csv"
```
- Best for: CSV read/write

---

## 🔒 Encryption

### **crypto** - Standard crypto (built-in)
```go
import "crypto/aes"
import "crypto/rsa"
import "crypto/sha256"
```
- Best for: AES, RSA, hashing

### **age** - File encryption
```go
import "filippo.io/age"
```
- Best for: Modern file encryption

---

## 📡 gRPC & Protocol Buffers

### **grpc-go** - gRPC framework
```go
import "google.golang.org/grpc"
```
- Best for: High-performance RPC
- Microservices communication

### **protobuf** - Protocol Buffers
```go
import "google.golang.org/protobuf"
```
- Best for: Data serialization

---

## 🖥️ System Programming

### **sys** - System calls (built-in extension)
```go
import "golang.org/x/sys/windows"
import "golang.org/x/sys/unix"
```
- Best for: Low-level OS operations

### **gopsutil** - System info (cross-platform)
```go
import "github.com/shirou/gopsutil/v3/cpu"
import "github.com/shirou/gopsutil/v3/mem"
import "github.com/shirou/gopsutil/v3/disk"
```
- Best for: CPU, memory, disk stats
- Like: psutil in Python

### **service** - System services
```go
import "github.com/kardianos/service"
```
- Best for: Windows services, Linux daemons
- Cross-platform service management

---

## 🔍 Monitoring & Metrics

### **prometheus/client_golang** - Prometheus metrics
```go
import "github.com/prometheus/client_golang/prometheus"
```
- Best for: Application metrics
- Standard monitoring

---

## 📋 Scraping & Parsing

### **goquery** - HTML parsing (like jQuery)
```go
import "github.com/PuerkitoBio/goquery"
```
- Best for: Web scraping
- CSS selectors

### **colly** - Web scraping framework
```go
import "github.com/gocolly/colly/v2"
```
- Best for: Advanced scraping
- Concurrent scraping

---

## 🔄 Serialization

### **json** - JSON (built-in)
```go
import "encoding/json"
```
- Best for: JSON marshaling

### **yaml.v3** - YAML
```go
import "gopkg.in/yaml.v3"
```
- Best for: YAML config files

### **toml** - TOML
```go
import "github.com/BurntSushi/toml"
```
- Best for: TOML config files

---

## 🧰 Rate Limiting & Circuit Breaker

### **rate** - Rate limiting (built-in)
```go
import "golang.org/x/time/rate"
```
- Best for: Token bucket rate limiting

### **tollbooth** - HTTP rate limiter
```go
import "github.com/didip/tollbooth/v7"
```
- Best for: API rate limiting middleware

### **gobreaker** - Circuit breaker
```go
import "github.com/sony/gobreaker"
```
- Best for: Fault tolerance

---

## 🔌 API Documentation

### **swag** - Swagger/OpenAPI
```go
import "github.com/swaggo/swag"
import "github.com/swaggo/gin-swagger"
```
- Best for: Auto-generate API docs
- Works with Gin

---

## 🎯 Dependency Injection

### **wire** - Compile-time DI (Google)
```go
import "github.com/google/wire"
```
- Best for: Type-safe dependency injection

### **dig** - Runtime DI (Uber)
```go
import "go.uber.org/dig"
```
- Best for: Dependency injection container

---

## 📦 Package Management Commands

```bash
# Initialize module
go mod init github.com/yourusername/project

# Add dependency
go get github.com/gin-gonic/gin

# Add specific version
go get github.com/gin-gonic/gin@v1.9.0

# Update all dependencies
go get -u ./...

# Tidy dependencies
go mod tidy

# Download dependencies
go mod download

# Verify dependencies
go mod verify
```

---

## 🎓 Quick Start Template

```go
package main

import (
    "github.com/gin-gonic/gin"           // Web framework
    "gorm.io/gorm"                        // ORM
    "gorm.io/driver/postgres"             // PostgreSQL
    "github.com/redis/go-redis/v9"        // Redis
    "github.com/spf13/viper"              // Config
    "github.com/sirupsen/logrus"          // Logging
    "github.com/golang-jwt/jwt/v5"        // JWT
)

func main() {
    // Your app here
}
```

---

## 📚 Essential Reading

- **Awesome Go**: github.com/avelino/awesome-go (comprehensive list)
- **Go Packages**: pkg.go.dev (official package search)
- **Go by Example**: gobyexample.com
- **Effective Go**: go.dev/doc/effective_go

---

## 🎯 Your Project Stack Recommendations

### For **Docker Reseller Panel**:
- Gin/Echo (API)
- GORM (database)
- docker/docker (Docker SDK)
- client-go (Kubernetes)
- jwt-go (auth)
- redis (caching)
- zap (logging)
- swagger (docs)

### For **Windows Service Daemon**:
- cobra (CLI)
- gopsutil (system info)
- service (Windows service)
- logrus (logging)
- viper (config)

### For **WebSocket App**:
- Gin (API)
- gorilla/websocket (WebSocket)
- redis (pub/sub)
- GORM (persistence)
- jwt-go (auth)

---

**Pro Tip**: Start with fewer packages, add as needed. Go's standard library is powerful!
