package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"job-post/logger"
	"job-post/models"
	"job-post/service"
	"job-post/service/helper"
	"net/http"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/log/level"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

type Jobs interface {
	PostJob(service.Service) endpoint.Endpoint
	DecodePostJobRequest(context.Context, *http.Request) (interface{}, error)

	GetAllJobs(service.Service) endpoint.Endpoint
	DecodeGetAllJobsRequest(context.Context, *http.Request) (interface{}, error)

	UpdateJob(service.Service) endpoint.Endpoint
	DecodeUpdateJobRequest(context.Context, *http.Request) (interface{}, error)

	DeleteJob(service.Service) endpoint.Endpoint
	DecodeDeleteJobRequest(context.Context, *http.Request) (interface{}, error)

	GetJob(service.Service) endpoint.Endpoint
	DecodeGetID(context.Context, *http.Request) (interface{}, error)
}

// Helps to post new job
func (e Endpoints) PostJob(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		post, ok := request.(models.Post)
		if !ok {
			level.Error(logger.GokitLogger(fmt.Errorf("can't assert the request"))).Log()
			response = models.ResponseMessage{Data: request, Error: "Invalid Request", Code: http.StatusBadRequest, Message: ""}
			return response, nil
		}
		post.PostID = helper.UniqueID()
		validate := validator.New()
		if err := validate.Struct(post); err != nil {
			level.Error(logger.GokitLogger(err)).Log()
			response = models.ResponseMessage{Data: "", Error: err.Error(), Code: http.StatusBadRequest, Message: "Invalid request"}
			return response, nil
		}

		if err := e.DB.CreateJob(post); err != nil {
			level.Error(logger.GokitLogger(err)).Log()
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
	return jobs, nil
}

// Helps to read all the jobs from the portal
func (e Endpoints) GetAllJobs(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		params, ok := request.(map[string]string)
		if !ok {
			level.Error(logger.GokitLogger(fmt.Errorf("can't assert the request"))).Log()
			response = models.ResponseMessage{Data: request, Error: "Invalid Request", Code: http.StatusBadRequest, Message: ""}
			return response, nil
		}

		var jobs []models.Post
		if err = e.DB.ReadJobs(params["keyword"], params["country"], params["jobtype"], &jobs); err != nil {
			level.Error(logger.GokitLogger(err)).Log()
			response = models.ResponseMessage{Data: "", Error: "Server Error", Code: http.StatusInternalServerError, Message: "Try again later"}
			return response, nil
		}

		var jobResponse models.PostSearch
		jobResponse.Post = make([]models.PostResponse, len(jobs))
		
		for index, value := range jobs {
			marshalJob, err := json.Marshal(value)
			if err != nil {
				level.Error(logger.GokitLogger(err)).Log()
				response = models.ResponseMessage{Data: "", Error: err.Error(), Code: http.StatusInternalServerError, Message: "Try again later"}
				return response, nil
			}

			if err = json.Unmarshal(marshalJob, &jobResponse.Post[index]); err != nil {
				level.Error(logger.GokitLogger(err)).Log()
				response = models.ResponseMessage{Data: "", Error: err.Error(), Code: http.StatusInternalServerError, Message: "Try again later"}
				return response, nil
			}
		}

		if err := svc.Summary(jobResponse.Post, &jobResponse.Summary); err != nil {
			level.Error(logger.GokitLogger(err)).Log()
			response = models.ResponseMessage{Data: "", Error: "Server Error", Code: http.StatusInternalServerError, Message: "Try again later"}
			return response, nil
		}
		response = models.ResponseMessage{Data: jobResponse, Error: "", Code: http.StatusOK, Message: "All Posts in the portal"}
		return
	}
}

func (e Endpoints) DecodeGetAllJobsRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	params := make(map[string]string)
	params["keyword"] = r.URL.Query().Get("keyword")
	params["country"] = r.URL.Query().Get("country")
	params["jobtype"] = r.URL.Query().Get("jobtype")
	return params, nil
}

// Helps to update the existing job
func (e Endpoints) UpdateJob(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		post, ok := request.(models.Post)
		if !ok {
			level.Error(logger.GokitLogger(fmt.Errorf("can't assert the request"))).Log()
			response = models.ResponseMessage{Data: request, Error: "Invalid Request", Code: http.StatusBadRequest, Message: ""}
			return response, nil
		}

		var existingPost models.Post
		if err = e.DB.ReadJob(post.PostID, &existingPost); err != nil {
			level.Error(logger.GokitLogger(err)).Log()
			response = models.ResponseMessage{Data: "", Error: err.Error(), Code: http.StatusInternalServerError, Message: "Try again later"}
			return response, nil
		}

		if existingPost.UsersID != post.UsersID {
			response = models.ResponseMessage{Data: "", Error: "Unauthorized to update this post", Code: http.StatusBadRequest, Message: "Invalid Request"}
			return response, nil
		}

		if err := e.DB.UpdateJob(post.PostID, post); err != nil {
			level.Error(logger.GokitLogger(err)).Log()
			response = models.ResponseMessage{Data: "", Error: err.Error(), Code: http.StatusBadRequest, Message: "Invalid request"}
			return response, nil
		}

		if err := e.DB.ReadJob(post.PostID, &post); err != nil {
			level.Error(logger.GokitLogger(err)).Log()
			response = models.ResponseMessage{Data: "", Error: "Server Error", Code: http.StatusInternalServerError, Message: "Try again later"}
			return response, nil
		}

		response = models.ResponseMessage{Data: post, Error: "", Code: http.StatusOK, Message: "Updated Successfully"}
		return
	}
}

func (e Endpoints) DecodeUpdateJobRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
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
	jobs.PostID = mux.Vars(r)["id"]
	return jobs, nil
}

// Helps to delete the existing job
func (e Endpoints) DeleteJob(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		data, ok := request.(map[string]string)
		if !ok {
			level.Error(logger.GokitLogger(fmt.Errorf("can't assert the request"))).Log()
			response = models.ResponseMessage{Data: request, Error: "Invalid Request", Code: http.StatusBadRequest, Message: ""}
			return response, nil
		}
		var post models.Post
		if err = e.DB.ReadJob(data["post_id"], &post); err != nil {
			level.Error(logger.GokitLogger(err)).Log()
			response = models.ResponseMessage{Data: "", Error: err.Error(), Code: http.StatusBadRequest, Message: "Invalid request"}
			return response, nil
		}

		if data["user_id"] != post.UsersID {
			response = models.ResponseMessage{Data: "", Error: fmt.Sprint("You are not authorized to delete this post ", data["name"]), Code: http.StatusBadRequest, Message: "Invalid request"}
			return response, nil
		}

		if err = e.DB.DeleteJobByID(post); err != nil {
			level.Error(logger.GokitLogger(err)).Log()
			response = models.ResponseMessage{Data: "", Error: err.Error(), Code: http.StatusBadRequest, Message: "Invalid request"}
			return response, nil
		}
		response = models.ResponseMessage{Data: "", Error: "", Code: http.StatusOK, Message: "Deleted Successfully"}
		return
	}
}

func (e Endpoints) DecodeDeleteJobRequest(_ context.Context, r *http.Request) (request interface{}, err error) {

	var claims models.Claims
	var service service.Service
	if err := service.Claims(r.Header.Get("Authorization"), &claims); err != nil {
		return err.Error(), nil
	}

	data := make(map[string]string)
	data["user_id"] = claims.UsersID
	data["name"] = claims.Name
	data["post_id"] = mux.Vars(r)["id"]
	return data, nil
}

// Helps to get the job by id
func (e Endpoints) GetJob(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		id, ok := request.(string)
		if !ok {
			level.Error(logger.GokitLogger(fmt.Errorf("can't assert the request"))).Log()
			response = models.ResponseMessage{Data: request, Error: "Invalid Request", Code: http.StatusBadRequest, Message: ""}
			return response, nil
		}
		var post models.Post
		if err := e.DB.ReadJob(id, &post); err != nil {
			level.Error(logger.GokitLogger(err)).Log()
			response = models.ResponseMessage{Data: "", Error: err.Error(), Code: http.StatusBadRequest, Message: "Invalid request"}
			return response, nil
		}

		response = models.ResponseMessage{Data: post, Error: "", Code: http.StatusOK, Message: "Successful request"}
		return
	}
}

func (e Endpoints) DecodeGetID(_ context.Context, r *http.Request) (request interface{}, err error) {
	postID := mux.Vars(r)["id"]
	return postID, nil
}
