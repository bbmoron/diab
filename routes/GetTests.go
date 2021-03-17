package diab

import (
	schemas "diab/database"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// GetTests handles GET request for tests history per user
func GetTests(c *gin.Context, db *gorm.DB) {
	uid := c.Query("uid")
	var tests []schemas.Test
	db.Find(&tests, "owner_id = ?", uid)
	response, _ := json.Marshal(tests)
	c.String(http.StatusOK, string(response))
}
