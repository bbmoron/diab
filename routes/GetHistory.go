package diab

import (
	schemas "diab/database"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// GetHistory handles GET request for message history in chat
func GetHistory(c *gin.Context, db *gorm.DB) {
	chatID := c.Query("chatId")
	var messages []schemas.Message
	db.Find(&messages, "chat_id = ?", chatID)
	response, _ := json.Marshal(messages)
	c.String(http.StatusOK, string(response))
}
