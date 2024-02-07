package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"online-purchase/logger"
	"online-purchase/models"
	"online-purchase/service"
	"online-purchase/service/helper"

	"github.com/astaxie/beego/context"

	"github.com/go-playground/validator/v10"
)

// Helps to create the profile to the application
func (h *Handlers) Register(ctx *context.Context) {
	ctx.Output.Header("content-Type", "application/json")
	var user models.Users
	if err := json.NewDecoder(ctx.Request.Body).Decode(&user); err != nil {
		logger.ZapLog().Error(err.Error())
		service.SendResponse(ctx, http.StatusInternalServerError, err.Error(), "Please try again later", "")
		return
	}

	user.UserID = helper.UniqueID()

	validate := validator.New()
	if err := validate.Struct(user); err != nil {
		logger.ZapLog().Error(err.Error())
		service.SendResponse(ctx, http.StatusBadRequest, err.Error(), "Invalid request", "")
		return
	}

	if err := helper.GenerateHash(&user.Password); err != nil {
		logger.ZapLog().Error(err.Error())
		service.SendResponse(ctx, http.StatusInternalServerError, err.Error(), "Please try again later", "")
		return
	}

	if err := helper.EmailAndNameValidation(user, h.Database); err != nil {
		logger.ZapLog().Error(err.Error())
		service.SendResponse(ctx, http.StatusBadRequest, err.Error(), "Invalid request", "")
		return
	}

	if err := h.CreateUser(user); err != nil {
		logger.ZapLog().Error(err.Error())
		service.SendResponse(ctx, http.StatusBadRequest, err.Error(), "Invalid request", "")
		return
	}

}

// Helps to login the existing account
func (h *Handlers) Login(ctx *context.Context) {
	ctx.Output.Header("content-Type", "application/json")
	var credentials models.Login
	if err := json.NewDecoder(ctx.Request.Body).Decode(&credentials); err != nil {
		logger.ZapLog().Error(err.Error())
		service.SendResponse(ctx, http.StatusInternalServerError, err.Error(), "Please try again later", "")
		return
	}

	validate := validator.New()
	if err := validate.Struct(credentials); err != nil {
		logger.ZapLog().Error(err.Error())
		service.SendResponse(ctx, http.StatusBadRequest, err.Error(), "Invalid request", "")
		return
	}

	var Profile models.Users
	if err := h.User(credentials.Name, &Profile); err != nil {
		logger.ZapLog().Error(err.Error())
		service.SendResponse(ctx, http.StatusBadRequest, err.Error(), "Invalid request", "")
		return
	}

	if result := helper.CompareHashPassword(credentials.Password, Profile.Password); !result {
		logger.ZapLog().Error("Invalid password")
		service.SendResponse(ctx, http.StatusBadRequest, "Invalid password", "Invalid request", "")
		return
	}

	token, err := helper.CreateToken(Profile.UserName, Profile.Email, Profile.RolesID, Profile.UserID)
	if err != nil {
		logger.ZapLog().Error(err.Error())
		service.SendResponse(ctx, http.StatusBadRequest, err.Error(), "Invalid request", "")
		return

	}

	ctx.ResponseWriter.Header().Set("authorization", token)
	service.SendResponse(ctx, http.StatusOK, "", "Logged in Successfully", "")

}

// Helps to get the user profile
func (h *Handlers) GetProfile(ctx *context.Context) {
	var claims models.Claims

	if err := helper.Claims(ctx.Request.Header["Authorization"][0], &claims); err != nil {
		logger.ZapLog().Error(err.Error())
		service.SendResponse(ctx, http.StatusBadRequest, err.Error(), "Invalid request", "")
		return
	}
	var profile models.Users
	if err := h.ReadProfile(claims.UsersID, &profile); err != nil {
		logger.ZapLog().Error(err.Error())
		service.SendResponse(ctx, http.StatusBadRequest, err.Error(), "Invalid request", "")
		return
	}
// profile.Password = ""
	service.SendResponse(ctx, http.StatusOK, "", fmt.Sprintf("Profile of %v", claims.Name), profile)
}

// Helps to update the user profile
func (h *Handlers) UpdateProfile(ctx *context.Context) {
	var claims models.Claims
	if err := helper.Claims(ctx.Request.Header["Authorization"][0], &claims); err != nil {
		logger.ZapLog().Error(err.Error())
		service.SendResponse(ctx, http.StatusBadRequest, err.Error(), "Invalid request", "")
		return
	}

	var userProfile models.Users
	if err := json.NewDecoder(ctx.Request.Body).Decode(&userProfile); err != nil {
		logger.ZapLog().Error(err.Error())
		service.SendResponse(ctx, http.StatusInternalServerError, err.Error(), "Please try again later", "")
		return
	}
	userProfile.UserID = claims.UsersID

	if err := helper.EmailAndNameValidation(userProfile, h.Database); err != nil {
		logger.ZapLog().Error(err.Error())
		service.SendResponse(ctx, http.StatusBadRequest, err.Error(), "Invalid request", "")
		return
	}

	if userProfile.Password != "" {
		helper.GenerateHash(&userProfile.Password)
	}

	if err := h.UpdateUserProfile(userProfile); err != nil {
		logger.ZapLog().Error(err.Error())
		service.SendResponse(ctx, http.StatusBadRequest, err.Error(), "Invalid request", "")
		return
	}

	if err := h.ReadProfile(claims.UsersID, &userProfile); err != nil {
		logger.ZapLog().Error(err.Error())
		service.SendResponse(ctx, http.StatusBadRequest, err.Error(), "Invalid request", "")
		return
	}
	// userProfile.Password = ""
	service.SendResponse(ctx, http.StatusOK, "", fmt.Sprintf("Profile of %v", claims.Name), userProfile)
}
