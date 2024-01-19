package repository

import (
	"blog_post/adaptor"
	"blog_post/models"
	"time"
)

func DateFilter(fromDate, toDate time.Time, Posts *[]models.Post) error {
	if result := adaptor.GetConn().Where("created_at between ? and ?", fromDate, toDate).Find(&Posts); result.Error != nil {
		return result.Error
	}
	return nil
}

func CategoryFilter(category string, Posts *[]models.Post) error {
	if result := adaptor.GetConn().Where("? = ANY(category_id) ", category).Find(&Posts); result.Error != nil {
		return result.Error
	}
	return nil
}
