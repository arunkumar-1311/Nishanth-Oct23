package handler

import (
	"blog_post/logger"
	"blog_post/models"
	"blog_post/repository"
	"blog_post/service"
	"blog_post/service/helper"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func Register(c *fiber.Ctx) error {
	c.Accepts("application/json")
	var user models.Users
	validate := validator.New()

	if err := c.BodyParser(&user); err != nil {
		logger.Logging().Error(err)
		service.SendResponse(c, fiber.StatusInternalServerError, err.Error(), "Try Again Later", fiber.MethodPost, "")
		return nil
	}
	user.RolesID = "US1"
	user.UserID = helper.UniqueID()
	if err := validate.Struct(user); err != nil {
		logger.Logging().Error(err)
		service.SendResponse(c, fiber.StatusBadRequest, err.Error(), "Give all required fields", fiber.MethodPost, "")
		return nil
	}

	if err := helper.GenerateHash(&user.Password); err != nil {
		logger.Logging().Error(err)
		service.SendResponse(c, fiber.StatusInternalServerError, err.Error(), "Try Again Later", fiber.MethodPost, "")
		return nil
	}

	if result := helper.EmailAndNameValidation(user); result != nil {
		logger.Logging().Error(result)
		service.SendResponse(c, fiber.StatusBadRequest, result.Error(), "Oops error occurs", fiber.MethodPost, "")
		return nil
	}

	if err := repository.Create(user); err != nil {
		logger.Logging().Error(err)
		service.SendResponse(c, fiber.StatusBadRequest, err.Error(), "Oops error occurs", fiber.MethodPost, "")
		return nil
	}

	service.SendResponse(c, fiber.StatusOK, "", "Please login", fiber.MethodPost, "User Created successfully")
	return nil
}
