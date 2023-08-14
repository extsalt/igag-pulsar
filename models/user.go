package models

import (
	"time"
)

type User struct {
	ID        uint
	Username  string
	Password  string
	Email     string
	Active    bool
	CreatedAt time.Time
}
