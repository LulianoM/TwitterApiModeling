package routes

import (
	"apiposterr/src/controllers/healtcheck"
	"apiposterr/src/controllers/homepage"
	"apiposterr/src/controllers/post"
	"apiposterr/src/controllers/user"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	api := app.Group("api")
	admin := api.Group("admin")

	api.Get("ping", healtcheck.Ping)

	api.Post("create", user.CreateUser)
	api.Post("/follow", user.FollowUser)
	admin.Get("user", user.GetUser)
	admin.Get("user/:id", user.GetUserByID)

	api.Post("post", post.CreatePost)
	admin.Get("post", post.GetPosts)
	admin.Get("post/:id", post.GetPostByID)

	api.Get("homepage", homepage.ShowHomepage)
}
