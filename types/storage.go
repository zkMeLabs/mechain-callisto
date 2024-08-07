package types

import "time"

type Groups struct {
	GroupId    uint64
	GroupName  string
	Owner      string
	SourceType string
	Extra      string
	Tags       string
}

func NewGroups(
	groupId uint64,
	groupName string,
	owner string,
	sourceType string,
	extra string,
	tags string,
) Groups {
	return Groups{
		GroupId:    groupId,
		GroupName:  groupName,
		Owner:      owner,
		SourceType: sourceType,
		Extra:      extra,
		Tags:       tags,
	}
}
func (g Groups) Equal(v Groups) bool {
	return g.GroupId == v.GroupId &&
		g.GroupName == v.GroupName &&
		g.Owner == v.Owner &&
		g.SourceType == v.SourceType &&
		g.Extra == v.Extra &&
		g.Tags == v.Tags
}

type Buckets struct {
	BucketId                   uint64
	BucketName                 string
	Owner                      string
	Visibility                 string
	SourceType                 string
	CreateAt                   time.Time
	PaymentAddress             string
	BucketStatus               string
	Tags                       string
	ChargedReadQuota           uint64
	GlobalVirtualGroupFamilyId uint32
	SpAsDelegatedAgentDisabled bool
}

func NewBuckets(
	bucketId uint64,
	bucketName string,
	owner string,
	visibility string,
	sourceType string,
	createAt time.Time,
	paymentAddress string,
	bucketStatus string,
	tags string,
	chargedReadQuota uint64,
	globalVirtualGroupFamilyId uint32,
	spAsDelegatedAgentDisabled bool,
) Buckets {
	return Buckets{
		BucketId:                   bucketId,
		BucketName:                 bucketName,
		Owner:                      owner,
		Visibility:                 visibility,
		SourceType:                 sourceType,
		CreateAt:                   createAt,
		PaymentAddress:             paymentAddress,
		BucketStatus:               bucketStatus,
		Tags:                       tags,
		ChargedReadQuota:           chargedReadQuota,
		GlobalVirtualGroupFamilyId: globalVirtualGroupFamilyId,
		SpAsDelegatedAgentDisabled: spAsDelegatedAgentDisabled,
	}
}

func (b Buckets) Equal(v Buckets) bool {
	return b.BucketId == v.BucketId &&
		b.BucketName == v.BucketName &&
		b.Owner == v.Owner &&
		b.Visibility == v.Visibility &&
		b.SourceType == v.SourceType &&
		b.CreateAt.Equal(v.CreateAt) &&
		b.PaymentAddress == v.PaymentAddress &&
		b.BucketStatus == v.BucketStatus &&
		b.Tags == v.Tags &&
		b.ChargedReadQuota == v.ChargedReadQuota &&
		b.GlobalVirtualGroupFamilyId == v.GlobalVirtualGroupFamilyId &&
		b.SpAsDelegatedAgentDisabled == v.SpAsDelegatedAgentDisabled

}

type Objects struct {
	ObjectId            uint64
	ObjectName          string
	BucketName          string
	Owner               string
	Creator             string
	PayloadSize         uint64
	Visibility          string
	ContentType         string
	ObjectStatus        string
	RedundancyType      string
	SourceType          string
	Checksums           string
	Tags                string
	IsUpdating          bool
	CreateAt            time.Time
	UpdatedAt           time.Time
	UpdatedBy           string
	Version             int64
	LocalVirtualGroupId uint32
}

func NewObjects(
	objectId uint64,
	objectName string,
	bucketName string,
	owner string,
	creator string,
	payloadSize uint64,
	visibility string,
	contentType string,
	objectStatus string,
	redundancyType string,
	sourceType string,
	checksums string,
	tags string,
	isUpdating bool,
	createAt time.Time,
	updateAt time.Time,
	updatedBy string,
	version int64,
	localVirtualGroupId uint32,
) Objects {
	return Objects{
		ObjectId:            objectId,
		ObjectName:          objectName,
		BucketName:          bucketName,
		Owner:               owner,
		Creator:             creator,
		PayloadSize:         payloadSize,
		Visibility:          visibility,
		ContentType:         contentType,
		ObjectStatus:        objectStatus,
		RedundancyType:      redundancyType,
		SourceType:          sourceType,
		Checksums:           checksums,
		Tags:                tags,
		IsUpdating:          isUpdating,
		CreateAt:            createAt,
		UpdatedAt:           updateAt,
		UpdatedBy:           updatedBy,
		Version:             version,
		LocalVirtualGroupId: localVirtualGroupId,
	}
}

func (o Objects) Equal(v Objects) bool {
	return o.ObjectId == v.ObjectId &&
		o.ObjectName == v.ObjectName &&
		o.BucketName == v.BucketName &&
		o.Owner == v.Owner &&
		o.Creator == v.Creator &&
		o.PayloadSize == v.PayloadSize &&
		o.Visibility == v.Visibility &&
		o.ContentType == v.ContentType &&
		o.ObjectStatus == v.ObjectStatus &&
		o.RedundancyType == v.RedundancyType &&
		o.SourceType == v.SourceType &&
		o.Checksums == v.Checksums &&
		o.Tags == v.Tags &&
		o.IsUpdating == v.IsUpdating &&
		o.CreateAt.Equal(v.CreateAt) &&
		o.UpdatedAt.Equal(v.UpdatedAt) &&
		o.UpdatedBy == v.UpdatedBy &&
		o.Version == v.Version &&
		o.LocalVirtualGroupId == v.LocalVirtualGroupId
}
