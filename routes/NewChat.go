package diab

import (
	schemas "diab/database"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// NewChat handles POST request for creating new chat (chat init)
func NewChat(c *gin.Context, db *gorm.DB) {
	patientID := c.PostForm("patientId")
	doctorID := c.PostForm("doctorId")
	// Creating new chat
	chat := schemas.Chat{
		PatientID: patientID,
		DoctorID:  doctorID,
	}
	db.Create(&chat)
	response, _ := json.Marshal(chat)
	c.String(http.StatusOK, string(response))
}
