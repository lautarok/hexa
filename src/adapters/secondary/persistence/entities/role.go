package entities

import (
	"context"
	"time"

	"github.com/lautarok/hexa/src/app/domain"
	"github.com/uptrace/bun"
)

type Role struct {
	ID          string        `bun:"id,type:uuid,notnull,pk,default:gen_random_uuid()"`
	NameEn      string        `bun:"name_en,type:varchar(30)"`
	NameEs      string        `bun:"name_es,type:varchar(30)"`
	NameFr      string        `bun:"name_fr,type:varchar(30)"`
	NamePt      string        `bun:"name_pt,type:varchar(30)"`
	Users       []*User       `bun:"rel:has-many,join:id=role_id"`
	Permissions []*Permission `bun:"m2m:role_permissions,join:Role=Permission"`
	CreatedAt   time.Time     `bun:"created_at,type:timestamp,default:current_timestamp,notnull"`
	UpdatedAt   time.Time     `bun:"updated_at,type:timestamp,default:current_timestamp,notnull"`
}

func (entity *Role) BeforeUpdate(ctx context.Context, query bun.Query) error {
	entity.UpdatedAt = time.Now()
	return nil
}

func (entity *Role) ToDomain() *domain.Role {
	role := &domain.Role{
		ID:        entity.ID,
		NameEn:    entity.NameEn,
		NameEs:    entity.NameEs,
		NameFr:    entity.NameFr,
		NamePt:    entity.NamePt,
		CreatedAt: entity.CreatedAt,
		UpdatedAt: entity.UpdatedAt,
	}

	if entity.Users != nil {
		for _, user := range entity.Users {
			role.Users = append(role.Users, user.ToDomain())
		}
	}

	if entity.Permissions != nil {
		for _, permission := range entity.Permissions {
			role.Permissions = append(role.Permissions, permission.ToDomain())
		}
	}

	return role
}
