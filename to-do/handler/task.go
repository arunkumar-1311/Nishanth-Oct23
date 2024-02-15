package handler

import (
	"net/http"
	"to-do/logger"
	"to-do/models"
	"to-do/service/helper"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo"
)

type Task interface {
	AddTask(c echo.Context) error
	DeleteTask(c echo.Context) error
	UpdateTask(c echo.Context) error
	UpdateAllTaskStatus(c echo.Context) error
}

// Helps to add the task
func (e *EndPoint) AddTask(c echo.Context) error {
	var task models.Tasks
	if err := c.Bind(&task); err != nil {
		logger.ZeroLogger().Msg(err.Error())
		return helper.SendResponse(c, "", echo.ErrBadRequest.Code, err.Error(), "Invalid request")
	}
	task.TaskID = e.Service.UniqueID()
	task.Active = true
	task.UsersID = c.Get("userid").(string)

	validate := validator.New()
	if err := validate.Struct(task); err != nil {
		logger.ZeroLogger().Msg(err.Error())
		return helper.SendResponse(c, "", echo.ErrBadRequest.Code, err.Error(), "Invalid request")
	}

	if err := e.DB.AddTask(task); err != nil {
		logger.ZeroLogger().Msg(err.Error())
		return helper.SendResponse(c, "", echo.ErrBadRequest.Code, err.Error(), "Invalid request")
	}

	return helper.SendResponse(c, "Task created successfully", http.StatusOK, "", "Task Added")
}

// Helps to delete the existing task
func (e *EndPoint) DeleteTask(c echo.Context) error {
	taskID := c.Param("id")
	var task models.Tasks
	if err := e.DB.ReadTaskByID(taskID, &task); err != nil {
		logger.ZeroLogger().Msg(err.Error())
		return helper.SendResponse(c, "", echo.ErrBadRequest.Code, err.Error(), "Invalid request")
	}

	if task.UsersID != c.Get("userid").(string) {
		return helper.SendResponse(c, "", echo.ErrBadRequest.Code, "Unauthorized to delete this task", "Invalid request")
	}

	if err := e.DB.DeleteTask(taskID); err != nil {
		logger.ZeroLogger().Msg(err.Error())
		return helper.SendResponse(c, "", echo.ErrBadRequest.Code, "Unauthorized to delete this task", "Invalid request")
	}

	return helper.SendResponse(c, "", http.StatusOK, "Task Deleted Successfully", "task deleted")
}

// Helps to update the task
func (e *EndPoint) UpdateTask(c echo.Context) error {
	var updateTask map[string]string
	if err := c.Bind(&updateTask); err != nil {
		logger.ZeroLogger().Msg(err.Error())
		return helper.SendResponse(c, "", echo.ErrBadRequest.Code, err.Error(), "Invalid request")
	}

	if updateTask["task"] == "" {
		return helper.SendResponse(c, "", echo.ErrBadRequest.Code, "content is empty", "Invalid request")
	}
	taskID := c.Param("id")
	var task models.Tasks
	if err := e.DB.ReadTaskByID(taskID, &task); err != nil {
		logger.ZeroLogger().Msg(err.Error())
		return helper.SendResponse(c, "", echo.ErrBadRequest.Code, err.Error(), "Invalid request")
	}

	if task.UsersID != c.Get("userid").(string) {
		return helper.SendResponse(c, "", echo.ErrBadRequest.Code, "Unauthorized to delete this task", "Invalid request")
	}

	if err := e.DB.UpdateTask(taskID, updateTask["task"]); err != nil {
		logger.ZeroLogger().Msg(err.Error())
		return helper.SendResponse(c, "", echo.ErrBadRequest.Code, err.Error(), "Invalid request")
	}

	if err := e.DB.ReadTaskByID(taskID, &task); err != nil {
		logger.ZeroLogger().Msg(err.Error())
		return helper.SendResponse(c, "", echo.ErrBadRequest.Code, err.Error(), "Invalid request")
	}

	return helper.SendResponse(c, task, http.StatusOK, "Task updated Successfully", "task updated")
}

// Helps to update all tasks status
func (e *EndPoint) UpdateAllTaskStatus(c echo.Context) error {

	if err := e.DB.UpdateAllStatus(c.Get("userid").(string)); err != nil {
		logger.ZeroLogger().Msg(err.Error())
		return helper.SendResponse(c, "", echo.ErrInternalServerError.Code, err.Error(), "Try Again later")
	}
	return helper.SendResponse(c, "", http.StatusOK, "Tasks status updated Successfully", "task updated")
}
