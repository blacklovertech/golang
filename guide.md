# Go Project Structure & Sample Projects

## 📁 Standard Go Project Structure

### **Simple API Project Structure**
```
my-go-api/
├── cmd/
│   └── api/
│       └── main.go              # Application entry point
├── internal/                     # Private app code (can't be imported by others)
│   ├── handlers/                # HTTP handlers (controllers)
│   │   ├── user.go
│   │   └── auth.go
│   ├── models/                  # Database models
│   │   └── user.go
│   ├── middleware/              # Middleware functions
│   │   ├── auth.go
│   │   └── logger.go
│   ├── database/                # Database connection
│   │   └── postgres.go
│   └── config/                  # Configuration
│       └── config.go
├── pkg/                         # Public libraries (can be imported)
│   └── utils/
│       └── response.go
├── migrations/                  # Database migrations
│   ├── 001_create_users.sql
│   └── 002_create_posts.sql
├── scripts/                     # Utility scripts
│   └── seed.sh
├── docker/
│   ├── Dockerfile
│   └── docker-compose.yml
├── .env                         # Environment variables
├── .env.example
├── .gitignore
├── go.mod                       # Dependencies
├── go.sum                       # Dependency checksums
├── Makefile                     # Build automation
└── README.md
```

### **Microservice Project Structure**
```
my-microservice/
├── cmd/
│   ├── api/                     # API service
│   │   └── main.go
│   └── worker/                  # Background worker
│       └── main.go
├── internal/
│   ├── api/
│   │   ├── handlers/
│   │   ├── middleware/
│   │   └── routes.go
│   ├── domain/                  # Business logic
│   │   ├── user/
│   │   │   ├── service.go
│   │   │   ├── repository.go
│   │   │   └── model.go
│   │   └── auth/
│   ├── infrastructure/          # External services
│   │   ├── database/
│   │   ├── cache/
│   │   └── queue/
│   └── config/
├── pkg/
│   ├── logger/
│   ├── validator/
│   └── errors/
├── api/                         # API definitions
│   ├── openapi.yaml
│   └── proto/                   # gRPC proto files
├── deployments/
│   ├── docker-compose.yml
│   └── kubernetes/
│       ├── deployment.yaml
│       └── service.yaml
├── tests/
│   ├── integration/
│   └── e2e/
├── go.mod
└── README.md
```

### **CLI Tool Structure**
```
my-cli-tool/
├── cmd/
│   ├── root.go                  # Root command
│   ├── start.go                 # Subcommands
│   ├── stop.go
│   └── status.go
├── internal/
│   ├── service/
│   └── config/
├── pkg/
│   └── utils/
├── go.mod
└── README.md
```

---

## 🎯 Starter Template Projects (GitHub)

### **1. REST API with PostgreSQL (Gin + GORM)**
**Repository**: https://github.com/gothinkster/golang-gin-realworld-example-app
- **Features**: JWT auth, CRUD operations, PostgreSQL
- **Stack**: Gin, GORM, PostgreSQL
- **Perfect for**: Learning REST API patterns

### **2. Microservices Example**
**Repository**: https://github.com/micro/examples
- **Features**: Multiple microservices, service discovery
- **Stack**: Go Micro framework
- **Perfect for**: Microservice architecture

### **3. Clean Architecture Example**
**Repository**: https://github.com/bxcodec/go-clean-arch
- **Features**: Clean architecture, dependency injection
- **Stack**: Echo, MySQL
- **Perfect for**: Enterprise-level structure

### **4. Docker Management API**
**Repository**: https://github.com/portainer/portainer
- **Features**: Docker container management
- **Stack**: Go + Docker SDK
- **Perfect for**: Your Docker reseller panel idea

### **5. Kubernetes Client Example**
**Repository**: https://github.com/kubernetes/client-go/tree/master/examples
- **Features**: Official K8s examples
- **Stack**: client-go
- **Perfect for**: Kubernetes automation

### **6. WebSocket Chat Application**
**Repository**: https://github.com/gorilla/websocket/tree/master/examples/chat
- **Features**: Real-time chat with rooms
- **Stack**: Gorilla WebSocket
- **Perfect for**: Your WebSocket project

### **7. CLI Tool (Cobra)**
**Repository**: https://github.com/spf13/cobra-cli
- **Features**: Full CLI with subcommands
- **Stack**: Cobra
- **Perfect for**: System tools

### **8. Background Jobs (Asynq)**
**Repository**: https://github.com/hibiken/asynq/tree/master/example
- **Features**: Task queues, cron jobs
- **Stack**: Asynq + Redis
- **Perfect for**: Background processing

### **9. gRPC Microservice**
**Repository**: https://github.com/grpc/grpc-go/tree/master/examples
- **Features**: gRPC communication
- **Stack**: gRPC
- **Perfect for**: High-performance APIs

### **10. Full-Stack Go + React**
**Repository**: https://github.com/gothinkster/realworld
- **Features**: Full-stack example (backend in Go)
- **Stack**: Multiple implementations
- **Perfect for**: Reference implementation

---

## 🚀 Sample Project Ideas with Code

### **Project 1: Simple REST API (Todo App)**

**File: `cmd/api/main.go`**
```go
package main

import (
    "log"
    "github.com/gin-gonic/gin"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
)

type Todo struct {
    ID        uint   `json:"id" gorm:"primaryKey"`
    Title     string `json:"title"`
    Completed bool   `json:"completed"`
}

func main() {
    // Database connection
    dsn := "host=localhost user=postgres password=postgres dbname=todos port=5432"
    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatal(err)
    }
    
    // Auto migrate
    db.AutoMigrate(&Todo{})
    
    // Router
    r := gin.Default()
    
    // Routes
    r.GET("/todos", func(c *gin.Context) {
        var todos []Todo
        db.Find(&todos)
        c.JSON(200, todos)
    })
    
    r.POST("/todos", func(c *gin.Context) {
        var todo Todo
        if err := c.ShouldBindJSON(&todo); err != nil {
            c.JSON(400, gin.H{"error": err.Error()})
            return
        }
        db.Create(&todo)
        c.JSON(201, todo)
    })
    
    r.PUT("/todos/:id", func(c *gin.Context) {
        var todo Todo
        if err := db.First(&todo, c.Param("id")).Error; err != nil {
            c.JSON(404, gin.H{"error": "Todo not found"})
            return
        }
        c.ShouldBindJSON(&todo)
        db.Save(&todo)
        c.JSON(200, todo)
    })
    
    r.DELETE("/todos/:id", func(c *gin.Context) {
        db.Delete(&Todo{}, c.Param("id"))
        c.JSON(204, nil)
    })
    
    r.Run(":8080")
}
```

**File: `docker-compose.yml`**
```yaml
version: '3.8'

services:
  postgres:
    image: postgres:15
    environment:
      POSTGRES_USER: postgres
      POSTGRES_PASSWORD: postgres
      POSTGRES_DB: todos
    ports:
      - "5432:5432"
    volumes:
      - postgres_data:/var/lib/postgresql/data

  redis:
    image: redis:7-alpine
    ports:
      - "6379:6379"

  app:
    build: .
    ports:
      - "8080:8080"
    depends_on:
      - postgres
      - redis
    environment:
      DATABASE_URL: postgres://postgres:postgres@postgres:5432/todos
      REDIS_URL: redis://redis:6379

volumes:
  postgres_data:
```

**File: `Dockerfile`**
```dockerfile
FROM golang:1.22-alpine AS builder

WORKDIR /app

# Copy go mod files
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build
RUN CGO_ENABLED=0 GOOS=linux go build -o main ./cmd/api

# Final stage
FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy binary from builder
COPY --from=builder /app/main .

EXPOSE 8080

CMD ["./main"]
```

**File: `Makefile`**
```makefile
.PHONY: run build test docker-up docker-down

run:
	go run cmd/api/main.go

build:
	go build -o bin/api cmd/api/main.go

test:
	go test -v ./...

docker-up:
	docker-compose up -d

docker-down:
	docker-compose down

migrate-up:
	migrate -path migrations -database "postgresql://postgres:postgres@localhost:5432/todos?sslmode=disable" up

migrate-down:
	migrate -path migrations -database "postgresql://postgres:postgres@localhost:5432/todos?sslmode=disable" down

deps:
	go mod download
	go mod tidy

.DEFAULT_GOAL := run
```

---

### **Project 2: JWT Authentication API**

**File: `internal/handlers/auth.go`**
```go
package handlers

import (
    "time"
    "github.com/gin-gonic/gin"
    "github.com/golang-jwt/jwt/v5"
    "golang.org/x/crypto/bcrypt"
)

var jwtSecret = []byte("your-secret-key")

type Claims struct {
    UserID uint   `json:"user_id"`
    Email  string `json:"email"`
    jwt.RegisteredClaims
}

type LoginRequest struct {
    Email    string `json:"email" binding:"required,email"`
    Password string `json:"password" binding:"required"`
}

type RegisterRequest struct {
    Email    string `json:"email" binding:"required,email"`
    Password string `json:"password" binding:"required,min=6"`
    Name     string `json:"name" binding:"required"`
}

func Register(c *gin.Context) {
    var req RegisterRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(400, gin.H{"error": err.Error()})
        return
    }
    
    // Hash password
    hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
    
    // Save user to database (pseudo-code)
    // user := User{Email: req.Email, Password: string(hashedPassword), Name: req.Name}
    // db.Create(&user)
    
    c.JSON(201, gin.H{"message": "User registered successfully"})
}

func Login(c *gin.Context) {
    var req LoginRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(400, gin.H{"error": err.Error()})
        return
    }
    
    // Find user (pseudo-code)
    // var user User
    // db.Where("email = ?", req.Email).First(&user)
    
    // Verify password
    // if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
    //     c.JSON(401, gin.H{"error": "Invalid credentials"})
    //     return
    // }
    
    // Generate JWT
    claims := &Claims{
        UserID: 1,
        Email:  req.Email,
        RegisteredClaims: jwt.RegisteredClaims{
            ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
            IssuedAt:  jwt.NewNumericDate(time.Now()),
        },
    }
    
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    tokenString, _ := token.SignedString(jwtSecret)
    
    c.JSON(200, gin.H{
        "token": tokenString,
        "user": gin.H{
            "email": req.Email,
        },
    })
}
```

**File: `internal/middleware/auth.go`**
```go
package middleware

import (
    "strings"
    "github.com/gin-gonic/gin"
    "github.com/golang-jwt/jwt/v5"
)

var jwtSecret = []byte("your-secret-key")

type Claims struct {
    UserID uint   `json:"user_id"`
    Email  string `json:"email"`
    jwt.RegisteredClaims
}

func AuthMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        authHeader := c.GetHeader("Authorization")
        if authHeader == "" {
            c.JSON(401, gin.H{"error": "Authorization header required"})
            c.Abort()
            return
        }
        
        tokenString := strings.Replace(authHeader, "Bearer ", "", 1)
        
        claims := &Claims{}
        token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
            return jwtSecret, nil
        })
        
        if err != nil || !token.Valid {
            c.JSON(401, gin.H{"error": "Invalid token"})
            c.Abort()
            return
        }
        
        // Set user info in context
        c.Set("user_id", claims.UserID)
        c.Set("email", claims.Email)
        
        c.Next()
    }
}
```

---

### **Project 3: Docker Container Manager (CLI)**

**File: `cmd/docker-cli/main.go`**
```go
package main

import (
    "context"
    "fmt"
    "os"
    
    "github.com/docker/docker/api/types"
    "github.com/docker/docker/client"
    "github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
    Use:   "docker-cli",
    Short: "Docker container management CLI",
}

var listCmd = &cobra.Command{
    Use:   "list",
    Short: "List all containers",
    Run: func(cmd *cobra.Command, args []string) {
        cli, err := client.NewClientWithOpts(client.FromEnv)
        if err != nil {
            panic(err)
        }
        
        containers, err := cli.ContainerList(context.Background(), types.ContainerListOptions{All: true})
        if err != nil {
            panic(err)
        }
        
        for _, container := range containers {
            fmt.Printf("ID: %s | Name: %s | Status: %s\n", 
                container.ID[:12], 
                container.Names[0], 
                container.Status)
        }
    },
}

var startCmd = &cobra.Command{
    Use:   "start [container_id]",
    Short: "Start a container",
    Args:  cobra.ExactArgs(1),
    Run: func(cmd *cobra.Command, args []string) {
        cli, _ := client.NewClientWithOpts(client.FromEnv)
        err := cli.ContainerStart(context.Background(), args[0], types.ContainerStartOptions{})
        if err != nil {
            fmt.Println("Error:", err)
            return
        }
        fmt.Println("Container started:", args[0])
    },
}

var stopCmd = &cobra.Command{
    Use:   "stop [container_id]",
    Short: "Stop a container",
    Args:  cobra.ExactArgs(1),
    Run: func(cmd *cobra.Command, args []string) {
        cli, _ := client.NewClientWithOpts(client.FromEnv)
        timeout := 10
        err := cli.ContainerStop(context.Background(), args[0], container.StopOptions{Timeout: &timeout})
        if err != nil {
            fmt.Println("Error:", err)
            return
        }
        fmt.Println("Container stopped:", args[0])
    },
}

func main() {
    rootCmd.AddCommand(listCmd)
    rootCmd.AddCommand(startCmd)
    rootCmd.AddCommand(stopCmd)
    
    if err := rootCmd.Execute(); err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
}
```

---

### **Project 4: WebSocket Chat Server**

**File: `cmd/websocket-chat/main.go`**
```go
package main

import (
    "log"
    "net/http"
    
    "github.com/gin-gonic/gin"
    "github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
    CheckOrigin: func(r *http.Request) bool {
        return true // Allow all origins in development
    },
}

type Client struct {
    conn *websocket.Conn
    send chan []byte
}

type Hub struct {
    clients    map[*Client]bool
    broadcast  chan []byte
    register   chan *Client
    unregister chan *Client
}

func newHub() *Hub {
    return &Hub{
        clients:    make(map[*Client]bool),
        broadcast:  make(chan []byte),
        register:   make(chan *Client),
        unregister: make(chan *Client),
    }
}

func (h *Hub) run() {
    for {
        select {
        case client := <-h.register:
            h.clients[client] = true
            log.Println("Client connected. Total:", len(h.clients))
            
        case client := <-h.unregister:
            if _, ok := h.clients[client]; ok {
                delete(h.clients, client)
                close(client.send)
                log.Println("Client disconnected. Total:", len(h.clients))
            }
            
        case message := <-h.broadcast:
            for client := range h.clients {
                select {
                case client.send <- message:
                default:
                    close(client.send)
                    delete(h.clients, client)
                }
            }
        }
    }
}

func (c *Client) readPump(hub *Hub) {
    defer func() {
        hub.unregister <- c
        c.conn.Close()
    }()
    
    for {
        _, message, err := c.conn.ReadMessage()
        if err != nil {
            break
        }
        hub.broadcast <- message
    }
}

func (c *Client) writePump() {
    defer c.conn.Close()
    
    for message := range c.send {
        err := c.conn.WriteMessage(websocket.TextMessage, message)
        if err != nil {
            break
        }
    }
}

func handleWebSocket(hub *Hub, c *gin.Context) {
    conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
    if err != nil {
        log.Println(err)
        return
    }
    
    client := &Client{
        conn: conn,
        send: make(chan []byte, 256),
    }
    hub.register <- client
    
    go client.writePump()
    go client.readPump(hub)
}

func main() {
    hub := newHub()
    go hub.run()
    
    r := gin.Default()
    
    r.GET("/ws", func(c *gin.Context) {
        handleWebSocket(hub, c)
    })
    
    r.GET("/", func(c *gin.Context) {
        c.JSON(200, gin.H{"message": "WebSocket server running on /ws"})
    })
    
    log.Println("Server starting on :8080")
    r.Run(":8080")
}
```

---

## 📦 Quick Setup Commands

```bash
# Initialize new project
mkdir my-go-project
cd my-go-project
go mod init github.com/yourusername/my-go-project

# Install common dependencies
go get github.com/gin-gonic/gin
go get gorm.io/gorm
go get gorm.io/driver/postgres
go get github.com/redis/go-redis/v9
go get github.com/golang-jwt/jwt/v5
go get github.com/spf13/viper
go get github.com/joho/godotenv

# Create basic structure
mkdir -p cmd/api internal/{handlers,models,middleware,database,config} pkg/utils migrations scripts docker

# Run with hot reload (install air first: go install github.com/cosmtrek/air@latest)
air

# Build
go build -o bin/app cmd/api/main.go

# Run
./bin/app
```

---

## 🎯 Best GitHub Repos to Clone and Learn From

1. **Standard Go Project Layout**: https://github.com/golang-standards/project-layout
2. **Go Web Examples**: https://github.com/gowebexamples/gowebexamples
3. **Awesome Go**: https://github.com/avelino/awesome-go
4. **Go Design Patterns**: https://github.com/tmrts/go-patterns
5. **Go Kit** (microservices): https://github.com/go-kit/kit
6. **Buffalo** (full framework): https://github.com/gobuffalo/buffalo

---

## 🔥 Next Steps

1. Clone one of the sample repos above
2. Read the code and understand the structure
3. Build the Todo API example I provided
4. Add JWT authentication
5. Containerize with Docker
6. Move to your custom projects

**Pro Tip**: Don't overcomplicate the structure initially. Start simple, refactor as you grow!
