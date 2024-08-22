package storage

import (
	"context"
	"errors"

	abci "github.com/cometbft/cometbft/abci/types"
	tmctypes "github.com/cometbft/cometbft/rpc/core/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	storagetypes "github.com/forbole/bdjuno/v4/modules/storage/types"
	junotypes "github.com/forbole/juno/v5/types"
	"github.com/rs/zerolog/log"
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
	block *tmctypes.ResultBlock, results *tmctypes.ResultBlockResults, txs []*junotypes.Tx, vals *tmctypes.ResultValidators,
) error {
	ctx := context.Background()
	allSQL, err := m.ExportEventsInTxs(ctx, block, txs)
	if err != nil {
		return err
	}
	_ = allSQL
	// sqlCount := len(allSQL)

	// step := 0
	// for step < sqlCount {
	// 	finalSQL := ""
	// 	finalVal := make([]interface{}, 0)
	// 	left := step
	// 	right := step + int(i.CommitNumber)
	// 	if right > sqlCount {
	// 		right = sqlCount
	// 	}
	// 	for _, m := range allSQL[left:right] {
	// 		for k, v := range m {
	// 			finalSQL += fmt.Sprintf("%s;   ", k)
	// 			finalVal = append(finalVal, v...)
	// 		}
	// 	}
	// 	tx := m.db.Begin(context.TODO())
	// 	if txErr := m.db.G.Session(&gorm.Session{DryRun: false}).Exec(finalSQL, finalVal...).Error; txErr != nil {
	// 		tx.Rollback()
	// 		return txErr
	// 	}

	// 	if txErr := tx.Commit(); txErr != nil {
	// 		return txErr
	// 	}
	// 	step = right
	// }
	return nil
}

// ExportEventsInTxs accepts a slice of events in tx in order to save in database.
func (m Module) ExportEventsInTxs(ctx context.Context, block *tmctypes.ResultBlock, txs []*junotypes.Tx) ([]map[string][]interface{}, error) {
	allSQL := make([]map[string][]interface{}, 0)
	for _, tx := range txs {
		sqls, err := m.ExtractEvent(ctx, block, tx)
		if err != nil {
			log.Err(err)
			continue
		}
		if len(sqls) != 0 {
			allSQL = append(allSQL, sqls)
		}
	}
	return allSQL, nil
}

// ExtractEvent accepts the transaction and handles events contained inside the transaction.
func (m *Module) ExtractEvent(ctx context.Context, block *tmctypes.ResultBlock, tx *junotypes.Tx) (map[string][]interface{}, error) {
	allSQL := make(map[string][]interface{})
	for _, event := range tx.Events {
		sqls, err := m.ExtractEventStatements(ctx, block, tx.TxHash, sdk.Event(event))
		if err != nil {
			log.Err(err)
			continue
		}
		for k, v := range sqls {
			allSQL[k] = v
		}
	}
	return allSQL, nil
}

func (m *Module) ExtractEventStatements(ctx context.Context, block *tmctypes.ResultBlock, txHash string, event sdk.Event) (map[string][]interface{}, error) {
	if !BucketEvents[event.Type] {
		return nil, nil
	}

	typedEvent, err := sdk.ParseTypedEvent(abci.Event(event))
	if err != nil {
		return nil, err
	}

	switch event.Type {
	// bucket event
	case EventCreateBucket:
		createBucket, ok := typedEvent.(*storagetypes.EventCreateBucket)
		if !ok {
			return nil, errors.New("create bucket event assert error")
		}
		return m.handleCreateBucket(ctx, block, txHash, createBucket), nil
	case EventDeleteBucket:
		deleteBucket, ok := typedEvent.(*storagetypes.EventDeleteBucket)
		if !ok {
			return nil, errors.New("delete bucket event assert error")
		}
		return m.handleDeleteBucket(ctx, block, txHash, deleteBucket), nil
	case EventUpdateBucketInfo:
		updateBucketInfo, ok := typedEvent.(*storagetypes.EventUpdateBucketInfo)
		if !ok {
			return nil, errors.New("update bucket event assert error")
		}
		return m.handleUpdateBucketInfo(ctx, block, txHash, updateBucketInfo), nil
	case EventDiscontinueBucket:
		discontinueBucket, ok := typedEvent.(*storagetypes.EventDiscontinueBucket)
		if !ok {
			return nil, errors.New("discontinue bucket event assert error")
		}
		return m.handleDiscontinueBucket(ctx, block, txHash, discontinueBucket), nil
	case EventMigrationBucket:
		migrationBucket, ok := typedEvent.(*storagetypes.EventMigrationBucket)
		if !ok {
			return nil, errors.New("migration bucket event assert error")
		}
		return m.handleEventMigrationBucket(ctx, block, txHash, migrationBucket), nil
	case EventCancelMigrationBucket:
		cancelMigrationBucket, ok := typedEvent.(*storagetypes.EventCancelMigrationBucket)
		if !ok {
			return nil, errors.New("cancel migration bucket event assert error")
		}
		return m.handleEventCancelMigrationBucket(ctx, block, txHash, cancelMigrationBucket), nil
	case EventRejectMigrateBucket:
		rejectMigrateBucket, ok := typedEvent.(*storagetypes.EventRejectMigrateBucket)
		if !ok {
			return nil, errors.New("reject migration bucket event assert error")
		}
		return m.handleEventRejectMigrateBucket(ctx, block, txHash, rejectMigrateBucket), nil
	case EventCompleteMigrationBucket:
		completeMigrationBucket, ok := typedEvent.(*storagetypes.EventCompleteMigrationBucket)
		if !ok {
			return nil, errors.New("complete migrate bucket event assert error")
		}
		return m.handleCompleteMigrationBucket(ctx, block, txHash, completeMigrationBucket), nil

		// object event
	}

	return nil, nil
}
