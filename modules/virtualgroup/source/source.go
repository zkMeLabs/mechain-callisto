package source

import (
	vgtypes "github.com/forbole/bdjuno/v4/modules/virtualgroup/types"
)

type Source interface {
	GlobalVirtualGroup(height int64, globalVirtualGroupId uint32) (vgtypes.GlobalVirtualGroup, error)
	GlobalVirtualGroupByFamilyID(height int64, globalVirtualGroupFamilyId uint32) ([]*vgtypes.GlobalVirtualGroup, error)
	GlobalVirtualGroupFamily(height int64, familyId uint32) (vgtypes.GlobalVirtualGroupFamily, error)
}
