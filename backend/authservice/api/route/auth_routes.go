package route

import (
	"authservice/api/handler"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App, authHandler *handler.AuthHandler) {
	apiRoutes := app.Group("api/v1/auth")

	apiRoutes.Post("/register", authHandler.CreateAuth)
	apiRoutes.Post("/login", authHandler.Login)
	fmt.Println("HEEEELLOO")
	// grab the userId from expired
	apiRoutes.Post("/refresh", authHandler.Refresh)
}
