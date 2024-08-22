package source

import (
	sptypes "github.com/forbole/bdjuno/v4/modules/sp/types"
)

type Source interface {
	StorageProvider(height int64, id uint32) (sptypes.StorageProvider, error)
}
