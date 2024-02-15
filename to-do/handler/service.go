package handler

import (
	"to-do/adaptor"
	"to-do/service"
)

type EndPoint struct {
	DB      adaptor.Database
	Service service.ServiceMethods
}

type API interface {
	Account
}

// Helps to set the struct to interface
func AcqurieAPI(endpoint EndPoint) API {
	return &endpoint
}
