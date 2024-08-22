package storage

import (
	"context"
	"errors"
	"strings"
	"time"

	abci "github.com/cometbft/cometbft/abci/types"
	tmctypes "github.com/cometbft/cometbft/rpc/core/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/gogoproto/proto"
	"github.com/forbole/bdjuno/v4/database/models"
	storagetypes "github.com/forbole/bdjuno/v4/modules/storage/types"
)

var (
	EventCreateObject               = proto.MessageName(&storagetypes.EventCreateObject{})
	EventCancelCreateObject         = proto.MessageName(&storagetypes.EventCancelCreateObject{})
	EventSealObject                 = proto.MessageName(&storagetypes.EventSealObject{})
	EventCopyObject                 = proto.MessageName(&storagetypes.EventCopyObject{})
	EventDeleteObject               = proto.MessageName(&storagetypes.EventDeleteObject{})
	EventRejectSealObject           = proto.MessageName(&storagetypes.EventRejectSealObject{})
	EventDiscontinueObject          = proto.MessageName(&storagetypes.EventDiscontinueObject{})
	EventUpdateObjectInfo           = proto.MessageName(&storagetypes.EventUpdateObjectInfo{})
	EventUpdateObjectContent        = proto.MessageName(&storagetypes.EventUpdateObjectContent{})
	EventUpdateObjectContentSuccess = proto.MessageName(&storagetypes.EventUpdateObjectContentSuccess{})
	EventCancelUpdateObjectContent  = proto.MessageName(&storagetypes.EventCancelUpdateObjectContent{})
)

var ObjectEvents = map[string]bool{
	EventCreateObject:               true,
	EventCancelCreateObject:         true,
	EventSealObject:                 true,
	EventCopyObject:                 true,
	EventDeleteObject:               true,
	EventRejectSealObject:           true,
	EventDiscontinueObject:          true,
	EventUpdateObjectInfo:           true,
	EventUpdateObjectContent:        true,
	EventUpdateObjectContentSuccess: true,
	EventCancelUpdateObjectContent:  true,
}

func (m *Module) ExtractObjectEventStatements(ctx context.Context, block *tmctypes.ResultBlock, txHash, evmTxHash string, event sdk.Event) (map[string][]interface{}, error) {
	typedEvent, err := sdk.ParseTypedEvent(abci.Event(event))
	if err != nil {
		return nil, err
	}

	switch event.Type {
	case EventCreateObject:
		createObject, ok := typedEvent.(*storagetypes.EventCreateObject)
		if !ok {
			return nil, errors.New("create object event assert error")
		}
		return m.handleCreateObject(ctx, block, txHash, evmTxHash, createObject), nil
	case EventCancelCreateObject:
		cancelCreateObject, ok := typedEvent.(*storagetypes.EventCancelCreateObject)
		if !ok {
			return nil, errors.New("cancel create object event assert error")
		}
		return m.handleCancelCreateObject(ctx, block, txHash, evmTxHash, cancelCreateObject), nil
	case EventSealObject:
		sealObject, ok := typedEvent.(*storagetypes.EventSealObject)
		if !ok {
			return nil, errors.New("seal object event assert error")
		}
		return m.handleSealObject(ctx, block, txHash, evmTxHash, sealObject), nil
	case EventCopyObject:
		copyObject, ok := typedEvent.(*storagetypes.EventCopyObject)
		if !ok {
			return nil, errors.New("copy object event assert error")
		}
		return m.handleCopyObject(ctx, block, txHash, evmTxHash, copyObject)
	case EventDeleteObject:
		deleteObject, ok := typedEvent.(*storagetypes.EventDeleteObject)
		if !ok {
			return nil, errors.New("delete object event assert error")
		}
		return m.handleDeleteObject(ctx, block, txHash, evmTxHash, deleteObject), nil
	case EventRejectSealObject:
		rejectSealObject, ok := typedEvent.(*storagetypes.EventRejectSealObject)
		if !ok {
			return nil, errors.New("reject seal object event assert error")
		}
		return m.handleRejectSealObject(ctx, block, txHash, evmTxHash, rejectSealObject), nil
	case EventDiscontinueObject:
		discontinueObject, ok := typedEvent.(*storagetypes.EventDiscontinueObject)
		if !ok {
			return nil, errors.New("discontinue object event assert error")
		}
		return m.handleEventDiscontinueObject(ctx, block, txHash, evmTxHash, discontinueObject), nil
	case EventUpdateObjectInfo:
		updateObjectInfo, ok := typedEvent.(*storagetypes.EventUpdateObjectInfo)
		if !ok {
			return nil, errors.New("update object event assert error")
		}
		return m.handleUpdateObjectInfo(ctx, block, txHash, evmTxHash, updateObjectInfo), nil
	case EventUpdateObjectContent:
		updateObjectContent, ok := typedEvent.(*storagetypes.EventUpdateObjectContent)
		if !ok {
			return nil, errors.New("update object event assert error")
		}
		return m.handleUpdateObjectContent(ctx, block, txHash, evmTxHash, updateObjectContent), nil
	case EventUpdateObjectContentSuccess:
		updateObjectContent, ok := typedEvent.(*storagetypes.EventUpdateObjectContentSuccess)
		if !ok {
			return nil, errors.New("update object success event assert error")
		}
		return m.handleUpdateObjectContentSuccess(ctx, block, txHash, evmTxHash, updateObjectContent), nil
	case EventCancelUpdateObjectContent:
		cancelUpdateObjectContent, ok := typedEvent.(*storagetypes.EventCancelUpdateObjectContent)
		if !ok {
			return nil, errors.New("cancel update object event assert error")
		}
		return m.handleCancelUpdateObjectContent(ctx, block, txHash, evmTxHash, cancelUpdateObjectContent), nil
	}
	return nil, nil
}

func convertByteArrayToStrArray(ba [][]byte) []string {
	sa := make([]string, len(ba))
	for i, b := range ba {
		sa[i] = string(b)
	}
	return sa
}

func (m *Module) handleCreateObject(ctx context.Context, block *tmctypes.ResultBlock, txHash, evmTxHash string, createObject *storagetypes.EventCreateObject) map[string][]interface{} {
	object := &models.Object{
		BucketID:        createObject.BucketId.BigInt().String(),
		BucketName:      createObject.BucketName,
		ObjectID:        createObject.ObjectId.BigInt().String(),
		ObjectName:      createObject.ObjectName,
		CreatorAddress:  createObject.Creator,
		OwnerAddress:    createObject.Owner,
		PayloadSize:     createObject.PayloadSize,
		Visibility:      createObject.Visibility.String(),
		ContentType:     createObject.ContentType,
		Status:          createObject.Status.String(),
		RedundancyType:  createObject.RedundancyType.String(),
		SourceType:      createObject.SourceType.String(),
		CheckSums:       convertByteArrayToStrArray(createObject.Checksums),
		CreateTxHash:    txHash,
		CreateEVMTxHash: evmTxHash,
		CreateAt:        block.Block.Height,
		CreateTime:      time.Unix(createObject.CreateAt, 0),
		UpdateAt:        block.Block.Height,
		UpdateTxHash:    txHash,
		UpdateEVMTxHash: evmTxHash,
		UpdateTime:      time.Unix(createObject.CreateAt, 0),
		Removed:         false,
	}

	res := make(map[string][]interface{})

	k, v := m.db.SaveObjectToSQL(ctx, object)
	res[k] = v

	return res
}

func (m *Module) handleSealObject(ctx context.Context, block *tmctypes.ResultBlock, txHash, evmTxHash string, sealObject *storagetypes.EventSealObject) map[string][]interface{} {
	object := &models.Object{
		BucketName:          sealObject.BucketName,
		ObjectName:          sealObject.ObjectName,
		ObjectID:            sealObject.ObjectId.BigInt().String(),
		OperatorAddress:     sealObject.Operator,
		LocalVirtualGroupID: sealObject.LocalVirtualGroupId,
		Status:              sealObject.Status.String(),
		SealedTxHash:        txHash,
		SealedEvmTxHash:     evmTxHash,
		CheckSums:           convertByteArrayToStrArray(sealObject.GetChecksums()),
		UpdateAt:            block.Block.Height,
		UpdateTxHash:        txHash,
		UpdateTime:          block.Block.Time,
		Removed:             false,
	}

	res := make(map[string][]interface{})

	k, v := m.db.UpdateStorageSizeToSQL(ctx, sealObject.ObjectId.BigInt().String(), sealObject.BucketName, "+")
	res[k] = v

	k, v = m.db.UpdateChargeSizeToSQL(ctx, sealObject.ObjectId.BigInt().String(), sealObject.BucketName, "+")
	res[k] = v

	k, v = m.db.UpdateObjectToSQL(ctx, object)
	res[k] = v

	return res
}

func (m *Module) handleCancelCreateObject(ctx context.Context, block *tmctypes.ResultBlock, txHash, evmTxHash string, cancelCreateObject *storagetypes.EventCancelCreateObject) map[string][]interface{} {
	object := &models.Object{
		BucketName:      cancelCreateObject.BucketName,
		ObjectName:      cancelCreateObject.ObjectName,
		ObjectID:        cancelCreateObject.ObjectId.BigInt().String(),
		OperatorAddress: cancelCreateObject.Operator,
		UpdateAt:        block.Block.Height,
		UpdateTxHash:    txHash,
		UpdateEVMTxHash: evmTxHash,
		UpdateTime:      block.Block.Time,
		Removed:         true,
	}

	res := make(map[string][]interface{})

	k, v := m.db.UpdateObjectToSQL(ctx, object)
	res[k] = v

	return res
}

func (m *Module) handleCopyObject(ctx context.Context, block *tmctypes.ResultBlock, txHash, evmTxHash string, copyObject *storagetypes.EventCopyObject) (map[string][]interface{}, error) {
	destObject, err := m.db.GetObject(ctx, copyObject.SrcObjectId.BigInt().String())
	if err != nil {
		return nil, err
	}

	destObject.ObjectID = copyObject.DstObjectId.BigInt().String()
	destObject.ObjectName = copyObject.DstObjectName
	destObject.BucketName = copyObject.DstBucketName
	destObject.OperatorAddress = copyObject.Operator
	destObject.CreateAt = block.Block.Height
	destObject.CreateTxHash = txHash
	destObject.CreateTime = block.Block.Time
	destObject.UpdateAt = block.Block.Height
	destObject.UpdateTxHash = txHash
	destObject.UpdateEVMTxHash = evmTxHash
	destObject.UpdateTime = block.Block.Time
	destObject.Removed = false
	if destObject.PayloadSize == 0 {
		destObject.Status = storagetypes.OBJECT_STATUS_SEALED.String()
	} else {
		destObject.Status = storagetypes.OBJECT_STATUS_CREATED.String()
	}

	res := make(map[string][]interface{})

	k, v := m.db.SaveObjectToSQL(ctx, destObject)
	res[k] = v

	if destObject.PayloadSize == 0 {
		k, v = m.db.UpdateChargeSizeToSQL(ctx, copyObject.DstObjectId.BigInt().String(), copyObject.DstBucketName, "+")
		res[k] = v
	}

	return res, nil
}

func (m *Module) handleDeleteObject(ctx context.Context, block *tmctypes.ResultBlock, txHash, evmTxHash string, deleteObject *storagetypes.EventDeleteObject) map[string][]interface{} {
	object := &models.Object{
		BucketName:          deleteObject.BucketName,
		ObjectName:          deleteObject.ObjectName,
		ObjectID:            deleteObject.ObjectId.BigInt().String(),
		LocalVirtualGroupID: deleteObject.LocalVirtualGroupId,
		UpdateAt:            block.Block.Height,
		UpdateTxHash:        txHash,
		UpdateEVMTxHash:     evmTxHash,
		UpdateTime:          block.Block.Time,
		Removed:             true,
	}

	res := make(map[string][]interface{})
	k, v := m.db.UpdateStorageSizeToSQL(ctx, deleteObject.ObjectId.BigInt().String(), deleteObject.BucketName, "-")
	res[k] = v

	k, v = m.db.UpdateChargeSizeToSQL(ctx, deleteObject.ObjectId.BigInt().String(), deleteObject.BucketName, "-")
	res[k] = v

	k, v = m.db.UpdateObjectToSQL(ctx, object)
	res[k] = v
	return res
}

// RejectSeal event won't emit a delete event, need to be deleted manually here in metadata service
// handle logic is set as removed, no need to set status
func (m *Module) handleRejectSealObject(ctx context.Context, block *tmctypes.ResultBlock, txHash, evmTxHash string, rejectSealObject *storagetypes.EventRejectSealObject) map[string][]interface{} {
	object := &models.Object{
		BucketName:      rejectSealObject.BucketName,
		ObjectName:      rejectSealObject.ObjectName,
		ObjectID:        rejectSealObject.ObjectId.BigInt().String(),
		OperatorAddress: rejectSealObject.Operator,
		UpdateAt:        block.Block.Height,
		UpdateTxHash:    txHash,
		UpdateEVMTxHash: evmTxHash,
		UpdateTime:      block.Block.Time,
	}
	if rejectSealObject.ForUpdate {
		object.IsUpdating = false
	} else {
		object.Removed = true
	}
	k, v := m.db.UpdateObjectToSQL(ctx, object)
	return map[string][]interface{}{
		k: v,
	}
}

func (m *Module) handleEventDiscontinueObject(ctx context.Context, block *tmctypes.ResultBlock, txHash, evmTxHash string, discontinueObject *storagetypes.EventDiscontinueObject) map[string][]interface{} {
	object := &models.Object{
		BucketName:      discontinueObject.BucketName,
		ObjectID:        discontinueObject.ObjectId.BigInt().String(),
		DeleteReason:    discontinueObject.Reason,
		DeleteAt:        discontinueObject.DeleteAt,
		Status:          storagetypes.OBJECT_STATUS_DISCONTINUED.String(),
		UpdateAt:        block.Block.Height,
		UpdateTxHash:    txHash,
		UpdateEVMTxHash: evmTxHash,
		UpdateTime:      block.Block.Time,
		Removed:         false,
	}

	k, v := m.db.UpdateObjectToSQL(ctx, object)
	return map[string][]interface{}{
		k: v,
	}
}

func (m *Module) handleUpdateObjectInfo(ctx context.Context, block *tmctypes.ResultBlock, txHash, evmTxHash string, updateObject *storagetypes.EventUpdateObjectInfo) map[string][]interface{} {
	object := &models.Object{
		BucketName:      updateObject.BucketName,
		ObjectID:        updateObject.ObjectId.BigInt().String(),
		ObjectName:      updateObject.ObjectName,
		OperatorAddress: updateObject.Operator,
		Visibility:      updateObject.Visibility.String(),
		UpdateAt:        block.Block.Height,
		UpdateTxHash:    txHash,
		UpdateEVMTxHash: evmTxHash,
		UpdateTime:      block.Block.Time,
	}

	k, v := m.db.UpdateObjectToSQL(ctx, object)
	return map[string][]interface{}{
		k: v,
	}
}

func (m *Module) handleUpdateObjectContent(ctx context.Context, block *tmctypes.ResultBlock, txHash, evmTxHash string, updateObject *storagetypes.EventUpdateObjectContent) map[string][]interface{} {
	object := &models.Object{
		BucketName:      updateObject.BucketName,
		ObjectID:        updateObject.ObjectId.BigInt().String(),
		ObjectName:      updateObject.ObjectName,
		OperatorAddress: updateObject.Operator,
		UpdateAt:        block.Block.Height,
		UpdateTxHash:    txHash,
		UpdateEVMTxHash: evmTxHash,
		UpdateTime:      block.Block.Time,

		IsUpdating: true,
	}

	k, v := m.db.UpdateObjectToSQL(ctx, object)
	return map[string][]interface{}{
		k: v,
	}
}

// handleUpdateObjectContentSuccess, when sealing an updated object, EventUpdateObjectContentSuccess will be emitted before EventSealObjet.
func (m *Module) handleUpdateObjectContentSuccess(ctx context.Context, block *tmctypes.ResultBlock, txHash, evmTxHash string, updateObject *storagetypes.EventUpdateObjectContentSuccess) map[string][]interface{} {
	object := &models.Object{
		BucketName:         updateObject.BucketName,
		ObjectID:           updateObject.ObjectId.BigInt().String(),
		ObjectName:         updateObject.ObjectName,
		OperatorAddress:    updateObject.Operator,
		UpdateAt:           block.Block.Height,
		UpdateTxHash:       txHash,
		UpdateEVMTxHash:    evmTxHash,
		UpdateTime:         block.Block.Time,
		ContentType:        updateObject.ContentType,
		IsUpdating:         false,
		ContentUpdatedTime: updateObject.UpdatedAt,
		Updater:            updateObject.Operator,
		PayloadSize:        updateObject.NewPayloadSize,
		CheckSums:          convertByteArrayToStrArray(updateObject.NewChecksums),
		Version:            updateObject.Version,
	}
	res := make(map[string][]interface{})
	vars := make([]interface{}, 0)
	// deduct the charged size of previous object.
	k1, v1 := m.db.UpdateStorageSizeToSQL(ctx, updateObject.ObjectId.BigInt().String(), updateObject.BucketName, "-")
	vars = append(vars, v1...)
	k2, v2 := m.db.UpdateChargeSizeToSQL(ctx, updateObject.ObjectId.BigInt().String(), updateObject.BucketName, "-")
	vars = append(vars, v2...)
	k3, v3 := m.db.UpdateObjectToSQL(ctx, object)
	vars = append(vars, v3...)
	k := strings.Join([]string{k1, k2, k3}, "; ")
	res[k] = vars
	return res
}

func (m *Module) handleCancelUpdateObjectContent(ctx context.Context, block *tmctypes.ResultBlock, txHash, evmTxHash string, cancelUpdateObject *storagetypes.EventCancelUpdateObjectContent) map[string][]interface{} {
	object := &models.Object{
		BucketName:      cancelUpdateObject.BucketName,
		ObjectName:      cancelUpdateObject.ObjectName,
		ObjectID:        cancelUpdateObject.ObjectId.BigInt().String(),
		OperatorAddress: cancelUpdateObject.Operator,
		UpdateAt:        block.Block.Height,
		UpdateTxHash:    txHash,
		UpdateEVMTxHash: evmTxHash,
		UpdateTime:      block.Block.Time,

		IsUpdating: false,
	}
	res := make(map[string][]interface{})
	k, v := m.db.UpdateObjectToSQL(ctx, object)
	res[k] = v

	return res
}
