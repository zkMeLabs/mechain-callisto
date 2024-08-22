package source

import (
	permission "github.com/forbole/bdjuno/v4/modules/storage/permission"
	sp "github.com/forbole/bdjuno/v4/modules/storage/sp"
	"github.com/forbole/bdjuno/v4/modules/storage/types"
	virtualgroup "github.com/forbole/bdjuno/v4/modules/storage/virtualgroup"
)

type Source interface {
	HeadBucket(height int64, bucketName string) (types.BucketInfo, types.BucketExtraInfo, error)
	HeadBucketById(height int64, bucketId string) (types.BucketInfo, types.BucketExtraInfo, error)
	HeadBucketExtra(height int64, bucketName string) (types.InternalBucketInfo, error)

	HeadGroup(height int64, groupOwner, groupName string) (types.GroupInfo, error)
	HeadGroupMember(height int64, member, groupOwner, groupName string) (permission.GroupMember, error)

	HeadObject(height int64, bucketName, objectName string) (types.ObjectInfo, virtualgroup.GlobalVirtualGroup, error)
	HeadObjectById(height int64, objectId string) (types.ObjectInfo, virtualgroup.GlobalVirtualGroup, error)
	HeadShadowObject(height int64, bucketName, objectName string) (types.ShadowObjectInfo, error)

	GlobalVirtualGroup(height int64, globalVirtualGroupId uint32) (virtualgroup.GlobalVirtualGroup, error)
	GlobalVirtualGroupByFamilyID(height int64, globalVirtualGroupFamilyId uint32) ([]*virtualgroup.GlobalVirtualGroup, error)
	GlobalVirtualGroupFamily(height int64, familyId uint32) (virtualgroup.GlobalVirtualGroupFamily, error)

	StorageProvider(height int64, id uint32) (sp.StorageProvider, error)
}
