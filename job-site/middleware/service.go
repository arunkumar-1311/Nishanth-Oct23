package middleware

import "job-post/service"

type Middleware struct {
	service.Service
}

func AcquireMiddleware() Middleware {
	return Middleware{}
}
