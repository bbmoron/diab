package diab

import (
	"gorm.io/gorm"
)

// Chat structure defines default chat schema
type Chat struct {
	gorm.Model
	PatientID string `json:"patientId"`
	DoctorID  string `json:"doctorId"`
}
