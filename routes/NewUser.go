package diab

import (
	schemas "diab/database"
	"encoding/json"
	"encoding/hex"
	"net/http"

	"crypto/sha256"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// NewSHA256 handles sha256-ing passwords
func NewSHA256(data []byte) []byte {
	hash := sha256.Sum256(data)
	return hash[:]
}

// NewUser handles POST request for creating new users (sign up)
func NewUser(c *gin.Context, db *gorm.DB) {
	name := c.PostForm("name")
	email := c.PostForm("email")
	password := c.PostForm("password")
	shaed := NewSHA256([]byte(password))
	// Creating new user
	user := schemas.User{
		Name:     name,
		Email:    email,
		Password: hex.EncodeToString(shaed),
	}
	db.Create(&user)
	response, _ := json.Marshal(user)
	c.String(http.StatusOK, string(response))
}
