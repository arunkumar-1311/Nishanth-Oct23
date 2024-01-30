package repository

import (
	"blog_post/models"
)

// Helps to create the user
// Find the user by email id
type Register interface {
	Create(models.Users) error
	FindUserAndEmail(string, string) (models.Users, error)
}

// Helps to create the user
func (d *GORM_Connection) Create(user models.Users) error {
	if err := d.DB.Create(&user).Error; err != nil {
		return err
	}
	return nil
}

// Helps to find the email id is already exist
func (d *GORM_Connection) FindUserAndEmail(email, name string) (user models.Users, err error) {
	if result := d.DB.Model(models.Users{}).Where("email = ?", email).Or("name = ?", name).Scan(&user); result.Error != nil {
		return user, result.Error
	}
	return
}
