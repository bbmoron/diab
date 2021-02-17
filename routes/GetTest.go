package diab

import (
	schemas "diab/database"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// GetTest handles GET request for test
func GetTest(c *gin.Context, db *gorm.DB) {
	id := c.Query("id")
	var test schemas.Test
	db.Find(&test, id)
	response, _ := json.Marshal(test)
	c.String(http.StatusOK, string(response))
}
