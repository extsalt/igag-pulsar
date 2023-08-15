package models

import (
	"gorm.io/gorm"
)

type Tag struct {
	gorm.Model
	ID  uint64
	Tag string
}
