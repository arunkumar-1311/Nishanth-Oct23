package repository

import (
	"fmt"
	"job-post/models"

	"gorm.io/gorm"
)

type Comments interface {
	CreateComment(models.Comment) error
	ReadComment(string, *models.Comment) error
	ReadCommentsByPost(string, *[]models.Comment) error
	UpdateComment(models.Comment) error
	DeleteCommentByID(models.Comment) error
}

// Helps to create the new comment the post
func (d *GORM_Connection) CreateComment(comment models.Comment) error {
	if err := d.DB.Create(&comment).Error; err != nil {
		return err
	}
	return nil
}

// Read comment by id
func (d *GORM_Connection) ReadComment(id string, comment *models.Comment) error {
	var result *gorm.DB
	if result = d.DB.Model(models.Comment{}).Preload("Users").Where("comment_id = ?", id).Find(&comment); result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("no such comment exist")
	}
	return nil
}

// Read comments by post
func (d *GORM_Connection) ReadCommentsByPost(id string, comment *[]models.Comment) error {
	var result *gorm.DB
	if result = d.DB.Where("post_id = ?", id).Find(&comment); result != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("no such post exist")
	}
	return nil
}

// helps to update the comment
func (d *GORM_Connection) UpdateComment(comment models.Comment) error {
	if err := d.DB.Model(models.Comment{}).Where("comment_id = ?", comment.CommentID).UpdateColumns(comment).Error; err != nil {
		return err
	}
	return nil
}

// Helps to delete the comment
func (d *GORM_Connection) DeleteCommentByID(comment models.Comment) error {
	if err := d.DB.Delete(&comment).Error; err != nil {
		return err
	}
	return nil
}
