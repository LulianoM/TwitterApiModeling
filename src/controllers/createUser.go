package controllers

import (
	"apiposterr/src/database"
	"apiposterr/src/structs"

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
		Username: data["username"],
	}

	if err := validadeUsername(user); err != nil {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "invalid username",
		})
	}

	database.DB.Create(&user)

	return c.JSON(user)
}

func validadeUsername(user structs.User) error {
	if len(user.Username) > 14 {
		return errors.New("username bigger than 14 characters")
	}
	if regexp.MustCompile(`^[a-zA-Z0-9]*$`).MatchString(user.Username) != true {
		return errors.New("username does not contain only alphanumeric characters")
	}
	return nil
}
