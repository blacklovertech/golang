package handlers

import (
	"github.com/blacklovertech/02-rest-api-notes/internal/database"
	"github.com/blacklovertech/02-rest-api-notes/internal/models"
	"github.com/gin-gonic/gin"
)

// ListNotes returns all notes
func ListNotes(c *gin.Context) {
	var notes []models.Note
	result := database.DB.Find(&notes)

	if result.Error != nil {
		c.JSON(500, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(200, notes)
}

// GetNote returns a single note by ID
func GetNote(c *gin.Context) {
	var note models.Note
	id := c.Param("id")

	if err := database.DB.First(&note, id).Error; err != nil {
		c.JSON(404, gin.H{"error": "Note not found"})
		return
	}

	c.JSON(200, note)
}

// CreateNote creates a new note
func CreateNote(c *gin.Context) {
	var note models.Note

	if err := c.ShouldBindJSON(&note); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := database.DB.Create(&note).Error; err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(201, note)
}

// UpdateNote updates an existing note
func UpdateNote(c *gin.Context) {
	var note models.Note
	id := c.Param("id")

	if err := database.DB.First(&note, id).Error; err != nil {
		c.JSON(404, gin.H{"error": "Note not found"})
		return
	}

	if err := c.ShouldBindJSON(&note); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := database.DB.Save(&note).Error; err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, note)
}

// DeleteNote deletes a note
func DeleteNote(c *gin.Context) {
	id := c.Param("id")

	if err := database.DB.Delete(&models.Note{}, id).Error; err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(204, nil)
}
