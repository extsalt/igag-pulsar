package models

import (
	"gorm.io/gorm"
)

type Resource uint

const (
	PostResource Resource = iota + 1
)

type Action uint

const (
	LikeAction Action = iota + 1
	DislikeAction
)

type LikeDislike struct {
	gorm.Model
	ID           uint64
	UserID       uint64
	ResourceID   uint64
	ResourceType Resource
	Action       Action
}

func (LikeDislike) TableName() string {
	return "likes_dislikes"
}
