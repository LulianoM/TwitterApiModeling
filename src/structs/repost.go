package structs

import "time"

type Respost struct {
	RepostID      uint `gorm:"id"`
	PostID        uint
	DataCreated   time.Time
	ContentText   string
	PostReference []Post
}
