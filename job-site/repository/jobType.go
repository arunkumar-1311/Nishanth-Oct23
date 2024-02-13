package repository

import "job-post/models"

type JobType interface {
	ReadAllJobTypes(*[]models.JobType) error
}

// Helps to read all the job types in the application
func (d *DB_Connection) ReadAllJobTypes(dest *[]models.JobType) error {
	if err := d.DB.Find(&dest).Error; err != nil {
		return err
	}
	return nil
}
