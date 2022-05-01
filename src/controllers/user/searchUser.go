package user

import (
	"apiposterr/src/database"
	"apiposterr/src/structs"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func GetUser(c *fiber.Ctx) error {
	var user []structs.User
	database.DB.Find(&user)
	return c.JSON(user)
}

func GetUserByID(c *fiber.Ctx) error {
	var user structs.User

	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return err
	}

	user.ID = id

	database.DB.Find(&user)

	return c.JSON(user)
}

func SearchUserByUsername(c *fiber.Ctx) structs.User {
	var user structs.User

	username := c.Params("username")

	database.DB.Find(&user, "username = ?", username)

	return user
}
