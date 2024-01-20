package handler

import (
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

// Helps to add a comment in the post
func AddComment(c *fiber.Ctx) error {

	var comment models.Comments
	if err := c.BodyParser(&comment); err != nil {
		logger.Logging().Print(err)
		service.SendResponse(c, http.StatusBadRequest, err.Error(), "Try again", http.MethodPost)
		return nil
	}

	var claimsDetails models.Claims
	err := helper.Claims(c.Get("Authorization"), &claimsDetails)
	if err != nil {
		logger.Logging().Error(err)
		service.SendResponse(c, http.StatusInternalServerError, err.Error(), "Try Again Later", http.MethodPost, "")
		return nil
	}

	comment.UsersID = claimsDetails.UsersID
	comment.PostID = c.Params("id")

	validate := validator.New()
	if err := validate.Struct(comment); err != nil {
		logger.Logging().Error(err)
		service.SendResponse(c, http.StatusBadRequest, err.Error(), "Give all required fields", http.MethodPost)
		return nil
	}

	if err := repository.CreateComment(comment); err != nil {
		logger.Logging().Print(err)
		service.SendResponse(c, http.StatusBadRequest, err.Error(), "Try again", http.MethodPost)
		return nil
	}
	service.SendResponse(c, http.StatusOK, "", "comments added", http.MethodPost, "")
	return nil
}

// Helps to Update Comment
func UpdateComment(c *fiber.Ctx) error {

	var updateComment models.Comments
	if err := c.BodyParser(&updateComment); err != nil {
		logger.Logging().Print(err)
		service.SendResponse(c, http.StatusBadRequest, err.Error(), "Try again", http.MethodPatch)
		return nil
	}

	var comment models.Comments
	if err := repository.ReadComment(c.Params("id"), &comment); err != nil {
		logger.Logging().Print(err)
		service.SendResponse(c, http.StatusBadRequest, err.Error(), "Try again", http.MethodPatch)
		return nil
	}

	var claims models.Claims
	if err := helper.Claims(c.Get("Authorization"), &claims); err != nil {
		logger.Logging().Print(err)
		service.SendResponse(c, http.StatusBadRequest, err.Error(), "Try again", http.MethodPatch)
		return nil
	}

	if comment.UsersID != claims.UsersID {
		service.SendResponse(c, http.StatusBadRequest, "Invalid authorization", "You can't update this comment", http.MethodPatch)
		return nil
	}

	if err := repository.UpdateComment(comment.CommentID, updateComment); err != nil {
		logger.Logging().Print(err)
		service.SendResponse(c, http.StatusBadRequest, err.Error(), "Try again", http.MethodPatch)
		return nil
	}

	service.SendResponse(c, http.StatusOK, "", "Updated Successfully", http.MethodPatch, "")
	return nil
}

// Helps to delete Comment
func DeleteComment(c *fiber.Ctx) error {

	var claims models.Claims
	if err := helper.Claims(c.Get("Authorization"), &claims); err != nil {
		logger.Logging().Print(err)
		service.SendResponse(c, http.StatusBadRequest, err.Error(), "Try again", http.MethodDelete)
		return nil
	}

	var comment models.Comments
	if err := repository.ReadComment(c.Params("id"), &comment); err != nil {
		logger.Logging().Print(err)
		service.SendResponse(c, http.StatusBadRequest, err.Error(), "Try again", http.MethodDelete)
		return nil
	}

	if comment.UsersID == claims.UsersID || claims.UsersID == "ADMIN1" {
		if err := repository.DeleteComment(comment.CommentID); err != nil {
			logger.Logging().Print(err)
			service.SendResponse(c, http.StatusBadRequest, err.Error(), "Try again", http.MethodDelete)
			return nil
		}
		service.SendResponse(c, http.StatusOK, "", "Comment Deleted Successfully", http.MethodDelete, "")
		return nil
	}
	service.SendResponse(c, http.StatusBadRequest, "Invalid authorization", "You can't delete this comment", http.MethodPatch)
	return nil

}

// Helps to Read all comments posted by single user
func ReadCommentByUser(c *fiber.Ctx) error {

	var comments []models.Comments

	if err := repository.ReadCommentByUser(c.Params("id"), &comments); err != nil {
		logger.Logging().Print(err)
		service.SendResponse(c, http.StatusBadRequest, err.Error(), "Try again", http.MethodGet)
		return nil
	}

	service.SendResponse(c, http.StatusOK, "", fmt.Sprintf("Fetching all comments of %v", c.Params("id")), http.MethodGet, comments)
	return nil
}

// Helps to read all the comments in the particular post
func ReadCommentByPost(c *fiber.Ctx) error {
	var comments []models.Comments

	if err := repository.PostComments(c.Params("id"), &comments); err != nil {
		logger.Logging().Print(err)
		service.SendResponse(c, http.StatusBadRequest, err.Error(), "Try again", http.MethodGet)
		return nil
	}

	service.SendResponse(c, http.StatusOK, "", fmt.Sprintf("Fetching all comments of %v Post", c.Params("id")), http.MethodGet, comments)
	return nil
}

// Helps to read the comment by its id
func ReadCommentByID(c *fiber.Ctx) error {
	var comment models.Comments

	if err := repository.ReadComment(c.Params("id"), &comment); err != nil {
		logger.Logging().Print(err)
		service.SendResponse(c, http.StatusBadRequest, err.Error(), "Try again", http.MethodGet)
		return nil
	}
	service.SendResponse(c, http.StatusOK, "", fmt.Sprintf("Fetching all comment of %v ", c.Params("id")), http.MethodGet, comment)
	return nil
}
