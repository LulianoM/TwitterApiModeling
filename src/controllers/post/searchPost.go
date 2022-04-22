package post

import (
	"apiposterr/src/database"
	"apiposterr/src/structs"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func GetPosts(c *fiber.Ctx) error {
	var posts []structs.Post
	database.DB.Find(&posts)
	return c.JSON(posts)
}

func GetPostByID(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	var post structs.Post

	post.PostID = uint(id)

	database.DB.Find(&post)

	return c.JSON(post)
}
