package routes

import (
	"github.com/labstack/echo/v4"
	"p2-livecode-3-JerSbs/handler"
)

func RegisterUserRoutes(e *echo.Echo) {
	userGroup := e.Group("/users")
	userGroup.POST("/register", handler.RegisterUserHandler)
}
