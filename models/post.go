package models

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	ID            uint64
	Title         string
	Body          string
	Slug          string
	OriginalImage string
	SmImage       string
	MdImage       string
	LgImage       string
	UserID        uint64
	Tags          []Tag `gorm:"many2many:posts_tags"`
}
