package handler

import (
	"job-post/adaptor"
	"job-post/middleware"
)

type Endpoints struct {
	DB            adaptor.Database
	Authorization middleware.Authorization
}


