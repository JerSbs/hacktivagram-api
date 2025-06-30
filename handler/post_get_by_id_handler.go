package handler

import (
	"net/http"
	"p2-livecode-3-JerSbs/dto"
	"p2-livecode-3-JerSbs/service"
	"strconv"

	"github.com/labstack/echo/v4"
)

// GetPostByIDHandler godoc
// @Summary Get post by ID
// @Description Retrieve a post with its comments and commenter info by post ID
// @Tags Posts
// @Security BearerAuth
// @Produce json
// @Param id path int true "Post ID"
// @Success 200 {object} dto.PostDetailResponse
// @Failure 400 {object} dto.ErrorResponse
// @Failure 401 {object} dto.ErrorResponse
// @Failure 404 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /posts/{id} [get]
func GetPostByIDHandler(c echo.Context) error {
	idParam := c.Param("id")
	postID, err := strconv.Atoi(idParam)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Error:   true,
			Message: "invalid post ID",
		})
	}

	result, err := service.GetPostByIDService(uint(postID))
	if err != nil {
		if err == service.ErrNotFound {
			return c.JSON(http.StatusNotFound, dto.ErrorResponse{
				Error:   true,
				Message: "post not found",
			})
		}
		return c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Error:   true,
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, result)
}
