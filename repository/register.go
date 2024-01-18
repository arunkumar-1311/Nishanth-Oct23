package repository

import (
	"blog_post/adaptor"
	"blog_post/models"
)

func Create(user models.Users) error {
	if err := adaptor.GetConn().Create(&user).Error; err != nil {
		return err
	}
	return nil
}

func FindEmail(email string) (emailID string, err error) {
	if result := adaptor.GetConn().Model(models.Users{}).Where("email = ?", email).Scan(&emailID); result.Error != nil {
		return emailID, result.Error
	}
	return
}
