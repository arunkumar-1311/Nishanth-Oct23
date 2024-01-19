package repository

import (
	"blog_post/adaptor"
	"blog_post/models"
	"fmt"

	"gorm.io/gorm"
)

// Helps to return the user with his name
func User(name string, dest *models.Users) error {
	var result *gorm.DB

	if result = adaptor.GetConn().Where("name", name).Find(&dest); result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("invalid Login Please create account to login")
	}
	return nil
}
