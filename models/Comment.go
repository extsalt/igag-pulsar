package models

import "gorm.io/gorm"

type Comment struct {
	gorm.Model
	ID            uint64
	PostID        uint64
	Body          string
	OriginalImage string
	SmImage       string
	MdImage       string
	LgImage       string
	UserID        uint64
}
