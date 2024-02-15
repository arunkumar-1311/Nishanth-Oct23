package middleware

import (
	"to-do/adaptor"
	"to-do/service"
)

type Middleware struct {
	service.Service
	DB adaptor.Database
}

type MiddleWareInterface interface {
	Authentication
}

func AcquireMiddleware(db adaptor.Database) MiddleWareInterface {
	return Middleware{DB:db,}
}
