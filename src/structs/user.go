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
