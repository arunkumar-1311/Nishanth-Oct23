package handler

import (
	"blog_post/logger"
	"blog_post/models"
	"blog_post/repository"
	"blog_post/service"
	"blog_post/service/helper"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

func Register(c *fiber.Ctx) error {
	c.Accepts("application/json")
	var user models.Users
	validate := validator.New()

	if err := c.BodyParser(&user); err != nil {
		logger.Logging().Error(err)
		service.SendResponse(c, http.StatusInternalServerError, err.Error(), "Try Again Later", http.MethodPost, "")
		return nil
	}
	user.RolesID = "US1"
	user.UserID = helper.UniqueID()
	if err := validate.Struct(user); err != nil {
		logger.Logging().Error(err)
		service.SendResponse(c, http.StatusBadRequest, err.Error(), "Give all required fields", http.MethodPost, "")
		return nil
	}

	if err := helper.GenerateHash(&user.Password); err != nil {
		logger.Logging().Error(err)
		service.SendResponse(c, http.StatusInternalServerError, err.Error(), "Try Again Later", http.MethodPost, "")
		return nil
	}

	if result := helper.EmailAndNameValidation(user); result != nil {
		logger.Logging().Error(result)
		service.SendResponse(c, http.StatusBadRequest, result.Error(), "Oops error occurs", http.MethodPost, "")
		return nil
	}

	if err := repository.Create(user); err != nil {
		logger.Logging().Error(err)
		service.SendResponse(c, http.StatusBadRequest, err.Error(), "Oops error occurs", http.MethodPost, "")
		return nil
	}

	service.SendResponse(c, http.StatusOK, "", "Please login", http.MethodPost, "User Created successfully")
	return nil
}
