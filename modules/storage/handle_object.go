package storage

import (
	"strconv"
	"time"

	abcitypes "github.com/cometbft/cometbft/abci/types"
	"github.com/forbole/bdjuno/v4/types"
	"github.com/rs/zerolog/log"
)

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
