package source

import (
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/forbole/bdjuno/v4/modules/sp/types"
)

type Source interface {
	StorageProvider(height int64, id uint32) (types.StorageProvider, error)
	StorageProviders(height int64, pageRequest query.PageRequest) ([]*types.StorageProvider, *query.PageResponse, error)
}
