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

// helps to authentuicate the user
func Authentication(c *fiber.Ctx) error {

	var credentials models.Login
	validate := validator.New()

	if err := c.BodyParser(&credentials); err != nil {
		logger.Logging().Error(err)
		service.SendResponse(c, fiber.StatusBadRequest, err.Error(), "Enter valid json", fiber.MethodPost, "")
		return nil
	}

	if err := validate.Struct(credentials); err != nil {
		logger.Logging().Error(err)
		service.SendResponse(c, fiber.StatusBadRequest, err.Error(), "Give all required fields", fiber.MethodPost, "")
		return nil
	}

	var user models.Users
	if err := repository.User(credentials.Name, credentials.Email, &user); err != nil {
		logger.Logging().Error(err)
		service.SendResponse(c, fiber.StatusInternalServerError, err.Error(), "Try Again Later", fiber.MethodPost, "")
		return nil
	}

	if result := helper.CompareHashPassword(credentials.Password, user.Password); !result {
		logger.Logging().Error("Invalid credentials")
		service.SendResponse(c, fiber.StatusBadRequest, "Invalid credentials ", "Check your password", fiber.MethodPost, "")
		return nil
	}

	token, err := helper.CreateToken(user.Name, user.Email, user.RolesID, user.UserID)
	if err != nil {
		logger.Logging().Error(err)
		service.SendResponse(c, fiber.StatusInternalServerError, err.Error(), "Try Again Later", fiber.MethodPost, "")
		return nil
	}

	c.Set("Authorization", token)
	service.SendResponse(c, fiber.StatusOK, "", "true", fiber.MethodPost, "Logged in successfully")
	return nil
}
