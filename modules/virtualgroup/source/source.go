package source

import (
	"context"

	"github.com/forbole/bdjuno/v4/modules/virtualgroup/types"
)

type Source interface {
	GlobalVirtualGroup(height int64, globalVirtualGroupId uint32) (*types.GlobalVirtualGroup, error)
	GlobalVirtualGroupByFamilyID(height int64, globalVirtualGroupFamilyId uint32) ([]*types.GlobalVirtualGroup, error)
	GlobalVirtualGroupFamily(height int64, familyId uint32) (*types.GlobalVirtualGroupFamily, error)
	GlobalVirtualGroupFamilies(height int64, ctx context.Context) ([]*types.GlobalVirtualGroupFamily, error)
}
