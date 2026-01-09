package domain

import (
	"time"

	"github.com/google/uuid"
)

type Permission struct {
	ID        uuid.UUID `validation:"uuid"`
	Alias     string    `validation:"min=1,max=30"`
	Roles     []*Role
	CreatedAt time.Time
	UpdatedAt time.Time
}
