package sp

import (
	tmctypes "github.com/cometbft/cometbft/rpc/core/types"
	junotypes "github.com/forbole/juno/v5/types"
)

// HandleBlock implements modules.BlockModule
func (m *Module) HandleBlock(
	block *tmctypes.ResultBlock, results *tmctypes.ResultBlockResults, txs []*junotypes.Tx, vals *tmctypes.ResultValidators,
) error {
	return nil
}
