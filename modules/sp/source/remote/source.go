package remote

import (
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/forbole/bdjuno/v4/modules/sp/source"
	"github.com/forbole/bdjuno/v4/modules/sp/types"
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

func (s Source) StorageProvider(height int64, id uint32) (types.StorageProvider, error) {
	res, err := s.Cli.StorageProvider(
		remote.GetHeightRequestContext(s.Ctx, height),
		&types.QueryStorageProviderRequest{Id: id},
	)
	if err != nil {
		return types.StorageProvider{}, err
	}

	return *res.StorageProvider, nil
}

func (s Source) StorageProviders(height int64, pageRequest query.PageRequest) ([]*types.StorageProvider, *query.PageResponse, error) {
	res, err := s.Cli.StorageProviders(
		remote.GetHeightRequestContext(s.Ctx, height),
		&types.QueryStorageProvidersRequest{Pagination: &pageRequest},
	)
	if err != nil {
		return nil, nil, err
	}

	return res.Sps, res.Pagination, nil
}
