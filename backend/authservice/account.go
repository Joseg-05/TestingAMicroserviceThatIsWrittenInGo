package main

import (
	"fmt"

	"authservice/api/dao"
	handler "authservice/api/handler"
	"authservice/api/route"
	"authservice/api/service"
	"authservice/internal/config"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	// port := config.SetupEnvironment()
	db := config.SetupDB()

	defer func() {
		dbInstance, _ := db.DB()
		_ = dbInstance.Close()
	}()

	app := fiber.New()
	// Initialize default config
	app.Use(cors.New())

	// Or extend your config for customization
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "*",
		AllowMethods: "GET, POST, PUT, DELETE",
	}))
	fmt.Println("PAPAPAPPAPA")

	dao := dao.Initialize(db)

	service := service.Initialize(dao)
	handler := handler.Initialize(service)
	route.Setup(app, handler)
	app.Listen("0.0.0.0:" + "8080")
	// app.Listen(":8000")
}
