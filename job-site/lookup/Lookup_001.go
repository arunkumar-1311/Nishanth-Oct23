package lookup

import (
	"job-post/models"

	"gorm.io/gorm"
)

func (Empty) Lookup_001(db *gorm.DB) error {
	if err := db.AutoMigrate(&models.Roles{}, models.Country{}, models.JobType{}, models.Users{}, models.Post{}, models.Comment{}); err != nil {
		return err
	}

	return nil
}
