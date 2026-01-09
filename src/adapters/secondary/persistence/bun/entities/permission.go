package entities

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/lautarok/hexa/src/app/domain"
	"github.com/uptrace/bun"
)

type Permission struct {
	ID        uuid.UUID `bun:"id,type:uuid,default:gen_random_uuid(),notnull,pk"`
	Alias     string    `bun:"alias,notnull"`
	Roles     []*Role   `bun:"m2m:role_permissions,join:Permission=Role"`
	CreatedAt time.Time `bun:"created_at,type:timestamp,default:current_timestamp,notnull"`
	UpdatedAt time.Time `bun:"updated_at,type:timestamp,default:current_timestamp,notnull"`
}

func (entity *Permission) BeforeUpdate(ctx context.Context, query bun.Query) error {
	entity.UpdatedAt = time.Now()
	return nil
}

func (entity *Permission) ToDomain() *domain.Permission {
	permission := &domain.Permission{
		ID:        entity.ID,
		Alias:     entity.Alias,
		CreatedAt: entity.CreatedAt,
		UpdatedAt: entity.UpdatedAt,
	}

	if entity.Roles != nil {
		for _, role := range entity.Roles {
			permission.Roles = append(permission.Roles, role.ToDomain())
		}
	}

	return permission
}

func (entity *Permission) FromDomain(domain *domain.Permission) {
	entity.ID = domain.ID
	entity.Alias = domain.Alias
	entity.CreatedAt = domain.CreatedAt
	entity.UpdatedAt = domain.UpdatedAt

	if domain.Roles != nil {
		for _, roleDomain := range domain.Roles {
			var role Role
			role.FromDomain(roleDomain)
			entity.Roles = append(entity.Roles, &role)
		}
	}
}
