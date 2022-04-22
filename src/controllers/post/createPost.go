package post

import (
	"apiposterr/src/database"
	"apiposterr/src/structs"
	"errors"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
)

func CreateUser(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	id, _ := strconv.Atoi(data["user_id"])

	post := structs.Post{
		UserID:      uint(id),
		DataCreated: time.Now(),
		ContentText: data["text"],
	}

	if err := ValidadePostText(post); err != nil {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "invalid text",
		})
	}

	if err := MoreThan5PostForDay(post); err != nil {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "exceeded the post limit",
		})
	}

	database.DB.Create(&post)

	return c.JSON(post)
}

func ValidadePostText(post structs.Post) error {
	if len(post.ContentText) > 777 {
		return errors.New("text bigger than 777 characters")
	}
	return nil
}

func MoreThan5PostForDay(post structs.Post) error {
	var posts []structs.Post

	database.DB.Where("user_id = ? AND data_created = ?", post.UserID, post.DataCreated).Find(&posts)

	if len(posts) > 5 {
		return errors.New("user posted the day's post limit")
	}
	return nil
}
