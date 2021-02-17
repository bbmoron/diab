package diab

import (
	"gorm.io/gorm"
)

// User structure defines default user schema
type User struct {
	gorm.Model
	Name       string `json:"name"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	TestResult string `json:"testResult"`
	Subscribed bool   `json:"subscribed"`
	Height     string `json:"height"`
	Weight     string `json:"weight"`
}
