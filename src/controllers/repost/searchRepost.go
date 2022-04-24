package repost

import (
	"apiposterr/src/database"
	"apiposterr/src/structs"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func GetReposts(c *fiber.Ctx) error {
	var reposts []structs.Respost
	database.DB.Find(&reposts)
	return c.JSON(reposts)
}

func GetRepostByID(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	var repost structs.Respost

	repost.RepostID = uint(id)

	database.DB.Find(&repost)

	return c.JSON(repost)
}
