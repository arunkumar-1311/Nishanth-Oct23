package lookup

import (
	"blog_post/adaptor"
	"blog_post/logger"
	"blog_post/models"
)

func (*Empty) Lookup_002() {

	roles := []models.Roles{
		{RoleID: "AD1", Role :"Admin"},
		{RoleID: "US1", Role :"User"},
	}
	if err := adaptor.GetConn().Create(roles); err != nil {
		logger.Logging().Error(err)
		panic(err)
	}

}
