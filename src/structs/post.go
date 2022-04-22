package structs

import (
	"time"
)

type Post struct {
	UserID      uint
	PostID      uint `gorm:"id"`
	DataCreated time.Time
	ContentText string
}
