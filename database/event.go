package database

import (
	"context"

	"github.com/forbole/bdjuno/v4/database/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func (db *DB) SaveBucketEventToSQL(ctx context.Context, be *models.BucketEvent) (string, []interface{}) {
	stat := db.G.Session(&gorm.Session{DryRun: true}).Table((&models.BucketEvent{}).TableName()).Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "bucket_id"}},
		UpdateAll: true,
	}).Create(be).Statement
	return stat.SQL.String(), stat.Vars
}

func (db *DB) SaveObjectEventToSQL(ctx context.Context, oe *models.ObjectEvent) (string, []interface{}) {
	stat := db.G.Session(&gorm.Session{DryRun: true}).Table((&models.ObjectEvent{}).TableName()).Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "object_id"}},
		UpdateAll: true,
	}).Create(oe).Statement
	return stat.SQL.String(), stat.Vars
}

func (db *DB) SaveGroupEventToSQL(ctx context.Context, ge *models.GroupEvent) (string, []interface{}) {
	stat := db.G.Session(&gorm.Session{DryRun: true}).Table((&models.GroupEvent{}).TableName()).Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "group_id"}},
		UpdateAll: true,
	}).Create(ge).Statement
	return stat.SQL.String(), stat.Vars
}
