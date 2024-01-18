package admin

import (
	"blog_post/logger"
	"blog_post/models"
	"blog_post/repository"
	"blog_post/service"
	"blog_post/service/helper"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

// Helps to create the category
func CreateCategory(c *fiber.Ctx) error {

	if err := helper.AdminAccess(c.Get("Authorization")); err != nil {
		logger.Logging().Error(err)
		service.SendResponse(c, http.StatusBadRequest, err.Error(), "Try Again Later", http.MethodPost, "")
		return nil
	}

	c.Accepts("application/json")
	var category models.Category
	validate := validator.New()

	if err := c.BodyParser(&category); err != nil {
		logger.Logging().Error(err)
		service.SendResponse(c, http.StatusInternalServerError, err.Error(), "Try Again Later", http.MethodPost, "")
		return nil
	}

	if err := validate.Struct(category); err != nil {
		logger.Logging().Error(err)
		service.SendResponse(c, http.StatusBadRequest, err.Error(), "Give all required fields", http.MethodPost, "")
		return nil
	}

	if err := repository.CreateCategory(category); err != nil {
		logger.Logging().Error(err)
		service.SendResponse(c, http.StatusBadRequest, err.Error(), "Oops error occurs", http.MethodPost, "")
		return nil
	}
	service.SendResponse(c, http.StatusOK, "", "Category Added", http.MethodPost, "Category Created successfully")
	return nil
}

// Helps to read all the category
func ReadAllCategories(c *fiber.Ctx) error {

	var categories []models.Category
	if err := repository.ReadCategories(&categories); err != nil {
		logger.Logging().Error(err)
		service.SendResponse(c, http.StatusBadRequest, err.Error(), "Oops error occurs", http.MethodGet, "")
		return nil
	}

	service.SendResponse(c, http.StatusOK, "", "All available categories", http.MethodGet, categories)
	return nil
}

// Helps to update the existing post
func UpdateCategory(c *fiber.Ctx) error {

	if err := helper.AdminAccess(c.Get("Authorization")); err != nil {
		logger.Logging().Error(err)
		service.SendResponse(c, http.StatusBadRequest, err.Error(), "Try Again Later", http.MethodPatch, "")
		return nil
	}

	c.Accepts("application/json")
	var category models.Category
	id := c.Params("id")
	if err := c.BodyParser(&category); err != nil {
		logger.Logging().Error(err)
		service.SendResponse(c, http.StatusInternalServerError, err.Error(), "Try Again Later", http.MethodPatch, "")
		return nil
	}
	if err := repository.UpdateCategory(id, category); err != nil {
		logger.Logging().Error(err)
		service.SendResponse(c, http.StatusBadRequest, err.Error(), "Use valid id", http.MethodPatch, "")
		return nil
	}

	service.SendResponse(c, http.StatusOK, "", "Category Updated successfully", http.MethodPatch, "updated successfully")
	return nil
}

// helps to delete the post with its id
func DeleteCategory(c *fiber.Ctx) error {

	if err := helper.AdminAccess(c.Get("Authorization")); err != nil {
		logger.Logging().Error(err)
		service.SendResponse(c, http.StatusBadRequest, err.Error(), "Try Again Later", http.MethodDelete, "")
		return nil
	}

	id := c.Params("id")
	if err := repository.DeleteCategory(id); err != nil {
		logger.Logging().Error(err)
		service.SendResponse(c, http.StatusBadRequest, err.Error(), "use valid id", http.MethodDelete, "")
		return nil
	}
	service.SendResponse(c, http.StatusOK, "", "Category Deleted successfully", http.MethodDelete, "")
	return nil
}
