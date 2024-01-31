package lookup

import (
	"blog_post/logger"
	"gorm.io/gorm"
)

// Helps to create the lookup table
func (*Empty) Lookup_000(db *gorm.DB) {

	if err := db.AutoMigrate(&Lookup{}); err != nil {
		logger.Logging().Error(err)
		panic(err)
	}
}
