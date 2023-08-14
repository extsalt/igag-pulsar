package models

import (
	"time"
)

type User struct {
	ID        uint64
	Username  string
	Password  string
	Email     string
	Active    bool
	CreatedAt time.Time
}
