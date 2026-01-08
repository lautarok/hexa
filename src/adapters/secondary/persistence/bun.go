package persistence

import (
	"context"
	"database/sql"
	"log"

	"github.com/lautarok/hexa/src/adapters/secondary/persistence/entities"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
)

type BunAdapter struct {
	db bun.IDB
}

type BunAdapterDeps struct {
	DSN string
}

func NewBunAdapter(deps *BunAdapterDeps) *BunAdapter {
	sqldb := sql.OpenDB(pgdriver.NewConnector(
		pgdriver.WithDSN(deps.DSN),
	))

	db := bun.NewDB(sqldb, pgdialect.New())

	db.RegisterModel((*entities.RolePermission)(nil))
	db.RegisterModel((*entities.Permission)(nil))
	db.RegisterModel((*entities.Role)(nil))
	db.RegisterModel((*entities.Credential)(nil))
	db.RegisterModel((*entities.User)(nil))

	ctx := context.Background()

	_, err := db.NewCreateTable().Model((*entities.Permission)(nil)).IfNotExists().Exec(ctx)
	if err != nil {
		log.Fatal(err)
	}
	db.NewCreateTable().Model((*entities.Role)(nil)).IfNotExists().Exec(ctx)
	db.NewCreateTable().Model((*entities.RolePermission)(nil)).IfNotExists().Exec(ctx)
	db.NewCreateTable().Model((*entities.User)(nil)).IfNotExists().Exec(ctx)
	db.NewCreateTable().Model((*entities.Credential)(nil)).IfNotExists().Exec(ctx)

	return &BunAdapter{
		db: db,
	}
}

func (adapter *BunAdapter) GetDB() bun.IDB {
	return adapter.db
}
