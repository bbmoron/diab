package diab

import (
	schemas "diab/database"
	"encoding/hex"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// DoctorAuth is used for doctor's access
func DoctorAuth(c *gin.Context, db *gorm.DB) {
	email := c.PostForm("email")
	passowrd := c.PostForm("password")
	shaed := NewSHA256([]byte(password))
	var doctor schemas.Doctor
	db.Find(&doctor, "email = ? AND password = ?", email, hex.EncodeToString(shaed))
	response, _ := json.Marshal(doctor)
	c.String(http.StatusOK, string(response))
}
