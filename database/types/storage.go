package types

import "time"

type GroupsRow struct {
	GroupId    uint64 `db:"id"`
	GroupName  string `db:"group_name"`
	Owner      string `db:"owner"`
	SourceType string `db:"source_type"`
	Extra      string `db:"extra"`
	Tags       string `db:"tags"`
}

func NewGroupsRow(
	groupId uint64,
	groupName string,
	owner string,
	sourceType string,
	extra string,
	tags string,
) GroupsRow {
	return GroupsRow{
		GroupId:    groupId,
		GroupName:  groupName,
		Owner:      owner,
		SourceType: sourceType,
		Extra:      extra,
		Tags:       tags,
	}
}
func (g GroupsRow) Equal(v GroupsRow) bool {
	return g.GroupId == v.GroupId &&
		g.GroupName == v.GroupName &&
		g.Owner == v.Owner &&
		g.SourceType == v.SourceType &&
		g.Extra == v.Extra &&
		g.Tags == v.Tags
}

type BucketsRow struct {
	BucketId                   uint64    `db:"id"`
	BucketName                 string    `db:"bucket_name"`
	Owner                      string    `db:"owner"`
	Visibility                 string    `db:"visibility"`
	SourceType                 string    `db:"source_type"`
	CreateAt                   time.Time `db:"create_at"`
	PaymentAddress             string    `db:"payment_address"`
	BucketStatus               string    `db:"bucket_status"`
	Tags                       string    `db:"tags"`
	ChargedReadQuota           uint64    `db:"charged_read_quota"`
	GlobalVirtualGroupFamilyId uint32    `db:"global_virtual_group_family_id"`
	SpAsDelegatedAgentDisabled bool      `db:"sp_as_delegated_agent_disabled"`
}

func NewBucketsRow(
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
) BucketsRow {
	return BucketsRow{
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

func (b BucketsRow) Equal(v BucketsRow) bool {
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

type ObjectsRow struct {
	ObjectId            uint64    `db:"id"`
	ObjectName          string    `db:"object_name"`
	BucketName          string    `db:"bucket_name"`
	Owner               string    `db:"owner"`
	Creator             string    `db:"creator"`
	PayloadSize         uint64    `db:"payload_size"`
	Visibility          string    `db:"visibility"`
	ContentType         string    `db:"content_type"`
	ObjectStatus        string    `db:"object_status"`
	RedundancyType      string    `db:"redundancy_type"`
	SourceType          string    `db:"source_type"`
	Checksums           string    `db:"checksums"`
	Tags                string    `db:"tags"`
	IsUpdating          bool      `db:"is_updating"`
	CreateAt            time.Time `db:"create_at"`
	UpdatedAt           time.Time `db:"update_at"`
	UpdatedBy           string    `db:"update_by"`
	Version             int64     `db:"version"`
	LocalVirtualGroupId uint32    `db:"local_virtual_group_id"`
}

func NewObjectsRow(
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
) ObjectsRow {
	return ObjectsRow{
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

func (o ObjectsRow) Equal(v ObjectsRow) bool {
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
