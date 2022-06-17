package entity

import "time"

type User struct {
	ID        uint
	Username  string
	Email     string
	Password  string
	Age       int
	CreatedAt time.Time
	UpdatedAt time.Time
}
