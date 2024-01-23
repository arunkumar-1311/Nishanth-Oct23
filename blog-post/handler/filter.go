package handler

import (
	"blog_post/logger"
	"blog_post/models"
	"blog_post/repository"
	"blog_post/service"
	"blog_post/service/helper"
	"time"

	"github.com/gofiber/fiber/v2"
)

// Helps to filter the post by its published date
func DateFilter(c *fiber.Ctx) error {
	var filter map[string]string
	var Posts models.AllPost

	if err := c.BodyParser(&filter); err != nil {
		logger.Logging().Error(err)
		service.SendResponse(c, fiber.StatusInternalServerError, err.Error(), "Try Again Later", fiber.MethodPost, "")
		return nil
	}

	fromDate, err := time.Parse("2006-01-02", filter["from"])
	if err != nil {
		logger.Logging().Error(err)
		service.SendResponse(c, fiber.StatusInternalServerError, err.Error(), "Try Again Later", fiber.MethodPost, "")
		return nil
	}

	toDate, err := time.Parse("2006-01-02", filter["to"])
	if err != nil {
		logger.Logging().Error(err)
		service.SendResponse(c, fiber.StatusInternalServerError, err.Error(), "Try Again Later", fiber.MethodPost, "")
		return nil
	}

	toDate = time.Date(toDate.Year(), toDate.Month(), toDate.Day(), 23, 59, 59, 999999999, time.Local)
	err = repository.DateFilter(fromDate, toDate, &Posts.Post)
	if err != nil {
		logger.Logging().Error(err)
		service.SendResponse(c, fiber.StatusInternalServerError, err.Error(), "Try Again Later", fiber.MethodPost, "")
		return nil
	}

	if err := helper.CommentsAndCategory(Posts.Post); err != nil {
		logger.Logging().Error(err)
		service.SendResponse(c, fiber.StatusBadRequest, err.Error(), "Oops error occurs", fiber.MethodPost, "")
		return nil
	}

	postResp := new([]models.PostResponse)
	if err := helper.PostResp(Posts, postResp); err != nil {
		logger.Logging().Error(err)
		service.SendResponse(c, fiber.StatusBadRequest, err.Error(), "Oops error occurs", fiber.MethodGet, "")
		return nil
	}
	service.SendResponse(c, fiber.StatusOK, "", "All available categories", fiber.MethodPost, postResp)

	return nil
}

// Helps to filter the post by its Category
func CategoryFilter(c *fiber.Ctx) error {

	var Posts models.AllPost

	if err := repository.CategoryFilter(c.Params("id"), &Posts.Post); err != nil {
		logger.Logging().Error(err)
		service.SendResponse(c, fiber.StatusBadRequest, err.Error(), "Oops error occurs", fiber.MethodGet, "")
		return nil
	}

	if err := helper.CommentsAndCategory(Posts.Post); err != nil {
		logger.Logging().Error(err)
		service.SendResponse(c, fiber.StatusBadRequest, err.Error(), "Oops error occurs", fiber.MethodGet, "")
		return nil
	}

	postResp := new([]models.PostResponse)
	if err := helper.PostResp(Posts, postResp); err != nil {
		logger.Logging().Error(err)
		service.SendResponse(c, fiber.StatusBadRequest, err.Error(), "Oops error occurs", fiber.MethodGet, "")
		return nil
	}
	service.SendResponse(c, fiber.StatusOK, "", "All available categories", fiber.MethodPost, postResp)
	return nil
}
