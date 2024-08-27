package storage

import (
	"context"
	"errors"
	"time"

	abci "github.com/cometbft/cometbft/abci/types"
	tmctypes "github.com/cometbft/cometbft/rpc/core/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/gogoproto/proto"
	"github.com/forbole/bdjuno/v4/database/models"
	storagetypes "github.com/forbole/bdjuno/v4/modules/storage/types"
)

var (
	EventCreateBucket             = proto.MessageName(&storagetypes.EventCreateBucket{})
	EventDeleteBucket             = proto.MessageName(&storagetypes.EventDeleteBucket{})
	EventUpdateBucketInfo         = proto.MessageName(&storagetypes.EventUpdateBucketInfo{})
	EventDiscontinueBucket        = proto.MessageName(&storagetypes.EventDiscontinueBucket{})
	EventMigrationBucket          = proto.MessageName(&storagetypes.EventMigrationBucket{})
	EventCancelMigrationBucket    = proto.MessageName(&storagetypes.EventCancelMigrationBucket{})
	EventRejectMigrateBucket      = proto.MessageName(&storagetypes.EventRejectMigrateBucket{})
	EventCompleteMigrationBucket  = proto.MessageName(&storagetypes.EventCompleteMigrationBucket{})
	EventToggleSPAsDelegatedAgent = proto.MessageName(&storagetypes.EventToggleSPAsDelegatedAgent{})
)

var BucketEvents = map[string]bool{
	EventCreateBucket:             true,
	EventDeleteBucket:             true,
	EventUpdateBucketInfo:         true,
	EventDiscontinueBucket:        true,
	EventMigrationBucket:          true,
	EventCancelMigrationBucket:    true,
	EventRejectMigrateBucket:      true,
	EventCompleteMigrationBucket:  true,
	EventToggleSPAsDelegatedAgent: true,
}

func (m *Module) ExtractBucketEventStatements(ctx context.Context, block *tmctypes.ResultBlock, txHash, evmTxHash string, event sdk.Event) (map[string][]interface{}, error) {
	typedEvent, err := sdk.ParseTypedEvent(abci.Event(event))
	if err != nil {
		return nil, err
	}

	switch event.Type {
	case EventCreateBucket:
		createBucket, ok := typedEvent.(*storagetypes.EventCreateBucket)
		if !ok {
			return nil, errors.New("create bucket event assert error")
		}
		return m.handleCreateBucket(ctx, block, txHash, evmTxHash, createBucket), nil
	case EventDeleteBucket:
		deleteBucket, ok := typedEvent.(*storagetypes.EventDeleteBucket)
		if !ok {
			return nil, errors.New("delete bucket event assert error")
		}
		return m.handleDeleteBucket(ctx, block, txHash, evmTxHash, deleteBucket), nil
	case EventUpdateBucketInfo:
		updateBucketInfo, ok := typedEvent.(*storagetypes.EventUpdateBucketInfo)
		if !ok {
			return nil, errors.New("update bucket event assert error")
		}
		return m.handleUpdateBucketInfo(ctx, block, txHash, evmTxHash, updateBucketInfo), nil
	case EventDiscontinueBucket:
		discontinueBucket, ok := typedEvent.(*storagetypes.EventDiscontinueBucket)
		if !ok {
			return nil, errors.New("discontinue bucket event assert error")
		}
		return m.handleDiscontinueBucket(ctx, block, txHash, evmTxHash, discontinueBucket), nil
	case EventMigrationBucket:
		migrationBucket, ok := typedEvent.(*storagetypes.EventMigrationBucket)
		if !ok {
			return nil, errors.New("migration bucket event assert error")
		}
		return m.handleEventMigrationBucket(ctx, block, txHash, evmTxHash, migrationBucket), nil
	case EventCancelMigrationBucket:
		cancelMigrationBucket, ok := typedEvent.(*storagetypes.EventCancelMigrationBucket)
		if !ok {
			return nil, errors.New("cancel migration bucket event assert error")
		}
		return m.handleEventCancelMigrationBucket(ctx, block, txHash, evmTxHash, cancelMigrationBucket), nil
	case EventRejectMigrateBucket:
		rejectMigrateBucket, ok := typedEvent.(*storagetypes.EventRejectMigrateBucket)
		if !ok {
			return nil, errors.New("reject migration bucket event assert error")
		}
		return m.handleEventRejectMigrateBucket(ctx, block, txHash, evmTxHash, rejectMigrateBucket), nil
	case EventCompleteMigrationBucket:
		completeMigrationBucket, ok := typedEvent.(*storagetypes.EventCompleteMigrationBucket)
		if !ok {
			return nil, errors.New("complete migrate bucket event assert error")
		}
		return m.handleCompleteMigrationBucket(ctx, block, txHash, evmTxHash, completeMigrationBucket), nil
	}

	return nil, nil
}

func (m *Module) handleCreateBucket(ctx context.Context, block *tmctypes.ResultBlock, txHash, evmTxHash string, createBucket *storagetypes.EventCreateBucket) map[string][]interface{} {
	b := &models.Bucket{
		BucketID:                   createBucket.BucketId.BigInt().String(),
		BucketName:                 createBucket.BucketName,
		OwnerAddress:               createBucket.Owner,
		PaymentAddress:             createBucket.PaymentAddress,
		GlobalVirtualGroupFamilyId: createBucket.GlobalVirtualGroupFamilyId,
		OperatorAddress:            createBucket.Owner,
		SourceType:                 createBucket.SourceType.String(),
		ChargedReadQuota:           createBucket.ChargedReadQuota,
		Visibility:                 createBucket.Visibility.String(),
		Status:                     createBucket.Status.String(),
		Removed:                    false,
		CreateAt:                   block.Block.Height,
		CreateTxHash:               txHash,
		CreateEVMTxHash:            evmTxHash,
		CreateTime:                 time.Unix(createBucket.CreateAt, 0),
		UpdateAt:                   block.Block.Height,
		UpdateTxHash:               txHash,
		UpdateEVMTxHash:            evmTxHash,
		UpdateTime:                 block.Block.Time,
	}
	k, v := m.db.SaveBucketToSQL(ctx, b)
	ek, ev := m.db.SaveBucketEventToSQL(ctx, b.ToBucketEvent(EventCreateBucket))
	return map[string][]interface{}{
		k:  v,
		ek: ev,
	}
}

func (m *Module) handleDeleteBucket(ctx context.Context, block *tmctypes.ResultBlock, txHash, evmTxHash string, deleteBucket *storagetypes.EventDeleteBucket) map[string][]interface{} {
	b := &models.Bucket{
		BucketID:                   deleteBucket.BucketId.BigInt().String(),
		BucketName:                 deleteBucket.BucketName,
		OwnerAddress:               deleteBucket.Owner,
		GlobalVirtualGroupFamilyId: deleteBucket.GlobalVirtualGroupFamilyId,
		Removed:                    true,
		UpdateAt:                   block.Block.Height,
		UpdateTxHash:               txHash,
		UpdateEVMTxHash:            evmTxHash,
		UpdateTime:                 block.Block.Time,
	}
	k, v := m.db.UpdateBucketToSQL(ctx, b)
	ek, ev := m.db.SaveBucketEventToSQL(ctx, b.ToBucketEvent(EventDeleteBucket))
	return map[string][]interface{}{
		k:  v,
		ek: ev,
	}
}

func (m *Module) handleDiscontinueBucket(ctx context.Context, block *tmctypes.ResultBlock, txHash, evmTxHash string, discontinueBucket *storagetypes.EventDiscontinueBucket) map[string][]interface{} {
	b := &models.Bucket{
		BucketID:        discontinueBucket.BucketId.BigInt().String(),
		BucketName:      discontinueBucket.BucketName,
		DeleteReason:    discontinueBucket.Reason,
		DeleteAt:        discontinueBucket.DeleteAt,
		Status:          storagetypes.BUCKET_STATUS_DISCONTINUED.String(),
		UpdateAt:        block.Block.Height,
		UpdateTxHash:    txHash,
		UpdateEVMTxHash: evmTxHash,
		UpdateTime:      block.Block.Time,
	}
	k, v := m.db.UpdateBucketToSQL(ctx, b)
	ek, ev := m.db.SaveBucketEventToSQL(ctx, b.ToBucketEvent(EventDiscontinueBucket))
	return map[string][]interface{}{
		k:  v,
		ek: ev,
	}
}

func (m *Module) handleUpdateBucketInfo(ctx context.Context, block *tmctypes.ResultBlock, txHash, evmTxHash string, updateBucket *storagetypes.EventUpdateBucketInfo) map[string][]interface{} {
	b := &models.Bucket{
		BucketName:                 updateBucket.BucketName,
		BucketID:                   updateBucket.BucketId.BigInt().String(),
		ChargedReadQuota:           updateBucket.ChargedReadQuota,
		PaymentAddress:             updateBucket.PaymentAddress,
		Visibility:                 updateBucket.Visibility.String(),
		GlobalVirtualGroupFamilyId: updateBucket.GlobalVirtualGroupFamilyId,
		UpdateAt:                   block.Block.Height,
		UpdateTxHash:               txHash,
		UpdateEVMTxHash:            evmTxHash,
		UpdateTime:                 block.Block.Time,
	}
	k, v := m.db.UpdateBucketToSQL(ctx, b)
	ek, ev := m.db.SaveBucketEventToSQL(ctx, b.ToBucketEvent(EventUpdateBucketInfo))
	return map[string][]interface{}{
		k:  v,
		ek: ev,
	}
}

func (m *Module) handleEventMigrationBucket(ctx context.Context, block *tmctypes.ResultBlock, txHash, evmTxHash string, migrationBucket *storagetypes.EventMigrationBucket) map[string][]interface{} {
	b := &models.Bucket{
		BucketID:        migrationBucket.BucketId.BigInt().String(),
		BucketName:      migrationBucket.BucketName,
		Status:          migrationBucket.Status.String(),
		UpdateAt:        block.Block.Height,
		UpdateTxHash:    txHash,
		UpdateEVMTxHash: evmTxHash,
		UpdateTime:      block.Block.Time,
	}

	k, v := m.db.UpdateBucketToSQL(ctx, b)
	ek, ev := m.db.SaveBucketEventToSQL(ctx, b.ToBucketEvent(EventMigrationBucket))
	return map[string][]interface{}{
		k:  v,
		ek: ev,
	}
}

func (m *Module) handleEventCancelMigrationBucket(ctx context.Context, block *tmctypes.ResultBlock, txHash, evmTxHash string, cancelMigrationBucket *storagetypes.EventCancelMigrationBucket) map[string][]interface{} {
	b := &models.Bucket{
		BucketID:        cancelMigrationBucket.BucketId.BigInt().String(),
		BucketName:      cancelMigrationBucket.BucketName,
		Status:          cancelMigrationBucket.Status.String(),
		UpdateAt:        block.Block.Height,
		UpdateTxHash:    txHash,
		UpdateEVMTxHash: evmTxHash,
		UpdateTime:      block.Block.Time,
	}
	k, v := m.db.UpdateBucketToSQL(ctx, b)
	ek, ev := m.db.SaveBucketEventToSQL(ctx, b.ToBucketEvent(EventCancelMigrationBucket))
	return map[string][]interface{}{
		k:  v,
		ek: ev,
	}
}

func (m *Module) handleEventRejectMigrateBucket(ctx context.Context, block *tmctypes.ResultBlock, txHash, evmTxHash string, rejectMigrateBucket *storagetypes.EventRejectMigrateBucket) map[string][]interface{} {
	b := &models.Bucket{
		BucketID:        rejectMigrateBucket.BucketId.BigInt().String(),
		BucketName:      rejectMigrateBucket.BucketName,
		Status:          rejectMigrateBucket.Status.String(),
		UpdateAt:        block.Block.Height,
		UpdateTxHash:    txHash,
		UpdateEVMTxHash: evmTxHash,
		UpdateTime:      block.Block.Time,
	}
	k, v := m.db.UpdateBucketToSQL(ctx, b)
	ek, ev := m.db.SaveBucketEventToSQL(ctx, b.ToBucketEvent(EventRejectMigrateBucket))
	return map[string][]interface{}{
		k:  v,
		ek: ev,
	}
}

func (m *Module) handleCompleteMigrationBucket(ctx context.Context, block *tmctypes.ResultBlock, txHash, evmTxHash string, completeMigrationBucket *storagetypes.EventCompleteMigrationBucket) map[string][]interface{} {
	b := &models.Bucket{
		BucketID:                   completeMigrationBucket.BucketId.BigInt().String(),
		BucketName:                 completeMigrationBucket.BucketName,
		GlobalVirtualGroupFamilyId: completeMigrationBucket.GlobalVirtualGroupFamilyId,
		Status:                     completeMigrationBucket.Status.String(),
		UpdateAt:                   block.Block.Height,
		UpdateTxHash:               txHash,
		UpdateEVMTxHash:            evmTxHash,
		UpdateTime:                 block.Block.Time,
	}
	k, v := m.db.UpdateBucketToSQL(ctx, b)
	ek, ev := m.db.SaveBucketEventToSQL(ctx, b.ToBucketEvent(EventCompleteMigrationBucket))
	return map[string][]interface{}{
		k:  v,
		ek: ev,
	}
}
