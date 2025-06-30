package handler

import (
	"net/http"
	"p2-livecode-3-JerSbs/dto"
	"p2-livecode-3-JerSbs/service"
	"strconv"

	"github.com/labstack/echo/v4"
)

// DeletePostHandler godoc
// @Summary Delete a post
// @Description Delete a post by ID, only if the user is the owner
// @Tags Posts
// @Security BearerAuth
// @Produce json
// @Param id path int true "Post ID"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} dto.ErrorResponse
// @Failure 401 {object} dto.ErrorResponse
// @Failure 403 {object} dto.ErrorResponse
// @Failure 404 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /posts/{id} [delete]
func DeletePostHandler(c echo.Context) error {
	idParam := c.Param("id")
	postID, err := strconv.Atoi(idParam)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Error:   true,
			Message: "invalid post ID",
		})
	}

	userID := c.Get("user_id").(uint)

	post, err := service.DeletePostService(uint(postID), userID)
	if err != nil {
		switch err {
		case service.ErrNotFound:
			return c.JSON(http.StatusNotFound, dto.ErrorResponse{Error: true, Message: "post not found"})
		case service.ErrForbidden:
			return c.JSON(http.StatusForbidden, dto.ErrorResponse{Error: true, Message: "you are not allowed to delete this post"})
		default:
			return c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: true, Message: err.Error()})
		}
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "post deleted successfully",
		"data":    post,
	})
}
