package repository

import (
	"fmt"
	"to-do/models"

	"gorm.io/gorm"
)

type Account interface {
	FindUserAndEmail(email, name string) (user models.Users, err error)
	NewUser(user models.Users) error
	ReadUser(user models.Login, dest *models.Users) error
}

// Helps to check if the user name and email id is already exist
func (d *DB_Connection) FindUserAndEmail(email, name string) (user models.Users, err error) {
	return user, d.DB.Model(models.Users{}).Where("email = ?", email).Or("name = ?", name).Scan(&user).Error
}

// Helps to create new user
func (d *DB_Connection) NewUser(user models.Users) error {
	return d.DB.Create(&user).Error
}

// helps to find the user profile
func (d *DB_Connection) ReadUser(user models.Login, dest *models.Users) error {
	var result *gorm.DB
	
	if result = d.DB.Where("name = ?", user.Name).Or("email = ?", user.Email).First(&dest); result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("no such user exist")
	}
	return nil
}
