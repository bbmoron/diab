package diab

import (
	schemas "diab/database"
	"encoding/json"
	"net/http"

	"crypto/sha256"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// NewUser handles POST request for creating new users (sign up)
func NewUser(c *gin.Context, db *gorm.DB) {
	name := c.PostForm("name")
	email := c.PostForm("email")
	password := c.PostForm("password")
	h := sha256.New()
	h.Write([]byte(password))
	// Creating new user
	user := schemas.User{
		Name:     name,
		Email:    email,
		Password: string(h.Sum(nil)),
	}
	db.Create(&user)
	response, _ := json.Marshal(user)
	c.String(http.StatusOK, string(response))
}
