package storage

import (
	"strconv"
	"time"

	abcitypes "github.com/cometbft/cometbft/abci/types"
	tmctypes "github.com/cometbft/cometbft/rpc/core/types"
	"github.com/forbole/bdjuno/v4/types"
	juno "github.com/forbole/juno/v5/types"
	"github.com/rs/zerolog/log"
)

const (
	EventTypeCreateGroup  = "greenfield.storage.EventCreateGroup"
	EventTypeDeleteGroup  = "greenfield.storage.EventDeleteGroup"
	EventTypeCreateBucket = "greenfield.storage.EventCreateBucket"
	EventTypeDeleteBucket = "greenfield.storage.EventDeleteBucket"
	EventTypeCreateObject = "greenfield.storage.EventCreateObject"
	EventTypeDeleteObject = "greenfield.storage.EventCancelCreateObject"
)

// HandleBlock implements modules.BlockModule
func (m *Module) HandleBlock(
	b *tmctypes.ResultBlock, results *tmctypes.ResultBlockResults, txs []*juno.Tx, vals *tmctypes.ResultValidators,
) error {

	if len(txs) > 0 {
		m.parseTransactionEvents(b, txs)
	}

	return nil
}
func (m *Module) parseTransactionEvents(b *tmctypes.ResultBlock, txs []*juno.Tx) {
	log.Debug().Str("module", "distribution").Int64("height", b.Block.Height)
	for _, tx := range txs {
		for _, event := range tx.Events {
			switch event.Type {
			case EventTypeCreateGroup:
				m.handleCreateGroup(b.Block.Height, event.Attributes)
			case EventTypeDeleteGroup:
				m.handleDeleteGroup(b.Block.Height, event.Attributes)
			case EventTypeCreateBucket:
				m.handleCreateBucket(b.Block.Height, event.Attributes)
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
		switch v.Key {
		case "group_id":
			msg.GroupId = v.Value[1 : len(v.Value)-1]
		case "group_name":
			msg.GroupName = v.Value
		case "owner":
			msg.Owner = v.Value
		case "source_type":
			msg.SourceType = v.Value
		case "extra":
			msg.Extra = v.Value
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
		switch v.Key {
		case "group_id":
			msg.GroupId = v.Value
		case "group_name":
			msg.GroupName = v.Value
		case "owner":
			msg.Owner = v.Value
		}
	}
}

func (m *Module) handleCreateBucket(height int64, values []abcitypes.EventAttribute) {
	msg := &createBucketEvent{}
	for _, v := range values {
		switch v.Key {
		case "bucket_id":
			msg.BucketId = v.Value[1 : len(v.Value)-1]
		case "bucket_name":
			msg.BucketName = v.Value
		case "owner":
			msg.Owner = v.Value
		case "visibility":
			msg.Visibility = v.Value
		case "create_at":
			createAt, err := strconv.ParseInt(v.Value, 10, 64)
			if err != nil {
				createAt, _ = strconv.ParseInt(v.Value[1:len(v.Value)-1], 10, 64)
			}
			msg.CreateAt = createAt
		case "source_type":
			msg.SourceType = v.Value
		case "charged_read_quota":
			chargedReadQuota, err := strconv.ParseUint(v.Value, 10, 64)
			if err != nil {
				chargedReadQuota, _ = strconv.ParseUint(v.Value[1:len(v.Value)-1], 10, 64)
			}
			msg.ChargedReadQuota = chargedReadQuota
		case "payment_address":
			msg.PaymentAddress = v.Value
		case "primary_sp_id":
			primarySpId, err := strconv.ParseUint(v.Value, 10, 32)
			if err != nil {
				primarySpId, _ = strconv.ParseUint(v.Value[1:len(v.Value)-1], 10, 32)
			}
			msg.PrimarySpId = uint32(primarySpId)
		case "global_virtual_group_family_id":
			globalVirtualGroupFamilyId, err := strconv.ParseUint(v.Value, 10, 32)
			if err != nil {
				globalVirtualGroupFamilyId, _ = strconv.ParseUint(v.Value[1:len(v.Value)-1], 10, 32)
			}
			msg.GlobalVirtualGroupFamilyId = uint32(globalVirtualGroupFamilyId)
		case "status":
			msg.Status = v.Value
		}
	}
	bucketId, _ := strconv.ParseUint(msg.BucketId, 10, 64)
	createAt := time.Unix(msg.CreateAt, 0)

	err := m.db.SaveBucket(height, []types.Bucket{
		{
			BucketId:                   bucketId,
			BucketName:                 msg.BucketName,
			Owner:                      msg.Owner,
			Visibility:                 msg.Visibility,
			SourceType:                 msg.SourceType,
			CreateAt:                   createAt,
			PaymentAddress:             msg.PaymentAddress,
			BucketStatus:               msg.Status,
			ChargedReadQuota:           msg.ChargedReadQuota,
			GlobalVirtualGroupFamilyId: msg.GlobalVirtualGroupFamilyId,
		},
	})
	if err != nil {
		log.Error().Str("module", "storage").Err(err)
	}
}

func (m *Module) handleDeleteBucket(height int64, values []abcitypes.EventAttribute) {
	msg := &deleteBucketEvent{}
	for _, v := range values {
		switch v.Key {
		case "bucket_id":
			msg.BucketId = v.Value
		case "bucket_name":
			msg.BucketName = v.Value
		case "owner":
			msg.Owner = v.Value
		case "operator":
			msg.Operator = v.Value
		case "global_virtual_group_family_id":
			globalVirtualGroupFamilyId, err := strconv.ParseUint(v.Value, 10, 32)
			if err != nil {
				globalVirtualGroupFamilyId, _ = strconv.ParseUint(v.Value[1:len(v.Value)-1], 10, 32)
			}
			msg.GlobalVirtualGroupFamilyId = uint32(globalVirtualGroupFamilyId)
		}
	}
}

func (m *Module) handleCreateObject(height int64, values []abcitypes.EventAttribute) {
	msg := &createObjectEvent{}
	for _, v := range values {
		switch v.Key {
		case "object_id":
			msg.ObjectId = v.Value[1 : len(v.Value)-1]
		case "object_name":
			msg.ObjectName = v.Value
		case "bucket_id":
			msg.BucketId = v.Value
		case "bucket_name":
			msg.BucketName = v.Value
		case "owner":
			msg.Owner = v.Value
		case "creator":
			msg.Creator = v.Value
		case "primary_sp_id":
			primarySpId, err := strconv.ParseUint(v.Value, 10, 32)
			if err != nil {
				primarySpId, _ = strconv.ParseUint(v.Value[1:len(v.Value)-1], 10, 32)
			}
			msg.PrimarySpId = uint32(primarySpId)
		case "payload_size":
			payloadSize, err := strconv.ParseUint(v.Value, 10, 64)
			if err != nil {
				payloadSize, _ = strconv.ParseUint(v.Value[1:len(v.Value)-1], 10, 64)
			}
			msg.PayloadSize = payloadSize
		case "visibility":
			msg.Visibility = v.Value
		case "content_type":
			msg.ContentType = v.Value
		case "create_at":
			createAt, err := strconv.ParseInt(v.Value, 10, 64)
			if err != nil {
				createAt, _ = strconv.ParseInt(v.Value[1:len(v.Value)-1], 10, 64)
			}
			msg.CreateAt = createAt
		case "status":
			msg.Status = v.Value
		case "redundancy_type":
			msg.RedundancyType = v.Value
		case "source_type":
			msg.SourceType = v.Value
		case "checksum":
			msg.Checksum = v.Value
		case "local_virtual_group_id":
			localVirtualGroupId, err := strconv.ParseUint(v.Value, 10, 32)
			if err != nil {
				localVirtualGroupId, _ = strconv.ParseUint(v.Value[1:len(v.Value)-1], 10, 32)
			}
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
		switch v.Key {
		case "object_id":
			msg.ObjectId = v.Value
		case "object_name":
			msg.ObjectName = v.Value
		case "bucket_name":
			msg.BucketName = v.Value
		case "operator":
			msg.Operator = v.Value
		case "primary_sp_id":
			primarySpId, _ := strconv.ParseUint(v.Value, 10, 32)
			msg.PrimarySpId = uint32(primarySpId)
		}
	}
}
