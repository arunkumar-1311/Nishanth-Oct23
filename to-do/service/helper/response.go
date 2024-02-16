package helper

import (
	"github.com/labstack/echo"
	"to-do/models"
)

// Helps to set header and send response
func SendResponse(c echo.Context, data interface{}, code int, err, message string) error {

	c.Response().Header().Add(echo.HeaderContentType, "application/json")
	c.Response().Header().Set(echo.HeaderAccessControlAllowOrigin, "*")
	response := models.ResponseMessage{
		Data:    data,
		Error:   err,
		Code:    code,
		Message: message,
	}
	return c.JSON(code, response)
}
