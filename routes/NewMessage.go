package diab

import (
	schemas "diab/database"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// NewMessage handles POST request for creating new message in chat
func NewMessage(c *gin.Context, db *gorm.DB) {
	chatID := c.PostForm("chatId")
	authorID := c.PostForm("authorId")
	content := c.PostForm("content")
	// Creating new message
	message := schemas.Message{
		ChatID:   chatID,
		AuthorID: authorID,
		Content:  content,
	}
	db.Create(&message)
	response, _ := json.Marshal(message)
	c.String(http.StatusOK, string(response))
}
