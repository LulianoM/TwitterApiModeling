package structs

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID        uuid.UUID `gorm:"primaryKey"`
	CreatedAt time.Time
	Username  string
	Followers []*User `gorm:"many2many:user_followers"`
	Following []*User `gorm:"many2many:user_following"`
}

type UserProfile struct {
	Username        string
	DataJoined      string
	FollowingNumber int
	PostNumber      int
	PostFeed        []Post
}
