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

type Summary interface {
	GetSummary(service.Service) endpoint.Endpoint
}

// Helps to get the summary of the application
func (e Endpoints) GetSummary(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		summary, err := e.DB.ReadSummary()
		if err != nil {
			level.Error(logger.GokitLogger(err)).Log()
			response = models.ResponseMessage{Data: "", Error: err.Error(), Code: http.StatusInternalServerError, Message: "Try again later"}
			return response, nil
		}
		response = models.ResponseMessage{Data: summary, Error: "", Code: http.StatusOK, Message: "Summary of the application"}
		return
	}
}
