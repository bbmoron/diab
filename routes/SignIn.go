package diab

import (
	schemas "diab/database"
	"encoding/json"
	"encoding/hex"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// SignIn handles POST request for user auth
func SignIn(c *gin.Context, db *gorm.DB) {
	email := c.PostForm("email")
	password := c.PostForm("password")
	shaed := NewSHA256([]byte(password))
	var user schemas.User
	db.Find(&user, "email = ? AND password = ?", email, hex.EncodeToString(shaed))
	response, _ := json.Marshal(user)
	c.String(http.StatusOK, string(response))
}
