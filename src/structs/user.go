package structs

import (
	"time"
)

type User struct {
	ID        uint `gorm:"primaryKey;unique"`
	CreatedAt time.Time
	Username  string
}

type Followers struct {
	ID             uint
	FriendID       []uint
	FriendUsername []string
	User           User `gorm:"foreignKey:ID"`
}

type Following struct {
	ID             uint
	FriendID       []uint
	FriendUsername []string
	User           User `gorm:"foreignKey:ID"`
}
