package database

import (
	"fmt"

	"github.com/forbole/juno/v5/database"
	"github.com/forbole/juno/v5/database/postgresql"
	"github.com/jmoiron/sqlx"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var _ database.Database = &Db{}

// Db represents a PostgreSQL database with expanded features.
// so that it can properly store custom BigDipper-related data.
type Db struct {
	G *gorm.DB
	*postgresql.Database
	Sqlx *sqlx.DB
}

// Builder allows to create a new Db instance implementing the db.Builder type
func Builder(ctx *database.Context) (database.Database, error) {
	database, err := postgresql.Builder(ctx)
	if err != nil {
		return nil, err
	}

	psqlDb, ok := (database).(*postgresql.Database)
	if !ok {
		return nil, fmt.Errorf("invalid configuration database, must be PostgreSQL")
	}
	gdb, err := gorm.Open(postgres.Open(ctx.Cfg.URL), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("error while opening the connection to the database: %s", err)
	}
	return &Db{
		G:        gdb,
		Database: psqlDb,
		Sqlx:     sqlx.NewDb(psqlDb.SQL.DB, "postgresql"),
	}, nil
}

// Cast allows to cast the given db to a Db instance
func Cast(db database.Database) *Db {
	bdDatabase, ok := db.(*Db)
	if !ok {
		panic(fmt.Errorf("given database instance is not a Db"))
	}
	return bdDatabase
}
