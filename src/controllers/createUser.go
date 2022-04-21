package controllers

import (
	"apiposterr/src/structs"

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

	return c.JSON(user)
}
