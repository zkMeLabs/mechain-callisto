package upgrade

import (
	"github.com/forbole/bdjuno/v4/database"

	"github.com/forbole/juno/v5/modules"
)

var (
	_ modules.Module      = &Module{}
	_ modules.BlockModule = &Module{}
)

// Module represents the x/upgrade module
type Module struct {
	db            *database.DB
	stakingModule StakingModule
}

// NewModule builds a new Module instance
func NewModule(db *database.DB, stakingModule StakingModule) *Module {
	return &Module{
		stakingModule: stakingModule,
		db:            db,
	}
}

// Name implements modules.Module
func (m *Module) Name() string {
	return "upgrade"
}
