package database

import (
	"context"

	"github.com/forbole/bdjuno/v4/database/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func (db *DB) CreateGroup(ctx context.Context, groupMembers []*models.Group) error {
	return nil
}

func (db *DB) UpdateGroup(ctx context.Context, group *models.Group) error {
	return nil
}

func (db *DB) CreateGroupToSQL(ctx context.Context, groupMembers []*models.Group) (string, []interface{}) {
	stat := db.G.Session(&gorm.Session{DryRun: true}).Table((&models.Group{}).TableName()).Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "account_address"}, {Name: "group_id"}},
		UpdateAll: true,
	}).Create(groupMembers).Statement
	return stat.SQL.String(), stat.Vars
}

func (db *DB) UpdateGroupToSQL(ctx context.Context, group *models.Group) (string, []interface{}) {
	stat := db.G.Session(&gorm.Session{DryRun: true}).Table((&models.Group{}).TableName()).Where("account_address = ? AND group_id = ? ", group.AccountAddress, group.GroupID).Updates(group).Statement
	return stat.SQL.String(), stat.Vars
}

func (db *DB) UpdateGroupByOwnerAndNameToSQL(ctx context.Context, group *models.Group) (string, []interface{}) {
	stat := db.G.Session(&gorm.Session{DryRun: true}).Table((&models.Group{}).TableName()).Where("account_address = ? AND owner = ? AND group_name = ?  AND removed=false", group.AccountAddress, group.OwnerAddress, group.GroupName).Updates(group).Statement
	return stat.SQL.String(), stat.Vars
}

func (db *DB) DeleteGroupToSQL(ctx context.Context, group *models.Group) (string, []interface{}) {
	stat := db.G.Session(&gorm.Session{DryRun: true}).Table((&models.Group{}).TableName()).Where("group_id = ?", group.GroupID).Updates(group).Statement
	return stat.SQL.String(), stat.Vars
}

func (db *DB) BatchDeleteGroupMemberToSQL(ctx context.Context, group *models.Group, groupID string, accountIDs []string) (string, []interface{}) {
	stat := db.G.Session(&gorm.Session{DryRun: true}).Table((&models.Group{}).TableName()).Where("account_address IN ? AND group_id = ? ", accountIDs, groupID).UpdateColumns(group).Statement
	return stat.SQL.String(), stat.Vars
}
