package user

import (
	"apiposterr/src/controllers/post"
	"apiposterr/src/structs"
	"time"

	"github.com/gofiber/fiber/v2"
)

func UserProfile(c *fiber.Ctx) error {

	var (
		user  structs.User
		posts []structs.Post
	)

	user = SearchUserByUsername(c)

	posts = post.GetFeedPostByUserID(user.ID)

	userProfile := structs.UserProfile{
		Username:        user.Username,
		DataJoined:      user.CreatedAt.Format(time.ANSIC),
		FollowingNumber: len(user.Following),
		PostNumber:      len(posts),
		PostFeed:        posts,
	}

	return c.JSON(userProfile)
}
