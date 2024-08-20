package database

import (
	"context"
	"fmt"

	"github.com/forbole/bdjuno/v4/database/models"
	dbtypes "github.com/forbole/bdjuno/v4/database/types"
	"github.com/forbole/bdjuno/v4/types"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func (db *Db) SaveStorageGroup(height int64, groups []types.StorageGroup) error {
	if len(groups) == 0 {
		return nil
	}

	var accounts []types.Account

	groupQuery := `
INSERT INTO storage_group(
	id, group_name, owner, source_type, extra, height, tags
) VALUES`
	var groupParams []interface{}

	for i, group := range groups {
		// Prepare the account query
		accounts = append(accounts, types.NewAccount(group.Owner))

		// Prepare the group query
		vi := i * 7
		groupQuery += fmt.Sprintf("($%d,$%d,$%d,$%d,$%d,$%d,$%d),",
			vi+1, vi+2, vi+3, vi+4, vi+5, vi+6, vi+7)

		groupParams = append(groupParams,
			group.GroupId,
			group.GroupName,
			group.Owner,
			group.SourceType,
			group.Extra,
			height,
			group.Tags,
		)
	}

	// Store the accounts
	err := db.SaveAccounts(accounts)
	if err != nil {
		return fmt.Errorf("error while storing owner accounts: %s", err)
	}

	// Store the groups
	groupQuery = groupQuery[:len(groupQuery)-1] // Remove trailing ","
	groupQuery += " ON CONFLICT DO NOTHING"
	_, err = db.SQL.Exec(groupQuery, groupParams...)
	if err != nil {
		return fmt.Errorf("error while storing groups: %s", err)
	}

	return nil
}

func (db *Db) GetGroup(id uint64) (*types.StorageGroup, error) {
	var rows []*dbtypes.StorageGroupRow
	err := db.Sqlx.Select(&rows, `SELECT * FROM storage_group WHERE id = $1`, id)
	if err != nil {
		return nil, err
	}

	if len(rows) == 0 {
		return nil, nil
	}

	row := rows[0]

	group := types.NewStorageGroup(
		row.GroupId,
		row.GroupName,
		row.Owner,
		row.SourceType,
		row.Extra,
		row.Height,
		row.Tags,
	)
	return &group, nil
}

func (db *Db) SaveBucket(ctx context.Context, bucket *models.Bucket) (string, []interface{}) {
	stat := db.Db.Session(&gorm.Session{DryRun: true}).Table((&models.Bucket{}).TableName()).Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "bucket_id"}},
		UpdateAll: true,
	}).Create(bucket).Statement

	return stat.SQL.String(), stat.Vars
}

func (db *Db) GetBucket(id uint64) (*types.Bucket, error) {
	var rows []*dbtypes.BucketRow
	err := db.Sqlx.Select(&rows, `SELECT * FROM bucket WHERE id = $1`, id)
	if err != nil {
		return nil, err
	}

	if len(rows) == 0 {
		return nil, nil
	}

	row := rows[0]

	bucket := types.NewBucket(
		row.BucketId,
		row.BucketName,
		row.Owner,
		row.Visibility,
		row.SourceType,
		row.CreateAt,
		row.PaymentAddress,
		row.BucketStatus,
		row.Tags,
		row.ChargedReadQuota,
		row.GlobalVirtualGroupFamilyId,
		row.SpAsDelegatedAgentDisabled,
		row.Height,
	)
	return &bucket, nil
}

func (db *Db) SaveObject(height int64, objects []types.Object) error {
	if len(objects) == 0 {
		return nil
	}

	var accounts []types.Account

	objectQuery := `
INSERT INTO object(
	id, object_name, bucket_name, owner, creator, payload_size, visibility, content_type,object_status,redundancy_type,
    source_type, checksums, create_at, local_virtual_group_id, height, tags, is_updating, updated_at, updated_by, version
) VALUES`
	var objectParams []interface{}

	for i, object := range objects {
		// Prepare the account query
		accounts = append(accounts, types.NewAccount(object.Owner))

		// Prepare the object query
		vi := i * 20
		objectQuery += fmt.Sprintf("($%d,$%d,$%d,$%d,$%d,$%d,$%d,$%d,$%d,$%d,$%d,$%d,$%d,$%d,$%d,$%d,$%d,$%d,$%d,$%d),",
			vi+1, vi+2, vi+3, vi+4, vi+5, vi+6, vi+7, vi+8, vi+9, vi+10, vi+11, vi+12, vi+13, vi+14, vi+15, vi+16, vi+17, vi+18, vi+19, vi+20)

		objectParams = append(objectParams,
			object.ObjectId,
			object.ObjectName,
			object.BucketName,
			object.Owner,
			object.Creator,
			object.PayloadSize,
			object.Visibility,
			object.ContentType,
			object.ObjectStatus,
			object.RedundancyType,
			object.SourceType,
			object.Checksums,
			object.CreateAt,
			object.LocalVirtualGroupId,
			height,
			object.Tags,
			object.IsUpdating,
			object.UpdatedAt,
			object.UpdatedBy,
			object.Version,
		)
	}

	// Store the accounts
	err := db.SaveAccounts(accounts)
	if err != nil {
		return fmt.Errorf("error while storing owner accounts: %s", err)
	}

	// Store the objects
	objectQuery = objectQuery[:len(objectQuery)-1] // Remove trailing ","
	objectQuery += " ON CONFLICT DO NOTHING"
	_, err = db.SQL.Exec(objectQuery, objectParams...)
	if err != nil {
		return fmt.Errorf("error while storing objects: %s", err)
	}

	return nil
}

func (db *Db) GetObject(id uint64) (*types.Object, error) {
	var rows []*dbtypes.ObjectRow
	err := db.Sqlx.Select(&rows, `SELECT * FROM object WHERE id = $1`, id)
	if err != nil {
		return nil, err
	}

	if len(rows) == 0 {
		return nil, nil
	}

	row := rows[0]

	object := types.NewObject(
		row.ObjectId,
		row.ObjectName,
		row.BucketName,
		row.Owner,
		row.Creator,
		row.PayloadSize,
		row.Visibility,
		row.ContentType,
		row.ObjectStatus,
		row.RedundancyType,
		row.SourceType,
		row.Checksums,
		row.Tags,
		row.IsUpdating,
		row.CreateAt,
		row.UpdatedAt,
		row.UpdatedBy,
		row.Version,
		row.LocalVirtualGroupId,
		row.Height,
	)
	return &object, nil
}
