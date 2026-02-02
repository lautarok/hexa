package entities

import (
	"time"

	"github.com/google/uuid"
	"github.com/lautarok/hexa/src/core/domain"
)

type GoogleIdentity struct {
	ID        uuid.UUID `bun:"id,type:uuid,notnull,default:gen_random_uuid(),notnull,nullzero,pk"`
	GoogleID  string    `bun:"google_id,type:varchar(255),notnull,nullzero"`
	UserID    uuid.UUID `bun:"user_id,type:uuid,notnull,nullzero"`
	User      User      `bun:"rel:belongs-to,join=user_id=id,on_delete:cascade"`
	CreatedAt time.Time `bun:"created_at,type:timestamp,notnull,nullzero,default:current_timestamp"`
}

func (entity *GoogleIdentity) ToDomain() *domain.GoogleIdentity {
	domain := domain.GoogleIdentity{
		ID:        entity.ID,
		GoogleID:  entity.GoogleID,
		User:      *entity.User.ToDomain(),
		CreatedAt: entity.CreatedAt,
	}

	return &domain
}

func (entity *GoogleIdentity) FromDomain(domain *domain.GoogleIdentity) {
	user := User{}
	user.FromDomain(&domain.User)

	entity.ID = domain.ID
	entity.GoogleID = domain.GoogleID
	entity.UserID = domain.User.ID
	entity.User = user
	entity.CreatedAt = domain.CreatedAt
}
