package main

import (
	"log"
	"os"

	"p2-livecode-3-JerSbs/config"
	"p2-livecode-3-JerSbs/routes"

	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

// @title Hacktivagram API Docs
// @version 1.0
// @description RESTful API for Hacktivagram using Golang Echo Framework
// @host localhost:8080
// @BasePath /
// @schemes http https

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
	// Init DB
	config.InitDB()

	// Init Echo
	e := echo.New()

	// Swagger docs endpoint
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	// Register all routes
	routes.RegisterUserRoutes(e)
	routes.RegisterPostRoutes(e)

	// Run server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	if err := e.Start(":" + port); err != nil {
		log.Fatal(err)
	}
}
