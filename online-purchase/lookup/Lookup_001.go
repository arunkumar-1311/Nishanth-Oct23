package lookup

import (
	"online-purchase/models"

	"gorm.io/gorm"
)

// Create all needed table to work with the application
func (*Empty) Lookup_001(db *gorm.DB) error {

	if err := db.AutoMigrate(&models.Roles{}, &models.OrderStatus{}, &models.Ram{}, &models.Brand{}, &models.Users{}, &models.Orders{}); err != nil {
		return err
	}
	return nil
}
