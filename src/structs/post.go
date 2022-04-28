package structs

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	ID uuid.UUID
	//User       *User `gorm:"foreignKey:ID"`
	PostID     uuid.UUID
	TextPost   string
	TextRepost string
	IsRepost   bool
}
