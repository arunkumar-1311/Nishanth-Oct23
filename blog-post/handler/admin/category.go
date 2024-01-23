package admin

import (
	"blog_post/logger"
	"blog_post/models"
	"blog_post/repository"
	"blog_post/service"
	"blog_post/service/helper"


	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

// Helps to create the category
func CreateCategory(c *fiber.Ctx) error {

	c.Accepts("application/json")
	var category models.Category
	validate := validator.New()

	if err := c.BodyParser(&category); err != nil {
		logger.Logging().Error(err)
		service.SendResponse(c, fiber.StatusInternalServerError, err.Error(), "Try Again Later", fiber.MethodPost, "")
		return nil
	}
	category.CategoryID = helper.UniqueID()
	if err := validate.Struct(category); err != nil {
		logger.Logging().Error(err)
		service.SendResponse(c, fiber.StatusBadRequest, err.Error(), "Give all required fields", fiber.MethodPost, "")
		return nil
	}

	if err := repository.CreateCategory(category); err != nil {
		logger.Logging().Error(err)
		service.SendResponse(c, fiber.StatusBadRequest, err.Error(), "Oops error occurs", fiber.MethodPost, "")
		return nil
	}
	service.SendResponse(c, fiber.StatusOK, "", "Category Added", fiber.MethodPost, "Category Created successfully")
	return nil
}

// Helps to read all the category
func ReadAllCategories(c *fiber.Ctx) error {

	var categories []models.CategoryResponse
	if err := repository.ReadCategories(&categories); err != nil {
		logger.Logging().Error(err)
		service.SendResponse(c, fiber.StatusBadRequest, err.Error(), "Oops error occurs", fiber.MethodGet, "")
		return nil
	}

	
	if err := helper.Categories(categories); err != nil {
		logger.Logging().Error(err)
		service.SendResponse(c, fiber.StatusBadRequest, err.Error(), "Oops error occurs", fiber.MethodGet, "")
		return nil
	}
	service.SendResponse(c, fiber.StatusOK, "", "All available categories", fiber.MethodGet, categories)
	return nil
}

// Helps to update the existing post
func UpdateCategory(c *fiber.Ctx) error {

	c.Accepts("application/json")
	var category models.Category
	id := c.Params("id")
	if err := c.BodyParser(&category); err != nil {
		logger.Logging().Error(err)
		service.SendResponse(c, fiber.StatusInternalServerError, err.Error(), "Try Again Later", fiber.MethodPatch, "")
		return nil
	}
	if err := repository.UpdateCategory(id, category); err != nil {
		logger.Logging().Error(err)
		service.SendResponse(c, fiber.StatusBadRequest, err.Error(), "Use valid id", fiber.MethodPatch, "")
		return nil
	}

	service.SendResponse(c, fiber.StatusOK, "", "Category Updated successfully", fiber.MethodPatch, "updated successfully")
	return nil
}

// helps to delete the post with its id
func DeleteCategory(c *fiber.Ctx) error {

	id := c.Params("id")
	if err := repository.DeleteCategory(id); err != nil {
		logger.Logging().Error(err)
		service.SendResponse(c, fiber.StatusBadRequest, err.Error(), "use valid id", fiber.MethodDelete, "")
		return nil
	}
	service.SendResponse(c, fiber.StatusOK, "", "Category Deleted successfully", fiber.MethodDelete, "")
	return nil
}
