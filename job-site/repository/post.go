package repository

import "job-post/models"

type Post interface {
	CreateJob(models.Post) error
}

// Helps to create the new job post
func (d *GORM_Connection) CreateJob(post models.Post) error {
	if err := d.DB.Create(post).Error; err != nil {
		return err
	}
	return nil
}
