package entities

import (
	"context"
	"time"

	"github.com/uptrace/bun"
)

type User struct {
	ID        string    `bun:"ID,type:uuid,pk,default:uuid_generate_v4()"`
	Name      string    `bun:"Name,notnull,type:VARCHAR(40)"`
	Surname   string    `bun:"Surname,notnull,type:VARCHAR(40)"`
	RoleID    string    `bun:"Role_ID,type:uuid,notnull"`
	Role      *Role     `bun:"rel:belongs-to,join:Role_ID=ID,on_delete:CASCADE"`
	CreatedAt time.Time `bun:"Created_At,type:timestamp,notnull,default:current_timestamp"`
	UpdatedAt time.Time `bun:"Updated_At,type:timestamp,notnull,default:current_timestamp"`
}

func (entity *User) BeforeUpdate(ctx context.Context, query bun.Query) error {
	entity.UpdatedAt = time.Now()
	return nil
}
