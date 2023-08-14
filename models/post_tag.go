package models

import (
	"gorm.io/gorm"
)

type PostTag struct {
	gorm.Model
	PostID uint64
	TagID  uint64
}

func (PostTag) TableName() string {
	return "posts_tags"
}
