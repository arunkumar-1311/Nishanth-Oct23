package handler

import (
	"blog_post/logger"
	"blog_post/models"
	"blog_post/repository"
	"blog_post/service"
	"blog_post/service/helper"
	"fmt"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
)

// Helps to filter the post by its published date
func DateFilter(c *fiber.Ctx) error {
	var filter map[string]string
	var Posts models.AllPost

	if err := c.BodyParser(&filter); err != nil {
		logger.Logging().Error(err)
		service.SendResponse(c, http.StatusInternalServerError, err.Error(), "Try Again Later", http.MethodPost, "")
		return nil
	}

	fromDate, err := time.Parse("2006-01-02", fmt.Sprint(filter["from"], "-01"))
	if err != nil {
		logger.Logging().Error(err)
		service.SendResponse(c, http.StatusInternalServerError, err.Error(), "Try Again Later", http.MethodPost, "")
		return nil
	}

	toDate, err := time.Parse("2006-01-02", fmt.Sprint(filter["to"], "-01"))
	if err != nil {
		logger.Logging().Error(err)
		service.SendResponse(c, http.StatusInternalServerError, err.Error(), "Try Again Later", http.MethodPost, "")
		return nil
	}

	toDate = time.Date(toDate.Year(), toDate.Month() +1, toDate.Day() - 1, 23, 59, 59, 999999999, time.Local)
	err = repository.DateFilter(fromDate, toDate, &Posts.Post)
	if err != nil {
		logger.Logging().Error(err)
		service.SendResponse(c, http.StatusInternalServerError, err.Error(), "Try Again Later", http.MethodPost, "")
		return nil
	}

	if err := helper.CommentsAndCategory(Posts.Post); err != nil {
		logger.Logging().Error(err)
		service.SendResponse(c, http.StatusBadRequest, err.Error(), "Oops error occurs", http.MethodPost, "")
		return nil
	}
	service.SendResponse(c, http.StatusOK, "", "All available categories", http.MethodPost, Posts)

	return nil
}

// Helps to filter the post by its Category
func CategoryFilter(c *fiber.Ctx) error {

	var Posts models.AllPost

	if err := repository.CategoryFilter(c.Params("id"), &Posts.Post); err != nil {
		logger.Logging().Error(err)
		service.SendResponse(c, http.StatusBadRequest, err.Error(), "Oops error occurs", http.MethodGet, "")
		return nil
	}

	if err := helper.CommentsAndCategory(Posts.Post); err != nil {
		logger.Logging().Error(err)
		service.SendResponse(c, http.StatusBadRequest, err.Error(), "Oops error occurs", http.MethodGet, "")
		return nil
	}
	service.SendResponse(c, http.StatusOK, "", "All available categories", http.MethodPost, Posts)
	return nil
}
