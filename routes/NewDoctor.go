package diab

import (
	schemas "diab/database"
	"encoding/hex"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// NewDoctor creates new doctor entity
func NewDoctor(c *gin.Context, db *gorm.DB) {
	name := c.PostForm("name")
	email := c.PostForm("email")
	password := c.PostForm("password")
	shaed := NewSHA256([]byte(password))
	// Creating new doctor
	doctor := schemas.Doctor{
		Name:     name,
		Email:    email,
		Password: hex.EncodeToString(shaed),
		Ready:    "true",
	}
	db.Create(&doctor)
	response, _ := json.Marshal(doctor)
	c.String(http.StatusOK, string(response))
}
