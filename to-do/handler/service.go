package handler

import (
	"to-do/adaptor"
	"to-do/middleware"
	"to-do/service"
)

type EndPoint struct {
	DB      adaptor.Database
	Service service.ServiceMethods
	MW      middleware.MiddleWareInterface
}

type API interface {
	Account
}

// Helps to set the struct to interface
func AcqurieAPI(endpoint EndPoint) API {
	return &endpoint
}
