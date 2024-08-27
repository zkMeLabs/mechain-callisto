package database

import (
	"context"

	"github.com/forbole/bdjuno/v4/database/models"
	"gorm.io/gorm"
)

func (db *DB) SaveBucketEventToSQL(ctx context.Context, be *models.BucketEvent) (string, []interface{}) {
	stat := db.G.Session(&gorm.Session{DryRun: true}).Table((&models.BucketEvent{}).TableName()).Create(be).Statement
	return stat.SQL.String(), stat.Vars
}

func (db *DB) SaveObjectEventToSQL(ctx context.Context, oe *models.ObjectEvent) (string, []interface{}) {
	stat := db.G.Session(&gorm.Session{DryRun: true}).Table((&models.ObjectEvent{}).TableName()).Create(oe).Statement
	return stat.SQL.String(), stat.Vars
}

func (db *DB) SaveGroupEventToSQL(ctx context.Context, ge *models.GroupEvent) (string, []interface{}) {
	stat := db.G.Session(&gorm.Session{DryRun: true}).Table((&models.GroupEvent{}).TableName()).Create(ge).Statement
	return stat.SQL.String(), stat.Vars
}
