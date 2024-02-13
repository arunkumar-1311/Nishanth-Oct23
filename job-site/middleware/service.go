package middleware

import (
	"job-post/adaptor"
	"job-post/service"
)

type Middleware struct {
	service.Service
	DB adaptor.Database
}
