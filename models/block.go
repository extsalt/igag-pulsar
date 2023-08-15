package models

import (
	"gorm.io/gorm"
)

type Block struct {
	gorm.Model
	ID           uint64
	UserId       uint64
	ResourceID   uint64
	ResourceType string
}
