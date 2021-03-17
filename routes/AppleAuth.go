package diab

import (
	schemas "diab/database"
	"encoding/hex"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// AppleAuth handles POST request for user auth
func AppleAuth(c *gin.Context, db *gorm.DB) {
	password := c.PostForm("password")
	shaed := NewSHA256([]byte(password))
	var user schemas.User
	db.Find(&user, "password = ?", hex.EncodeToString(shaed))
	response, _ := json.Marshal(user)
	c.String(http.StatusOK, string(response))
}
