package middleware

import (
	"to-do/models"
	"to-do/service/helper"

	"github.com/labstack/echo"
)

type Authentication interface {
	Authentication(next echo.HandlerFunc) echo.HandlerFunc
}

// Helps to autheticate the user
func (m Middleware) Authentication(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		token := ctx.Request().Header.Get(echo.HeaderAuthorization)
		if token == "" {
			return helper.SendResponse(ctx, "", echo.ErrBadRequest.Code, "token not found", "!!!Please do login or sign up move forward!!!")
		}

		var claims models.Claims
		if err := m.Service.Claims(token[7:], &claims); err != nil {
			return helper.SendResponse(ctx, "", echo.ErrBadRequest.Code, err.Error(), "Login again")
		}
		user, err := m.DB.GetCache(claims.UUID)

		if err != nil {
			return helper.SendResponse(ctx, "", echo.ErrBadRequest.Code, err.Error(), "Invalid request")
		}
		ctx.Set("userid", user["userid"])
		ctx.Set("name", user["name"])
		ctx.Set("email", user["email"])
		ctx.Set("uuid", claims.UUID)

		return next(ctx)
	}
}
