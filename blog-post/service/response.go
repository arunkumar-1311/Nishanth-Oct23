package service

import (
	"blog_post/logger"
	"blog_post/models"

	"github.com/gofiber/fiber/v2"
)

// Helps to set header and send response
func SendResponse(c *fiber.Ctx, status int, err string, message, method string, data ...interface{}) {

	resp := models.Message{Data: data, Error: err, Code: status, Message: message}
	c.SendStatus(status)
	resultErr := c.JSON(resp)
	if resultErr != nil {
		logger.Logging().Error(err)
		return
	}
	c.Context().Done()

}
