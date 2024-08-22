package storage

import (
	"context"
	"time"

	tmctypes "github.com/cometbft/cometbft/rpc/core/types"
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

func (m *Module) handleCreateBucket(ctx context.Context, block *tmctypes.ResultBlock, txHash string, createBucket *storagetypes.EventCreateBucket) map[string][]interface{} {
	bucket := &models.Bucket{
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
		CreateTime:                 time.Unix(createBucket.CreateAt, 0),
		UpdateAt:                   block.Block.Height,
		UpdateTxHash:               txHash,
		UpdateTime:                 block.Block.Time,
	}
	k, v := m.db.SaveBucketToSQL(ctx, bucket)
	return map[string][]interface{}{
		k: v,
	}
}

func (m *Module) handleDeleteBucket(ctx context.Context, block *tmctypes.ResultBlock, txHash string, deleteBucket *storagetypes.EventDeleteBucket) map[string][]interface{} {
	bucket := &models.Bucket{
		BucketID:                   deleteBucket.BucketId.BigInt().String(),
		BucketName:                 deleteBucket.BucketName,
		OwnerAddress:               deleteBucket.Owner,
		GlobalVirtualGroupFamilyId: deleteBucket.GlobalVirtualGroupFamilyId,
		Removed:                    true,
		UpdateAt:                   block.Block.Height,
		UpdateTxHash:               txHash,
		UpdateTime:                 block.Block.Time,
	}

	k, v := m.db.UpdateBucketToSQL(ctx, bucket)
	return map[string][]interface{}{
		k: v,
	}
}

func (m *Module) handleDiscontinueBucket(ctx context.Context, block *tmctypes.ResultBlock, txHash string, discontinueBucket *storagetypes.EventDiscontinueBucket) map[string][]interface{} {
	bucket := &models.Bucket{
		BucketID:     discontinueBucket.BucketId.BigInt().String(),
		BucketName:   discontinueBucket.BucketName,
		DeleteReason: discontinueBucket.Reason,
		DeleteAt:     discontinueBucket.DeleteAt,
		Status:       storagetypes.BUCKET_STATUS_DISCONTINUED.String(),
		UpdateAt:     block.Block.Height,
		UpdateTxHash: txHash,
		UpdateTime:   block.Block.Time,
	}

	k, v := m.db.UpdateBucketToSQL(ctx, bucket)
	return map[string][]interface{}{
		k: v,
	}
}

func (m *Module) handleUpdateBucketInfo(ctx context.Context, block *tmctypes.ResultBlock, txHash string, updateBucket *storagetypes.EventUpdateBucketInfo) map[string][]interface{} {
	bucket := &models.Bucket{
		BucketName:                 updateBucket.BucketName,
		BucketID:                   updateBucket.BucketId.BigInt().String(),
		ChargedReadQuota:           updateBucket.ChargedReadQuota,
		PaymentAddress:             updateBucket.PaymentAddress,
		Visibility:                 updateBucket.Visibility.String(),
		GlobalVirtualGroupFamilyId: updateBucket.GlobalVirtualGroupFamilyId,
		UpdateAt:                   block.Block.Height,
		UpdateTxHash:               txHash,
		UpdateTime:                 block.Block.Time,
	}

	k, v := m.db.UpdateBucketToSQL(ctx, bucket)
	return map[string][]interface{}{
		k: v,
	}
}

func (m *Module) handleEventMigrationBucket(ctx context.Context, block *tmctypes.ResultBlock, txHash string, migrationBucket *storagetypes.EventMigrationBucket) map[string][]interface{} {
	bucket := &models.Bucket{
		BucketID:     migrationBucket.BucketId.BigInt().String(),
		BucketName:   migrationBucket.BucketName,
		Status:       migrationBucket.Status.String(),
		UpdateAt:     block.Block.Height,
		UpdateTxHash: txHash,
		UpdateTime:   block.Block.Time,
	}

	k, v := m.db.UpdateBucketToSQL(ctx, bucket)
	return map[string][]interface{}{
		k: v,
	}
}

func (m *Module) handleEventCancelMigrationBucket(ctx context.Context, block *tmctypes.ResultBlock, txHash string, cancelMigrationBucket *storagetypes.EventCancelMigrationBucket) map[string][]interface{} {
	bucket := &models.Bucket{
		BucketID:   cancelMigrationBucket.BucketId.BigInt().String(),
		BucketName: cancelMigrationBucket.BucketName,
		Status:     cancelMigrationBucket.Status.String(),

		UpdateAt:     block.Block.Height,
		UpdateTxHash: txHash,
		UpdateTime:   block.Block.Time,
	}

	k, v := m.db.UpdateBucketToSQL(ctx, bucket)
	return map[string][]interface{}{
		k: v,
	}
}

func (m *Module) handleEventRejectMigrateBucket(ctx context.Context, block *tmctypes.ResultBlock, txHash string, rejectMigrateBucket *storagetypes.EventRejectMigrateBucket) map[string][]interface{} {
	bucket := &models.Bucket{
		BucketID:   rejectMigrateBucket.BucketId.BigInt().String(),
		BucketName: rejectMigrateBucket.BucketName,
		Status:     rejectMigrateBucket.Status.String(),

		UpdateAt:     block.Block.Height,
		UpdateTxHash: txHash,
		UpdateTime:   block.Block.Time,
	}

	k, v := m.db.UpdateBucketToSQL(ctx, bucket)
	return map[string][]interface{}{
		k: v,
	}
}

func (m *Module) handleCompleteMigrationBucket(ctx context.Context, block *tmctypes.ResultBlock, txHash string, completeMigrationBucket *storagetypes.EventCompleteMigrationBucket) map[string][]interface{} {
	bucket := &models.Bucket{
		BucketID:                   completeMigrationBucket.BucketId.BigInt().String(),
		BucketName:                 completeMigrationBucket.BucketName,
		GlobalVirtualGroupFamilyId: completeMigrationBucket.GlobalVirtualGroupFamilyId,
		Status:                     completeMigrationBucket.Status.String(),

		UpdateAt:     block.Block.Height,
		UpdateTxHash: txHash,
		UpdateTime:   block.Block.Time,
	}

	k, v := m.db.UpdateBucketToSQL(ctx, bucket)
	return map[string][]interface{}{
		k: v,
	}
}
