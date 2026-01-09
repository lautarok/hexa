package entities

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/lautarok/hexa/src/app/domain"
	"github.com/uptrace/bun"
)

type User struct {
	ID         uuid.UUID   `bun:"id,type:uuid,pk,default:gen_random_uuid()"`
	Name       string      `bun:"name,notnull,type:varchar(40)"`
	Surname    string      `bun:"surname,notnull,type:varchar(40)"`
	RoleID     uuid.UUID   `bun:"role_id,type:uuid,notnull"`
	Role       *Role       `bun:"rel:belongs-to,join:role_id=id,on_delete:CASCADE"`
	Credential *Credential `bun:"rel:has-one,join:id=user_id"`
	CreatedAt  time.Time   `bun:"created_at,type:timestamp,notnull,default:current_timestamp"`
	UpdatedAt  time.Time   `bun:"updated_at,type:timestamp,notnull,default:current_timestamp"`
}

func (entity *User) BeforeUpdate(ctx context.Context, query bun.Query) error {
	entity.UpdatedAt = time.Now()
	return nil
}

func (entity *User) ToDomain() *domain.User {
	user := &domain.User{
		ID:        entity.ID,
		Name:      entity.Name,
		Surname:   entity.Surname,
		CreatedAt: entity.CreatedAt,
		UpdatedAt: entity.UpdatedAt,
	}

	if entity.Role != nil {
		user.Role = entity.Role.ToDomain()
	}

	if entity.Credential != nil {
		user.Credential = entity.Credential.ToDomain()
	}

	return user
}

func (entity *User) FromDomain(domain *domain.User) {
	entity.ID = domain.ID
	entity.Name = domain.Name
	entity.Surname = domain.Surname
	entity.CreatedAt = domain.CreatedAt
	entity.UpdatedAt = domain.UpdatedAt

	if domain.Role != nil {
		var role Role
		role.FromDomain(domain.Role)
		entity.Role = &role
	}

	if domain.Credential != nil {
		var credential Credential
		credential.FromDomain(domain.Credential)
		entity.Credential = &credential
	}
}
