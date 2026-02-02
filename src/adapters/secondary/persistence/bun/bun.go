package bun

import (
	"context"
	"database/sql"
	"errors"
	"strings"

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

func (adapter *BunAdapter) IsUniqueViolation(err error) bool {
	return strings.Contains(err.Error(), "23505")
}

func (adapter *BunAdapter) IsErrNotFound(err error) bool {
	return errors.Is(err, sql.ErrNoRows)
}
