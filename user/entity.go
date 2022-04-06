package user

import "time"

type User struct {
	ID         int
	Name       string
	Email      string
	Password   string
	Occupation string
	Roles      string
	Avatar     string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}