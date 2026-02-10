package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Note struct {
	ID      uint   `json:"id" gorm:"primaryKey"`
	Title   string `json:"title" binding:"required"`
	Content string `json:"content"`
}

var db *gorm.DB

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system environment variables")
	}

	// Database connection
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		dsn = "host=localhost user=postgres password=postgres dbname=notes port=5432 sslmode=disable"
	}

	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Auto migrate the schema
	db.AutoMigrate(&Note{})
	log.Println("Database migration completed")

	// Setup router
	r := gin.Default()

	// Routes
	r.GET("/notes", listNotes)
	r.GET("/notes/:id", getNote)
	r.POST("/notes", createNote)
	r.PUT("/notes/:id", updateNote)
	r.DELETE("/notes/:id", deleteNote)

	// Health check
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	// Start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server starting on port %s", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}

// List all notes
func listNotes(c *gin.Context) {
	var notes []Note
	result := db.Find(&notes)
	
	if result.Error != nil {
		c.JSON(500, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(200, notes)
}

// Get single note by ID
func getNote(c *gin.Context) {
	var note Note
	id := c.Param("id")

	if err := db.First(&note, id).Error; err != nil {
		c.JSON(404, gin.H{"error": "Note not found"})
		return
	}

	c.JSON(200, note)
}

// Create new note
func createNote(c *gin.Context) {
	var note Note

	if err := c.ShouldBindJSON(&note); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := db.Create(&note).Error; err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(201, note)
}

// Update existing note
func updateNote(c *gin.Context) {
	var note Note
	id := c.Param("id")

	if err := db.First(&note, id).Error; err != nil {
		c.JSON(404, gin.H{"error": "Note not found"})
		return
	}

	if err := c.ShouldBindJSON(&note); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := db.Save(&note).Error; err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, note)
}

// Delete note
func deleteNote(c *gin.Context) {
	id := c.Param("id")

	if err := db.Delete(&Note{}, id).Error; err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(204, nil)
}
