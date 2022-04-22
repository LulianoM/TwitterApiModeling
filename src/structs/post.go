package structs

import (
	"time"
)

type Post struct {
	UserID      uint `gorm:"userid"`
	PostID      uint `gorm:"postid"`
	DataCreated time.Time
	ContentText string
}
