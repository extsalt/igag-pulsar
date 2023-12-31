package models

import "gorm.io/gorm"

type Reply struct {
	gorm.Model
	ID            uint64
	CommentID     uint64
	Body          string
	OriginalImage string
	LikeCount     uint64
	DislikeCount  uint64
	SmImage       string
	MdImage       string
	LgImage       string
	UserID        uint64
}
