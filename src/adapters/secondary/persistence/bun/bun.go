package bun

import (
	"context"
	"database/sql"
	"log"

	"github.com/lautarok/hexa/src/adapters/secondary/persistence/bun/entities"
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

type txKey struct{}

func (adapter *BunAdapter) GetDB(ctx context.Context) bun.IDB {
	if tx, ok := ctx.Value(txKey{}).(bun.IDB); ok {
		return tx
	}
	return adapter.db
}

func (adapter *BunAdapter) Transaction(
	ctx context.Context,
	function func(ctx context.Context) error,
) error {
	return adapter.db.RunInTx(
		ctx,
		&sql.TxOptions{},
		func(ctx context.Context, tx bun.Tx) error {
			ctxWithTx := context.WithValue(ctx, txKey{}, tx)
			return function(ctxWithTx)
		},
	)
}
