package user

import (
	"apiposterr/src/database"
	"apiposterr/src/structs"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func GetUser(c *fiber.Ctx) error {
	var users []structs.User
	database.DB.Find(&users)
	return c.JSON(users)
}

func GetUserByID(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	var user structs.User

	user.ID = uint(id)

	database.DB.Find(&user)

	return c.JSON(user)
}
