package handler

import (
	"net/http"
	"strconv"
	"to-do/logger"
	"to-do/models"
	"to-do/service/helper"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo"
)

type Task interface {
	AddTask(c echo.Context) error

	ReadAllTasks(c echo.Context) (err error)
	GetDeletedTasks(c echo.Context) error

	UpdateTask(c echo.Context) error
	UpdateAllTaskStatus(c echo.Context) error
	UpdateTaskStatus(c echo.Context) error

	DeleteTask(c echo.Context) error
	ClearCompletedTasks(c echo.Context) error
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

// Helps to update the status of single task
func (e *EndPoint) UpdateTaskStatus(c echo.Context) error {

	var task models.Tasks
	taskID := c.Param("id")
	if err := e.DB.ReadTaskByID(taskID, &task); err != nil {
		logger.ZeroLogger().Msg(err.Error())
		return helper.SendResponse(c, "", echo.ErrBadRequest.Code, err.Error(), "Invalid request")
	}

	if c.Get("userid").(string) != task.UsersID {
		return helper.SendResponse(c, "", echo.ErrBadRequest.Code, "Not authorized to update the status", "Invalid request")
	}

	if err := e.DB.UpdateTaskStatus(taskID, task.Active); err != nil {
		logger.ZeroLogger().Msg(err.Error())
		return helper.SendResponse(c, "", echo.ErrInternalServerError.Code, err.Error(), "Try Again later")
	}
	return helper.SendResponse(c, "", http.StatusOK, "Task status updated Successfully", "task updated")
}

// helps to read the task by its status
func (e *EndPoint) ReadAllTasks(c echo.Context) (err error) {

	var taskResponse models.TaskResponse
	status := c.QueryParam("status")
	var taskStatus bool
	if status != "" {
		if taskStatus, err = strconv.ParseBool(status); err != nil {
			return helper.SendResponse(c, "", echo.ErrBadRequest.Code, err.Error(), "status must be a true or false")
		}

		if err := e.DB.ReadAllTask(c.Get("userid").(string), &taskResponse.Task, taskStatus); err != nil {
			logger.ZeroLogger().Msg(err.Error())
			return helper.SendResponse(c, "", echo.ErrInternalServerError.Code, err.Error(), "Try Again later")
		}
	} else {
		if err := e.DB.ReadAllTask(c.Get("userid").(string), &taskResponse.Task); err != nil {
			logger.ZeroLogger().Msg(err.Error())
			return helper.SendResponse(c, "", echo.ErrInternalServerError.Code, err.Error(), "Try Again later")
		}
	}

	if err := e.DB.CountActiveTasks(c.Get("userid").(string), &taskResponse.Active); err != nil {
		logger.ZeroLogger().Msg(err.Error())
		return helper.SendResponse(c, "", echo.ErrInternalServerError.Code, err.Error(), "Try Again later")
	}

	return helper.SendResponse(c, taskResponse, http.StatusOK, "Fetching all tasks", "read task by status")
}

// Helps to clear all completed tasks
func (e *EndPoint) ClearCompletedTasks(c echo.Context) error {

	if err := e.DB.ClearCompleted(c.Get("userid").(string)); err != nil {
		logger.ZeroLogger().Msg(err.Error())
		return helper.SendResponse(c, "", echo.ErrInternalServerError.Code, err.Error(), "Try Again later")
	}
	return helper.SendResponse(c, "", http.StatusOK, "Deleted all completed tasks", "All completed task deleted")
}

// Helps to read all deleted tasks
func (e *EndPoint) GetDeletedTasks(c echo.Context) error {
	var tasks []models.Tasks
	if err := e.DB.ReadDeletedTask(c.Get("userid").(string), &tasks); err != nil {
		logger.ZeroLogger().Msg(err.Error())
		return helper.SendResponse(c, "", echo.ErrInternalServerError.Code, err.Error(), "Try Again later")
	}
	return helper.SendResponse(c, tasks, http.StatusOK, "Fetching all deleted tasks", "read all deleted")
}
