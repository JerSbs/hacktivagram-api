package handler

import (
	"net/http"
	"p2-livecode-3-JerSbs/dto"
	"p2-livecode-3-JerSbs/service"

	"github.com/labstack/echo/v4"
)

// CreatePostHandler godoc
// @Summary Create a new post
// @Description Create a new post by the logged-in user. If content is empty, a joke will be auto-filled.
// @Tags Posts
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param request body dto.CreatePostRequest true "Create Post Payload"
// @Success 201 {object} map[string]interface{}
// @Failure 400 {object} dto.ErrorResponse
// @Failure 401 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /posts [post]
func CreatePostHandler(c echo.Context) error {
	var input dto.CreatePostRequest

	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Error:   true,
			Message: "invalid input",
		})
	}

	if err := c.Validate(&input); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Error:   true,
			Message: err.Error(),
		})
	}

	userID := c.Get("user_id").(uint)

	post, err := service.CreatePostService(input, userID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Error:   true,
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, echo.Map{
		"message": "post created successfully",
		"data":    post,
	})
}
