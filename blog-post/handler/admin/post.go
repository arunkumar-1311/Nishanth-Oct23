package admin

import (
	"blog_post/adaptor"
	"blog_post/logger"
	"blog_post/models"
	"blog_post/repository"
	"blog_post/service"
	"blog_post/service/helper"
	"fmt"
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
	var allPosts models.AllPost

	if err := repository.ReadPosts(&allPosts.Post); err != nil {
		logger.Logging().Error(err)
		service.SendResponse(c, http.StatusBadRequest, err.Error(), "Oops error occurs", http.MethodGet, "")
		return nil
	}

	category := make(map[string]int)
	archieve := make(map[string]int)

	for i := 0; i < len(allPosts.Post); i++ {
		var count int64

		if err := repository.NumberOfComments(allPosts.Post[i].PostID, &count); err != nil {
			logger.Logging().Error(err)
			service.SendResponse(c, http.StatusBadRequest, err.Error(), "Oops error occurs", http.MethodGet, "")
			return nil
		}
		allPosts.Post[i].Comments = int(count)

		year, month, _ := allPosts.Post[i].CreatedAt.Date()
		archieve[fmt.Sprint(month.String(), "-", year)] = 1

		for key, value := range allPosts.Post[i].CategoryID {
			name, err := repository.ReadCategoryByID(value)
			if err != nil {
				logger.Logging().Error(err)
				service.SendResponse(c, http.StatusInternalServerError, err.Error(), "Oops error occurs", http.MethodGet, "")
				return nil
			}
			category[name] = category[name] + 1
			allPosts.Post[i].CategoryID[key] = name
		}

		if err := repository.PostComments(allPosts.Post[i].PostID, &allPosts.Post[i].PostComments); err != nil {
			logger.Logging().Error(err)
			service.SendResponse(c, http.StatusInternalServerError, err.Error(), "Oops error occurs", http.MethodGet, "")
			return nil
		}
	}

	for key := range category {
		data := models.CategoriesCount{CategoryName: key, Total: category[key]}
		allPosts.CategoriesCount = append(allPosts.CategoriesCount, data)
	}

	for key := range archieve {
		allPosts.Archieves = append(allPosts.Archieves, key)
	}
	service.SendResponse(c, http.StatusOK, "", "All available categories", http.MethodGet, allPosts)
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
