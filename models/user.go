package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID       uint64
	Username string
	Password string
	Email    string
	Active   bool
	Posts    []Post // User can have many posts
}
