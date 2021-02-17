package diab

import (
	schemas "diab/database"
	"encoding/json"
	"net/http"
	"strconv"

	"crypto/sha256"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// EditUser handles POST request for editing users
func EditUser(c *gin.Context, db *gorm.DB) {
	UID := c.PostForm("uid")
	name := c.PostForm("name")
	email := c.PostForm("email")
	testResult := c.PostForm("testResult")
	subscribed, _ := strconv.ParseBool(c.PostForm("subscribed"))
	height := c.PostForm("height")
	weight := c.PostForm("weight")
	password := c.PostForm("password")
	h := sha256.New()
	h.Write([]byte(password))
	// Editing existing user
	var user schemas.User
	newData := schemas.User{
		Name:       name,
		Email:      email,
		Password:   string(h.Sum(nil)),
		TestResult: testResult,
		Subscribed: subscribed,
		Height:     height,
		Weight:     weight,
	}
	db.Model(&user).Select(UID).Updates(newData)
	response, _ := json.Marshal(newData)
	c.String(http.StatusOK, string(response))
}
