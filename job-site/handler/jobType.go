package handler

import (
	"context"
	"job-post/logger"
	"job-post/models"
	"job-post/service"
	"net/http"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/log/level"
)

type JobType interface {
	GetAllJobType(service.Service) endpoint.Endpoint
}

// Helps to get all the Job Type available in the application
func (e Endpoints) GetAllJobType(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		var jobType []models.JobType

		if err = e.DB.ReadAllJobTypes(&jobType); err != nil {
			level.Error(logger.GokitLogger(err)).Log()
			response = models.ResponseMessage{Data: "", Error: err.Error(), Code: http.StatusBadRequest, Message: "Invalid request"}
			return response, nil
		}
		response = models.ResponseMessage{Data: jobType, Error: "", Code: http.StatusOK, Message: "Fetching all Job Types"}
		return
	}
}


