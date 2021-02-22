package diab

import (
	schemas "diab/database"
	"encoding/json"
	"net/http"

	"crypto/sha256"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// SignIn handles POST request for user auth
func SignIn(c *gin.Context, db *gorm.DB) {
	email := c.PostForm("email")
	password := c.PostForm("password")
	h := sha256.New()
	h.Write([]byte(password))
	var user schemas.User
	db.Find(&user, "email = ? AND password = ?", email, password)
	response, _ := json.Marshal(user)
	c.String(http.StatusOK, string(response))
}
