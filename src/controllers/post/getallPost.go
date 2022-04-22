package post

import (
	"apiposterr/src/database"
	"apiposterr/src/structs"

	"github.com/gofiber/fiber/v2"
)

func GetPosts(c *fiber.Ctx) error {
	var posts []structs.Post
	database.DB.Find(&posts)
	return c.JSON(posts)
}
