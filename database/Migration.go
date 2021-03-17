package diab

import (
	"gorm.io/gorm"
)

// InitAutoMigrate allows quick migration from golang structs to SQL
func InitAutoMigrate(db *gorm.DB) {
	db.AutoMigrate(&Message{})
	db.AutoMigrate(&Chat{})
	db.AutoMigrate(&User{})
	db.AutoMigrate(&Test{})
}
