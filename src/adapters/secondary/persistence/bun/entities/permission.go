package entities

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/lautarok/hexa/src/domain/models"
	"github.com/uptrace/bun"
)

type Permission struct {
	ID        uuid.UUID `bun:"id,type:uuid,default:gen_random_uuid(),notnull,nullzero,pk"`
	Alias     string    `bun:"alias,notnull,nullzero,pk"`
	Roles     []*Role   `bun:"m2m:role_permissions"`
	CreatedAt time.Time `bun:"created_at,type:timestamp,default:current_timestamp,notnull,nullzero"`
	UpdatedAt time.Time `bun:"updated_at,type:timestamp,default:current_timestamp,notnull,nullzero"`
}

func (entity *Permission) BeforeUpdate(ctx context.Context, query bun.Query) error {
	entity.UpdatedAt = time.Now()
	return nil
}

func (entity *Permission) ToDomainModel() *models.Permission {
	permission := &models.Permission{
		ID:        entity.ID,
		Alias:     entity.Alias,
		CreatedAt: entity.CreatedAt,
		UpdatedAt: entity.UpdatedAt,
	}

	if entity.Roles != nil {
		for _, role := range entity.Roles {
			permission.Roles = append(permission.Roles, role.ToDomainModel())
		}
	}

	return permission
}

func (entity *Permission) FromDomainModel(model *models.Permission) {
	entity.ID = model.ID
	entity.Alias = model.Alias
	entity.CreatedAt = model.CreatedAt
	entity.UpdatedAt = model.UpdatedAt

	if model.Roles != nil {
		for _, roleDomain := range model.Roles {
			var role Role
			role.FromDomainModel(roleDomain)
			entity.Roles = append(entity.Roles, &role)
		}
	}
}
