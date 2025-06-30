package handler

import (
	"net/http"

	"p2-livecode-3-JerSbs/dto"
	"p2-livecode-3-JerSbs/service"

	"github.com/labstack/echo/v4"
)

// LoginUserHandler godoc
// @Summary Login user
// @Description Authenticate user and return JWT token
// @Tags Users
// @Accept json
// @Produce json
// @Param request body dto.UserLoginRequest true "User Login Payload"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} dto.ErrorResponse
// @Failure 401 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /users/login [post]
func LoginUserHandler(c echo.Context) error {
	var input dto.UserLoginRequest
	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Error:   true,
			Message: "invalid request body",
		})
	}

	token, err := service.LoginUser(input)
	if err != nil {
		switch err {
		case service.ErrInvalidInput:
			return c.JSON(http.StatusBadRequest, dto.ErrorResponse{
				Error:   true,
				Message: err.Error(),
			})
		case service.ErrUserNotFound, service.ErrWrongPassword:
			return c.JSON(http.StatusUnauthorized, dto.ErrorResponse{
				Error:   true,
				Message: err.Error(),
			})
		default:
			return c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
				Error:   true,
				Message: "internal server error",
			})
		}
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "login successful",
		"token":   token,
	})
}
