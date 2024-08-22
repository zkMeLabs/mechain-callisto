package remote

import (
	"github.com/forbole/bdjuno/v4/modules/virtualgroup/source"
	"github.com/forbole/bdjuno/v4/modules/virtualgroup/types"
	"github.com/forbole/juno/v5/node/remote"
)

var _ source.Source = &Source{}

// Source implements storage.Source using a remote node
type Source struct {
	*remote.Source
	virtualGroupClient types.QueryClient
}

// NewSource returns a new Source implementation
func NewSource(source *remote.Source, virtualGroupClient types.QueryClient) *Source {
	return &Source{
		Source:             source,
		virtualGroupClient: virtualGroupClient,
	}
}

func (s Source) GlobalVirtualGroup(height int64, globalVirtualGroupId uint32) (types.GlobalVirtualGroup, error) {
	res, err := s.virtualGroupClient.GlobalVirtualGroup(
		remote.GetHeightRequestContext(s.Ctx, height),
		&types.QueryGlobalVirtualGroupRequest{GlobalVirtualGroupId: globalVirtualGroupId},
	)
	if err != nil {
		return types.GlobalVirtualGroup{}, err
	}

	return *res.GlobalVirtualGroup, nil
}

func (s Source) GlobalVirtualGroupByFamilyID(height int64, globalVirtualGroupFamilyId uint32) ([]*types.GlobalVirtualGroup, error) {
	res, err := s.virtualGroupClient.GlobalVirtualGroupByFamilyID(
		remote.GetHeightRequestContext(s.Ctx, height),
		&types.QueryGlobalVirtualGroupByFamilyIDRequest{GlobalVirtualGroupFamilyId: globalVirtualGroupFamilyId},
	)
	if err != nil {
		return nil, err
	}

	return res.GlobalVirtualGroups, nil
}

func (s Source) GlobalVirtualGroupFamily(height int64, familyId uint32) (types.GlobalVirtualGroupFamily, error) {
	res, err := s.virtualGroupClient.GlobalVirtualGroupFamily(
		remote.GetHeightRequestContext(s.Ctx, height),
		&types.QueryGlobalVirtualGroupFamilyRequest{FamilyId: familyId},
	)
	if err != nil {
		return types.GlobalVirtualGroupFamily{}, err
	}

	return *res.GlobalVirtualGroupFamily, nil
}
