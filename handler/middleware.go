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

// helps to authentuicate the user
func Authentication() fiber.Handler {
	return func(c *fiber.Ctx) error {

		var credentials models.Login
		validate := validator.New()

		if err := c.BodyParser(&credentials); err != nil {
			logger.Logging().Error(err)
			service.SendResponse(c, http.StatusBadRequest, err.Error(), "Enter valid body", http.MethodPost, "")
			return nil
		}

		if err := validate.Struct(credentials); err != nil {
			logger.Logging().Error(err)
			service.SendResponse(c, http.StatusBadRequest, err.Error(), "Give all required fields", http.MethodPost, "")
			return nil
		}

		var user models.Users
		if err := repository.User(credentials.Name, &user); err != nil {
			logger.Logging().Error(err)
			service.SendResponse(c, http.StatusInternalServerError, err.Error(), "Try Again Later", http.MethodPost, "")
			return nil
		}

		if result := helper.CompareHashPassword(credentials.Password, user.Password); !result {
			logger.Logging().Error("Invalid credentials")
			service.SendResponse(c, http.StatusBadRequest, "Invalid credentials ", "Check your password", http.MethodPost, "")
			return nil
		}

		token, err := helper.CreateToken(user.Name, user.Email, user.RolesID, user.UserID)
		if err != nil {
			logger.Logging().Error(err)
			service.SendResponse(c, http.StatusInternalServerError, err.Error(), "Try Again Later", http.MethodPost, "")
			return nil
		}

		c.Set("Authorization", token)
		service.SendResponse(c, http.StatusOK, "", "true", http.MethodPost, "Logged in successfully")
		return c.Next()

	}
}


// helps to authorize the admin
func Authorization() fiber.Handler {
	return func(c *fiber.Ctx) error {
		
		tokenString := c.Get("Authorization")
		if tokenString == "" {
			data := map[string]interface{}{
				"Regiter": "localhost:8000/register",
				"Login":   "localhost:8000/login",
			}
			service.SendResponse(c, http.StatusBadRequest, "Please login or register", "Login or register to make this action", http.MethodGet, data)
			return nil
		}

		if _, err := helper.VerifyToken(tokenString[7:]); err != nil {
			logger.Logging().Error(err)
			service.SendResponse(c, http.StatusBadRequest, err.Error(), "invalid token", http.MethodGet, "")
			return nil
		}
		return c.Next()

	}
}
