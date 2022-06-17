package entity

import "time"

type Comment struct {
	ID        uint
	UserID    int
	PhotoID   int
	Message   string
	CreatedAt time.Time
	UpdatedAt time.Time
}
