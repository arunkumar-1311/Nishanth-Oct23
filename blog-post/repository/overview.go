package repository

import (
	"blog_post/models"
	"blog_post/service"
	"time"
)

// Helps to see the overview of the profile
type Overview interface {
	Overview(*models.Overview) error
}

// Helps to return the overview of the blog post profile
func (d *GORM_Connection) Overview(data *models.Overview) (err error) {

	if result := d.DB.Model(&models.Post{}).Count(&data.TotalPost); result.Error != nil {
		return result.Error
	}

	if result := d.DB.Model(&models.Comments{}).Count(&data.TotalComments); result.Error != nil {
		return result.Error
	}

	var postDate time.Time
	if result := d.DB.Model(&models.Post{}).Select("created_at").Order("created_at ASC").First(&postDate); result.Error != nil {
		return result.Error
	}

	if err := service.TimeDifference(postDate, time.Now(), &data.OldestPost); err != nil {
		return err
	}
	return nil
}
