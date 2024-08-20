package database

import (
	"fmt"

	"github.com/forbole/juno/v5/database"
	"github.com/forbole/juno/v5/database/postgresql"
	"github.com/jmoiron/sqlx"
	"gorm.io/gorm"
)

var _ database.Database = &Db{}

// Db represents a PostgreSQL database with expanded features.
// so that it can properly store custom BigDipper-related data.
type Db struct {
	Db *gorm.DB
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

	return &Db{
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
