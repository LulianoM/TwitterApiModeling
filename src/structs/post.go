package structs

import (
	"time"
)

type Post struct {
	UserID      uint `gorm:"user_id"`
	PostID      uint `gorm:"id"`
	DataCreated time.Time
	ContentText string
}
