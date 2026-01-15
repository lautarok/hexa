package dtos

import (
	"time"

	"github.com/google/uuid"
)

type PermissionOutputDto struct {
	ID        uuid.UUID `json:"id"`
	Alias     string    `json:"alias"`
	CreatedAt time.Time `json:"createdAt"`
}
