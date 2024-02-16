package lookup

import (
	"gorm.io/gorm"
	"to-do/models"
)

func (Empty) Lookup_001(db *gorm.DB) error {
	if err := db.AutoMigrate(&models.Users{}, &models.Tasks{}); err != nil {
		return err
	}

	return nil
}
