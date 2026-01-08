package entities

import (
	"context"
	"time"

	"github.com/uptrace/bun"
)

type RolePermission struct {
	ID           string      `bun:"id,type:uuid,default:gen_random_uuid(),notnull,pk"`
	RoleID       string      `bun:"role_id,type:uuid,notnull,pk"`
	Role         *Role       `bun:"rel:belongs-to,join:role_id=id"`
	PermissionID string      `bun:"permission_id,type:uuid,notnull,pk"`
	Permission   *Permission `bun:"rel:belongs-to,join:permission_id=id"`
	CreatedAt    time.Time   `bun:"created_at,notnull,default:current_timestamp,type:timestamp"`
	UpdatedAt    time.Time   `bun:"updated_at,notnull,default:current_timestamp,type:timestamp"`
}

func (entity *RolePermission) BeforeUpdate(ctx context.Context, query bun.Query) error {
	entity.UpdatedAt = time.Now()
	return nil
}
