package domain

import "time"

type Permission struct {
	ID        string `validate:"uuid"`
	Alias     string `validate:"min=1,max=30"`
	Roles     []*Role
	CreatedAt time.Time
	UpdatedAt time.Time
}
