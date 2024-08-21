package storage

import (
	tmctypes "github.com/cometbft/cometbft/rpc/core/types"
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
				m.handleCreateBucket(b, tx.TxHash, event.Attributes)
			case EventTypeDeleteBucket:
				m.handleDeleteBucket(b, tx.TxHash, event.Attributes)
			case EventTypeCreateObject:
				m.handleCreateObject(b.Block.Height, event.Attributes)
			case EventTypeDeleteObject:
				m.handleDeleteObject(b.Block.Height, event.Attributes)
			}
		}
	}
}
