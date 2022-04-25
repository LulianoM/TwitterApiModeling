package structs

import "time"

type Post struct {
	ID        uint `gorm:"unique"`
	User      User `gorm:"foreignKey:ID"`
	PostID    uint `gorm:"unique"`
	CreatedAt time.Time
	Text      string
}
