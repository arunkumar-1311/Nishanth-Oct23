package handler

import (
	"blog_post/logger"
	"blog_post/service"
	"blog_post/service/helper"
	"net/http"
	"strings"

	"github.com/gofiber/fiber/v2"
)

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

		if err := helper.AdminAccess(tokenString); strings.Contains(c.Path(), "admin") && err != nil {
			logger.Logging().Error(err)
			service.SendResponse(c, http.StatusBadRequest, err.Error(), "Try Again", http.MethodGet, "")
			return nil
		}
		return c.Next()
	}
}
