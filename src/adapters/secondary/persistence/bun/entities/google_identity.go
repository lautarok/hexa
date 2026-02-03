package entities

import (
	"time"

	"github.com/google/uuid"
	"github.com/lautarok/hexa/src/domain/models"
)

type GoogleIdentity struct {
	ID        uuid.UUID `bun:"id,type:uuid,default:gen_random_uuid(),notnull,nullzero,pk"`
	GoogleID  string    `bun:"google_id,type:varchar(255),notnull,nullzero"`
	UserID    uuid.UUID `bun:"user_id,type:uuid,notnull,nullzero"`
	CreatedAt time.Time `bun:"created_at,type:timestamp,notnull,nullzero,default:current_timestamp"`
}

func (entity *GoogleIdentity) ToDomainModel() *models.GoogleIdentity {
	domain := models.GoogleIdentity{
		ID:        entity.ID,
		GoogleID:  entity.GoogleID,
		CreatedAt: entity.CreatedAt,
	}

	return &domain
}

func (entity *GoogleIdentity) FromDomainModel(model *models.GoogleIdentity) {
	entity.ID = model.ID
	entity.GoogleID = model.GoogleID
	entity.UserID = model.UserID
	entity.CreatedAt = model.CreatedAt
}
