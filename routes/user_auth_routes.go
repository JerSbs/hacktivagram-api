package routes

import (
	"p2-livecode-3-JerSbs/handler"

	"github.com/labstack/echo/v4"
)

func RegisterAuthRoutes(e *echo.Echo) {
	authGroup := e.Group("/users")
	authGroup.POST("/login", handler.LoginUserHandler)
}
