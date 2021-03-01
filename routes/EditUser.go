package diab

import (
	schemas "diab/database"
	"encoding/json"
	"encoding/hex"
	"net/http"
	"strconv"

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
	shaed := NewSHA256([]byte(password))
	// Editing existing user
	var user schemas.User
	newData := schemas.User{
		Name:       name,
		Email:      email,
		Password:   hex.EncodeToString(shaed),
		TestResult: testResult,
		Subscribed: subscribed,
		Height:     height,
		Weight:     weight,
	}
	db.Model(&user).Select(UID).Updates(newData)
	response, _ := json.Marshal(newData)
	c.String(http.StatusOK, string(response))
}
