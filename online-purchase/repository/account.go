package repository

import (
	"fmt"
	"online-purchase/models"

	"gorm.io/gorm"
)

// Helps to create the user
// Find the user by email id
type Account interface {
	CreateUser(models.Users) error
	FindUserAndEmail(string, string) (models.Users, error)
	User(string, *models.Users) error
}

// Helps to create the user
func (d *GORM_Connection) CreateUser(user models.Users) error {
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

// Helps to return the user with his name
func (d *GORM_Connection) User(name string, dest *models.Users) error {
	var result *gorm.DB

	if result = d.DB.Where("name = ?", name).Find(&dest); result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("invalid Login Please create account to login")
	}
	return nil
}
