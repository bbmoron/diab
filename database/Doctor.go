package diab

import (
	"gorm.io/gorm"
)

// Doctor structure defines default user schema
type Doctor struct {
	gorm.Model
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Ready    string `json:"ready"`
}
