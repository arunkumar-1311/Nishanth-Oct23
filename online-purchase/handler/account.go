package handler

import (
	"encoding/json"
	"github.com/astaxie/beego/context"
	"github.com/go-playground/validator/v10"
	"net/http"
	"online-purchase/adaptor"
	"online-purchase/logger"
	"online-purchase/models"
	"online-purchase/service"
	"online-purchase/service/helper"
)

type Handlers struct {
	adaptor.Database
}

// Helps to create the profile to the application
func (h *Handlers) Register(ctx *context.Context) {
	ctx.Output.Header("content-Type", "application/json")
	var user models.Users
	if err := json.NewDecoder(ctx.Request.Body).Decode(&user); err != nil {
		logger.ZapLog().Error(err.Error())
		service.SendResponse(ctx, http.StatusInternalServerError, err.Error(), "Please try again later", ctx.Request.Method, "")
		return
	}

	user.UserID = helper.UniqueID()

	validate := validator.New()
	if err := validate.Struct(user); err != nil {
		logger.ZapLog().Error(err.Error())
		service.SendResponse(ctx, http.StatusBadRequest, err.Error(), "Invalid request", ctx.Request.Method, "")
		return
	}

	if err := helper.GenerateHash(&user.Password); err != nil {
		logger.ZapLog().Error(err.Error())
		service.SendResponse(ctx, http.StatusInternalServerError, err.Error(), "Please try again later", ctx.Request.Method, "")
		return
	}
	
	if err := helper.EmailAndNameValidation(user, h); err != nil {
		logger.ZapLog().Error(err.Error())
		service.SendResponse(ctx, http.StatusBadRequest, err.Error(), "Invalid request", ctx.Request.Method, "")
		return
	}

	if err := h.CreateUser(user); err != nil {
		logger.ZapLog().Error(err.Error())
		service.SendResponse(ctx, http.StatusBadRequest, err.Error(), "Invalid request", ctx.Request.Method, "")
		return
	}

	service.SendResponse(ctx, http.StatusOK, "", "Profile Created Successfully", ctx.Request.Method, "")

}

// Helps to login the existing account
func (h *Handlers) Login(ctx *context.Context) {
	ctx.Output.Header("content-Type", "application/json")
	var credentials models.Login
	if err := json.NewDecoder(ctx.Request.Body).Decode(&credentials); err != nil {
		logger.ZapLog().Error(err.Error())
		service.SendResponse(ctx, http.StatusInternalServerError, err.Error(), "Please try again later", ctx.Request.Method, "")
		return
	}

	validate := validator.New()
	if err := validate.Struct(credentials); err != nil {
		logger.ZapLog().Error(err.Error())
		service.SendResponse(ctx, http.StatusBadRequest, err.Error(), "Invalid request", ctx.Request.Method, "")
		return
	}

	var Profile models.Users
	if err := h.User(credentials.Name, &Profile); err != nil {
		logger.ZapLog().Error(err.Error())
		service.SendResponse(ctx, http.StatusBadRequest, err.Error(), "Invalid request", ctx.Request.Method, "")
		return
	}

	if result := helper.CompareHashPassword(credentials.Password, Profile.Password); !result {
		logger.ZapLog().Error("Invalid password")
		service.SendResponse(ctx, http.StatusBadRequest, "Invalid password", "Invalid request", ctx.Request.Method, "")
		return
	}

	token, err := helper.CreateToken(Profile.UserName, Profile.Email, Profile.RolesID, Profile.UserID)
	if err != nil {
		logger.ZapLog().Error(err.Error())
		service.SendResponse(ctx, http.StatusBadRequest, err.Error(), "Invalid request", ctx.Request.Method, "")
		return

	}

	ctx.ResponseWriter.Header().Set("authorization", token)
	service.SendResponse(ctx, http.StatusOK, "", "Logged in Successfully", ctx.Request.Method, "")

}
