package routes

import (
	"apiposterr/src/controllers/healtcheck"
	"apiposterr/src/controllers/post"
	"apiposterr/src/controllers/user"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	api := app.Group("api")

	api.Get("ping", healtcheck.Ping)

	api.Post("create", user.CreateUser)
	api.Get("user", user.GetUser)
	api.Get("user/:id", user.GetUserByID)

	api.Post("post", post.CreateUser)
	api.Get("post", post.GetPosts)
}
