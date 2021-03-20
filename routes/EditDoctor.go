package diab

import (
	schemas "diab/database"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// EditDoctor handles ready status update for doctors
func EditDoctor(c *gin.Context, db *gorm.DB) {
	UID := c.PostForm("uid")
	ready := c.PostForm("ready")
	db.Model(&schemas.Doctor).Where("id = ?", UID).Update("ready", ready)
	response, _ := json.Marshal(schemas.Doctor{
		Ready: ready,
	})
	c.String(http.StatusOK, string(response))
}
