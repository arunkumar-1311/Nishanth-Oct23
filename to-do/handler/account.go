package handler

import (
	"net/http"
	"to-do/logger"
	"to-do/models"
	"to-do/service/helper"

	"github.com/go-playground/validator"
	"github.com/labstack/echo"
)

type Account interface {
	SignIn(c echo.Context) error
	Login(c echo.Context) error
}

// Helps to create new account
func (e *EndPoint) SignIn(c echo.Context) error {

	var user models.Users

	if err := c.Bind(&user); err != nil {
		logger.ZeroLogger().Msg(err.Error())
		return helper.SendResponse(c, "", echo.ErrBadRequest.Code, err.Error(), "Invalid request")
	}

	user.UserID = e.Service.UniqueID()
	if err := e.Service.GenerateHash(&user.Password); err != nil {
		logger.ZeroLogger().Msg(err.Error())
		return helper.SendResponse(c, "", echo.ErrInternalServerError.Code, err.Error(), "Try again later")
	}

	if err := e.Service.EmailAndNameValidation(user, e.DB); err != nil {
		logger.ZeroLogger().Msg(err.Error())
		return helper.SendResponse(c, "", echo.ErrBadRequest.Code, err.Error(), "Invalid request")
	}

	validate := validator.New()
	if err := validate.Struct(user); err != nil {
		logger.ZeroLogger().Msg(err.Error())
		return helper.SendResponse(c, "", echo.ErrBadRequest.Code, err.Error(), "Invalid request")
	}

	if err := e.DB.NewUser(user); err != nil {
		logger.ZeroLogger().Msg(err.Error())
		return helper.SendResponse(c, "", echo.ErrBadRequest.Code, err.Error(), "Invalid request")
	}

	return helper.SendResponse(c, "Profile Created Successfully", http.StatusOK, "", "Success Message")
}

// Helps to login to the existing account
func (e *EndPoint) Login(c echo.Context) error {

	var user models.Login

	if err := c.Bind(&user); err != nil {
		logger.ZeroLogger().Msg(err.Error())
		return helper.SendResponse(c, "", echo.ErrBadRequest.Code, err.Error(), "Invalid request")
	}

	validate := validator.New()
	if err := validate.Struct(user); err != nil {
		logger.ZeroLogger().Msg(err.Error())
		return helper.SendResponse(c, "", echo.ErrBadRequest.Code, err.Error(), "Invalid request")
	}

	var userProfile models.Users
	if err := e.DB.ReadUser(user, &userProfile); err != nil {
		logger.ZeroLogger().Msg(err.Error())
		return helper.SendResponse(c, "", echo.ErrBadRequest.Code, err.Error(), "Invalid request")
	}

	if !e.Service.CompareHashPassword(user.Password, userProfile.Password) {
		return helper.SendResponse(c, "", echo.ErrBadRequest.Code, "Wrong password", "Invalid request")
	}
	
	uuid := e.Service.UniqueID()
	token, err := e.Service.CreateToken(uuid)
	if err != nil {
		logger.ZeroLogger().Msg(err.Error())
		return helper.SendResponse(c, "", echo.ErrInternalServerError.Code, err.Error(), "Try Again Later")
	}

	if err := e.DB.SetCache(uuid, userProfile.UserID, userProfile.UserName, userProfile.Email); err != nil {
		logger.ZeroLogger().Msg(err.Error())
		return helper.SendResponse(c, "", echo.ErrInternalServerError.Code, err.Error(), "Try Again Later")
	}

	c.Response().Header().Set(echo.HeaderAuthorization, token)
	return helper.SendResponse(c, "Logged in Successfully", http.StatusOK, "", "Success Message")
}
