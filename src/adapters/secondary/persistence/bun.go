package persistence

import (
	"database/sql"

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

	return &BunAdapter{
		db: db,
	}
}
