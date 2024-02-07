package lookup

import (
	"online-purchase/models"

	"gorm.io/gorm"
)

func (*Empty) Lookup_002(db *gorm.DB) error {
	roles := []models.Roles{
		{RoleID: "AD1", Role: "Admin"},
		{RoleID: "USER1", Role: "User"},
	}

	if err := db.Create(roles).Error; err != nil {
		return err
	}
	return nil
}
