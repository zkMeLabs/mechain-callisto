package database

import (
	"fmt"

	dbtypes "github.com/forbole/bdjuno/v4/database/types"
	"github.com/forbole/bdjuno/v4/types"
)

func (db *Db) SaveGroups(groups []types.Groups) error {
	if len(groups) == 0 {
		return nil
	}

	var accounts []types.Account

	groupsQuery := `
INSERT INTO groups(
	id, group_name, owner, source_type, extra, tags
) VALUES`
	var groupsParams []interface{}

	for i, group := range groups {
		// Prepare the account query
		accounts = append(accounts, types.NewAccount(group.Owner))

		// Prepare the group query
		vi := i * 6
		groupsQuery += fmt.Sprintf("($%d,$%d,$%d,$%d,$%d,$%d),",
			vi+1, vi+2, vi+3, vi+4, vi+5, vi+6)

		groupsParams = append(groupsParams,
			group.GroupId,
			group.GroupName,
			group.Owner,
			group.SourceType,
			group.Extra,
			group.Tags,
		)
	}

	// Store the accounts
	err := db.SaveAccounts(accounts)
	if err != nil {
		return fmt.Errorf("error while storing owner accounts: %s", err)
	}

	// Store the groups
	groupsQuery = groupsQuery[:len(groupsQuery)-1] // Remove trailing ","
	groupsQuery += " ON CONFLICT DO NOTHING"
	_, err = db.SQL.Exec(groupsQuery, groupsParams...)
	if err != nil {
		return fmt.Errorf("error while storing groups: %s", err)
	}

	return nil
}

func (db *Db) GetGroups(id uint64) (*types.Groups, error) {
	var rows []*dbtypes.GroupsRow
	err := db.Sqlx.Select(&rows, `SELECT * FROM groups WHERE id = $1`, id)
	if err != nil {
		return nil, err
	}

	if len(rows) == 0 {
		return nil, nil
	}

	row := rows[0]

	groups := types.NewGroups(
		row.GroupId,
		row.GroupName,
		row.Owner,
		row.SourceType,
		row.Extra,
		row.Tags,
	)
	return &groups, nil
}

func (db *Db) SaveBuckets(buckets []types.Buckets) error {
	if len(buckets) == 0 {
		return nil
	}

	var accounts []types.Account

	bucketsQuery := `
INSERT INTO buckets(
	id, bucket_name, owner, visibility, source_type, create_at, payment_address, bucket_status, 
    charged_read_quota, global_virtual_group_family_id, sp_as_delegated_agent_disabled, tags
) VALUES`
	var bucketsParams []interface{}

	for i, bucket := range buckets {
		// Prepare the account query
		accounts = append(accounts, types.NewAccount(bucket.Owner))

		// Prepare the bucket query
		vi := i * 12
		bucketsQuery += fmt.Sprintf("($%d,$%d,$%d,$%d,$%d,$%d,$%d,$%d,$%d,$%d,$%d,$%d),",
			vi+1, vi+2, vi+3, vi+4, vi+5, vi+6, vi+7, vi+8, vi+9, vi+10, vi+11, vi+12)

		bucketsParams = append(bucketsParams,
			bucket.BucketId,
			bucket.BucketName,
			bucket.Owner,
			bucket.Visibility,
			bucket.SourceType,
			bucket.CreateAt,
			bucket.PaymentAddress,
			bucket.BucketStatus,
			bucket.ChargedReadQuota,
			bucket.GlobalVirtualGroupFamilyId,
			bucket.SpAsDelegatedAgentDisabled,
			bucket.Tags,
		)
	}

	// Store the accounts
	err := db.SaveAccounts(accounts)
	if err != nil {
		return fmt.Errorf("error while storing owner accounts: %s", err)
	}

	// Store the buckets
	bucketsQuery = bucketsQuery[:len(bucketsQuery)-1] // Remove trailing ","
	bucketsQuery += " ON CONFLICT DO NOTHING"
	_, err = db.SQL.Exec(bucketsQuery, bucketsParams...)
	if err != nil {
		return fmt.Errorf("error while storing buckets: %s", err)
	}

	return nil
}

func (db *Db) GetBuckets(id uint64) (*types.Buckets, error) {
	var rows []*dbtypes.BucketsRow
	err := db.Sqlx.Select(&rows, `SELECT * FROM buckets WHERE id = $1`, id)
	if err != nil {
		return nil, err
	}

	if len(rows) == 0 {
		return nil, nil
	}

	row := rows[0]

	bucket := types.NewBuckets(
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
	)
	return &bucket, nil
}

func (db *Db) SaveObjects(objects []types.Objects) error {
	if len(objects) == 0 {
		return nil
	}

	var accounts []types.Account

	objectsQuery := `
INSERT INTO objects(
	id, object_name, bucket_name, owner, creator, payload_size, visibility, content_type,object_status,redundancy_type,
    source_type, checksums, create_at, local_virtual_group_id, tags, is_updating, updated_at, updated_by, version
) VALUES`
	var objectsParams []interface{}

	for i, object := range objects {
		// Prepare the account query
		accounts = append(accounts, types.NewAccount(object.Owner))

		// Prepare the object query
		vi := i * 19
		objectsQuery += fmt.Sprintf("($%d,$%d,$%d,$%d,$%d,$%d,$%d,$%d,$%d,$%d,$%d,$%d,$%d,$%d,$%d,$%d,$%d,$%d,$%d),",
			vi+1, vi+2, vi+3, vi+4, vi+5, vi+6, vi+7, vi+8, vi+9, vi+10, vi+11, vi+12, vi+13, vi+14, vi+15, vi+16, vi+17, vi+18, vi+19)

		objectsParams = append(objectsParams,
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
	objectsQuery = objectsQuery[:len(objectsQuery)-1] // Remove trailing ","
	objectsQuery += " ON CONFLICT DO NOTHING"
	_, err = db.SQL.Exec(objectsQuery, objectsParams...)
	if err != nil {
		return fmt.Errorf("error while storing objects: %s", err)
	}

	return nil
}

func (db *Db) GetObjects(id uint64) (*types.Objects, error) {
	var rows []*dbtypes.ObjectsRow
	err := db.Sqlx.Select(&rows, `SELECT * FROM objects WHERE id = $1`, id)
	if err != nil {
		return nil, err
	}

	if len(rows) == 0 {
		return nil, nil
	}

	row := rows[0]

	object := types.NewObjects(
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
	)
	return &object, nil
}
