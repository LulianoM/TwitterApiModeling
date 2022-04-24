package structs

import "time"

type Respost struct {
	RepostID    uint `gorm:"id"`
	PostID      uint
	UserID      uint
	DataCreated time.Time
	RepostText  string
}
