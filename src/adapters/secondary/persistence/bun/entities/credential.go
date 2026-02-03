package entities

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/lautarok/hexa/src/domain/models"
	"github.com/uptrace/bun"
)

type Credential struct {
	ID        uuid.UUID `bun:"id,type:uuid,pk,default:gen_random_uuid(),notnull,nullzero"`
	UserID    uuid.UUID `bun:"user_id,type:uuid,notnull,nullzero,unique"`
	Email     string    `bun:"email,type:varchar(200),notnull,nullzero,unique"`
	Username  string    `bun:"username,type:varchar(25),nullzero,unique"`
	Password  string    `bun:"password,type:varchar(255),nullzero"`
	CreatedAt time.Time `bun:"created_at,type:timestamp,notnull,nullzero,default:current_timestamp"`
	UpdatedAt time.Time `bun:"updated_at,type:timestamp,notnull,nullzero,default:current_timestamp"`
}

func (entity *Credential) BeforeUpdate(ctx context.Context, query bun.Query) error {
	entity.UpdatedAt = time.Now()
	return nil
}

func (entity *Credential) ToDomainModel() *models.Credential {
	credential := &models.Credential{
		ID:        entity.ID,
		UserID:    entity.UserID,
		Email:     entity.Email,
		Username:  entity.Username,
		Password:  entity.Password,
		CreatedAt: entity.CreatedAt,
		UpdatedAt: entity.UpdatedAt,
	}

	return credential
}

func (entity *Credential) FromDomainModel(model *models.Credential) {
	entity.ID = model.ID
	entity.UserID = model.UserID
	entity.Email = model.Email
	entity.Username = model.Username
	entity.Password = model.Password
	entity.CreatedAt = model.CreatedAt
	entity.UpdatedAt = model.UpdatedAt
}
