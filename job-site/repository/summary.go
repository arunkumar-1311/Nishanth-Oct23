package repository

import "job-post/models"

type Summary interface {
	ReadSummary() (models.Summary, error)
}

func (d *GORM_Connection) ReadSummary() (models.Summary, error) {

	var result models.Summary

	if err := d.DB.Model(models.Post{}).Select("job_title as job, count(job_title) as total").Group("job_title").Find(&result.Jobs).Error; err != nil {
		return result, err
	}

	if err := d.DB.Model(models.Country{}).Select(" count(country) as totalCountries").Scan(&result.Countries).Error; err != nil {
		return result, err
	}

	
	if err := d.DB.Model(models.Post{}).Count(&result.TotalJobs).Error; err != nil {
		return result, err
	}

	return result, nil
}
