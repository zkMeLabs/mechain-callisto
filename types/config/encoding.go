package config

import (
	"cosmossdk.io/simapp/params"
	"github.com/cosmos/cosmos-sdk/types/module"
	ethermint "github.com/evmos/evmos/v12/encoding"
	evmtypes "github.com/evmos/evmos/v12/x/evm/types"
)

// MakeEncodingConfig creates an EncodingConfig to properly handle all the messages
func MakeEncodingConfig(managers []module.BasicManager) func() params.EncodingConfig {
	return func() params.EncodingConfig {
		manager := mergeBasicManagers(managers)
		encodingConfig := ethermint.MakeConfig(manager)
		evmtypes.RegisterInterfaces(encodingConfig.InterfaceRegistry)
		return encodingConfig
	}
}

// mergeBasicManagers merges the given managers into a single module.BasicManager
func mergeBasicManagers(managers []module.BasicManager) module.BasicManager {
	union := module.BasicManager{}
	for _, manager := range managers {
		for k, v := range manager {
			union[k] = v
		}
	}
	return union
}
