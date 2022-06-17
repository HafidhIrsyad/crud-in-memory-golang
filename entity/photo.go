package entity

import "time"

type Photo struct {
	ID        uint
	UserID    int
	Title     string
	Caption   string
	PhotoUrl  string
	CreatedAt time.Time
	UpdatedAt time.Time
}
