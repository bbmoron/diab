package diab

import (
	"gorm.io/gorm"
)

// Message structure defines default message schema
type Message struct {
	gorm.Model
	ChatID   string `json:"chatId"`
	AuthorID string `json:"authorId"`
	Content  string `json:"content"`
}
