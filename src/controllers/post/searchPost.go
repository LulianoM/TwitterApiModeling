package post

import (
	"apiposterr/src/database"
	"apiposterr/src/structs"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func GetPosts(c *fiber.Ctx) error {
	var posts []structs.Post
	database.DB.Find(&posts)
	return c.JSON(posts)
}

func GetPostByPostID(c *fiber.Ctx) error {

	var post structs.Post

	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return err
	}
	post.PostID = id

	database.DB.Find(&post)

	return c.JSON(post)
}

func GetFeedPostByUserID(UserID uuid.UUID) []structs.Post {

	var posts []structs.Post

	database.DB.Limit(5).Order("created_at").Find(&posts, "ID=?", UserID)

	return posts
}
