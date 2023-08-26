package models

import "gorm.io/gorm"

type LikeDislike struct {
	gorm.Model
	ID           uint64
	UserID       uint64
	ResourceID   uint64
	ResourceType string
	Action       uint8
}
