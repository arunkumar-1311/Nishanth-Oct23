package handler

import (
	"job-post/adaptor"
	"job-post/middleware"
)

type Endpoints struct {
	DB            adaptor.Database
	Authorization middleware.Authorization
}

type API interface {
	Account
	Comments
	Country
	Jobs
	JobType
}

type PageNotFound struct{}
// Helps to set the struct to interface
func AcqurieAPI(endpoint Endpoints) API {
	return &endpoint
}
