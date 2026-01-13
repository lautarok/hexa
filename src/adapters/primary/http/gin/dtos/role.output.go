package dtos

import (
	"time"

	"github.com/google/uuid"
)

type RoleOutputDto struct {
	ID        uuid.UUID `json:"id"`
	NameEn    string    `json:"nameEn"`
	NameEs    string    `json:"nameEs"`
	NameFr    string    `json:"nameFr"`
	NamePt    string    `json:"namePt"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
