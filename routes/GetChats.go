package diab

import (
	schemas "diab/database"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// GetChats handles GET request for getting chats
func GetChats(c *gin.Context, db *gorm.DB) {
	id := c.Query("id")
	var chats []schemas.Chat
	db.Find(&chats, "patientId = ? OR doctorId = ?", id, id)
	response, _ := json.Marshal(chats)
	c.String(http.StatusOK, string(response))
}
