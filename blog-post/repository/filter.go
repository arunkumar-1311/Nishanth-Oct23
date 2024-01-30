package repository

import (
	"blog_post/models"
	"time"
)

// Helps to filter the post content
type Filters interface {
	DateFilter(time.Time, time.Time, *[]models.Post) error
	CategoryFilter(string, *[]models.Post) error
}

// Helps to filter the post with its publishered date
func (d *GORM_Connection) DateFilter(fromDate, toDate time.Time, Posts *[]models.Post) error {
	if result := d.DB.Where("created_at between ? and ?", fromDate, toDate).Find(&Posts); result.Error != nil {
		return result.Error
	}
	return nil
}

// Helps to filter the posts with its category
func (d *GORM_Connection) CategoryFilter(category string, Posts *[]models.Post) error {
	if result := d.DB.Where("? = ANY(category_id) ", category).Find(&Posts); result.Error != nil {
		return result.Error
	}
	return nil
}
