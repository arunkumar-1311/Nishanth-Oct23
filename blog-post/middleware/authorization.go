package middleware



import (
	"blog_post/logger"
	"blog_post/service"
	"blog_post/service/helper"
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
			service.SendResponse(c, fiber.StatusBadRequest, "Please login or register", "Login or register to make this action", fiber.MethodGet, data)
			return nil
		}

		if _, err := helper.VerifyToken(tokenString[7:]); err != nil {
			logger.Logging().Error(err)
			service.SendResponse(c, fiber.StatusBadRequest, err.Error(), "invalid token", fiber.MethodGet, "")
			return nil
		}

		if err := helper.AdminAccess(tokenString); strings.Contains(c.Path(), "admin") && err != nil {
			logger.Logging().Error(err)
			service.SendResponse(c, fiber.StatusBadRequest, err.Error(), "Try Again", fiber.MethodGet, "")
			return nil
		}
		return c.Next()
	}
}

