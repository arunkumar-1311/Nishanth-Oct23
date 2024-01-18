package repository

import (
	"blog_post/adaptor"
	"blog_post/models"
	"fmt"

	"gorm.io/gorm"
)

// Helps to add comment
func CreateComment(data models.Comments) error {

	if result := adaptor.GetConn().Create(&data); result.Error != nil {
		return result.Error
	}
	return nil

}

// Helps to delete the comment
func DeleteComment(id string) error {
	var result *gorm.DB
	if result = adaptor.GetConn().Model(&models.Comments{}).Where("comment_id = ?", id).Unscoped().Delete(&models.Comments{}); result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("invalid comment id")
	}
	return nil
}

// helps to read all the comments by single user
func ReadCommentByUser(id string, dest *[]models.Comments) error {
	var result *gorm.DB
	if result = adaptor.GetConn().Model(&models.Comments{}).Where("user_id = ?", id).Find(&dest); result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("no comment added by %v", id)
	}
	return nil
}

// Helps to update the comment
func UpdateComment(id string, content models.Comments) error {
	var result *gorm.DB
	if result = adaptor.GetConn().Where("comment_id = ?", id).Updates(&content); result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("invalid comment id")
	}
	return nil
}

// Helps to read the comment with its unique id
func ReadComment(id string, dest *models.Comments) error {
	var result *gorm.DB
	if result = adaptor.GetConn().Model(&models.Comments{}).Where("comment_id = ?", id).Find(&dest); result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("invalid user id")
	}
	return nil
}
