package models

import "gorm.io/gorm"

type Comment struct {
	gorm.Model
	ID            uint64
	ResourceID    uint64
	ResourceType  Resource
	Body          string
	LikesCount    uint64
	DislikesCount uint64
	OriginalImage string
	SmImage       string
	MdImage       string
	LgImage       string
	UserID        uint64
}
