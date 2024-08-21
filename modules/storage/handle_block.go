package storage

import (
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
	info, extra, err := m.source.HeadBucket(190, "mechain")
	if err != nil {
		_ = info
		_ = extra
	}
	if len(txs) > 0 {
		m.parseTransactionEvents(b, txs)
	}
	return nil
}

func (m *Module) parseTransactionEvents(b *tmctypes.ResultBlock, txs []*junotypes.Tx) {
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
				m.handleCreateGroup(b.Block.Height, event.Attributes)
			case EventTypeDeleteGroup:
				m.handleDeleteGroup(b.Block.Height, event.Attributes)
			case EventTypeCreateBucket:
				m.handleCreateBucket(tx, event.Attributes)
			case EventTypeDeleteBucket:
				m.handleDeleteBucket(b.Block.Height, event.Attributes)
			case EventTypeCreateObject:
				m.handleCreateObject(b.Block.Height, event.Attributes)
			case EventTypeDeleteObject:
				m.handleDeleteObject(b.Block.Height, event.Attributes)
			}
		}
	}
}

func (m *Module) handleCreateGroup(height int64, values []abcitypes.EventAttribute) {
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
}

func (m *Module) handleDeleteGroup(height int64, values []abcitypes.EventAttribute) {
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
}

func (m *Module) handleCreateBucket(tx *junotypes.Tx, values []abcitypes.EventAttribute) {
	b := &models.Bucket{
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
			b.BucketID = common.HexToHash(strV)
		case "bucket_name":
			b.BucketName = strV
		case "owner":
			b.OwnerAddress = common.HexToAddress(strV)
		case "payment_address":
			b.PaymentAddress = common.HexToAddress(strV)
		case "global_virtual_group_family_id":
			globalVirtualGroupFamilyId, _ := strconv.ParseUint(strV, 10, 32)
			b.GlobalVirtualGroupFamilyId = uint32(globalVirtualGroupFamilyId)
		case "source_type":
			b.SourceType = strV
		case "charged_read_quota":
			chargedReadQuota, _ := strconv.ParseUint(strV, 10, 64)
			b.ChargedReadQuota = chargedReadQuota
		case "visibility":
			b.Visibility = strV
		case "status":
			b.Status = strV
		case "create_at":
			createTime, _ := strconv.ParseInt(strV, 10, 64)
			b.CreateTime = time.Unix(createTime, 0)

		}
	}
	b.OperatorAddress = b.OwnerAddress
	b.UpdateTxHash = b.CreateTxHash
	b.UpdateTime = b.CreateTime

	// 迁移 schema (自动创建表)
	dsn := "postgresql://postgres:postgres_mechain@localhost:5432/bdjuno?sslmode=disable&search_path=public"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	db.AutoMigrate(&models.Bucket{})
	if err != nil {
		log.Fatal().Str("module", "storage").Err(err)
	}
}

func (m *Module) handleDeleteBucket(height int64, values []abcitypes.EventAttribute) {
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
}

func (m *Module) handleCreateObject(height int64, values []abcitypes.EventAttribute) {
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
}

func (m *Module) handleDeleteObject(height int64, values []abcitypes.EventAttribute) {
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
}
