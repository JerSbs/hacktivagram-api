package handler

import (
	"net/http"
	"p2-livecode-3-JerSbs/dto"
	"p2-livecode-3-JerSbs/service"

	"github.com/labstack/echo/v4"
)

// GetAllPostsHandler godoc
// @Summary Get all posts
// @Description Retrieve all posts from all users
// @Tags Posts
// @Security BearerAuth
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Failure 401 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /posts [get]
func GetAllPostsHandler(c echo.Context) error {
	posts, err := service.GetAllPostsService()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Error:   true,
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"data": posts,
	})
}
