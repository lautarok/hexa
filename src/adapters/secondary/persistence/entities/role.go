package entities

import (
	"context"
	"time"

	"github.com/uptrace/bun"
)

type Role struct {
	ID          string        `bun:"id,type:uuid,notnull,pk,default:uuid_generate_v4()"`
	NameEn      string        `bun:"name_en,type:varchar(30)"`
	NameEs      string        `bun:"name_es,type:varchar(30)"`
	NameFr      string        `bun:"name_fr,type:varchar(30)"`
	NamePt      string        `bun:"name_pt,type:varchar(30)"`
	Users       []*User       `bun:"rel:has-many,join:ID=Role_ID"`
	Permissions []*Permission `bun:"m2m:role_permissions,join:Role=Permission"`
	CreatedAt   time.Time     `bun:"created_at,type:timestamp,default:current_timestamp,notnull"`
	UpdatedAt   time.Time     `bun:"updated_at,type:timestamp,default:current_timestamp,notnull"`
}

func (entity *Role) BeforeUpdate(ctx context.Context, query bun.Query) error {
	entity.UpdatedAt = time.Now()
	return nil
}
