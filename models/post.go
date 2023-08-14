package models

import "time"

type Post struct {
	ID        uint64
	Title     string
	Body      string
	UserID    uint64
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}
