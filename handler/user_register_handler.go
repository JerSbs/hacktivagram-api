package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"p2-livecode-3-JerSbs/dto"
	"p2-livecode-3-JerSbs/service"
)

// RegisterUserHandler godoc
// @Summary Register a new user
// @Description Create a new user account
// @Tags Users
// @Accept json
// @Produce json
// @Param request body dto.UserRegisterRequest true "User Register Payload"
// @Success 201 {object} map[string]interface{}
// @Failure 400 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /users/register [post]
func RegisterUserHandler(c echo.Context) error {
	var input dto.UserRegisterRequest
	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Error:   true,
			Message: "invalid request body",
		})
	}

	user, err := service.RegisterUser(input)
	if err != nil {
		if err == service.ErrEmailAlreadyExists {
			return c.JSON(http.StatusBadRequest, dto.ErrorResponse{
				Error:   true,
				Message: err.Error(),
			})
		}
		return c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Error:   true,
			Message: "internal server error",
		})
	}

	return c.JSON(http.StatusCreated, echo.Map{
		"message": "user registered successfully",
		"data":    user,
	})
}
