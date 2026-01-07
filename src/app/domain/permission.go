package domain

import "time"

type Permission struct {
	ID        string
	Alias     string
	Roles     []*Role
	CreatedAt time.Time
	UpdatedAt time.Time
}
