package repository

import (
	"fmt"
	"job-post/models"

	"gorm.io/gorm"
)

type Post interface {
	CreateJob(models.Post) error
	ReadJobs(string, *[]models.Post) error
	ReadJob(string, *models.Post) error
	UpdateJob(string, models.Post) error
	DeleteJobByID(models.Post) error
}

// Helps to create the new job post
func (d *DB_Connection) CreateJob(post models.Post) error {
	if err := d.DB.Create(&post).Error; err != nil {
		return err
	}
	return nil
}

// Helps to read all the jobs in the portal
func (d *DB_Connection) ReadJobs(condition string, dest *[]models.Post) error {

	if err := d.DB.Model(models.Post{}).Preload("JobType").Preload("Country").Where(condition).Order("id DESC").Find(&dest).Error; err != nil {
		return err
	}
	return nil
}

// helps to read the post by its ID
func (d *DB_Connection) ReadJob(id string, dest *models.Post) error {
	var result *gorm.DB
	if result = d.DB.Model(models.Post{}).Preload("JobType").Preload("Country").Where("post_id = ?", id).Find(&dest); result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("invalid job id")
	}
	return nil
}

// Helps to update the existing post
func (d *DB_Connection) UpdateJob(id string, dest models.Post) error {
	var result *gorm.DB
	if result = d.DB.Model(models.Post{}).Where("post_id = ?", id).UpdateColumns(dest); result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("no such job post exist")
	}
	return nil
}

// Helps to delete the job post
func (d *DB_Connection) DeleteJobByID(post models.Post) error {
	var result *gorm.DB

	if result = d.DB.Delete(&post); result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("no such job post exist")
	}
	return nil
}
