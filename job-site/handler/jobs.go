package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"job-post/logger"
	"job-post/models"
	"job-post/service"
	"net/http"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/log/level"
	"github.com/go-playground/validator/v10"
)

type Jobs interface {
	PostJob(service.Service) endpoint.Endpoint
	DecodePostJobRequest(context.Context, *http.Request) (interface{}, error)
}

// Helps to post new job
func (e Endpoints) PostJob(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		post, ok := request.(models.Post)
		if !ok {
			level.Debug(logger.GokitLogger(fmt.Errorf("can't assert the request"))).Log()
			response = models.ResponseMessage{Data: request, Error: "Invalid Request", Code: http.StatusBadRequest, Message: ""}
			return response, nil
		}

		validate := validator.New()
		if err := validate.Struct(post); err != nil {
			level.Debug(logger.GokitLogger(err)).Log()
			response = models.ResponseMessage{Data: "", Error: err.Error(), Code: http.StatusBadRequest, Message: "Invalid request"}
			return response, nil
		}

		if err := e.DB.CreateJob(post); err != nil {
			level.Debug(logger.GokitLogger(err)).Log()
			response = models.ResponseMessage{Data: "", Error: err.Error(), Code: http.StatusBadRequest, Message: "Invalid request"}
			return response, nil
		}
		response = models.ResponseMessage{Data: "", Error: "", Code: http.StatusOK, Message: "Post created successfully"}
		return
	}
}

func (e Endpoints) DecodePostJobRequest(_ context.Context, r *http.Request) (request interface{}, err error) {

	var claims models.Claims
	var service service.Service
	if err := service.Claims(r.Header.Get("Authorization"), &claims); err != nil {
		return err.Error(), nil
	}

	var jobs models.Post
	if err := json.NewDecoder(r.Body).Decode(&jobs); err != nil {
		return err.Error(), nil
	}
	jobs.UsersID = claims.UsersID
	return
}
