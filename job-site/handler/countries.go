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

type Country interface {
	GetAllCountries(service.Service) endpoint.Endpoint
}

// Helps to fetch all the countries in the application
func (e Endpoints) GetAllCountries(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		var countries []models.Country
		if err := e.DB.ReadCountries(&countries); err != nil {
			level.Error(logger.GokitLogger(err)).Log()
			response = models.ResponseMessage{Data: "", Error: err.Error(), Code: http.StatusBadRequest, Message: "Invalid request"}
			return response, nil
		}
		response = models.ResponseMessage{Data: countries, Error: "", Code: http.StatusOK, Message: "Fetching all countries"}
		return
	}
}


