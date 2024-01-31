package repository

import (
	"blog_post/models"
	"fmt"
	"gorm.io/gorm"
)

// Methods helps to manipulate with comments
type Comments interface {
	CreateComment(models.Comments) error
	DeleteComment(string) error
	ReadCommentByUser(string, *[]models.Comments) error
	UpdateComment(string, models.Comments) error
	ReadComment(string, *models.Comments) error
	PostComments(string, *[]models.Comments) error
}

// Helps to add comment
func (d *GORM_Connection) CreateComment(data models.Comments) error {

	if result := d.DB.Create(&data); result.Error != nil {
		return result.Error
	}
	return nil

}

// Helps to delete the comment
func (d *GORM_Connection) DeleteComment(id string) error {
	var result *gorm.DB
	if result = d.DB.Model(&models.Comments{}).Where("comment_id = ?", id).Unscoped().Delete(&models.Comments{}); result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("invalid comment id")
	}
	return nil
}

// helps to read all the comments by single user
func (d *GORM_Connection) ReadCommentByUser(id string, dest *[]models.Comments) error {

	if result := d.DB.Preload("Users").Where("user_id = ?", id).Find(&dest); result.Error != nil {
		return result.Error
	}

	return nil
}

// Helps to update the comment
func (d *GORM_Connection) UpdateComment(id string, content models.Comments) error {
	var result *gorm.DB
	if result = d.DB.Where("comment_id = ?", id).Updates(&content); result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("invalid comment id")
	}
	return nil
}

// Helps to read the comment with its unique id
func (d *GORM_Connection) ReadComment(id string, dest *models.Comments) error {
	var result *gorm.DB
	if result = d.DB.Model(&models.Comments{}).Preload("Users").Where("comment_id = ?", id).Find(&dest); result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("invalid user id")
	}
	return nil
}

// Helps to retrive the comments by its post
func (d *GORM_Connection) PostComments(id string, dest *[]models.Comments) error {
	var result *gorm.DB
	if result = d.DB.Model(&models.Comments{}).Preload("Users").Where("post_id = ?", id).Find(&dest); result.Error != nil {
		return result.Error
	}

	return nil
}
