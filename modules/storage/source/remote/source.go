package remote

import (
	permission "github.com/forbole/bdjuno/v4/modules/storage/permission"
	storagesource "github.com/forbole/bdjuno/v4/modules/storage/source"
	sp "github.com/forbole/bdjuno/v4/modules/storage/sp"
	types "github.com/forbole/bdjuno/v4/modules/storage/types"
	virtualgroup "github.com/forbole/bdjuno/v4/modules/storage/virtualgroup"
	"github.com/forbole/juno/v5/node/remote"
)

var (
	_ storagesource.Source = &Source{}
)

// Source implements storage.Source using a remote node
type Source struct {
	*remote.Source
	storageClient      types.QueryClient
	virtualGroupClient virtualgroup.QueryClient
	spClient           sp.QueryClient
}

// NewSource returns a new Source implementation
func NewSource(source *remote.Source, storageClient types.QueryClient, virtualGroupClient virtualgroup.QueryClient, spClient sp.QueryClient) *Source {
	return &Source{
		Source:             source,
		storageClient:      storageClient,
		virtualGroupClient: virtualGroupClient,
		spClient:           spClient,
	}
}

// HeadBucket implements storage source.Source
func (s Source) HeadBucket(height int64, bucketName string) (types.BucketInfo, types.BucketExtraInfo, error) {
	res, err := s.storageClient.HeadBucket(
		remote.GetHeightRequestContext(s.Ctx, height),
		&types.QueryHeadBucketRequest{BucketName: bucketName},
	)
	if err != nil {
		return types.BucketInfo{}, types.BucketExtraInfo{}, err
	}

	return *res.BucketInfo, *res.ExtraInfo, nil
}

func (s Source) HeadBucketById(height int64, bucketId string) (types.BucketInfo, types.BucketExtraInfo, error) {
	res, err := s.storageClient.HeadBucketById(
		remote.GetHeightRequestContext(s.Ctx, height),
		&types.QueryHeadBucketByIdRequest{BucketId: bucketId},
	)
	if err != nil {
		return types.BucketInfo{}, types.BucketExtraInfo{}, err
	}

	return *res.BucketInfo, *res.ExtraInfo, nil
}

func (s Source) HeadBucketExtra(height int64, bucketName string) (types.InternalBucketInfo, error) {
	res, err := s.storageClient.HeadBucketExtra(
		remote.GetHeightRequestContext(s.Ctx, height),
		&types.QueryHeadBucketExtraRequest{BucketName: bucketName},
	)
	if err != nil {
		return types.InternalBucketInfo{}, err
	}

	return *res.ExtraInfo, nil
}

func (s Source) HeadGroup(height int64, groupOwner, groupName string) (types.GroupInfo, error) {
	res, err := s.storageClient.HeadGroup(
		remote.GetHeightRequestContext(s.Ctx, height),
		&types.QueryHeadGroupRequest{
			GroupOwner: groupOwner,
			GroupName:  groupName,
		},
	)
	if err != nil {
		return types.GroupInfo{}, err
	}

	return *res.GroupInfo, nil
}

func (s Source) HeadGroupMember(height int64, member, groupOwner, groupName string) (permission.GroupMember, error) {
	res, err := s.storageClient.HeadGroupMember(
		remote.GetHeightRequestContext(s.Ctx, height),
		&types.QueryHeadGroupMemberRequest{
			Member:     member,
			GroupOwner: groupOwner,
			GroupName:  groupName,
		},
	)
	if err != nil {
		return permission.GroupMember{}, err
	}

	return *res.GroupMember, nil
}

func (s Source) HeadObject(height int64, bucketName, objectName string) (types.ObjectInfo, virtualgroup.GlobalVirtualGroup, error) {
	res, err := s.storageClient.HeadObject(
		remote.GetHeightRequestContext(s.Ctx, height),
		&types.QueryHeadObjectRequest{BucketName: bucketName, ObjectName: objectName},
	)
	if err != nil {
		return types.ObjectInfo{}, virtualgroup.GlobalVirtualGroup{}, err
	}

	return *res.ObjectInfo, *res.GlobalVirtualGroup, nil
}

func (s Source) HeadObjectById(height int64, objectId string) (types.ObjectInfo, virtualgroup.GlobalVirtualGroup, error) {
	res, err := s.storageClient.HeadObjectById(
		remote.GetHeightRequestContext(s.Ctx, height),
		&types.QueryHeadObjectByIdRequest{ObjectId: objectId},
	)
	if err != nil {
		return types.ObjectInfo{}, virtualgroup.GlobalVirtualGroup{}, err
	}

	return *res.ObjectInfo, *res.GlobalVirtualGroup, nil
}

func (s Source) HeadShadowObject(height int64, bucketName, objectName string) (types.ShadowObjectInfo, error) {
	res, err := s.storageClient.HeadShadowObject(
		remote.GetHeightRequestContext(s.Ctx, height),
		&types.QueryHeadShadowObjectRequest{BucketName: bucketName, ObjectName: objectName},
	)
	if err != nil {
		return types.ShadowObjectInfo{}, err
	}

	return *res.ObjectInfo, nil
}

func (s Source) GlobalVirtualGroup(height int64, globalVirtualGroupId uint32) (virtualgroup.GlobalVirtualGroup, error) {
	res, err := s.virtualGroupClient.GlobalVirtualGroup(
		remote.GetHeightRequestContext(s.Ctx, height),
		&virtualgroup.QueryGlobalVirtualGroupRequest{GlobalVirtualGroupId: globalVirtualGroupId},
	)
	if err != nil {
		return virtualgroup.GlobalVirtualGroup{}, err
	}

	return *res.GlobalVirtualGroup, nil
}

func (s Source) GlobalVirtualGroupByFamilyID(height int64, globalVirtualGroupFamilyId uint32) ([]*virtualgroup.GlobalVirtualGroup, error) {
	res, err := s.virtualGroupClient.GlobalVirtualGroupByFamilyID(
		remote.GetHeightRequestContext(s.Ctx, height),
		&virtualgroup.QueryGlobalVirtualGroupByFamilyIDRequest{GlobalVirtualGroupFamilyId: globalVirtualGroupFamilyId},
	)
	if err != nil {
		return nil, err
	}

	return res.GlobalVirtualGroups, nil
}

func (s Source) GlobalVirtualGroupFamily(height int64, familyId uint32) (virtualgroup.GlobalVirtualGroupFamily, error) {
	res, err := s.virtualGroupClient.GlobalVirtualGroupFamily(
		remote.GetHeightRequestContext(s.Ctx, height),
		&virtualgroup.QueryGlobalVirtualGroupFamilyRequest{FamilyId: familyId},
	)
	if err != nil {
		return virtualgroup.GlobalVirtualGroupFamily{}, err
	}

	return *res.GlobalVirtualGroupFamily, nil
}

func (s Source) StorageProvider(height int64, id uint32) (sp.StorageProvider, error) {
	res, err := s.spClient.StorageProvider(
		remote.GetHeightRequestContext(s.Ctx, height),
		&sp.QueryStorageProviderRequest{Id: id},
	)
	if err != nil {
		return sp.StorageProvider{}, err
	}

	return *res.StorageProvider, nil
}
