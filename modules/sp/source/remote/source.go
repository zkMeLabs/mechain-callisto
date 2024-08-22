package remote

import (
	"github.com/forbole/bdjuno/v4/modules/sp/source"
	"github.com/forbole/bdjuno/v4/modules/sp/types"
	"github.com/forbole/juno/v5/node/remote"
)

var _ source.Source = &Source{}

// Source implements storage.Source using a remote node
type Source struct {
	*remote.Source
	spClient types.QueryClient
}

// NewSource returns a new Source implementation
func NewSource(source *remote.Source, spClient types.QueryClient) *Source {
	return &Source{
		Source:   source,
		spClient: spClient,
	}
}

func (s Source) StorageProvider(height int64, id uint32) (types.StorageProvider, error) {
	res, err := s.spClient.StorageProvider(
		remote.GetHeightRequestContext(s.Ctx, height),
		&types.QueryStorageProviderRequest{Id: id},
	)
	if err != nil {
		return types.StorageProvider{}, err
	}

	return *res.StorageProvider, nil
}
