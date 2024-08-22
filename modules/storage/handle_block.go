package storage

import (
	"context"

	tmctypes "github.com/cometbft/cometbft/rpc/core/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
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
	statements, err := m.ExportEventsInTxs(ctx, block, txs)
	if err != nil {
		return err
	}
	return m.ExecuteStatements(statements)
}

func (m *Module) ExecuteStatements(statements map[string][]interface{}) error {
	tx := m.db.G.Begin()
	if tx.Error != nil {
		return tx.Error
	}
	for sql, vars := range statements {
		if err := tx.Exec(sql, vars...).Error; err != nil {
			tx.Rollback()
			return err
		}
	}
	if err := tx.Commit().Error; err != nil {
		return err
	}
	return nil
}

// ExportEventsInTxs accepts a slice of events in tx in order to save in database.
func (m Module) ExportEventsInTxs(ctx context.Context, block *tmctypes.ResultBlock, txs []*junotypes.Tx) (map[string][]interface{}, error) {
	allSQL := make(map[string][]interface{}, 0)
	for _, tx := range txs {
		sqls, err := m.ExtractEvent(ctx, block, tx)
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

// ExtractEvent accepts the transaction and handles events contained inside the transaction.
func (m *Module) ExtractEvent(ctx context.Context, block *tmctypes.ResultBlock, tx *junotypes.Tx) (map[string][]interface{}, error) {
	allSQL := make(map[string][]interface{})
	for _, event := range tx.Events {
		e := sdk.Event(event)
		h := m.getExtractEventFunc(e)
		if h == nil {
			continue
		}
		sqls, err := h(ctx, block, tx.TxHash, e)
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

type ExtractFunc func(ctx context.Context, block *tmctypes.ResultBlock, txHash string, event sdk.Event) (map[string][]interface{}, error)

func (m *Module) getExtractEventFunc(event sdk.Event) ExtractFunc {
	if BucketEvents[event.Type] {
		return m.ExtractBucketEventStatements
	}
	if ObjectEvents[event.Type] {
		return m.ExtractObjectEventStatements
	}
	if GroupEvents[event.Type] {
		return m.ExtractGroupEventStatements
	}
	return nil
}
