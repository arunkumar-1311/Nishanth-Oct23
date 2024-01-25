package lookup

import (
	"blog_post/adaptor"
	"blog_post/logger"
)

func (*Empty) Lookup_000() {

	if err := adaptor.GetConn().AutoMigrate(&Lookup{}); err != nil {
		logger.Logging().Error(err)
		panic(err)
	}
}
