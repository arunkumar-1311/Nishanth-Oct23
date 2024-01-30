package lookup

import (
	"blog_post/logger"
	"blog_post/models"
	"gorm.io/gorm"
)

// Create all needed table to work with the application
func (*Empty) Lookup_001(db *gorm.DB) {

	if err := db.AutoMigrate(&models.Post{}, &models.Category{}, &models.Roles{}, &models.Users{}, &models.Comments{}); err != nil {
		logger.Logging().Error(err)
		panic(err)
	}

}
