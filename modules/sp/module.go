package sp

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/forbole/bdjuno/v4/database"
	storagesource "github.com/forbole/bdjuno/v4/modules/storage/source"
	"github.com/forbole/juno/v5/modules"
)

var (
	_ modules.Module      = &Module{}
	_ modules.BlockModule = &Module{}
)

// Module represents the x/distr module
type Module struct {
	source storagesource.Source
	cdc    codec.Codec
	db     *database.DB
}

// NewModule returns a new Module instance
func NewModule(source storagesource.Source, cdc codec.Codec, db *database.DB) *Module {
	return &Module{
		source: source,
		cdc:    cdc,
		db:     db,
	}
}

// Name implements modules.Module
func (m *Module) Name() string {
	return "sp"
}
