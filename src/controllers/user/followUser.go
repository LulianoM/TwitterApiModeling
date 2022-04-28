package user

import (
	"apiposterr/src/database"
	"apiposterr/src/structs"
	"errors"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func FollowUser(c *fiber.Ctx) error {

	var data map[string]string
	if err := c.BodyParser(&data); err != nil {
		return err
	}

	id, err := uuid.Parse(data["user_id"])
	if err != nil {
		return err
	}

	user_id_follow, err := uuid.Parse(data["user_id_follow"])
	if err != nil {
		return err
	}

	var user structs.User
	var user_follow structs.User

	user.ID = id
	user_follow.ID = user_id_follow

	ValidadeSameUser(id, user_id_follow)

	errorFind := database.DB.First(&user, id)
	if errorFind.Error != nil {
		return errorFind.Error
	}

	errorFind = database.DB.First(&user_follow, user_id_follow)
	if errorFind.Error != nil {
		return errorFind.Error
	}

	is_following, err := strconv.ParseBool(data["following"])
	if err != nil {
		return err
	}

	if is_following {
		user.Following = append(user.Following, &user_follow)
		database.DB.Session(&gorm.Session{FullSaveAssociations: true}).Model(&user).Association("Following").Append(&user_follow)
	} else {
		database.DB.Session(&gorm.Session{FullSaveAssociations: true}).Model(&user).Association("Followers").Append(&user_follow)
	}

	return c.JSON(user)
}

func ValidadeSameUser(user uuid.UUID, user_id_follow uuid.UUID) error {
	if user == user_id_follow {
		return errors.New("you cant follow yourself")
	}
	return nil
}
