package domain

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type GoogleIdentity struct {
	ID        uuid.UUID
	GoogleID  string
	User      User
	CreatedAt time.Time
}

type IGoogleIdentitiesRepository interface {
	FindOneByGoogleID(ctx context.Context, googleId string) (*GoogleIdentity, error)
	CreateOne(ctx context.Context, googleIdentity *GoogleIdentity) error
}
