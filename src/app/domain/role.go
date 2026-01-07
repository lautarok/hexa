package domain

import "time"

type Role struct {
	ID          string
	NameFr      string
	NameEs      string
	NameEn      string
	NamePt      string
	Permissions []*Permission
	Users       []*User
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
