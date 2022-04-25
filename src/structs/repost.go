package structs

import "time"

type Respost struct {
	RepostID  uint `gorm:"id;unique"`
	PostID    uint
	Post      Post `gorm:"foreignKey:PostID"`
	CreatedAt time.Time
	Text      string
}
