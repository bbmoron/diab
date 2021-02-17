package diab

import (
	"gorm.io/gorm"
)

// Test structure defines default test schema
type Test struct {
	gorm.Model
	OwnerID       string `json:"ownerId"`
	Age           string `json:"age"`
	Gender        string `json:"gender"`
	Race          string `json:"race"`
	PrevDiagnosed string `json:"prevDiagnosed"`
	Relatives     string `json:"relatives"`
	BloodPressure string `json:"bloodPressure"`
	Active        string `json:"active"`
	Height        string `json:"height"`
	Weight        string `json:"weight"`
	Score         string `json:"score"`
}
