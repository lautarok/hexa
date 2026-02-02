package entities

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/lautarok/hexa/src/core/domain"
	"github.com/uptrace/bun"
)

type Role struct {
	ID          uuid.UUID    `bun:"id,type:uuid,notnull,nullzero,pk,default:gen_random_uuid()"`
	NameEn      string       `bun:"name_en,type:varchar(30),unique,nullzero"`
	NameEs      string       `bun:"name_es,type:varchar(30),unique,nullzero"`
	NameFr      string       `bun:"name_fr,type:varchar(30),unique,nullzero"`
	NamePt      string       `bun:"name_pt,type:varchar(30),unique,nullzero"`
	NameNl      string       `bun:"name_nl,type:varchar(30),unique,nullzero"`
	Lock        bool         `bun:"lock,notnull,default:false"`
	Users       []*User      `bun:"rel:has-many,join:id=role_id"`
	Permissions []Permission `bun:"m2m:role_permissions,join:Role=Permission"`
	CreatedAt   time.Time    `bun:"created_at,type:timestamp,default:current_timestamp,notnull,nullzero"`
	UpdatedAt   time.Time    `bun:"updated_at,type:timestamp,default:current_timestamp,notnull,nullzero"`
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
			role.Permissions = append(role.Permissions, *permission.ToDomain())
		}
	}

	return role
}

func (entity *Role) FromDomain(domain *domain.Role) {
	entity.ID = domain.ID
	entity.NameEn = domain.NameEn
	entity.NameEs = domain.NameEs
	entity.NameFr = domain.NameFr
	entity.NamePt = domain.NamePt
	entity.NameNl = domain.NameNl
	entity.CreatedAt = domain.CreatedAt
	entity.UpdatedAt = domain.UpdatedAt

	if domain.Users != nil {
		for _, userDomain := range domain.Users {
			var user User
			user.FromDomain(userDomain)
			entity.Users = append(entity.Users, &user)
		}
	}

	if domain.Permissions != nil {
		for _, permissionDomain := range domain.Permissions {
			var permission Permission
			permission.FromDomain(&permissionDomain)
			entity.Permissions = append(entity.Permissions, permission)
		}
	}
}
