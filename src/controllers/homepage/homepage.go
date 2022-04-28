package homepage

import (
	"apiposterr/src/database"
	"apiposterr/src/structs"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func ShowHomepage(c *fiber.Ctx) error {

	type QueryParsers struct {
		onlyFollowing bool   `query:"only_following"`
		id            string `query:"id"`
	}

	q := new(QueryParsers)
	if err := c.QueryParser(q); err != nil {
		return err
	}

	var homepageView []structs.Post

	if q.onlyFollowing {
		homepageView = HomepageViewJustFollowing(uuid.MustParse(q.id))
	} else {
		homepageView = HomepageViewAll()
	}

	return c.JSON(homepageView)
}

func HomepageViewAll() []structs.Post {

	var homepage []structs.Post

	database.DB.Limit(10).Order("created_at").Find(&homepage)

	return homepage
}

func HomepageViewJustFollowing(UserID uuid.UUID) []structs.Post {
	var homepage []structs.Post

	database.DB.Limit(10).Order("created_at").Find(&homepage)

	return homepage
}
