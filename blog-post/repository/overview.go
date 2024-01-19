package repository

import (
	"blog_post/adaptor"
	"blog_post/models"
	"blog_post/service"
	"time"
)

func Overview(data *models.Overview) (err error) {

	if result := adaptor.GetConn().Model(&models.Post{}).Count(&data.TotalPost); result.Error != nil {
		return result.Error
	}

	if result := adaptor.GetConn().Model(&models.Comments{}).Count(&data.TotalComments); result.Error != nil {
		return result.Error
	}

	var postDate time.Time
	if result := adaptor.GetConn().Model(&models.Post{}).Select("created_at").Order("created_at ASC").First(&postDate); result.Error != nil {
		return result.Error
	}
	
	if err := service.TimeDifference(postDate, time.Now(), &data.OldestPost); err != nil {
		return err
	}
	return nil
}
