package repository

import (
	"blog_post/models"
	"fmt"
	"gorm.io/gorm"
)

// Helps to manipulate with all posts
type Post interface {
	ReadPosts(*[]models.Post) error
	NumberOfComments(string, *int64) error
	DeletePost(string) error
	UpdatePost(string, models.Post) error
	CreatePost(models.Post) error
}

// Helps to read all data
func (d *GORM_Connection) ReadPosts(dest *[]models.Post) error {
	if result := d.DB.Find(&dest); result.Error != nil {
		return result.Error
	}
	return nil
}

// Helps to find all the post comments
func (d *GORM_Connection) NumberOfComments(id string, dest *int64) error {
	var result *gorm.DB
	if result = d.DB.Model(&models.Comments{}).Where("post_id = ?", id).Count(dest); result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("invalid post id")
	}
	return nil
}

// Helps to delete the post
func (d *GORM_Connection) DeletePost(id string) error {
	var result *gorm.DB
	if result = d.DB.Model(&models.Post{}).Where("post_id = ?", id).Unscoped().Delete(&models.Post{}); result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("invalid post id")
	}
	return nil
}

// Helps to update the post
func (d *GORM_Connection) UpdatePost(id string, content models.Post) error {
	var result *gorm.DB
	if result = d.DB.Where("post_id = ?", id).Updates(content); result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("invalid post id")
	}
	return nil
}

// Helps to create the post
func (d *GORM_Connection) CreatePost(post models.Post) error {
	if err := d.DB.Create(&post).Error; err != nil {
		return err
	}
	return nil
}
