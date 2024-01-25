package route

import (
	"apigateway/api/handler"

	"github.com/gofiber/fiber/v2"
)

func SetupAuth(app *fiber.App, handler *handler.AuthHandler) {
	apiRoutes := app.Group("api/v1/auth")
	apiRoutes.Post("/register", handler.RegisterHandler)
	apiRoutes.Post("/login", handler.LoginHandler)
	apiRoutes.Post("/refresh", handler.RefreshHandler)
}
