package storage

import (
	"context"
	"strconv"
	"time"

	abcitypes "github.com/cometbft/cometbft/abci/types"
	tmctypes "github.com/cometbft/cometbft/rpc/core/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/forbole/bdjuno/v4/database/models"
	"github.com/forbole/bdjuno/v4/types"
	junotypes "github.com/forbole/juno/v5/types"
	"github.com/rs/zerolog/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	EventTypeCreateGroup  = "greenfield.storage.EventCreateGroup"
	EventTypeDeleteGroup  = "greenfield.storage.EventDeleteGroup"
	EventTypeCreateBucket = "greenfield.storage.EventCreateBucket"
	EventTypeDeleteBucket = "greenfield.storage.EventDeleteBucket"
	EventTypeCreateObject = "greenfield.storage.EventCreateObject"
	EventTypeDeleteObject = "greenfield.storage.EventCancelCreateObject"
	EventTypeEthereumTx   = "ethereum_tx"
)

// HandleBlock implements modules.BlockModule
func (m *Module) HandleBlock(
	b *tmctypes.ResultBlock, results *tmctypes.ResultBlockResults, txs []*junotypes.Tx, vals *tmctypes.ResultValidators,
) error {
	dsn := "host=localhost user=postgres password=yourpassword dbname=yourdb port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	// 迁移 schema (自动创建表)
	db.AutoMigrate(&models.Bucket{})
	info, extra, err := m.source.HeadBucket(190, "mechain")
	if err != nil {
		_ = info
		_ = extra
	}
	if len(txs) > 0 {
		sqls, err := m.parseTransactionEvents(context.Background(), b, txs)
		if err != nil {
			return err
		}

	}

	return nil
}

func (m *Module) parseTransactionEvents(ctx context.Context, b *tmctypes.ResultBlock, txs []*junotypes.Tx) (map[string][]interface{}, error) {
	for _, tx := range txs {
		for _, event := range tx.Events {
			switch event.Type {
			case EventTypeEthereumTx:
				for _, attribute := range event.Attributes {
					if attribute.Key == "ethereumTxHash" {
						log.Debug().Str("module", "storage").Str(attribute.Key, attribute.Value)
					}
				}
			case EventTypeCreateGroup:
				return m.handleCreateGroup(b.Block.Height, event.Attributes), nil
			case EventTypeDeleteGroup:
				return m.handleDeleteGroup(b.Block.Height, event.Attributes), nil
			case EventTypeCreateBucket:
				return m.handleCreateBucket(ctx, tx, event.Attributes), nil
			case EventTypeDeleteBucket:
				return m.handleDeleteBucket(b.Block.Height, event.Attributes), nil
			case EventTypeCreateObject:
				return m.handleCreateObject(b.Block.Height, event.Attributes), nil
			case EventTypeDeleteObject:
				return m.handleDeleteObject(b.Block.Height, event.Attributes), nil
			}
		}
	}
}

func (m *Module) handleCreateGroup(height int64, values []abcitypes.EventAttribute) map[string][]interface{} {
	msg := &createGroupEvent{}
	for _, v := range values {
		strV := ""
		if len(v.Value) > 2 {
			strV = v.Value[1 : len(v.Value)-1]
		}
		switch v.Key {
		case "group_id":
			msg.GroupId = strV
		case "group_name":
			msg.GroupName = strV
		case "owner":
			msg.Owner = strV
		case "source_type":
			msg.SourceType = strV
		case "extra":
			msg.Extra = strV
		}
	}

	groupId, _ := strconv.ParseUint(msg.GroupId, 10, 64)

	err := m.db.SaveStorageGroup(height, []types.StorageGroup{
		{
			GroupId:    groupId,
			GroupName:  msg.GroupName,
			Owner:      msg.Owner,
			SourceType: msg.SourceType,
			Extra:      msg.Extra,
		},
	})
	if err != nil {
		log.Error().Str("module", "storage").Err(err)
	}
	return nil
}

func (m *Module) handleDeleteGroup(height int64, values []abcitypes.EventAttribute) map[string][]interface{} {
	msg := &deleteGroupEvent{}
	for _, v := range values {
		strV := ""
		if len(v.Value) > 2 {
			strV = v.Value[1 : len(v.Value)-1]
		}
		switch v.Key {
		case "group_id":
			msg.GroupId = strV
		case "group_name":
			msg.GroupName = strV
		case "owner":
			msg.Owner = strV
		}
	}
	return nil
}

func (m *Module) handleCreateBucket(ctx context.Context, tx *junotypes.Tx, values []abcitypes.EventAttribute) map[string][]interface{} {
	msg := &models.Bucket{
		Removed:      false,
		CreateAt:     tx.Height,
		CreateTxHash: common.HexToHash(tx.TxHash),
		UpdateAt:     tx.Height,
	}
	for _, v := range values {
		strV := ""
		if len(v.Value) > 2 {
			strV = v.Value[1 : len(v.Value)-1]
		}
		switch v.Key {
		case "bucket_id":
			msg.BucketID = common.HexToHash(strV)
		case "bucket_name":
			msg.BucketName = strV
		case "owner":
			msg.OwnerAddress = common.HexToAddress(strV)
		case "payment_address":
			msg.PaymentAddress = common.HexToAddress(strV)
		case "global_virtual_group_family_id":
			globalVirtualGroupFamilyId, _ := strconv.ParseUint(strV, 10, 32)
			msg.GlobalVirtualGroupFamilyId = uint32(globalVirtualGroupFamilyId)
		case "source_type":
			msg.SourceType = strV
		case "charged_read_quota":
			chargedReadQuota, _ := strconv.ParseUint(strV, 10, 64)
			msg.ChargedReadQuota = chargedReadQuota
		case "visibility":
			msg.Visibility = strV
		case "status":
			msg.Status = strV
		case "create_at":
			createTime, _ := strconv.ParseInt(strV, 10, 64)
			msg.CreateTime = time.Unix(createTime, 0)

		}
	}
	msg.OperatorAddress = msg.OwnerAddress
	msg.UpdateTxHash = msg.CreateTxHash
	msg.UpdateTime = msg.CreateTime

	k, v := m.db.SaveBucket(ctx, bucket)
	return map[string][]interface{}{
		k: v,
	}
}

func (m *Module) handleDeleteBucket(height int64, values []abcitypes.EventAttribute) map[string][]interface{} {
	msg := &deleteBucketEvent{}
	for _, v := range values {
		strV := ""
		if len(v.Value) > 2 {
			strV = v.Value[1 : len(v.Value)-1]
		}
		switch v.Key {
		case "bucket_id":
			msg.BucketId = strV
		case "bucket_name":
			msg.BucketName = strV
		case "owner":
			msg.Owner = strV
		case "operator":
			msg.Operator = strV
		case "global_virtual_group_family_id":
			globalVirtualGroupFamilyId, _ := strconv.ParseUint(strV, 10, 32)
			msg.GlobalVirtualGroupFamilyId = uint32(globalVirtualGroupFamilyId)
		}
	}
	return nil
}

func (m *Module) handleCreateObject(height int64, values []abcitypes.EventAttribute) map[string][]interface{} {
	msg := &createObjectEvent{}
	for _, v := range values {
		strV := ""
		if len(v.Value) > 2 {
			strV = v.Value[1 : len(v.Value)-1]
		}
		switch v.Key {
		case "object_id":
			msg.ObjectId = strV
		case "object_name":
			msg.ObjectName = strV
		case "bucket_id":
			msg.BucketId = strV
		case "bucket_name":
			msg.BucketName = strV
		case "owner":
			msg.Owner = strV
		case "creator":
			msg.Creator = strV
		case "primary_sp_id":
			primarySpId, _ := strconv.ParseUint(strV, 10, 32)
			msg.PrimarySpId = uint32(primarySpId)
		case "payload_size":
			payloadSize, _ := strconv.ParseUint(strV, 10, 64)
			msg.PayloadSize = payloadSize
		case "visibility":
			msg.Visibility = strV
		case "content_type":
			msg.ContentType = strV
		case "create_at":
			createAt, _ := strconv.ParseInt(strV, 10, 64)
			msg.CreateAt = createAt
		case "status":
			msg.Status = strV
		case "redundancy_type":
			msg.RedundancyType = strV
		case "source_type":
			msg.SourceType = strV
		case "checksum":
			msg.Checksum = strV
		case "local_virtual_group_id":
			localVirtualGroupId, _ := strconv.ParseUint(strV, 10, 32)
			msg.LocalVirtualGroupId = uint32(localVirtualGroupId)
		}
	}

	objectId, _ := strconv.ParseUint(msg.ObjectId, 10, 64)
	createAt := time.Unix(msg.CreateAt, 0)

	err := m.db.SaveObject(height, []types.Object{
		{
			ObjectId:            objectId,
			ObjectName:          msg.ObjectName,
			BucketName:          msg.BucketName,
			Owner:               msg.Owner,
			Creator:             msg.Creator,
			PayloadSize:         msg.PayloadSize,
			Visibility:          msg.Visibility,
			ContentType:         msg.ContentType,
			ObjectStatus:        msg.Status,
			RedundancyType:      msg.RedundancyType,
			SourceType:          msg.SourceType,
			Checksums:           msg.Checksum,
			CreateAt:            createAt,
			LocalVirtualGroupId: msg.LocalVirtualGroupId,
		},
	})
	if err != nil {
		log.Error().Str("module", "storage").Err(err)
	}
	return nil
}

func (m *Module) handleDeleteObject(height int64, values []abcitypes.EventAttribute) map[string][]interface{} {
	msg := &deleteObjectEvent{}
	for _, v := range values {
		strV := ""
		if len(v.Value) > 2 {
			strV = v.Value[1 : len(v.Value)-1]
		}
		switch v.Key {
		case "object_id":
			msg.ObjectId = strV
		case "object_name":
			msg.ObjectName = strV
		case "bucket_name":
			msg.BucketName = strV
		case "operator":
			msg.Operator = strV
		case "primary_sp_id":
			primarySpId, _ := strconv.ParseUint(strV, 10, 32)
			msg.PrimarySpId = uint32(primarySpId)
		}
	}
	return nil
}
