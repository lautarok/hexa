package domain

import "time"

type Role struct {
	ID          string        `validate:"uuid"`
	NameFr      string        `validate:"min=2,max=30"`
	NameEs      string        `validate:"min=2,max=30"`
	NameEn      string        `validate:"min=2,max=30"`
	NamePt      string        `validate:"min=2,max=30"`
	Permissions []*Permission `validate:"required,min=1"`
	Users       []*User
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
