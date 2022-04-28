package db_testing_base

import (
	"context"
	_ "github.com/mattn/go-sqlite3"
	"github.com/shifty11/cosmos-gov/database"
	"github.com/shifty11/cosmos-gov/ent"
	"github.com/shifty11/cosmos-gov/ent/enttest"
	"testing"
)

func GetBase(t *testing.T) (*ent.Client, context.Context, database.DbManagers) {
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	ctx := context.Background()

	var m = database.NewCustomDbManagers(client, ctx)
	return client, ctx, m
}
