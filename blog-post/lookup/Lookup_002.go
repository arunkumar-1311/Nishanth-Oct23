package lookup

import (
	"blog_post/logger"
	"blog_post/models"
	"gorm.io/gorm"
)

// Insert data to the assert table
func (*Empty) Lookup_002(db *gorm.DB) {

	roles := []models.Roles{
		{RoleID: "AD1", Role: "Admin"},
		{RoleID: "US1", Role: "User"},
	}
	if err := db.Create(roles); err != nil {
		logger.Logging().Error(err)
		panic(err)
	}

}
