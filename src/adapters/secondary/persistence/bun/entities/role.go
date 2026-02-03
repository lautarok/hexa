package entities

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/lautarok/hexa/src/domain/models"
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

func (entity *Role) ToDomainModel() *models.Role {
	role := &models.Role{
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
			role.Users = append(role.Users, user.ToDomainModel())
		}
	}

	if entity.Permissions != nil {
		for _, permission := range entity.Permissions {
			role.Permissions = append(role.Permissions, *permission.ToDomainModel())
		}
	}

	return role
}

func (entity *Role) FromDomainModel(model *models.Role) {
	entity.ID = model.ID
	entity.NameEn = model.NameEn
	entity.NameEs = model.NameEs
	entity.NameFr = model.NameFr
	entity.NamePt = model.NamePt
	entity.NameNl = model.NameNl
	entity.CreatedAt = model.CreatedAt
	entity.UpdatedAt = model.UpdatedAt

	if model.Users != nil {
		for _, userDomain := range model.Users {
			var user User
			user.FromDomainModel(userDomain)
			entity.Users = append(entity.Users, &user)
		}
	}

	if model.Permissions != nil {
		for _, permissionDomain := range model.Permissions {
			var permission Permission
			permission.FromDomainModel(&permissionDomain)
			entity.Permissions = append(entity.Permissions, permission)
		}
	}
}
