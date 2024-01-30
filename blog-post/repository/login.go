package repository

import (
	"blog_post/models"
	"fmt"
	"gorm.io/gorm"
)

// Return the profile of the user
type User interface {
	User(string, string, *models.Users) error
}

// Helps to return the user with his name
func (d *GORM_Connection) User(name string, email string, dest *models.Users) error {
	var result *gorm.DB

	if result = d.DB.Where("name = ?", name).Or("email = ?", email).Find(&dest); result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("invalid Login Please create account to login")
	}
	return nil
}
