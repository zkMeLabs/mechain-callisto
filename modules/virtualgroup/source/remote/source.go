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
	Cli types.QueryClient
}

// NewSource returns a new Source implementation
func NewSource(source *remote.Source, cli types.QueryClient) *Source {
	return &Source{
		Source: source,
		Cli:    cli,
	}
}

func (s Source) GlobalVirtualGroup(height int64, globalVirtualGroupID uint32) (types.GlobalVirtualGroup, error) {
	res, err := s.Cli.GlobalVirtualGroup(
		remote.GetHeightRequestContext(s.Ctx, height),
		&types.QueryGlobalVirtualGroupRequest{GlobalVirtualGroupId: globalVirtualGroupID},
	)
	if err != nil {
		return types.GlobalVirtualGroup{}, err
	}

	return *res.GlobalVirtualGroup, nil
}

func (s Source) GlobalVirtualGroupByFamilyID(height int64, globalVirtualGroupFamilyID uint32) ([]*types.GlobalVirtualGroup, error) {
	res, err := s.Cli.GlobalVirtualGroupByFamilyID(
		remote.GetHeightRequestContext(s.Ctx, height),
		&types.QueryGlobalVirtualGroupByFamilyIDRequest{GlobalVirtualGroupFamilyId: globalVirtualGroupFamilyID},
	)
	if err != nil {
		return nil, err
	}

	return res.GlobalVirtualGroups, nil
}

func (s Source) GlobalVirtualGroupFamily(height int64, familyID uint32) (types.GlobalVirtualGroupFamily, error) {
	res, err := s.Cli.GlobalVirtualGroupFamily(
		remote.GetHeightRequestContext(s.Ctx, height),
		&types.QueryGlobalVirtualGroupFamilyRequest{FamilyId: familyID},
	)
	if err != nil {
		return types.GlobalVirtualGroupFamily{}, err
	}

	return *res.GlobalVirtualGroupFamily, nil
}
