package admin

import (
	"blog_post/adaptor"
	"blog_post/logger"
	"blog_post/models"
	"blog_post/repository"
	"blog_post/service"
	"blog_post/service/helper"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

// Helps to create the post
func CreatePost(c *fiber.Ctx) error {

	c.Accepts("application/json")
	var post models.Post
	validate := validator.New()

	if err := c.BodyParser(&post); err != nil {
		logger.Logging().Error(err)
		service.SendResponse(c, fiber.StatusInternalServerError, err.Error(), "Try Again Later", fiber.MethodPost, "")
		return nil
	}

	for _, value := range post.CategoryID {
		if _, err := repository.ReadCategoryByID(value); err != nil {
			logger.Logging().Error(err)
			service.SendResponse(c, fiber.StatusBadRequest, err.Error(), "Try Again Later", fiber.MethodPost, "")
			return nil
		}
	}
	post.PostID = helper.UniqueID()
	if err := validate.Struct(post); err != nil {
		logger.Logging().Error(err)
		service.SendResponse(c, fiber.StatusBadRequest, err.Error(), "Give all required fields", fiber.MethodPost, "")
		return nil
	}

	if err := adaptor.GetConn().Create(&post).Error; err != nil {
		logger.Logging().Error(err)
		service.SendResponse(c, fiber.StatusBadRequest, err.Error(), "User Valid Values", fiber.MethodPost, "")
		return nil
	}
	service.SendResponse(c, fiber.StatusOK, "", "Post Added", fiber.MethodPost, "Post Created successfully")

	return nil
}

// Helps to read all the post
func ReadAllPosts(c *fiber.Ctx) error {

	var allPosts models.AllPost
	if err := repository.ReadPosts(&allPosts.Post); err != nil {
		logger.Logging().Error(err)
		service.SendResponse(c, fiber.StatusBadRequest, err.Error(), "Oops error occurs", fiber.MethodGet, "")
		return nil
	}

	if err := helper.CommentsAndCategory(allPosts.Post); err != nil {
		logger.Logging().Error(err)
		service.SendResponse(c, fiber.StatusBadRequest, err.Error(), "Oops error occurs", fiber.MethodGet, "")
		return nil
	}

	var postResp []models.PostResponse

	if err := helper.PostResp(allPosts, &postResp); err != nil {
		logger.Logging().Error(err)
		service.SendResponse(c, fiber.StatusBadRequest, err.Error(), "Oops error occurs", fiber.MethodGet, "")
		return nil
	}

	service.SendResponse(c, fiber.StatusOK, "", "All available Posts", fiber.MethodGet, postResp)
	return nil
}

// Helps to update the existing post
func UpdatePost(c *fiber.Ctx) error {

	c.Accepts("application/json")
	var post models.Post
	id := c.Params("id")
	if err := c.BodyParser(&post); err != nil {
		logger.Logging().Error(err)
		service.SendResponse(c, fiber.StatusInternalServerError, err.Error(), "Try Again Later", fiber.MethodPatch, "")
		return nil
	}

	if err := repository.UpdatePost(id, post); err != nil {
		logger.Logging().Error(err)
		service.SendResponse(c, fiber.StatusBadRequest, err.Error(), "Use valid id", fiber.MethodPatch, "")
		return nil
	}

	service.SendResponse(c, fiber.StatusOK, "", "Post Updated successfully", fiber.MethodPatch, "updated successfully")
	return nil
}

// helps to delete the post with its id
func DeletePost(c *fiber.Ctx) error {

	id := c.Params("id")
	if err := repository.DeletePost(id); err != nil {
		logger.Logging().Error(err)
		service.SendResponse(c, fiber.StatusBadRequest, err.Error(), "use valid id", fiber.MethodDelete, "")
		return nil
	}
	service.SendResponse(c, fiber.StatusOK, "", "Post Deleted successfully", fiber.MethodDelete, "")
	return nil
}

// Its shows overview statistics of the website
func Overview(c *fiber.Ctx) error {

	var overview models.Overview

	if err := repository.Overview(&overview); err != nil {
		logger.Logging().Error(err)
		service.SendResponse(c, fiber.StatusInternalServerError, err.Error(), "Try Again Later", fiber.MethodGet, "")
		return nil
	}

	service.SendResponse(c, fiber.StatusOK, "", "Blog Created successfully", fiber.MethodGet, overview)
	return nil
}
