package repository

import (
	"blog_post/adaptor"
	"blog_post/models"
)

func Overview(data *models.Overview) error {

	if result := adaptor.GetConn().Model(&models.Post{}).Count(&data.TotalPost); result.Error != nil {
		return result.Error
	}

	if result := adaptor.GetConn().Model(&models.Comments{}).Count(&data.TotalComments); result.Error != nil {
		return result.Error
	}

	if result := adaptor.GetConn().Model(&models.Post{}).Select("created_at").Order("created_at ASC").First(&data.OldestPost); result.Error != nil {
		return result.Error
	}
	return nil
}
