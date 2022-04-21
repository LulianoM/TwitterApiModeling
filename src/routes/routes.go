package routes

import (
	"apiposterr/src/controllers"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	api := app.Group("api")

	api.Get("ping", controllers.Ping)
}
