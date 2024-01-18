package repository

import (
	"blog_post/adaptor"
	"blog_post/models"
	"fmt"

	"gorm.io/gorm"
)

// Helps to read all data
func ReadPosts(dest *[]models.Post) error {
	if result := adaptor.GetConn().Find(&dest); result.Error != nil {
		return result.Error
	}
	return nil
}

// Helps to find all the post comments
func NumberOfComments(id string, dest *int64) error {
	var result *gorm.DB
	if result = adaptor.GetConn().Model(&models.Comments{}).Where("post_id = ?", id).Count(dest); result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("invalid post id")
	}
	return nil
}

// Helps to delete the post
func DeletePost(id string) error {
	var result *gorm.DB
	if result = adaptor.GetConn().Model(&models.Post{}).Where("post_id = ?", id).Unscoped().Delete(&models.Post{}); result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("invalid post id")
	}
	return nil
}

// Helps to update the post
func UpdatePost(id string, content models.Post) error {
	var result *gorm.DB
	if result = adaptor.GetConn().Where("post_id = ?", id).Updates(content); result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("invalid post id")
	}
	return nil
}
