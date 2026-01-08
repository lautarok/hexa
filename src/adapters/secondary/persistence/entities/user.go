package entities

import (
	"context"
	"time"

	"github.com/uptrace/bun"
)

type User struct {
	ID         string      `bun:"id,type:uuid,pk,default:uuid_generate_v4()"`
	Name       string      `bun:"name,notnull,type:varchar(40)"`
	Surname    string      `bun:"surname,notnull,type:varchar(40)"`
	RoleID     string      `bun:"role_id,type:uuid,notnull"`
	Role       *Role       `bun:"rel:belongs-to,join:role_id=id,on_delete:CASCADE"`
	Credential *Credential `bun:"rel:has-one,join:id=user_id"`
	CreatedAt  time.Time   `bun:"created_at,type:timestamp,notnull,default:current_timestamp"`
	UpdatedAt  time.Time   `bun:"updated_at,type:timestamp,notnull,default:current_timestamp"`
}

func (entity *User) BeforeUpdate(ctx context.Context, query bun.Query) error {
	entity.UpdatedAt = time.Now()
	return nil
}
