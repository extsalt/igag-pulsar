package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID            uint64
	Username      string
	Password      string
	Email         string
	Avatar        string
	Active        bool
	OauthProvider string
	Posts         []Post // User can have many posts
}
