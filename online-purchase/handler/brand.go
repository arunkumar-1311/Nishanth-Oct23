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

// Helps to create new brand
func (h *Handlers) CreateBrand(ctx *context.Context) {
	ctx.Output.Header("content-Type", "application/json")
	var brand models.Brand

	if err := json.NewDecoder(ctx.Request.Body).Decode(&brand); err != nil {
		logger.ZapLog().Error(err.Error())
		service.SendResponse(ctx, http.StatusInternalServerError, err.Error(), "Please try again later", "")
		return
	}

	brand.BrandID = helper.UniqueID()
	validate := validator.New()
	if err := validate.Struct(brand); err != nil {
		logger.ZapLog().Error(err.Error())
		service.SendResponse(ctx, http.StatusBadRequest, err.Error(), "Invalid request", "")
		return
	}

	if err := h.CreateNewBrand(brand); err != nil {
		logger.ZapLog().Error(err.Error())
		service.SendResponse(ctx, http.StatusBadRequest, err.Error(), "Invalid request", "")
		return
	}
	
	service.SendResponse(ctx, http.StatusOK, "", "Brand Created Successfully", "")
}

// Helps to get all brands in th application
func (h *Handlers) GetBrands(ctx *context.Context) {
	ctx.Output.Header("content-Type", "application/json")
	var brands []models.Brand

	if err := h.ReadBrands(&brands); err != nil {
		logger.ZapLog().Error(err.Error())
		service.SendResponse(ctx, http.StatusInternalServerError, err.Error(), "Please try again later", "")
		return
	}

	service.SendResponse(ctx, http.StatusOK, "", "Fetching all brands", brands)
}

// Helps to update the existing brand
func (h *Handlers) UpdateBrand(ctx *context.Context) {
	ctx.Output.Header("content-Type", "application/json")
	var brand models.Brand

	if err := json.NewDecoder(ctx.Request.Body).Decode(&brand); err != nil {
		logger.ZapLog().Error(err.Error())
		service.SendResponse(ctx, http.StatusInternalServerError, err.Error(), "Please try again later", "")
		return
	}

	
	if err := h.UpdateBrandByID(ctx.Input.Query(":id"), brand); err != nil {
		logger.ZapLog().Error(err.Error())
		service.SendResponse(ctx, http.StatusBadRequest, err.Error(), "Invalid request", "")
		return
	}

	if err := h.ReadBrandByID(ctx.Input.Query(":id"), &brand); err != nil {
		logger.ZapLog().Error(err.Error())
		service.SendResponse(ctx, http.StatusBadRequest, err.Error(), "Invalid request", "")
		return
	}

	service.SendResponse(ctx, http.StatusOK, "", fmt.Sprintf("Updated %v brand successfully", ctx.Input.Query(":id")), brand)
}

// Delete the brand by id
func (h *Handlers) DeleteBrand(ctx *context.Context) {
	ctx.Output.Header("content-Type", "application/json")

	var brand models.Brand
	if err := h.ReadBrandByID(ctx.Input.Query(":id"), &brand); err != nil {
		logger.ZapLog().Error(err.Error())
		service.SendResponse(ctx, http.StatusBadRequest, err.Error(), "Invalid request", "")
		return
	}

	if brand.Status {
		service.SendResponse(ctx, http.StatusBadRequest, "Can't delete active brand", "Invalid request", "")
		return
	}
	
	if err := h.DeleteBrandByID(ctx.Input.Query(":id")); err != nil {
		logger.ZapLog().Error(err.Error())
		service.SendResponse(ctx, http.StatusBadRequest, err.Error(), "Invalid request", "")
		return
	}
	service.SendResponse(ctx, http.StatusOK, "", fmt.Sprintf("Deleted %v brand successfully", ctx.Input.Query(":id")), "")
}
