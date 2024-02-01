package middleware

import (
	"net/http"
	"online-purchase/logger"
	"online-purchase/service"
	"online-purchase/service/helper"
	"strings"

	"github.com/astaxie/beego/context"
)

func Authorization(ctx *context.Context) {
	ctx.Output.Header("content-Type", "application/json")

	token := ctx.Request.Header["Authorization"]
	if token == nil {
		data := map[string]interface{}{
			"Regiter": "http://localhost:8000/user/new",
			"Login":   "http://localhost:8000/user",
		}
		service.SendResponse(ctx, http.StatusBadRequest, "No token found", "Invalid authorization", data)
		return
	}

	if _, err := helper.VerifyToken(token[0][7:]); err != nil {
		logger.ZapLog().Error(err.Error())
		service.SendResponse(ctx, http.StatusBadRequest, err.Error(), "Please try again later", "")
		return
	}

	if err := helper.AdminAccess(token[0]); strings.Contains(ctx.Request.RequestURI, "admin") && err != nil {
		logger.ZapLog().Error(err.Error())
		service.SendResponse(ctx, http.StatusBadRequest, err.Error(), "Please try again later", "")
		return
	}
}
