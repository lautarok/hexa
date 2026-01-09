package bun

import (
	"context"
	"log"

	"github.com/lautarok/hexa/src/adapters/secondary/persistence/bun/entities"
	"github.com/uptrace/bun"
)

func SetupInitialTables(db bun.IDB) {
	ctx := context.Background()

	_, err := db.NewCreateTable().Model((*entities.RolePermission)(nil)).IfNotExists().Exec(ctx)
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.NewCreateTable().Model((*entities.Permission)(nil)).IfNotExists().Exec(ctx)
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.NewCreateTable().Model((*entities.Role)(nil)).IfNotExists().Exec(ctx)
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.NewCreateTable().Model((*entities.User)(nil)).IfNotExists().Exec(ctx)
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.NewCreateTable().Model((*entities.Credential)(nil)).IfNotExists().Exec(ctx)
	if err != nil {
		log.Fatal(err)
	}
}
