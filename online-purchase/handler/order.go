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

	if !brand.Status {
		service.SendResponse(ctx, http.StatusOK, "Sorry this brand is out of stock", "Invalid request", "")
		return
	}

	if !ram.Status {
		service.SendResponse(ctx, http.StatusOK, "Sorry this ram is out of stock", "Invalid request", "")
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

	service.SendResponse(ctx, http.StatusOK, "", "Order Placed Successfully", fmt.Sprintf("To track your order : http://localhost:8000/order/%v", order.OrderID))
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

	if order.UserID != claims.UsersID && claims.RolesID != "AD1" {
		service.SendResponse(ctx, http.StatusBadRequest, "Unauthorized for this page", "Invalid request", "")
		return
	}
	service.SendResponse(ctx, http.StatusOK, "", "Here is your order details", order)

}

// Helps to read all the orders by the admin
func (h *Handlers) GetAllOrders(ctx *context.Context) {
	var orders []models.Orders

	if err := h.ReadOrders(&orders); err != nil {
		logger.ZapLog().Error(err.Error())
		service.SendResponse(ctx, http.StatusInternalServerError, err.Error(), "Please try again later", "")
		return
	}

	orderResponse := make([]models.OrderResponse, len(orders))
	for index, order := range orders {
		orderMarshal, err := json.Marshal(order)
		if err != nil {
			logger.ZapLog().Error(err.Error())
			service.SendResponse(ctx, http.StatusInternalServerError, err.Error(), "Please try again later", "")
			return
		}

		if err := json.Unmarshal(orderMarshal, &orderResponse[index]); err != nil {
			logger.ZapLog().Error(err.Error())
			service.SendResponse(ctx, http.StatusInternalServerError, err.Error(), "Please try again later", "")
			return
		}
	}

	service.SendResponse(ctx, http.StatusOK, "", "Fetching all orders", orderResponse)
}

// Helps to get all the order status
func (h *Handlers) GetAllOrderStatus(ctx *context.Context) {
	var orderStatus []models.OrderStatus
	if err := h.ReadAllOrderStatus(&orderStatus); err != nil {
		logger.ZapLog().Error(err.Error())
		service.SendResponse(ctx, http.StatusInternalServerError, err.Error(), "Please try again later", "")
		return
	}

	service.SendResponse(ctx, http.StatusOK, "", "Fetching all order status", orderStatus)
}

// Helps to cancel the order by the ordered user
func (h *Handlers) CancelOrder(ctx *context.Context) {
	var claims models.Claims
	var order models.Orders
	var cancelOrder map[string]interface{}

	if err := json.NewDecoder(ctx.Request.Body).Decode(&cancelOrder); err != nil {
		logger.ZapLog().Error(err.Error())
		service.SendResponse(ctx, http.StatusInternalServerError, err.Error(), "Please try again later", "")
		return
	}

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
		service.SendResponse(ctx, http.StatusBadRequest, "Unauthorized to cancel the order", "Invalid request", "")
		return
	}

	if cancelOrder["active"] == nil {
		logger.ZapLog().Error("Needed \"active\" field")
		service.SendResponse(ctx, http.StatusBadRequest, "field \"active\" is empty", "Invalid request", "")
		return
	}

	active, ok := cancelOrder["active"].(bool)
	if !ok {
		logger.ZapLog().Error("Invalid datatype expected boolean")
		service.SendResponse(ctx, http.StatusBadRequest, "Invalid datatype expected boolean", "Invalid request", "")
		return
	}

	if err := h.CancelOrderByID(ctx.Input.Query(":id"), active); err != nil {
		logger.ZapLog().Error(err.Error())
		service.SendResponse(ctx, http.StatusBadRequest, err.Error(), "Invalid request", "")
		return
	}

	service.SendResponse(ctx, http.StatusOK, "", "Order status changed", fmt.Sprintf("To track your order : http://localhost:8000/order/%v", order.OrderID))
}

// Helps to update the order
func (h *Handlers) UpdateStatus(ctx *context.Context) {

	var updateStatus map[string]interface{}
	var claims models.Claims
	var order models.Orders

	if err := json.NewDecoder(ctx.Request.Body).Decode(&updateStatus); err != nil {
		logger.ZapLog().Error(err.Error())
		service.SendResponse(ctx, http.StatusInternalServerError, err.Error(), "Please try again later", "")
		return
	}

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

	if updateStatus["order_status_id"] == nil {
		service.SendResponse(ctx, http.StatusBadRequest, "Requried \"order_status_id\" field", "Invalid request", "")
		return
	}

	status, ok := updateStatus["order_status_id"].(string)
	if !ok {
		logger.ZapLog().Error("Invalid datatype expected string")
		service.SendResponse(ctx, http.StatusBadRequest, "Invalid datatype expected string", "Invalid request", "")
		return
	}

	if err := h.UpdateOrderStatusByID(ctx.Input.Query(":id"), status); err != nil {
		logger.ZapLog().Error(err.Error())
		service.SendResponse(ctx, http.StatusBadRequest, err.Error(), "Invalid request", "")
		return
	}

	if status == "S-DD" {
		if err := h.CancelOrderByID(ctx.Input.Query(":id"), false); err != nil {
			logger.ZapLog().Error(err.Error())
			service.SendResponse(ctx, http.StatusBadRequest, err.Error(), "Invalid request", "")
			return
		}
	}

	if err := h.ReadOrderByID(ctx.Input.Query(":id"), &order); err != nil {
		logger.ZapLog().Error(err.Error())
		service.SendResponse(ctx, http.StatusBadRequest, err.Error(), "Invalid request", "")
		return
	}

	service.SendResponse(ctx, http.StatusOK, "", "Order status changed", order)
}

// Helps to delete the order
func (h *Handlers) DeleteOrder(ctx *context.Context) {
	var order models.Orders

	if err := h.ReadOrderByID(ctx.Input.Query(":id"), &order); err != nil {
		logger.ZapLog().Error(err.Error())
		service.SendResponse(ctx, http.StatusBadRequest, err.Error(), "Invalid request", "")
		return
	}

	if order.Active {
		service.SendResponse(ctx, http.StatusBadRequest, "Can't Delete the active order", "Invalid request", "")
		return
	}

	if err := h.DeleteOrderByID(ctx.Input.Query(":id")); err != nil {
		logger.ZapLog().Error(err.Error())
		service.SendResponse(ctx, http.StatusBadRequest, err.Error(), "Invalid request", "")
		return
	}

	service.SendResponse(ctx, http.StatusOK, "", fmt.Sprintf("Deleted the order %v", ctx.Input.Query(":id")), "")
}
