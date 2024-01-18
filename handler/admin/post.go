package admin

import (
	"blog_post/adaptor"
	"blog_post/logger"
	"blog_post/models"
	"blog_post/repository"
	"blog_post/service"
	"blog_post/service/helper"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

// Helps to create the post
func CreatePost(c *fiber.Ctx) error {

	if err := helper.AdminAccess(c.Get("Authorization")); err != nil {
		logger.Logging().Error(err)
		service.SendResponse(c, http.StatusBadRequest, err.Error(), "Try Again Later", http.MethodPatch, "")
		return nil
	}

	c.Accepts("application/json")
	var post models.Post
	validate := validator.New()

	if err := c.BodyParser(&post); err != nil {
		logger.Logging().Error(err)
		service.SendResponse(c, http.StatusInternalServerError, err.Error(), "Try Again Later", http.MethodPost, "")
		return nil
	}

	if err := validate.Struct(post); err != nil {
		logger.Logging().Error(err)
		service.SendResponse(c, http.StatusBadRequest, err.Error(), "Give all required fields", http.MethodPost, "")
		return nil
	}

	if err := adaptor.GetConn().Create(&post).Error; err != nil {
		logger.Logging().Error(err)
		service.SendResponse(c, http.StatusBadRequest, err.Error(), "User Valid Values", http.MethodPost, "")
		return nil
	}
	service.SendResponse(c, http.StatusOK, "", "Post Added", http.MethodPost, "Post Created successfully")

	return nil
}

// Helps to read all the post
func ReadAllPosts(c *fiber.Ctx) error {
	var posts []models.Post

	if err := repository.ReadPosts(&posts); err != nil {
		logger.Logging().Error(err)
		service.SendResponse(c, http.StatusBadRequest, err.Error(), "Oops error occurs", http.MethodGet, "")
		return nil
	}

	for i := 0; i < len(posts); i++ {
		var count int64
		if err := repository.NumberOfComments(posts[i].PostID, &count); err != nil {
			logger.Logging().Error(err)
			service.SendResponse(c, http.StatusBadRequest, err.Error(), "Oops error occurs", http.MethodGet, "")
			return nil
		}

		posts[i].Comments = int(count)
	}

	service.SendResponse(c, http.StatusOK, "", "All available categories", http.MethodGet, posts)
	return nil
}

// Helps to update the existing post
func UpdatePost(c *fiber.Ctx) error {

	if err := helper.AdminAccess(c.Get("Authorization")); err != nil {
		logger.Logging().Error(err)
		service.SendResponse(c, http.StatusBadRequest, err.Error(), "Try Again Later", http.MethodPatch, "")
		return nil
	}

	c.Accepts("application/json")
	var post models.Post
	id := c.Params("id")
	if err := c.BodyParser(&post); err != nil {
		logger.Logging().Error(err)
		service.SendResponse(c, http.StatusInternalServerError, err.Error(), "Try Again Later", http.MethodPatch, "")
		return nil
	}

	if err := repository.UpdatePost(id, post); err != nil {
		logger.Logging().Error(err)
		service.SendResponse(c, http.StatusBadRequest, err.Error(), "Use valid id", http.MethodPatch, "")
		return nil
	}

	service.SendResponse(c, http.StatusOK, "", "Post Updated successfully", http.MethodPatch, "updated successfully")
	return nil
}

// helps to delete the post with its id
func DeletePost(c *fiber.Ctx) error {

	if err := helper.AdminAccess(c.Get("Authorization")); err != nil {
		logger.Logging().Error(err)
		service.SendResponse(c, http.StatusBadRequest, err.Error(), "Try Again Later", http.MethodPatch, "")
		return nil
	}

	id := c.Params("id")
	if err := repository.DeletePost(id); err != nil {
		logger.Logging().Error(err)
		service.SendResponse(c, http.StatusBadRequest, err.Error(), "use valid id", http.MethodDelete, "")
		return nil
	}
	service.SendResponse(c, http.StatusOK, "", "Post Deleted successfully", http.MethodDelete, "")
	return nil
}
