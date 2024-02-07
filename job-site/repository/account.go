package repository

import (
	"fmt"
	"job-post/models"

	"gorm.io/gorm"
)

type Account interface {
	CreateUser(models.Users) error
	FindUserAndEmail(string, string) (models.Users, error)
	ReadUser(string, *models.Users) error
	ReadProfile(name string, user *models.Users) error
	UpdateUser(user *models.Users) error
}

// Helps to create the user profile
func (d *GORM_Connection) CreateUser(user models.Users) error {

	if err := d.DB.Create(user).Error; err != nil {
		return err
	}
	return nil
}

// Helps to check if the user name and email id is already exist
func (d *GORM_Connection) FindUserAndEmail(email, name string) (user models.Users, err error) {
	if result := d.DB.Model(models.Users{}).Where("email = ?", email).Or("name = ?", name).Scan(&user); result.Error != nil {
		return user, result.Error
	}
	return
}

// Helps to read the user account
func (d *GORM_Connection) ReadUser(name string, user *models.Users) error {
	var result *gorm.DB
	if result = d.DB.Preload("Roles").Where("name = ?", name).Find(&user); result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("no such user exist")
	}
	return nil
}

// Helps to update the user profile
func (d *GORM_Connection) UpdateUser(user *models.Users) error {
	if err := d.DB.Where("user_id = ?", user.UserID).UpdateColumns(user).Error; err != nil {
		return err
	}
	return nil
}

// Helps to read the profile
func (d *GORM_Connection) ReadProfile(id string, user *models.Users) error {
	var result *gorm.DB
	if result = d.DB.Model(models.Users{}).Omit("users.password").Where("user_id = ?", id).Find(&user); result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("no such user exist")
	}
	return nil
}
