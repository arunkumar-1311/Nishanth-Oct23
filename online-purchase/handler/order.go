package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"online-purchase/logger"
	"online-purchase/models"
	"online-purchase/service"
	"online-purchase/service/helper"
	"time"

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
	order.OrderStatusID = "S-UP"
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
	ctx.Output.Header("content-Type", "application/json")
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
	ctx.Output.Header("content-Type", "application/json")
	var orders []models.Orders
	var claims models.Claims

	if err := helper.Claims(ctx.Request.Header["Authorization"][0], &claims); err != nil {
		logger.ZapLog().Error(err.Error())
		service.SendResponse(ctx, http.StatusBadRequest, err.Error(), "Invalid request", "")
		return
	}

	if claims.RolesID == "USER1" {
		if err := h.ReadOrdersByUser(claims.UsersID, &orders); err != nil {
			logger.ZapLog().Error(err.Error())
			service.SendResponse(ctx, http.StatusInternalServerError, err.Error(), "Please try again later", "")
			return
		}
	} else {
		if err := h.ReadOrders(&orders); err != nil {
			logger.ZapLog().Error(err.Error())
			service.SendResponse(ctx, http.StatusInternalServerError, err.Error(), "Please try again later", "")
			return
		}
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
	ctx.Output.Header("content-Type", "application/json")
	var orderStatus []models.OrderStatus
	if err := h.ReadAllOrderStatus(&orderStatus); err != nil {
		logger.ZapLog().Error(err.Error())
		service.SendResponse(ctx, http.StatusInternalServerError, err.Error(), "Please try again later", "")
		return
	}

	service.SendResponse(ctx, http.StatusOK, "", "Fetching all order status", orderStatus)
}

// Helps to get the order by its status and its time
func (h *Handlers) GetOrderByStatus(ctx *context.Context) {
	ctx.Output.Header("content-Type", "application/json")
	var orders []models.Orders

	startDate := ctx.Input.Query("start")
	endDate := ctx.Input.Query("end")
	orderStatusID := ctx.Input.Query("order_status_id")
	if orderStatusID == "" {
		service.SendResponse(ctx, http.StatusBadRequest, "Expected order_status_id id", "Please try again later", "")
		return
	}

	from := time.Date(0, 0, 0, 0, 0, 0, 0, time.Local)
	if startDate != "" {
		var err error
		from, err = time.Parse("2006-01-02", startDate)
		if err != nil {
			logger.ZapLog().Error(err.Error())
			service.SendResponse(ctx, http.StatusBadRequest, err.Error(), "Please try again later", "")
			return
		}
	}

	to := time.Now()
	if endDate != "" {
		var err error
		to, err = time.Parse("2006-01-02", endDate)
		if err != nil {
			logger.ZapLog().Error(err.Error())
			service.SendResponse(ctx, http.StatusBadRequest, err.Error(), "Please try again later", "")
			return
		}
	}

	if err := h.ReadByStatus(orderStatusID, from, to, &orders); err != nil {
		logger.ZapLog().Error(err.Error())
		service.SendResponse(ctx, http.StatusBadRequest, err.Error(), "Please try again later", "")
		return
	}
	service.SendResponse(ctx, http.StatusOK, "", "Fetching all orders by status", orders)

}

// Helps to update the order
func (h *Handlers) UpdateStatus(ctx *context.Context) {
	ctx.Output.Header("content-Type", "application/json")
	var updateStatus map[string]string
	var order models.Orders

	if err := json.NewDecoder(ctx.Request.Body).Decode(&updateStatus); err != nil {
		logger.ZapLog().Error(err.Error())
		service.SendResponse(ctx, http.StatusInternalServerError, err.Error(), "Please try again later", "")
		return
	}

	if updateStatus["order_status_id"] == "" {
		service.SendResponse(ctx, http.StatusBadRequest, "Requried \"order_status_id\" field", "Invalid request", "")
		return
	}

	if updateStatus["order_status_id"] == "S-CAN" {
		service.SendResponse(ctx, http.StatusBadRequest, "You can't cancel the user order", "Invalid request", "")
		return
	}

	if err := h.UpdateOrderStatusByID(ctx.Input.Query(":id"), updateStatus["order_status_id"]); err != nil {
		logger.ZapLog().Error(err.Error())
		service.SendResponse(ctx, http.StatusBadRequest, err.Error(), "Invalid request", "")
		return
	}

	if err := h.ReadOrderByID(ctx.Input.Query(":id"), &order); err != nil {
		logger.ZapLog().Error(err.Error())
		service.SendResponse(ctx, http.StatusBadRequest, err.Error(), "Invalid request", "")
		return
	}

	service.SendResponse(ctx, http.StatusOK, "", "Order status changed", order)
}

// Helps to delete the order
func (h *Handlers) CancelOrder(ctx *context.Context) {
	var order models.Orders
	var claims models.Claims
	if err := h.ReadOrderByID(ctx.Input.Query(":id"), &order); err != nil {
		logger.ZapLog().Error(err.Error())
		service.SendResponse(ctx, http.StatusBadRequest, err.Error(), "Invalid request", "")
		return
	}

	if err := helper.Claims(ctx.Request.Header["Authorization"][0], &claims); err != nil {
		logger.ZapLog().Error(err.Error())
		service.SendResponse(ctx, http.StatusBadRequest, err.Error(), "Invalid request", "")
		return
	}

	if order.UserID != claims.UsersID {
		service.SendResponse(ctx, http.StatusBadRequest, "Unauthorized to cancel this order", "Invalid request", "")
		return
	}

	if order.OrderStatusID == "S-SH" || order.OrderStatusID == "S-DD" {
		service.SendResponse(ctx, http.StatusBadRequest, "You can't cancel this order product is already shipped", "Invalid request", "")
		return
	}

	if err := h.CancelOrderByID(ctx.Input.Query(":id")); err != nil {
		logger.ZapLog().Error(err.Error())
		service.SendResponse(ctx, http.StatusBadRequest, err.Error(), "Invalid request", "")
		return
	}

	service.SendResponse(ctx, http.StatusOK, "", fmt.Sprintf("Cancelled the order %v", ctx.Input.Query(":id")), "")
}
