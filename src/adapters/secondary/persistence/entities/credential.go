package entities

import (
	"context"
	"time"

	"github.com/lautarok/hexa/src/app/domain"
	"github.com/uptrace/bun"
)

type Credential struct {
	ID           string    `bun:"id,type:uuid,pk,default:gen_random_uuid(),notnull"`
	UserID       string    `bun:"user_id,type:uuid,notnull,unique"`
	User         *User     `bun:"rel:belongs-to,join:user_id=id"`
	Email        string    `bun:"email,type:varchar(200),notnull,unique"`
	Username     string    `bun:"username,type:varchar(25),notnull,unique"`
	PasswordHash string    `bun:"password,type:varchar(255),notnull"`
	CreatedAt    time.Time `bun:"created_at,type:timestamp,notnull,default:current_timestamp"`
	UpdatedAt    time.Time `bun:"updated_at,type:timestamp,notnull,default:current_timestamp"`
}

func (entity *Credential) BeforeUpdate(ctx context.Context, query bun.Query) error {
	entity.UpdatedAt = time.Now()
	return nil
}

func (entity *Credential) ToDomain() *domain.Credential {
	credential := &domain.Credential{
		ID:           entity.ID,
		Email:        entity.Email,
		Username:     entity.Username,
		PasswordHash: entity.PasswordHash,
		CreatedAt:    entity.CreatedAt,
		UpdatedAt:    entity.UpdatedAt,
	}

	if entity.User != nil {
		credential.User = entity.User.ToDomain()
	}

	return credential
}
