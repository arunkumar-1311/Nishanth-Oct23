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
)

type Account interface {
	Register(svc service.Service) endpoint.Endpoint
	DecodeRegisterRequest(context.Context, *http.Request) (interface{}, error)

	GetProfile(svc service.Service) endpoint.Endpoint
	DecodeGetProfileRequest(context.Context, *http.Request) (interface{}, error)

	UpdateProfile(svc service.Service) endpoint.Endpoint
	DecodeUpdateProfileRequest(context.Context, *http.Request) (interface{}, error)

	Login(svc service.Service) endpoint.Endpoint
	DecodeLoginRequest(context.Context, *http.Request) (interface{}, error)

	EncodeResponse(context.Context, http.ResponseWriter, interface{}) error
	DecodeRequest(context.Context, *http.Request) (interface{}, error)
}

// Helps to create a user profile for job-site application
func (e Endpoints) Register(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {

		user, ok := request.(models.Users)
		if !ok {
			level.Error(logger.GokitLogger(fmt.Errorf("can't assert the request"))).Log()
			response = models.ResponseMessage{Data: request, Error: "Invalid Request", Code: http.StatusBadRequest, Message: ""}
			return response, nil
		}
		user.UserID = helper.UniqueID()

		if err = svc.GenerateHash(&user.Password); err != nil {
			level.Error(logger.GokitLogger(err)).Log()
			response = models.ResponseMessage{Data: "", Error: err.Error(), Code: http.StatusBadRequest, Message: "Invalid request"}
			return response, nil
		}

		validate := validator.New()
		if err = validate.Struct(user); err != nil {
			level.Error(logger.GokitLogger(err)).Log()
			response = models.ResponseMessage{Data: "", Error: err.Error(), Code: http.StatusBadRequest, Message: "Invalid request"}
			return response, nil
		}

		if err = svc.EmailAndNameValidation(user, e.DB); err != nil {
			level.Error(logger.GokitLogger(err)).Log()
			response = models.ResponseMessage{Data: "", Error: err.Error(), Code: http.StatusBadRequest, Message: "Invalid request"}
			return response, nil
		}

		if err = e.DB.CreateUser(user); err != nil {
			level.Error(logger.GokitLogger(err)).Log()
			response = models.ResponseMessage{Data: "", Error: err.Error(), Code: http.StatusBadRequest, Message: "Invalid request"}
			return response, nil
		}

		response = models.ResponseMessage{Data: "", Error: "", Code: http.StatusOK, Message: "Profile created successfully"}
		return response, nil
	}
}

func (e Endpoints) DecodeRegisterRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	var user models.Users
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		return err.Error(), nil
	}
	return user, nil
}

// Helps get the profile of the user
func (e Endpoints) GetProfile(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		claims, ok := request.(models.Claims)
		if !ok {
			level.Error(logger.GokitLogger(fmt.Errorf("can't assert the request"))).Log()
			response = models.ResponseMessage{Data: request, Error: "Invalid Request", Code: http.StatusBadRequest, Message: ""}
			return response, nil
		}

		var user models.Users
		if err := e.DB.ReadProfile(claims.UsersID, &user); err != nil {
			level.Error(logger.GokitLogger(err)).Log()
			response = models.ResponseMessage{Data: "", Error: err.Error(), Code: http.StatusBadRequest, Message: "Invalid Request"}
			return response, nil
		}

		response = models.ResponseMessage{Data: user, Error: "", Code: http.StatusBadRequest, Message: "User Profile"}
		return
	}
}

func (e Endpoints) DecodeGetProfileRequest(ctx context.Context, r *http.Request) (request interface{}, err error) {
	var claims models.Claims
	var service service.Service
	if err := service.Claims(r.Header.Get("Authorization"), &claims); err != nil {
		return err.Error(), nil
	}
	return claims, nil
}

// Helps to update the user profile
func (e Endpoints) UpdateProfile(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		user, ok := request.(models.Users)
		if !ok {
			level.Error(logger.GokitLogger(fmt.Errorf("can't assert the request"))).Log()
			response = models.ResponseMessage{Data: request, Error: "Invalid Request", Code: http.StatusBadRequest, Message: ""}
			return response, nil
		}

		if user.Password != "" {
			if err := svc.GenerateHash(&user.Password); err != nil {
				level.Error(logger.GokitLogger(err)).Log()
				response = models.ResponseMessage{Data: "", Error: err.Error(), Code: http.StatusBadRequest, Message: "Invalid Request"}
				return response, nil
			}
		}

		if user.UserName != "" || user.Email != "" {
			if err = svc.EmailAndNameValidation(user, e.DB); err != nil {
				level.Error(logger.GokitLogger(err)).Log()
				response = models.ResponseMessage{Data: "", Error: err.Error(), Code: http.StatusBadRequest, Message: "Invalid request"}
				return response, nil
			}
		}
		if err = e.DB.UpdateUser(&user); err != nil {
			level.Error(logger.GokitLogger(err)).Log()
			response = models.ResponseMessage{Data: "", Error: err.Error(), Code: http.StatusBadRequest, Message: "Invalid Request"}
			return response, nil
		}

		if err = e.DB.ReadProfile(user.UserID, &user); err != nil {
			level.Error(logger.GokitLogger(err)).Log()
			response = models.ResponseMessage{Data: "", Error: err.Error(), Code: http.StatusBadRequest, Message: "Invalid Request"}
			return response, nil
		}
		user.Password = ""
		response = models.ResponseMessage{Data: user, Error: "", Code: http.StatusBadRequest, Message: "Profile after updation"}
		return
	}
}

func (e Endpoints) DecodeUpdateProfileRequest(ctx context.Context, r *http.Request) (request interface{}, err error) {

	var claims models.Claims
	var service service.Service
	if err := service.Claims(r.Header.Get("Authorization"), &claims); err != nil {
		return err.Error(), err
	}

	var user models.Users
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		return err.Error(), nil
	}
	user.UserID = claims.UsersID

	return user, nil
}

// Helps to login to the existing account
func (e Endpoints) Login(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {

		credentials, ok := request.(models.Login)
		if !ok {
			level.Error(logger.GokitLogger(fmt.Errorf("can't assert the request"))).Log()
			response = models.ResponseMessage{Data: request, Error: "Invalid Request", Code: http.StatusBadRequest, Message: "Please try again later"}
			return response, nil
		}

		validate := validator.New()
		if err = validate.Struct(credentials); err != nil {
			level.Error(logger.GokitLogger(err)).Log()
			response = models.ResponseMessage{Data: "", Error: err.Error(), Code: http.StatusBadRequest, Message: "Invalid request"}
			return response, nil
		}

		var profile models.Users
		if err = e.DB.ReadUser(credentials.Name, &profile); err != nil {
			level.Error(logger.GokitLogger(err)).Log()
			response = models.ResponseMessage{Data: "", Error: err.Error(), Code: http.StatusBadRequest, Message: "Invalid request"}
			return response, nil
		}

		if ok = svc.CompareHashPassword(credentials.Password, profile.Password); !ok {
			response = models.ResponseMessage{Data: "", Error: "Invalid credentials check your password", Code: http.StatusBadRequest, Message: "Invalid request"}
			return response, nil
		}

		token, err := svc.CreateToken(profile.UserName, profile.Email, profile.Roles.Role, profile.RolesID, profile.UserID)
		if err != nil {
			level.Error(logger.GokitLogger(err)).Log()
			response = models.ResponseMessage{Data: "", Error: err.Error(), Code: http.StatusBadRequest, Message: "Invalid request"}
			return response, nil
		}

		response = models.ResponseMessage{Data: "", Error: "", Code: http.StatusOK, Message: "Logged in successfully", Token: token}
		return response, nil
	}
}

func (e Endpoints) DecodeLoginRequest(ctx context.Context, r *http.Request) (request interface{}, err error) {

	var credentials models.Login
	if err = json.NewDecoder(r.Body).Decode(&credentials); err != nil {
		return err.Error(), nil
	}

	return credentials, nil

}

// helps to encode all response and send a json response
func (e Endpoints) EncodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {

	if err := helper.SendResponse(ctx, w, response); err != nil {
		level.Error(logger.GokitLogger(err)).Log()
		return err
	}
	return nil
}

// Helps to decode the plain request
func (e Endpoints) DecodeRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	return
}
