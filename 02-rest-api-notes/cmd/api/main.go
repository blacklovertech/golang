package main

import (
	"log"
	"os"

	"github.com/blacklovertech/02-rest-api-notes/internal/database"
	"github.com/blacklovertech/02-rest-api-notes/internal/handlers"
	"github.com/blacklovertech/02-rest-api-notes/internal/models"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Println("⚠️  No .env file found, using system environment variables")
	}

	// Get database connection string
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		dsn = "host=localhost user=postgres password=postgres dbname=notes port=5432 sslmode=disable"
		log.Println("⚠️  Using default database connection")
	}

	// Connect to database
	if err := database.Connect(dsn); err != nil {
		log.Fatal("❌ Failed to connect to database:", err)
	}

	// Auto migrate the schema
	database.DB.AutoMigrate(&models.Note{})
	log.Println("✅ Database migration completed")

	// Setup router
	r := gin.Default()
	r.SetTrustedProxies(nil)
	// Routes
	r.GET("/notes", handlers.ListNotes)
	r.GET("/notes/:id", handlers.GetNote)
	r.POST("/notes", handlers.CreateNote)
	r.PUT("/notes/:id", handlers.UpdateNote)
	r.DELETE("/notes/:id", handlers.DeleteNote)

	// Health check
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok", "message": "Notes API is running"})
	})

	// Start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("🚀 Server starting on port %s", port)
	log.Printf("📝 API endpoints:")
	log.Printf("   GET    http://localhost:%s/notes", port)
	log.Printf("   GET    http://localhost:%s/notes/:id", port)
	log.Printf("   POST   http://localhost:%s/notes", port)
	log.Printf("   PUT    http://localhost:%s/notes/:id", port)
	log.Printf("   DELETE http://localhost:%s/notes/:id", port)

	if err := r.Run(":" + port); err != nil {
		log.Fatal("❌ Failed to start server:", err)
	}
}
