package source

import (
	permission "github.com/forbole/bdjuno/v4/modules/storage/permission"
	"github.com/forbole/bdjuno/v4/modules/storage/types"
	vitualgroup "github.com/forbole/bdjuno/v4/modules/storage/vitualgroup"
)

type Source interface {
	HeadBucket(height int64, bucketName string) (types.BucketInfo, types.BucketExtraInfo, error)
	HeadBucketById(height int64, bucketId string) (types.BucketInfo, types.BucketExtraInfo, error)
	HeadBucketExtra(height int64, bucketName string) (types.InternalBucketInfo, error)

	HeadGroup(height int64, groupOwner, groupName string) (types.GroupInfo, error)
	HeadGroupMember(height int64, member, groupOwner, groupName string) (permission.GroupMember, error)

	HeadObject(height int64, bucketName, objectName string) (types.ObjectInfo, vitualgroup.GlobalVirtualGroup, error)
	HeadObjectById(height int64, objectId string) (types.ObjectInfo, vitualgroup.GlobalVirtualGroup, error)
	HeadShadowObject(height int64, bucketName, objectName string) (types.ShadowObjectInfo, error)
}
