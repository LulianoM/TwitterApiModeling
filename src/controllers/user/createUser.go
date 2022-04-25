package user

import (
	"apiposterr/src/database"
	"apiposterr/src/structs"
	"time"

	"errors"
	"regexp"

	"github.com/gofiber/fiber/v2"
)

func CreateUser(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	user := structs.User{
		Username:  data["username"],
		CreatedAt: time.Now(),
	}

	if err := ValidadeUsername(user); err != nil {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "invalid username",
		})
	}

	database.DB.Create(&user)

	return c.JSON(user)
}

func ValidadeUsername(user structs.User) error {
	if len(user.Username) > 14 {
		return errors.New("username bigger than 14 characters")
	}
	if !regexp.MustCompile(`^[a-zA-Z0-9]*$`).MatchString(user.Username) {
		return errors.New("username does not contain only alphanumeric characters")
	}
	return nil
}
