package entities

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/lautarok/hexa/src/domain/models"
	"github.com/uptrace/bun"
)

type User struct {
	ID         uuid.UUID  `bun:"id,type:uuid,pk,default:gen_random_uuid()"`
	Name       string     `bun:"name,notnull,nullzero,type:varchar(40)"`
	Surname    string     `bun:"surname,nullzero,type:varchar(40)"`
	RoleID     uuid.UUID  `bun:"role_id,type:uuid,notnull"`
	Role       Role       `bun:"rel:belongs-to,join:role_id=id,on_delete:cascade"`
	Credential Credential `bun:"rel:has-one,join:id=user_id"`
	CreatedAt  time.Time  `bun:"created_at,type:timestamp,notnull,nullzero,default:current_timestamp"`
	UpdatedAt  time.Time  `bun:"updated_at,type:timestamp,notnull,nullzero,default:current_timestamp"`
}

func (entity *User) BeforeUpdate(ctx context.Context, query bun.Query) error {
	entity.UpdatedAt = time.Now()
	return nil
}

func (entity *User) ToDomainModel() *models.User {
	user := &models.User{
		ID:        entity.ID,
		Name:      entity.Name,
		Surname:   entity.Surname,
		CreatedAt: entity.CreatedAt,
		UpdatedAt: entity.UpdatedAt,
	}

	user.Role = *entity.Role.ToDomainModel()
	user.Credential = *entity.Credential.ToDomainModel()

	return user
}

func (entity *User) FromDomainModel(model *models.User) {
	entity.ID = model.ID
	entity.Name = model.Name
	entity.Surname = model.Surname
	entity.CreatedAt = model.CreatedAt
	entity.UpdatedAt = model.UpdatedAt

	var role Role
	role.FromDomainModel(&model.Role)
	entity.Role = role
	entity.RoleID = entity.Role.ID

	var credential Credential
	credential.FromDomainModel(&model.Credential)
	entity.Credential = credential
}
