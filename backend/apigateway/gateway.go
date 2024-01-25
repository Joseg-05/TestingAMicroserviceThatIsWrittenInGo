package main

import (
	"apigateway/api/handler"
	"apigateway/api/route"

	"apigateway/internal/config"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {

	app := fiber.New()
	// Initialize default config
	_, _, hostIP, _, _, aMicro := config.Setup()
	app.Use(cors.New())
	// jwt := middleware.NewAuthMiddleware(secret)
	// Or extend your config for customization
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "*",
		AllowMethods: "GET, POST, PUT, DELETE",
	}))
	authHandler := handler.InitializeAuth("http://" + hostIP + ":" + aMicro + "/api/v1/auth")
	// authHandler := handler.InitializeAuth("http://host.docker.internal:8002/api/v1/auth")
	route.SetupAuth(app, authHandler)
	app.Listen(":" + "8080")
}
