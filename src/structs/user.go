package structs

import (
	"time"
)

type User struct {
	Username   string `json:"username"`
	DataJoined time.Time
	ID         uint `gorm:"id"`
}
