package repository

import (
	"blog_post/adaptor"
	"blog_post/models"
)

// Helps to create the user
func Create(user models.Users) error {
	if err := adaptor.GetConn().Create(&user).Error; err != nil {
		return err
	}
	return nil
}

// Helps to find the email id is already exist
func FindUserAndEmail(email string, name string) (user models.Users, err error) {
	if result := adaptor.GetConn().Model(models.Users{}).Where("email = ?", email).Or("name = ?", name).Scan(&user); result.Error != nil {
		return user, result.Error
	}
	return
}


