package models

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type GoogleIdentity struct {
	ID        uuid.UUID
	GoogleID  string
	UserID    uuid.UUID
	CreatedAt time.Time
}

type IGoogleIdentitiesRepository interface {
	FindOneByGoogleID(ctx context.Context, googleId string) (*GoogleIdentity, error)
	CreateOne(ctx context.Context, googleIdentity *GoogleIdentity) (*GoogleIdentity, error)
}
