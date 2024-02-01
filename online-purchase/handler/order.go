package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"online-purchase/logger"
	"online-purchase/models"
	"online-purchase/service"
	"online-purchase/service/helper"

	"github.com/astaxie/beego/context"
	"github.com/go-playground/validator/v10"
)

// Helps to create a new order
func (h *Handlers) CreateOrder(ctx *context.Context) {
	ctx.Output.Header("content-Type", "application/json")
	var order models.Orders

	if err := json.NewDecoder(ctx.Request.Body).Decode(&order); err != nil {
		logger.ZapLog().Error(err.Error())
		service.SendResponse(ctx, http.StatusInternalServerError, err.Error(), "Please try again later", "")
		return
	}

	order.OrderID = helper.UniqueID()
	order.Active = true
	order.OrderStatusID = "S-NULL"
	var brand models.Brand
	var ram models.Ram
	var claims models.Claims

	if err := h.ReadBrandByID(order.BrandID, &brand); err != nil {
		logger.ZapLog().Error(err.Error())
		service.SendResponse(ctx, http.StatusBadRequest, err.Error(), "Invalid request", "")
		return
	}

	if err := h.ReadRAMByID(order.RamID, &ram); err != nil {
		logger.ZapLog().Error(err.Error())
		service.SendResponse(ctx, http.StatusBadRequest, err.Error(), "Invalid request", "")
		return
	}

	if err := helper.Claims(ctx.Request.Header["Authorization"][0], &claims); err != nil {
		logger.ZapLog().Error(err.Error())
		service.SendResponse(ctx, http.StatusBadRequest, err.Error(), "Invalid request", "")
		return
	}

	order.Total = brand.Price + ram.Price
	order.UserID = claims.UsersID
	if order.DVD {
		order.Total = order.Total + 3000
	}

	validate := validator.New()
	if err := validate.Struct(&order); err != nil {
		logger.ZapLog().Error(err.Error())
		service.SendResponse(ctx, http.StatusBadRequest, err.Error(), "Invalid request", "")
		return
	}

	if err := h.CreateNewOrder(order); err != nil {
		logger.ZapLog().Error(err.Error())
		service.SendResponse(ctx, http.StatusBadRequest, err.Error(), "Invalid request", "")
		return
	}

	service.SendResponse(ctx, http.StatusOK, "", "Order Created Successfully", fmt.Sprintf("To track your order : http://localhost:8000/order/%v", order.OrderID))
}

// Helps to create a new order
func (h *Handlers) GetOrderByID(ctx *context.Context) {
	var order models.Orders
	var claims models.Claims

	if err := helper.Claims(ctx.Request.Header["Authorization"][0], &claims); err != nil {
		logger.ZapLog().Error(err.Error())
		service.SendResponse(ctx, http.StatusBadRequest, err.Error(), "Invalid request", "")
		return
	}

	if err := h.ReadOrderByID(ctx.Input.Query(":id"), &order); err != nil {
		logger.ZapLog().Error(err.Error())
		service.SendResponse(ctx, http.StatusBadRequest, err.Error(), "Invalid request", "")
		return
	}

	if order.UserID != claims.UsersID {
		service.SendResponse(ctx, http.StatusBadRequest, "Unauthorized for this page", "Invalid request", "")
		return
	}
	service.SendResponse(ctx, http.StatusOK, "", "Here is your order details", order)

}
