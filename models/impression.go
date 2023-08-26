package models

import (
	"gorm.io/gorm"
)

type Impression struct {
	gorm.Model
	ID     uint64
	postID uint64
}
