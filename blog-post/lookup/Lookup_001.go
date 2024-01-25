package lookup

import (
	"blog_post/adaptor"
	"blog_post/logger"
	"blog_post/models"
)

func (*Empty) Lookup_001() {

	if err := adaptor.GetConn().AutoMigrate(&models.Post{}, &models.Category{}, &models.Roles{}, &models.Users{}, &models.Comments{}); err != nil {
		logger.Logging().Error(err)
		panic(err)
	}

}
