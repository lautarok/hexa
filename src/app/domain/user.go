package domain

import "time"

type User struct {
	ID        string
	Role      *Role
	Name      string
	Surname   string
	CreatedAt time.Time
	UpdatedAt time.Time
}
