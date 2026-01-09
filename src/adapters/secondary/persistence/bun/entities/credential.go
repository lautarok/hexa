package entities

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/lautarok/hexa/src/app/domain"
	"github.com/uptrace/bun"
)

type Credential struct {
	ID        uuid.UUID `bun:"id,type:uuid,pk,default:gen_random_uuid(),notnull,nullzero"`
	UserID    uuid.UUID `bun:"user_id,type:uuid,notnull,nullzero,unique"`
	User      *User     `bun:"rel:belongs-to,join:user_id=id"`
	Email     string    `bun:"email,type:varchar(200),notnull,nullzero,unique"`
	Username  string    `bun:"username,type:varchar(25),notnull,nullzero,unique"`
	Password  string    `bun:"password,type:varchar(255),notnull,nullzero"`
	CreatedAt time.Time `bun:"created_at,type:timestamp,notnull,nullzero,default:current_timestamp"`
	UpdatedAt time.Time `bun:"updated_at,type:timestamp,notnull,nullzero,default:current_timestamp"`
}

func (entity *Credential) BeforeUpdate(ctx context.Context, query bun.Query) error {
	entity.UpdatedAt = time.Now()
	return nil
}

func (entity *Credential) ToDomain() *domain.Credential {
	credential := &domain.Credential{
		ID:        entity.ID,
		Email:     entity.Email,
		Username:  entity.Username,
		Password:  entity.Password,
		CreatedAt: entity.CreatedAt,
		UpdatedAt: entity.UpdatedAt,
	}

	if entity.User != nil {
		credential.User = entity.User.ToDomain()
	}

	return credential
}

func (entity *Credential) FromDomain(domain *domain.Credential) {
	entity.ID = domain.ID
	entity.Email = domain.Email
	entity.Username = domain.Username
	entity.Password = domain.Password
	entity.CreatedAt = domain.CreatedAt
	entity.UpdatedAt = domain.UpdatedAt

	if domain.User != nil {
		var user User
		user.FromDomain(domain.User)
		entity.User = &user
		entity.UserID = entity.User.ID
	}
}
