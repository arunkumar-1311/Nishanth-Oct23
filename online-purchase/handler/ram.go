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

// Helps to create new ram product
func (h *Handlers) CreateRAM(ctx *context.Context) {
	ctx.Output.Header("content-Type", "application/json")
	var ram models.Ram

	if err := json.NewDecoder(ctx.Request.Body).Decode(&ram); err != nil {
		logger.ZapLog().Error(err.Error())
		service.SendResponse(ctx, http.StatusInternalServerError, err.Error(), "Please try again later", "")
		return
	}

	ram.RamID = helper.UniqueID()
	validate := validator.New()
	if err := validate.Struct(ram); err != nil {
		logger.ZapLog().Error(err.Error())
		service.SendResponse(ctx, http.StatusBadRequest, err.Error(), "Invalid request", "")
		return
	}

	if err := h.CreateNewRam(ram); err != nil {
		logger.ZapLog().Error(err.Error())
		service.SendResponse(ctx, http.StatusBadRequest, err.Error(), "Invalid request", "")
		return
	}

	service.SendResponse(ctx, http.StatusOK, "", "Ram Created Successfully", "")
}

// Helps to get all RAM's in th application
func (h *Handlers) GetAllRAMs(ctx *context.Context) {
	ctx.Output.Header("content-Type", "application/json")
	var ram []models.Ram

	if err := h.ReadRAMs(&ram); err != nil {
		logger.ZapLog().Error(err.Error())
		service.SendResponse(ctx, http.StatusInternalServerError, err.Error(), "Please try again later", "")
		return
	}

	service.SendResponse(ctx, http.StatusOK, "", "Fetching all RAM's", ram)
}

// Helps to get all RAM's in th application
func (h *Handlers) GetRamByID(ctx *context.Context) {
	ctx.Output.Header("content-Type", "application/json")
	var ram models.Ram

	if err := h.ReadRAMByID(ctx.Input.Query(":id"), &ram); err != nil {
		logger.ZapLog().Error(err.Error())
		service.SendResponse(ctx, http.StatusInternalServerError, err.Error(), "Please try again later", "")
		return
	}

	service.SendResponse(ctx, http.StatusOK, "", "Fetching RAM by ID", ram)
}

// Helps to update the existing ram
func (h *Handlers) UpdateRAM(ctx *context.Context) {
	ctx.Output.Header("content-Type", "application/json")
	var ram models.Ram

	if err := json.NewDecoder(ctx.Request.Body).Decode(&ram); err != nil {
		logger.ZapLog().Error(err.Error())
		service.SendResponse(ctx, http.StatusInternalServerError, err.Error(), "Please try again later", "")
		return
	}

	if err := h.UpdateRAMByID(ctx.Input.Query(":id"), ram); err != nil {
		logger.ZapLog().Error(err.Error())
		service.SendResponse(ctx, http.StatusBadRequest, err.Error(), "Invalid request", "")
		return
	}

	if err := h.ReadRAMByID(ctx.Input.Query(":id"), &ram); err != nil {
		logger.ZapLog().Error(err.Error())
		service.SendResponse(ctx, http.StatusBadRequest, err.Error(), "Invalid request", "")
		return
	}

	service.SendResponse(ctx, http.StatusOK, "", fmt.Sprintf("Updated %v ram successfully", ctx.Input.Query(":id")), ram)
}

// Delete the ram by id
func (h *Handlers) DeleteRAM(ctx *context.Context) {
	ctx.Output.Header("content-Type", "application/json")

	var ram models.Ram

	if err := h.ReadRAMByID(ctx.Input.Query(":id"), &ram); err != nil {
		logger.ZapLog().Error(err.Error())
		service.SendResponse(ctx, http.StatusBadRequest, err.Error(), "Invalid request", "")
		return
	}

	if ram.Status {
		service.SendResponse(ctx, http.StatusBadRequest, "Can't delete active Product", "Invalid request", "")
		return
	}

	if err := h.DeleteRAMByID(ctx.Input.Query(":id")); err != nil {
		logger.ZapLog().Error(err.Error())
		service.SendResponse(ctx, http.StatusBadRequest, err.Error(), "Invalid request", "")
		return
	}
	service.SendResponse(ctx, http.StatusOK, "", fmt.Sprintf("Deleted %v ram successfully", ctx.Input.Query(":id")), "")
}
