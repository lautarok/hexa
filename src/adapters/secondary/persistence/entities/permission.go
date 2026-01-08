package entities

import (
	"context"
	"time"

	"github.com/uptrace/bun"
)

type Permission struct {
	ID        string    `bun:"id,type:uuid,default:uuid_generate_v4(),notnull,pk"`
	Alias     string    `bun:"alias,notnull"`
	Roles     []*Role   `bun:"m2m:role_permissions,join:Permission=Role"`
	CreatedAt time.Time `bun:"created_at,type:timestamp,default:current_timestamp,notnull"`
	UpdatedAt time.Time `bun:"updated_at,type:timestamp,default:current_timestamp,notnull"`
}

func (entity *Permission) BeforeUpdate(ctx context.Context, query bun.Query) error {
	entity.UpdatedAt = time.Now()
	return nil
}
