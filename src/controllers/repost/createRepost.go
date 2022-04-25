package repost

import (
	"apiposterr/src/database"
	"apiposterr/src/structs"
	"errors"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
)

func CreateRepost(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	repost := structs.Respost{
		PostID:    Strtouint(data["post_id"]),
		CreatedAt: time.Now(),
		Text:      data["text"],
	}

	if err := ValidadePostText(repost); err != nil {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "invalid text",
		})
	}

	if err := MoreThan5PostForDay(repost); err != nil {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "exceeded the post limit",
		})
	}

	database.DB.Create(&repost)

	return c.JSON(repost)
}

func ValidadePostText(repost structs.Respost) error {
	if len(repost.Text) > 777 {
		return errors.New("text bigger than 777 characters")
	}
	return nil
}

func MoreThan5PostForDay(repost structs.Respost) error {
	//var posts []structs.Respost

	//database.DB.Where("user_id = ? AND data_created = ?", repost.UserID, repost.DataCreated).Find(&posts)

	//if len(posts) > 5 {
	//	return errors.New("user posted the day's post limit")
	//}
	return nil
}

func Strtouint(data string) uint {
	id, _ := strconv.Atoi(data)
	return uint(id)
}
