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
func (d *DB_Connection) CreateComment(comment models.Comment) error {
	var result *gorm.DB
	if result = d.DB.Create(&comment); result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("can't create the account try again later")
	}
	return nil
}

// Read comment by id
func (d *DB_Connection) ReadComment(id string, comment *models.Comment) error {
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
func (d *DB_Connection) ReadCommentsByPost(id string, comment *[]models.Comment) error {
	var result *gorm.DB
	if result = d.DB.Model(models.Comment{}).Preload("Users").Where("post_id = ?", id).Find(&comment); result != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("no such post exist")
	}
	return nil
}

// helps to update the comment
func (d *DB_Connection) UpdateComment(comment models.Comment) error {
	if err := d.DB.Model(models.Comment{}).Where("comment_id = ?", comment.CommentID).UpdateColumns(comment).Error; err != nil {
		return err
	}
	return nil
}

// Helps to delete the comment
func (d *DB_Connection) DeleteCommentByID(comment models.Comment) error {
	if err := d.DB.Delete(&comment).Error; err != nil {
		return err
	}
	return nil
}
