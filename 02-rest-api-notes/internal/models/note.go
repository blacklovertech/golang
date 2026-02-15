package models

// Note represents a note in the database
type Note struct {
	ID      uint   `json:"id" gorm:"primaryKey"`
	Title   string `json:"title" binding:"required" gorm:"not null"`
	Content string `json:"content" gorm:"type:text"`
}
