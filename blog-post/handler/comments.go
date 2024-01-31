package handler

import (
	"blog_post/logger"
	"blog_post/models"
	"blog_post/service"
	"blog_post/service/helper"
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

// Helps to add a comment in the post
func (h *Handler) AddComment(c *fiber.Ctx) error {

	var comment models.Comments
	if err := c.BodyParser(&comment); err != nil {
		logger.Logging().Print(err)
		service.SendResponse(c, fiber.StatusBadRequest, err.Error(), "Try again", fiber.MethodPost)
		return nil
	}

	var claimsDetails models.Claims
	err := helper.Claims(c.Get("Authorization"), &claimsDetails)
	if err != nil {
		logger.Logging().Error(err)
		service.SendResponse(c, fiber.StatusInternalServerError, err.Error(), "Try Again Later", fiber.MethodPost, "")
		return nil
	}

	comment.UsersID = claimsDetails.UsersID
	comment.PostID = c.Params("id")
	comment.CommentID = helper.UniqueID()
	validate := validator.New()
	if err := validate.Struct(comment); err != nil {
		logger.Logging().Error(err)
		service.SendResponse(c, fiber.StatusBadRequest, err.Error(), "Give all required fields", fiber.MethodPost)
		return nil
	}

	if err := h.Method.CreateComment(comment); err != nil {
		logger.Logging().Print(err)
		service.SendResponse(c, fiber.StatusBadRequest, err.Error(), "Try again", fiber.MethodPost)
		return nil
	}
	service.SendResponse(c, fiber.StatusOK, "", "comments added", fiber.MethodPost, comment)
	return nil
}

// Helps to Update Comment
func (h *Handler) UpdateComment(c *fiber.Ctx) error {

	var updateComment models.Comments
	if err := c.BodyParser(&updateComment); err != nil {
		logger.Logging().Print(err)
		service.SendResponse(c, fiber.StatusBadRequest, err.Error(), "Try again", fiber.MethodPatch)
		return nil
	}

	var comment models.Comments
	if err := h.Method.ReadComment(c.Params("id"), &comment); err != nil {
		logger.Logging().Print(err)
		service.SendResponse(c, fiber.StatusBadRequest, err.Error(), "Try again", fiber.MethodPatch)
		return nil
	}

	var claims models.Claims
	if err := helper.Claims(c.Get("Authorization"), &claims); err != nil {
		logger.Logging().Print(err)
		service.SendResponse(c, fiber.StatusBadRequest, err.Error(), "Try again", fiber.MethodPatch)
		return nil
	}

	if comment.UsersID != claims.UsersID {
		service.SendResponse(c, fiber.StatusBadRequest, "Invalid authorization", "You can't update this comment", fiber.MethodPatch)
		return nil
	}

	if err := h.Method.UpdateComment(comment.CommentID, updateComment); err != nil {
		logger.Logging().Print(err)
		service.SendResponse(c, fiber.StatusBadRequest, err.Error(), "Try again", fiber.MethodPatch)
		return nil
	}

	service.SendResponse(c, fiber.StatusOK, "", "Updated Successfully", fiber.MethodPatch, "")
	return nil
}

// Helps to delete Comment
func (h *Handler) DeleteComment(c *fiber.Ctx) error {

	var claims models.Claims
	if err := helper.Claims(c.Get("Authorization"), &claims); err != nil {
		logger.Logging().Print(err)
		service.SendResponse(c, fiber.StatusBadRequest, err.Error(), "Try again", fiber.MethodDelete)
		return nil
	}

	var comment models.Comments
	if err := h.Method.ReadComment(c.Params("id"), &comment); err != nil {
		logger.Logging().Print(err)
		service.SendResponse(c, fiber.StatusBadRequest, err.Error(), "Try again", fiber.MethodDelete)
		return nil
	}

	if comment.UsersID == claims.UsersID || claims.RolesID == "AD1" {
		if err := h.Method.DeleteComment(comment.CommentID); err != nil {
			logger.Logging().Print(err)
			service.SendResponse(c, fiber.StatusBadRequest, err.Error(), "Try again", fiber.MethodDelete)
			return nil
		}
		service.SendResponse(c, fiber.StatusOK, "", "Comment Deleted Successfully", fiber.MethodDelete, "")
		return nil
	}
	service.SendResponse(c, fiber.StatusBadRequest, "Invalid authorization", "You can't delete this comment", fiber.MethodDelete)
	return nil

}

// Helps to Read all comments posted by single user
func (h *Handler) ReadCommentByUser(c *fiber.Ctx) error {

	var comments []models.Comments
	var commentsResp []models.CommentsResponse
	if err := h.Method.ReadCommentByUser(c.Params("id"), &comments); err != nil {
		logger.Logging().Print(err)
		service.SendResponse(c, fiber.StatusBadRequest, err.Error(), "Try again", fiber.MethodGet)
		return nil
	}

	for _, value := range comments {
		comment := models.CommentsResponse{
			CreatedAt: value.CreatedAt, CommentID: value.CommentID, Content: value.Content, Website: value.Website,
			UserID: value.Users.UserID, Email: value.Users.Email, Name: value.Users.Name,
		}
		commentsResp = append(commentsResp, comment)
	}
	service.SendResponse(c, fiber.StatusOK, "", fmt.Sprintf("Fetching all comments of %v", c.Params("id")), fiber.MethodGet, commentsResp)
	return nil
}

// Helps to read all the comments in the particular post
func (h *Handler) ReadCommentByPost(c *fiber.Ctx) error {
	var comments []models.Comments
	var commentsResp []models.CommentsResponse
	if err := h.Method.PostComments(c.Params("id"), &comments); err != nil {
		logger.Logging().Print(err)
		service.SendResponse(c, fiber.StatusBadRequest, err.Error(), "Try again", fiber.MethodGet)
		return nil
	}

	for _, value := range comments {
		comment := models.CommentsResponse{
			CreatedAt: value.CreatedAt, CommentID: value.CommentID, Content: value.Content, Website: value.Website,
			UserID: value.Users.UserID, Email: value.Users.Email, Name: value.Users.Name,
		}
		commentsResp = append(commentsResp, comment)
	}
	service.SendResponse(c, fiber.StatusOK, "", fmt.Sprintf("Fetching all comments of %v Post", c.Params("id")), fiber.MethodGet, commentsResp)
	return nil
}

// Helps to read the comment by its id
func (h *Handler) ReadCommentByID(c *fiber.Ctx) error {
	var comment models.Comments

	if err := h.Method.ReadComment(c.Params("id"), &comment); err != nil {
		logger.Logging().Print(err)
		service.SendResponse(c, fiber.StatusBadRequest, err.Error(), "Try again", fiber.MethodGet)
		return nil
	}
	service.SendResponse(c, fiber.StatusOK, "", fmt.Sprintf("Fetching all comment of %v ", c.Params("id")), fiber.MethodGet, comment)
	return nil
}
