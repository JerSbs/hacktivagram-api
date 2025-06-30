package routes

import (
	"p2-livecode-3-JerSbs/handler"
	"p2-livecode-3-JerSbs/middleware"

	"github.com/labstack/echo/v4"
)

func RegisterPostRoutes(e *echo.Echo) {
	posts := e.Group("/posts")
	posts.Use(middleware.JWTMiddleware)
	posts.POST("", handler.CreatePostHandler)
	posts.GET("", handler.GetAllPostsHandler)
	posts.GET("/:id", handler.GetPostByIDHandler)
	posts.DELETE("/:id", handler.DeletePostHandler)
}
