package entities

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type RolePermission struct {
	ID           uuid.UUID   `bun:"id,type:uuid,default:gen_random_uuid(),notnull,nullzero,pk"`
	RoleID       uuid.UUID   `bun:"role_id,type:uuid,notnull,nullzero,pk"`
	Role         *Role       `bun:"rel:belongs-to,join:role_id=id"`
	PermissionID uuid.UUID   `bun:"permission_id,type:uuid,notnull,nullzero,pk"`
	Permission   *Permission `bun:"rel:belongs-to,join:permission_id=id"`
	CreatedAt    time.Time   `bun:"created_at,notnull,nullzero,default:current_timestamp,type:timestamp"`
	UpdatedAt    time.Time   `bun:"updated_at,notnull,nullzero,default:current_timestamp,type:timestamp"`
}

func (entity *RolePermission) BeforeUpdate(ctx context.Context, query bun.Query) error {
	entity.UpdatedAt = time.Now()
	return nil
}
