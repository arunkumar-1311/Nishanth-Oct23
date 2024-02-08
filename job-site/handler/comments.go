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

type Comments interface {
	PostComments(service.Service) endpoint.Endpoint
	DecodePostCommentsRequest(context.Context, *http.Request) (interface{}, error)

	ReadCommentByID(service.Service) endpoint.Endpoint

	ReadCommentByPost(service.Service) endpoint.Endpoint

	UpdateCommentByID(service.Service) endpoint.Endpoint
	DecodeUpdateCommentByIDRequest(context.Context, *http.Request) (interface{}, error)

	DeleteComment(service.Service) endpoint.Endpoint
	DecodeDeleteCommentRequest(context.Context, *http.Request) (interface{}, error)
}

// Helps to add comment to specific job post
func (e Endpoints) PostComments(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		comment, ok := request.(models.Comment)
		if !ok {
			level.Error(logger.GokitLogger(fmt.Errorf("can't assert the request"))).Log()
			response = models.ResponseMessage{Data: request, Error: "Invalid Request", Code: http.StatusBadRequest, Message: ""}
			return response, nil
		}

		comment.CommentID = helper.UniqueID()
		validate := validator.New()
		if err := validate.Struct(comment); err != nil {
			level.Error(logger.GokitLogger(err)).Log()
			response = models.ResponseMessage{Data: "", Error: err.Error(), Code: http.StatusBadRequest, Message: "Invalid request"}
			return response, nil
		}

		if err := e.DB.CreateComment(comment); err != nil {
			level.Error(logger.GokitLogger(err)).Log()
			response = models.ResponseMessage{Data: "", Error: err.Error(), Code: http.StatusBadRequest, Message: "Invalid request"}
			return response, nil
		}

		response = models.ResponseMessage{Data: comment, Error: "", Code: http.StatusOK, Message: "Comment Added Successfully"}
		return
	}
}

func (e Endpoints) DecodePostCommentsRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	var claims models.Claims
	var service service.Service
	if err := service.Claims(r.Header.Get("Authorization"), &claims); err != nil {
		return err.Error(), nil
	}

	var comment models.Comment
	if err := json.NewDecoder(r.Body).Decode(&comment); err != nil {
		return err.Error(), nil
	}
	comment.UsersID = claims.UsersID
	comment.PostID = mux.Vars(r)["id"]
	return comment, nil
}

// Read comment by comment ID
func (e Endpoints) ReadCommentByID(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		id, ok := request.(string)
		if !ok {
			level.Error(logger.GokitLogger(fmt.Errorf("can't assert the request"))).Log()
			response = models.ResponseMessage{Data: request, Error: "Invalid Request", Code: http.StatusBadRequest, Message: ""}
			return response, nil
		}

		var comment models.Comment
		if err := e.DB.ReadComment(id, &comment); err != nil {
			level.Error(logger.GokitLogger(err)).Log()
			response = models.ResponseMessage{Data: "", Error: err.Error(), Code: http.StatusBadRequest, Message: "Invalid request"}
			return response, nil
		}

		response = models.ResponseMessage{Data: comment, Error: "", Code: http.StatusOK, Message: "Comment by its ID"}
		return
	}
}

// Read all comments by its post
func (e Endpoints) ReadCommentByPost(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		id, ok := request.(string)
		if !ok {
			level.Error(logger.GokitLogger(fmt.Errorf("can't assert the request"))).Log()
			response = models.ResponseMessage{Data: request, Error: "Invalid Request", Code: http.StatusBadRequest, Message: ""}
			return response, nil
		}

		var comment []models.Comment
		if err := e.DB.ReadCommentsByPost(id, &comment); err != nil {
			level.Error(logger.GokitLogger(err)).Log()
			response = models.ResponseMessage{Data: "", Error: err.Error(), Code: http.StatusBadRequest, Message: "Invalid request"}
			return response, nil
		}

		response = models.ResponseMessage{Data: comment, Error: "", Code: http.StatusOK, Message: "Comments by Post"}
		return
	}
}

// Helps to update the user comment
func (e Endpoints) UpdateCommentByID(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		comment, ok := request.(models.Comment)
		if !ok {
			level.Error(logger.GokitLogger(fmt.Errorf("can't assert the request"))).Log()
			response = models.ResponseMessage{Data: request, Error: "Invalid Request", Code: http.StatusBadRequest, Message: ""}
			return response, nil
		}

		var existingComment models.Comment
		if err = e.DB.ReadComment(comment.CommentID, &existingComment); err != nil {
			level.Error(logger.GokitLogger(err)).Log()
			response = models.ResponseMessage{Data: "", Error: err.Error(), Code: http.StatusBadRequest, Message: "Invalid request"}
			return response, nil
		}

		if existingComment.UsersID != comment.UsersID {
			response = models.ResponseMessage{Data: "", Error: "Unauthorized to update this comment", Code: http.StatusBadRequest, Message: "Invalid request"}
			return response, nil
		}
		comment.PostID = existingComment.PostID
		if err = e.DB.UpdateComment(comment); err != nil {
			level.Error(logger.GokitLogger(err)).Log()
			response = models.ResponseMessage{Data: "", Error: err.Error(), Code: http.StatusBadRequest, Message: "Invalid request"}
			return response, nil
		}

		response = models.ResponseMessage{Data: comment, Error: "", Code: http.StatusOK, Message: "Comment after updation"}
		return
	}
}

func (e Endpoints) DecodeUpdateCommentByIDRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	var claims models.Claims
	var service service.Service
	if err := service.Claims(r.Header.Get("Authorization"), &claims); err != nil {
		return err.Error(), nil
	}

	var comment models.Comment
	if err := json.NewDecoder(r.Body).Decode(&comment); err != nil {
		return err.Error(), nil
	}

	comment.CommentID = mux.Vars(r)["id"]
	comment.UsersID = claims.UsersID

	return comment, nil
}

// Helps to delete the comment
func (e Endpoints) DeleteComment(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		data, ok := request.(map[string]string)
		if !ok {
			level.Error(logger.GokitLogger(fmt.Errorf("can't assert the request"))).Log()
			response = models.ResponseMessage{Data: request, Error: "Invalid Request", Code: http.StatusBadRequest, Message: ""}
			return response, nil
		}

		var comment models.Comment
		if err := e.DB.ReadComment(data["comment_id"], &comment); err != nil {
			level.Error(logger.GokitLogger(err)).Log()
			response = models.ResponseMessage{Data: "", Error: err.Error(), Code: http.StatusBadRequest, Message: "Invalid request"}
			return response, nil
		}

		if comment.UsersID != data["user_id"] && data["role_id"] != "AD1" {
			level.Error(logger.GokitLogger(err)).Log()
			response = models.ResponseMessage{Data: "", Error: "Unauthorized to delete this comment", Code: http.StatusBadRequest, Message: "Invalid request"}
			return response, nil
		}

		if err = e.DB.DeleteCommentByID(comment); err != nil {
			level.Error(logger.GokitLogger(err)).Log()
			response = models.ResponseMessage{Data: "", Error: err.Error(), Code: http.StatusBadRequest, Message: "Invalid request"}
			return response, nil
		}

		response = models.ResponseMessage{Data: "", Error: "", Code: http.StatusOK, Message: "Comment Deleted Successfully"}
		return
	}
}

func (e Endpoints) DecodeDeleteCommentRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	var claims models.Claims
	var service service.Service
	if err := service.Claims(r.Header.Get("Authorization"), &claims); err != nil {
		return err.Error(), nil
	}
	response := make(map[string]string)
	response["user_id"] = claims.UsersID
	response["comment_id"] = mux.Vars(r)["id"]
	response["role_id"] = claims.RoleID

	return response, nil
}
