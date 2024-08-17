package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	query "github.com/cosmos/cosmos-sdk/types/query"
	_ "github.com/cosmos/gogoproto/gogoproto"
	grpc1 "github.com/cosmos/gogoproto/grpc"
	proto "github.com/cosmos/gogoproto/proto"
	types1 "github.com/forbole/bdjuno/v4/modules/storage/permission"
	types "github.com/forbole/bdjuno/v4/modules/storage/vitualgroup"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// QueryParamsRequest is request type for the Query/Params RPC method.
type QueryParamsRequest struct {
}

func (m *QueryParamsRequest) Reset()         { *m = QueryParamsRequest{} }
func (m *QueryParamsRequest) String() string { return proto.CompactTextString(m) }
func (*QueryParamsRequest) ProtoMessage()    {}
func (*QueryParamsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b1b80b580af04cb0, []int{0}
}
func (m *QueryParamsRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryParamsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryParamsRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryParamsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryParamsRequest.Merge(m, src)
}
func (m *QueryParamsRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryParamsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryParamsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryParamsRequest proto.InternalMessageInfo

type QueryHeadBucketRequest struct {
	BucketName string `protobuf:"bytes,1,opt,name=bucket_name,json=bucketName,proto3" json:"bucket_name,omitempty"`
}

func (m *QueryHeadBucketRequest) Reset()         { *m = QueryHeadBucketRequest{} }
func (m *QueryHeadBucketRequest) String() string { return proto.CompactTextString(m) }
func (*QueryHeadBucketRequest) ProtoMessage()    {}
func (*QueryHeadBucketRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b1b80b580af04cb0, []int{4}
}
func (m *QueryHeadBucketRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryHeadBucketRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryHeadBucketRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryHeadBucketRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryHeadBucketRequest.Merge(m, src)
}
func (m *QueryHeadBucketRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryHeadBucketRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryHeadBucketRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryHeadBucketRequest proto.InternalMessageInfo

func (m *QueryHeadBucketRequest) GetBucketName() string {
	if m != nil {
		return m.BucketName
	}
	return ""
}

type QueryHeadBucketByIdRequest struct {
	BucketId string `protobuf:"bytes,1,opt,name=bucket_id,json=bucketId,proto3" json:"bucket_id,omitempty"`
}

func (m *QueryHeadBucketByIdRequest) Reset()         { *m = QueryHeadBucketByIdRequest{} }
func (m *QueryHeadBucketByIdRequest) String() string { return proto.CompactTextString(m) }
func (*QueryHeadBucketByIdRequest) ProtoMessage()    {}
func (*QueryHeadBucketByIdRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b1b80b580af04cb0, []int{5}
}
func (m *QueryHeadBucketByIdRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryHeadBucketByIdRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryHeadBucketByIdRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryHeadBucketByIdRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryHeadBucketByIdRequest.Merge(m, src)
}
func (m *QueryHeadBucketByIdRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryHeadBucketByIdRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryHeadBucketByIdRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryHeadBucketByIdRequest proto.InternalMessageInfo

func (m *QueryHeadBucketByIdRequest) GetBucketId() string {
	if m != nil {
		return m.BucketId
	}
	return ""
}

type QueryHeadBucketResponse struct {
	BucketInfo *BucketInfo      `protobuf:"bytes,1,opt,name=bucket_info,json=bucketInfo,proto3" json:"bucket_info,omitempty"`
	ExtraInfo  *BucketExtraInfo `protobuf:"bytes,2,opt,name=extra_info,json=extraInfo,proto3" json:"extra_info,omitempty"`
}

func (m *QueryHeadBucketResponse) Reset()         { *m = QueryHeadBucketResponse{} }
func (m *QueryHeadBucketResponse) String() string { return proto.CompactTextString(m) }
func (*QueryHeadBucketResponse) ProtoMessage()    {}
func (*QueryHeadBucketResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b1b80b580af04cb0, []int{6}
}
func (m *QueryHeadBucketResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryHeadBucketResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryHeadBucketResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryHeadBucketResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryHeadBucketResponse.Merge(m, src)
}
func (m *QueryHeadBucketResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryHeadBucketResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryHeadBucketResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryHeadBucketResponse proto.InternalMessageInfo

func (m *QueryHeadBucketResponse) GetBucketInfo() *BucketInfo {
	if m != nil {
		return m.BucketInfo
	}
	return nil
}

func (m *QueryHeadBucketResponse) GetExtraInfo() *BucketExtraInfo {
	if m != nil {
		return m.ExtraInfo
	}
	return nil
}

type QueryHeadObjectRequest struct {
	BucketName string `protobuf:"bytes,1,opt,name=bucket_name,json=bucketName,proto3" json:"bucket_name,omitempty"`
	ObjectName string `protobuf:"bytes,2,opt,name=object_name,json=objectName,proto3" json:"object_name,omitempty"`
}

func (m *QueryHeadObjectRequest) Reset()         { *m = QueryHeadObjectRequest{} }
func (m *QueryHeadObjectRequest) String() string { return proto.CompactTextString(m) }
func (*QueryHeadObjectRequest) ProtoMessage()    {}
func (*QueryHeadObjectRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b1b80b580af04cb0, []int{7}
}
func (m *QueryHeadObjectRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryHeadObjectRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryHeadObjectRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryHeadObjectRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryHeadObjectRequest.Merge(m, src)
}
func (m *QueryHeadObjectRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryHeadObjectRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryHeadObjectRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryHeadObjectRequest proto.InternalMessageInfo

func (m *QueryHeadObjectRequest) GetBucketName() string {
	if m != nil {
		return m.BucketName
	}
	return ""
}

func (m *QueryHeadObjectRequest) GetObjectName() string {
	if m != nil {
		return m.ObjectName
	}
	return ""
}

type QueryHeadObjectByIdRequest struct {
	ObjectId string `protobuf:"bytes,1,opt,name=object_id,json=objectId,proto3" json:"object_id,omitempty"`
}

func (m *QueryHeadObjectByIdRequest) Reset()         { *m = QueryHeadObjectByIdRequest{} }
func (m *QueryHeadObjectByIdRequest) String() string { return proto.CompactTextString(m) }
func (*QueryHeadObjectByIdRequest) ProtoMessage()    {}
func (*QueryHeadObjectByIdRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b1b80b580af04cb0, []int{8}
}
func (m *QueryHeadObjectByIdRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryHeadObjectByIdRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryHeadObjectByIdRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryHeadObjectByIdRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryHeadObjectByIdRequest.Merge(m, src)
}
func (m *QueryHeadObjectByIdRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryHeadObjectByIdRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryHeadObjectByIdRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryHeadObjectByIdRequest proto.InternalMessageInfo

func (m *QueryHeadObjectByIdRequest) GetObjectId() string {
	if m != nil {
		return m.ObjectId
	}
	return ""
}

type QueryHeadShadowObjectRequest struct {
	BucketName string `protobuf:"bytes,1,opt,name=bucket_name,json=bucketName,proto3" json:"bucket_name,omitempty"`
	ObjectName string `protobuf:"bytes,2,opt,name=object_name,json=objectName,proto3" json:"object_name,omitempty"`
}

func (m *QueryHeadShadowObjectRequest) Reset()         { *m = QueryHeadShadowObjectRequest{} }
func (m *QueryHeadShadowObjectRequest) String() string { return proto.CompactTextString(m) }
func (*QueryHeadShadowObjectRequest) ProtoMessage()    {}
func (*QueryHeadShadowObjectRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b1b80b580af04cb0, []int{9}
}
func (m *QueryHeadShadowObjectRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryHeadShadowObjectRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryHeadShadowObjectRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryHeadShadowObjectRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryHeadShadowObjectRequest.Merge(m, src)
}
func (m *QueryHeadShadowObjectRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryHeadShadowObjectRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryHeadShadowObjectRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryHeadShadowObjectRequest proto.InternalMessageInfo

func (m *QueryHeadShadowObjectRequest) GetBucketName() string {
	if m != nil {
		return m.BucketName
	}
	return ""
}

func (m *QueryHeadShadowObjectRequest) GetObjectName() string {
	if m != nil {
		return m.ObjectName
	}
	return ""
}

type QueryHeadObjectResponse struct {
	ObjectInfo         *ObjectInfo               `protobuf:"bytes,1,opt,name=object_info,json=objectInfo,proto3" json:"object_info,omitempty"`
	GlobalVirtualGroup *types.GlobalVirtualGroup `protobuf:"bytes,2,opt,name=global_virtual_group,json=globalVirtualGroup,proto3" json:"global_virtual_group,omitempty"`
}

func (m *QueryHeadObjectResponse) Reset()         { *m = QueryHeadObjectResponse{} }
func (m *QueryHeadObjectResponse) String() string { return proto.CompactTextString(m) }
func (*QueryHeadObjectResponse) ProtoMessage()    {}
func (*QueryHeadObjectResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b1b80b580af04cb0, []int{10}
}
func (m *QueryHeadObjectResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryHeadObjectResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryHeadObjectResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryHeadObjectResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryHeadObjectResponse.Merge(m, src)
}
func (m *QueryHeadObjectResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryHeadObjectResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryHeadObjectResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryHeadObjectResponse proto.InternalMessageInfo

func (m *QueryHeadObjectResponse) GetObjectInfo() *ObjectInfo {
	if m != nil {
		return m.ObjectInfo
	}
	return nil
}

func (m *QueryHeadObjectResponse) GetGlobalVirtualGroup() *types.GlobalVirtualGroup {
	if m != nil {
		return m.GlobalVirtualGroup
	}
	return nil
}

type QueryHeadShadowObjectResponse struct {
	ObjectInfo *ShadowObjectInfo `protobuf:"bytes,1,opt,name=object_info,json=objectInfo,proto3" json:"object_info,omitempty"`
}

func (m *QueryHeadShadowObjectResponse) Reset()         { *m = QueryHeadShadowObjectResponse{} }
func (m *QueryHeadShadowObjectResponse) String() string { return proto.CompactTextString(m) }
func (*QueryHeadShadowObjectResponse) ProtoMessage()    {}
func (*QueryHeadShadowObjectResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b1b80b580af04cb0, []int{11}
}
func (m *QueryHeadShadowObjectResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryHeadShadowObjectResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryHeadShadowObjectResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryHeadShadowObjectResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryHeadShadowObjectResponse.Merge(m, src)
}
func (m *QueryHeadShadowObjectResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryHeadShadowObjectResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryHeadShadowObjectResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryHeadShadowObjectResponse proto.InternalMessageInfo

func (m *QueryHeadShadowObjectResponse) GetObjectInfo() *ShadowObjectInfo {
	if m != nil {
		return m.ObjectInfo
	}
	return nil
}

type QueryListBucketsRequest struct {
	Pagination *query.PageRequest `protobuf:"bytes,1,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (m *QueryListBucketsRequest) Reset()         { *m = QueryListBucketsRequest{} }
func (m *QueryListBucketsRequest) String() string { return proto.CompactTextString(m) }
func (*QueryListBucketsRequest) ProtoMessage()    {}
func (*QueryListBucketsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b1b80b580af04cb0, []int{12}
}
func (m *QueryListBucketsRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryListBucketsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryListBucketsRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryListBucketsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryListBucketsRequest.Merge(m, src)
}
func (m *QueryListBucketsRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryListBucketsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryListBucketsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryListBucketsRequest proto.InternalMessageInfo

func (m *QueryListBucketsRequest) GetPagination() *query.PageRequest {
	if m != nil {
		return m.Pagination
	}
	return nil
}

type QueryListBucketsResponse struct {
	BucketInfos []*BucketInfo       `protobuf:"bytes,1,rep,name=bucket_infos,json=bucketInfos,proto3" json:"bucket_infos,omitempty"`
	Pagination  *query.PageResponse `protobuf:"bytes,2,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (m *QueryListBucketsResponse) Reset()         { *m = QueryListBucketsResponse{} }
func (m *QueryListBucketsResponse) String() string { return proto.CompactTextString(m) }
func (*QueryListBucketsResponse) ProtoMessage()    {}
func (*QueryListBucketsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b1b80b580af04cb0, []int{13}
}
func (m *QueryListBucketsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryListBucketsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryListBucketsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryListBucketsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryListBucketsResponse.Merge(m, src)
}
func (m *QueryListBucketsResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryListBucketsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryListBucketsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryListBucketsResponse proto.InternalMessageInfo

func (m *QueryListBucketsResponse) GetBucketInfos() []*BucketInfo {
	if m != nil {
		return m.BucketInfos
	}
	return nil
}

func (m *QueryListBucketsResponse) GetPagination() *query.PageResponse {
	if m != nil {
		return m.Pagination
	}
	return nil
}

type QueryListObjectsRequest struct {
	Pagination *query.PageRequest `protobuf:"bytes,1,opt,name=pagination,proto3" json:"pagination,omitempty"`
	BucketName string             `protobuf:"bytes,2,opt,name=bucket_name,json=bucketName,proto3" json:"bucket_name,omitempty"`
}

func (m *QueryListObjectsRequest) Reset()         { *m = QueryListObjectsRequest{} }
func (m *QueryListObjectsRequest) String() string { return proto.CompactTextString(m) }
func (*QueryListObjectsRequest) ProtoMessage()    {}
func (*QueryListObjectsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b1b80b580af04cb0, []int{14}
}
func (m *QueryListObjectsRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryListObjectsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryListObjectsRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryListObjectsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryListObjectsRequest.Merge(m, src)
}
func (m *QueryListObjectsRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryListObjectsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryListObjectsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryListObjectsRequest proto.InternalMessageInfo

func (m *QueryListObjectsRequest) GetPagination() *query.PageRequest {
	if m != nil {
		return m.Pagination
	}
	return nil
}

func (m *QueryListObjectsRequest) GetBucketName() string {
	if m != nil {
		return m.BucketName
	}
	return ""
}

type QueryListObjectsByBucketIdRequest struct {
	Pagination *query.PageRequest `protobuf:"bytes,1,opt,name=pagination,proto3" json:"pagination,omitempty"`
	BucketId   string             `protobuf:"bytes,2,opt,name=bucket_id,json=bucketId,proto3" json:"bucket_id,omitempty"`
}

func (m *QueryListObjectsByBucketIdRequest) Reset()         { *m = QueryListObjectsByBucketIdRequest{} }
func (m *QueryListObjectsByBucketIdRequest) String() string { return proto.CompactTextString(m) }
func (*QueryListObjectsByBucketIdRequest) ProtoMessage()    {}
func (*QueryListObjectsByBucketIdRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b1b80b580af04cb0, []int{15}
}
func (m *QueryListObjectsByBucketIdRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryListObjectsByBucketIdRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryListObjectsByBucketIdRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryListObjectsByBucketIdRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryListObjectsByBucketIdRequest.Merge(m, src)
}
func (m *QueryListObjectsByBucketIdRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryListObjectsByBucketIdRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryListObjectsByBucketIdRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryListObjectsByBucketIdRequest proto.InternalMessageInfo

func (m *QueryListObjectsByBucketIdRequest) GetPagination() *query.PageRequest {
	if m != nil {
		return m.Pagination
	}
	return nil
}

func (m *QueryListObjectsByBucketIdRequest) GetBucketId() string {
	if m != nil {
		return m.BucketId
	}
	return ""
}

type QueryListObjectsResponse struct {
	ObjectInfos []*ObjectInfo       `protobuf:"bytes,1,rep,name=object_infos,json=objectInfos,proto3" json:"object_infos,omitempty"`
	Pagination  *query.PageResponse `protobuf:"bytes,2,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (m *QueryListObjectsResponse) Reset()         { *m = QueryListObjectsResponse{} }
func (m *QueryListObjectsResponse) String() string { return proto.CompactTextString(m) }
func (*QueryListObjectsResponse) ProtoMessage()    {}
func (*QueryListObjectsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b1b80b580af04cb0, []int{16}
}
func (m *QueryListObjectsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryListObjectsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryListObjectsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryListObjectsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryListObjectsResponse.Merge(m, src)
}
func (m *QueryListObjectsResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryListObjectsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryListObjectsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryListObjectsResponse proto.InternalMessageInfo

func (m *QueryListObjectsResponse) GetObjectInfos() []*ObjectInfo {
	if m != nil {
		return m.ObjectInfos
	}
	return nil
}

func (m *QueryListObjectsResponse) GetPagination() *query.PageResponse {
	if m != nil {
		return m.Pagination
	}
	return nil
}

type QueryNFTRequest struct {
	TokenId string `protobuf:"bytes,1,opt,name=token_id,json=tokenId,proto3" json:"token_id,omitempty"`
}

func (m *QueryNFTRequest) Reset()         { *m = QueryNFTRequest{} }
func (m *QueryNFTRequest) String() string { return proto.CompactTextString(m) }
func (*QueryNFTRequest) ProtoMessage()    {}
func (*QueryNFTRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b1b80b580af04cb0, []int{17}
}
func (m *QueryNFTRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryNFTRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryNFTRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryNFTRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryNFTRequest.Merge(m, src)
}
func (m *QueryNFTRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryNFTRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryNFTRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryNFTRequest proto.InternalMessageInfo

func (m *QueryNFTRequest) GetTokenId() string {
	if m != nil {
		return m.TokenId
	}
	return ""
}

type QueryBucketNFTResponse struct {
	MetaData *BucketMetaData `protobuf:"bytes,1,opt,name=meta_data,json=metaData,proto3" json:"meta_data,omitempty"`
}

func (m *QueryBucketNFTResponse) Reset()         { *m = QueryBucketNFTResponse{} }
func (m *QueryBucketNFTResponse) String() string { return proto.CompactTextString(m) }
func (*QueryBucketNFTResponse) ProtoMessage()    {}
func (*QueryBucketNFTResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b1b80b580af04cb0, []int{18}
}
func (m *QueryBucketNFTResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryBucketNFTResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryBucketNFTResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryBucketNFTResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryBucketNFTResponse.Merge(m, src)
}
func (m *QueryBucketNFTResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryBucketNFTResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryBucketNFTResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryBucketNFTResponse proto.InternalMessageInfo

func (m *QueryBucketNFTResponse) GetMetaData() *BucketMetaData {
	if m != nil {
		return m.MetaData
	}
	return nil
}

type QueryObjectNFTResponse struct {
	MetaData *ObjectMetaData `protobuf:"bytes,1,opt,name=meta_data,json=metaData,proto3" json:"meta_data,omitempty"`
}

func (m *QueryObjectNFTResponse) Reset()         { *m = QueryObjectNFTResponse{} }
func (m *QueryObjectNFTResponse) String() string { return proto.CompactTextString(m) }
func (*QueryObjectNFTResponse) ProtoMessage()    {}
func (*QueryObjectNFTResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b1b80b580af04cb0, []int{19}
}
func (m *QueryObjectNFTResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryObjectNFTResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryObjectNFTResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryObjectNFTResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryObjectNFTResponse.Merge(m, src)
}
func (m *QueryObjectNFTResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryObjectNFTResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryObjectNFTResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryObjectNFTResponse proto.InternalMessageInfo

func (m *QueryObjectNFTResponse) GetMetaData() *ObjectMetaData {
	if m != nil {
		return m.MetaData
	}
	return nil
}

type QueryGroupNFTResponse struct {
	MetaData *GroupMetaData `protobuf:"bytes,1,opt,name=meta_data,json=metaData,proto3" json:"meta_data,omitempty"`
}

func (m *QueryGroupNFTResponse) Reset()         { *m = QueryGroupNFTResponse{} }
func (m *QueryGroupNFTResponse) String() string { return proto.CompactTextString(m) }
func (*QueryGroupNFTResponse) ProtoMessage()    {}
func (*QueryGroupNFTResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b1b80b580af04cb0, []int{20}
}
func (m *QueryGroupNFTResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryGroupNFTResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryGroupNFTResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryGroupNFTResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryGroupNFTResponse.Merge(m, src)
}
func (m *QueryGroupNFTResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryGroupNFTResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryGroupNFTResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryGroupNFTResponse proto.InternalMessageInfo

func (m *QueryGroupNFTResponse) GetMetaData() *GroupMetaData {
	if m != nil {
		return m.MetaData
	}
	return nil
}

type QueryPolicyForAccountRequest struct {
	Resource         string `protobuf:"bytes,1,opt,name=resource,proto3" json:"resource,omitempty"`
	PrincipalAddress string `protobuf:"bytes,2,opt,name=principal_address,json=principalAddress,proto3" json:"principal_address,omitempty"`
}

func (m *QueryPolicyForAccountRequest) Reset()         { *m = QueryPolicyForAccountRequest{} }
func (m *QueryPolicyForAccountRequest) String() string { return proto.CompactTextString(m) }
func (*QueryPolicyForAccountRequest) ProtoMessage()    {}
func (*QueryPolicyForAccountRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b1b80b580af04cb0, []int{21}
}
func (m *QueryPolicyForAccountRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryPolicyForAccountRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryPolicyForAccountRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryPolicyForAccountRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryPolicyForAccountRequest.Merge(m, src)
}
func (m *QueryPolicyForAccountRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryPolicyForAccountRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryPolicyForAccountRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryPolicyForAccountRequest proto.InternalMessageInfo

func (m *QueryPolicyForAccountRequest) GetResource() string {
	if m != nil {
		return m.Resource
	}
	return ""
}

func (m *QueryPolicyForAccountRequest) GetPrincipalAddress() string {
	if m != nil {
		return m.PrincipalAddress
	}
	return ""
}

type QueryPolicyForAccountResponse struct {
	Policy *types1.Policy `protobuf:"bytes,1,opt,name=policy,proto3" json:"policy,omitempty"`
}

func (m *QueryPolicyForAccountResponse) Reset()         { *m = QueryPolicyForAccountResponse{} }
func (m *QueryPolicyForAccountResponse) String() string { return proto.CompactTextString(m) }
func (*QueryPolicyForAccountResponse) ProtoMessage()    {}
func (*QueryPolicyForAccountResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b1b80b580af04cb0, []int{22}
}
func (m *QueryPolicyForAccountResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryPolicyForAccountResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryPolicyForAccountResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryPolicyForAccountResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryPolicyForAccountResponse.Merge(m, src)
}
func (m *QueryPolicyForAccountResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryPolicyForAccountResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryPolicyForAccountResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryPolicyForAccountResponse proto.InternalMessageInfo

func (m *QueryPolicyForAccountResponse) GetPolicy() *types1.Policy {
	if m != nil {
		return m.Policy
	}
	return nil
}

type QueryVerifyPermissionRequest struct {
	Operator   string            `protobuf:"bytes,1,opt,name=operator,proto3" json:"operator,omitempty"`
	BucketName string            `protobuf:"bytes,2,opt,name=bucket_name,json=bucketName,proto3" json:"bucket_name,omitempty"`
	ObjectName string            `protobuf:"bytes,3,opt,name=object_name,json=objectName,proto3" json:"object_name,omitempty"`
	ActionType types1.ActionType `protobuf:"varint,4,opt,name=action_type,json=actionType,proto3,enum=greenfield.permission.ActionType" json:"action_type,omitempty"`
}

func (m *QueryVerifyPermissionRequest) Reset()         { *m = QueryVerifyPermissionRequest{} }
func (m *QueryVerifyPermissionRequest) String() string { return proto.CompactTextString(m) }
func (*QueryVerifyPermissionRequest) ProtoMessage()    {}
func (*QueryVerifyPermissionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b1b80b580af04cb0, []int{23}
}
func (m *QueryVerifyPermissionRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryVerifyPermissionRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryVerifyPermissionRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryVerifyPermissionRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryVerifyPermissionRequest.Merge(m, src)
}
func (m *QueryVerifyPermissionRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryVerifyPermissionRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryVerifyPermissionRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryVerifyPermissionRequest proto.InternalMessageInfo

func (m *QueryVerifyPermissionRequest) GetOperator() string {
	if m != nil {
		return m.Operator
	}
	return ""
}

func (m *QueryVerifyPermissionRequest) GetBucketName() string {
	if m != nil {
		return m.BucketName
	}
	return ""
}

func (m *QueryVerifyPermissionRequest) GetObjectName() string {
	if m != nil {
		return m.ObjectName
	}
	return ""
}

func (m *QueryVerifyPermissionRequest) GetActionType() types1.ActionType {
	if m != nil {
		return m.ActionType
	}
	return types1.ACTION_UNSPECIFIED
}

type QueryVerifyPermissionResponse struct {
	Effect types1.Effect `protobuf:"varint,1,opt,name=effect,proto3,enum=greenfield.permission.Effect" json:"effect,omitempty"`
}

func (m *QueryVerifyPermissionResponse) Reset()         { *m = QueryVerifyPermissionResponse{} }
func (m *QueryVerifyPermissionResponse) String() string { return proto.CompactTextString(m) }
func (*QueryVerifyPermissionResponse) ProtoMessage()    {}
func (*QueryVerifyPermissionResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b1b80b580af04cb0, []int{24}
}
func (m *QueryVerifyPermissionResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryVerifyPermissionResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryVerifyPermissionResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryVerifyPermissionResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryVerifyPermissionResponse.Merge(m, src)
}
func (m *QueryVerifyPermissionResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryVerifyPermissionResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryVerifyPermissionResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryVerifyPermissionResponse proto.InternalMessageInfo

func (m *QueryVerifyPermissionResponse) GetEffect() types1.Effect {
	if m != nil {
		return m.Effect
	}
	return types1.EFFECT_UNSPECIFIED
}

type QueryHeadGroupRequest struct {
	GroupOwner string `protobuf:"bytes,1,opt,name=group_owner,json=groupOwner,proto3" json:"group_owner,omitempty"`
	GroupName  string `protobuf:"bytes,2,opt,name=group_name,json=groupName,proto3" json:"group_name,omitempty"`
}

func (m *QueryHeadGroupRequest) Reset()         { *m = QueryHeadGroupRequest{} }
func (m *QueryHeadGroupRequest) String() string { return proto.CompactTextString(m) }
func (*QueryHeadGroupRequest) ProtoMessage()    {}
func (*QueryHeadGroupRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b1b80b580af04cb0, []int{25}
}
func (m *QueryHeadGroupRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryHeadGroupRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryHeadGroupRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryHeadGroupRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryHeadGroupRequest.Merge(m, src)
}
func (m *QueryHeadGroupRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryHeadGroupRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryHeadGroupRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryHeadGroupRequest proto.InternalMessageInfo

func (m *QueryHeadGroupRequest) GetGroupOwner() string {
	if m != nil {
		return m.GroupOwner
	}
	return ""
}

func (m *QueryHeadGroupRequest) GetGroupName() string {
	if m != nil {
		return m.GroupName
	}
	return ""
}

type QueryHeadGroupResponse struct {
	GroupInfo *GroupInfo `protobuf:"bytes,1,opt,name=group_info,json=groupInfo,proto3" json:"group_info,omitempty"`
}

func (m *QueryHeadGroupResponse) Reset()         { *m = QueryHeadGroupResponse{} }
func (m *QueryHeadGroupResponse) String() string { return proto.CompactTextString(m) }
func (*QueryHeadGroupResponse) ProtoMessage()    {}
func (*QueryHeadGroupResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b1b80b580af04cb0, []int{26}
}
func (m *QueryHeadGroupResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryHeadGroupResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryHeadGroupResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryHeadGroupResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryHeadGroupResponse.Merge(m, src)
}
func (m *QueryHeadGroupResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryHeadGroupResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryHeadGroupResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryHeadGroupResponse proto.InternalMessageInfo

func (m *QueryHeadGroupResponse) GetGroupInfo() *GroupInfo {
	if m != nil {
		return m.GroupInfo
	}
	return nil
}

type QueryListGroupsRequest struct {
	Pagination *query.PageRequest `protobuf:"bytes,1,opt,name=pagination,proto3" json:"pagination,omitempty"`
	GroupOwner string             `protobuf:"bytes,2,opt,name=group_owner,json=groupOwner,proto3" json:"group_owner,omitempty"`
}

func (m *QueryListGroupsRequest) Reset()         { *m = QueryListGroupsRequest{} }
func (m *QueryListGroupsRequest) String() string { return proto.CompactTextString(m) }
func (*QueryListGroupsRequest) ProtoMessage()    {}
func (*QueryListGroupsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b1b80b580af04cb0, []int{27}
}
func (m *QueryListGroupsRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryListGroupsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryListGroupsRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryListGroupsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryListGroupsRequest.Merge(m, src)
}
func (m *QueryListGroupsRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryListGroupsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryListGroupsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryListGroupsRequest proto.InternalMessageInfo

func (m *QueryListGroupsRequest) GetPagination() *query.PageRequest {
	if m != nil {
		return m.Pagination
	}
	return nil
}

func (m *QueryListGroupsRequest) GetGroupOwner() string {
	if m != nil {
		return m.GroupOwner
	}
	return ""
}

type QueryListGroupsResponse struct {
	Pagination *query.PageResponse `protobuf:"bytes,1,opt,name=pagination,proto3" json:"pagination,omitempty"`
	GroupInfos []*GroupInfo        `protobuf:"bytes,2,rep,name=group_infos,json=groupInfos,proto3" json:"group_infos,omitempty"`
}

func (m *QueryListGroupsResponse) Reset()         { *m = QueryListGroupsResponse{} }
func (m *QueryListGroupsResponse) String() string { return proto.CompactTextString(m) }
func (*QueryListGroupsResponse) ProtoMessage()    {}
func (*QueryListGroupsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b1b80b580af04cb0, []int{28}
}
func (m *QueryListGroupsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryListGroupsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryListGroupsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryListGroupsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryListGroupsResponse.Merge(m, src)
}
func (m *QueryListGroupsResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryListGroupsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryListGroupsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryListGroupsResponse proto.InternalMessageInfo

func (m *QueryListGroupsResponse) GetPagination() *query.PageResponse {
	if m != nil {
		return m.Pagination
	}
	return nil
}

func (m *QueryListGroupsResponse) GetGroupInfos() []*GroupInfo {
	if m != nil {
		return m.GroupInfos
	}
	return nil
}

type QueryHeadGroupMemberRequest struct {
	Member     string `protobuf:"bytes,1,opt,name=member,proto3" json:"member,omitempty"`
	GroupOwner string `protobuf:"bytes,2,opt,name=group_owner,json=groupOwner,proto3" json:"group_owner,omitempty"`
	GroupName  string `protobuf:"bytes,3,opt,name=group_name,json=groupName,proto3" json:"group_name,omitempty"`
}

func (m *QueryHeadGroupMemberRequest) Reset()         { *m = QueryHeadGroupMemberRequest{} }
func (m *QueryHeadGroupMemberRequest) String() string { return proto.CompactTextString(m) }
func (*QueryHeadGroupMemberRequest) ProtoMessage()    {}
func (*QueryHeadGroupMemberRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b1b80b580af04cb0, []int{29}
}
func (m *QueryHeadGroupMemberRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryHeadGroupMemberRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryHeadGroupMemberRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryHeadGroupMemberRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryHeadGroupMemberRequest.Merge(m, src)
}
func (m *QueryHeadGroupMemberRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryHeadGroupMemberRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryHeadGroupMemberRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryHeadGroupMemberRequest proto.InternalMessageInfo

func (m *QueryHeadGroupMemberRequest) GetMember() string {
	if m != nil {
		return m.Member
	}
	return ""
}

func (m *QueryHeadGroupMemberRequest) GetGroupOwner() string {
	if m != nil {
		return m.GroupOwner
	}
	return ""
}

func (m *QueryHeadGroupMemberRequest) GetGroupName() string {
	if m != nil {
		return m.GroupName
	}
	return ""
}

type QueryHeadGroupMemberResponse struct {
	GroupMember *types1.GroupMember `protobuf:"bytes,1,opt,name=group_member,json=groupMember,proto3" json:"group_member,omitempty"`
}

func (m *QueryHeadGroupMemberResponse) Reset()         { *m = QueryHeadGroupMemberResponse{} }
func (m *QueryHeadGroupMemberResponse) String() string { return proto.CompactTextString(m) }
func (*QueryHeadGroupMemberResponse) ProtoMessage()    {}
func (*QueryHeadGroupMemberResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b1b80b580af04cb0, []int{30}
}
func (m *QueryHeadGroupMemberResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryHeadGroupMemberResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryHeadGroupMemberResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryHeadGroupMemberResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryHeadGroupMemberResponse.Merge(m, src)
}
func (m *QueryHeadGroupMemberResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryHeadGroupMemberResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryHeadGroupMemberResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryHeadGroupMemberResponse proto.InternalMessageInfo

func (m *QueryHeadGroupMemberResponse) GetGroupMember() *types1.GroupMember {
	if m != nil {
		return m.GroupMember
	}
	return nil
}

type QueryPolicyForGroupRequest struct {
	Resource         string `protobuf:"bytes,1,opt,name=resource,proto3" json:"resource,omitempty"`
	PrincipalGroupId string `protobuf:"bytes,2,opt,name=principal_group_id,json=principalGroupId,proto3" json:"principal_group_id,omitempty"`
}

func (m *QueryPolicyForGroupRequest) Reset()         { *m = QueryPolicyForGroupRequest{} }
func (m *QueryPolicyForGroupRequest) String() string { return proto.CompactTextString(m) }
func (*QueryPolicyForGroupRequest) ProtoMessage()    {}
func (*QueryPolicyForGroupRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b1b80b580af04cb0, []int{31}
}
func (m *QueryPolicyForGroupRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryPolicyForGroupRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryPolicyForGroupRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryPolicyForGroupRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryPolicyForGroupRequest.Merge(m, src)
}
func (m *QueryPolicyForGroupRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryPolicyForGroupRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryPolicyForGroupRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryPolicyForGroupRequest proto.InternalMessageInfo

func (m *QueryPolicyForGroupRequest) GetResource() string {
	if m != nil {
		return m.Resource
	}
	return ""
}

func (m *QueryPolicyForGroupRequest) GetPrincipalGroupId() string {
	if m != nil {
		return m.PrincipalGroupId
	}
	return ""
}

type QueryPolicyForGroupResponse struct {
	Policy *types1.Policy `protobuf:"bytes,1,opt,name=policy,proto3" json:"policy,omitempty"`
}

func (m *QueryPolicyForGroupResponse) Reset()         { *m = QueryPolicyForGroupResponse{} }
func (m *QueryPolicyForGroupResponse) String() string { return proto.CompactTextString(m) }
func (*QueryPolicyForGroupResponse) ProtoMessage()    {}
func (*QueryPolicyForGroupResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b1b80b580af04cb0, []int{32}
}
func (m *QueryPolicyForGroupResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryPolicyForGroupResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryPolicyForGroupResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryPolicyForGroupResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryPolicyForGroupResponse.Merge(m, src)
}
func (m *QueryPolicyForGroupResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryPolicyForGroupResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryPolicyForGroupResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryPolicyForGroupResponse proto.InternalMessageInfo

func (m *QueryPolicyForGroupResponse) GetPolicy() *types1.Policy {
	if m != nil {
		return m.Policy
	}
	return nil
}

type QueryPolicyByIdRequest struct {
	PolicyId string `protobuf:"bytes,1,opt,name=policy_id,json=policyId,proto3" json:"policy_id,omitempty"`
}

func (m *QueryPolicyByIdRequest) Reset()         { *m = QueryPolicyByIdRequest{} }
func (m *QueryPolicyByIdRequest) String() string { return proto.CompactTextString(m) }
func (*QueryPolicyByIdRequest) ProtoMessage()    {}
func (*QueryPolicyByIdRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b1b80b580af04cb0, []int{33}
}
func (m *QueryPolicyByIdRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryPolicyByIdRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryPolicyByIdRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryPolicyByIdRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryPolicyByIdRequest.Merge(m, src)
}
func (m *QueryPolicyByIdRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryPolicyByIdRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryPolicyByIdRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryPolicyByIdRequest proto.InternalMessageInfo

func (m *QueryPolicyByIdRequest) GetPolicyId() string {
	if m != nil {
		return m.PolicyId
	}
	return ""
}

type QueryPolicyByIdResponse struct {
	Policy *types1.Policy `protobuf:"bytes,1,opt,name=policy,proto3" json:"policy,omitempty"`
}

func (m *QueryPolicyByIdResponse) Reset()         { *m = QueryPolicyByIdResponse{} }
func (m *QueryPolicyByIdResponse) String() string { return proto.CompactTextString(m) }
func (*QueryPolicyByIdResponse) ProtoMessage()    {}
func (*QueryPolicyByIdResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b1b80b580af04cb0, []int{34}
}
func (m *QueryPolicyByIdResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryPolicyByIdResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryPolicyByIdResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryPolicyByIdResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryPolicyByIdResponse.Merge(m, src)
}
func (m *QueryPolicyByIdResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryPolicyByIdResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryPolicyByIdResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryPolicyByIdResponse proto.InternalMessageInfo

func (m *QueryPolicyByIdResponse) GetPolicy() *types1.Policy {
	if m != nil {
		return m.Policy
	}
	return nil
}

type QueryLockFeeRequest struct {
	// primary_sp_address is the address of the primary sp.
	PrimarySpAddress string `protobuf:"bytes,1,opt,name=primary_sp_address,json=primarySpAddress,proto3" json:"primary_sp_address,omitempty"`
	// create_at define the block timestamp when the object created.
	CreateAt int64 `protobuf:"varint,2,opt,name=create_at,json=createAt,proto3" json:"create_at,omitempty"`
	// payloadSize is the total size of the object payload
	PayloadSize uint64 `protobuf:"varint,3,opt,name=payload_size,json=payloadSize,proto3" json:"payload_size,omitempty"`
}

func (m *QueryLockFeeRequest) Reset()         { *m = QueryLockFeeRequest{} }
func (m *QueryLockFeeRequest) String() string { return proto.CompactTextString(m) }
func (*QueryLockFeeRequest) ProtoMessage()    {}
func (*QueryLockFeeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b1b80b580af04cb0, []int{35}
}
func (m *QueryLockFeeRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryLockFeeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryLockFeeRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryLockFeeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryLockFeeRequest.Merge(m, src)
}
func (m *QueryLockFeeRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryLockFeeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryLockFeeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryLockFeeRequest proto.InternalMessageInfo

func (m *QueryLockFeeRequest) GetPrimarySpAddress() string {
	if m != nil {
		return m.PrimarySpAddress
	}
	return ""
}

func (m *QueryLockFeeRequest) GetCreateAt() int64 {
	if m != nil {
		return m.CreateAt
	}
	return 0
}

func (m *QueryLockFeeRequest) GetPayloadSize() uint64 {
	if m != nil {
		return m.PayloadSize
	}
	return 0
}

type QueryLockFeeResponse struct {
	Amount github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,1,opt,name=amount,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"amount"`
}

func (m *QueryLockFeeResponse) Reset()         { *m = QueryLockFeeResponse{} }
func (m *QueryLockFeeResponse) String() string { return proto.CompactTextString(m) }
func (*QueryLockFeeResponse) ProtoMessage()    {}
func (*QueryLockFeeResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b1b80b580af04cb0, []int{36}
}
func (m *QueryLockFeeResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryLockFeeResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryLockFeeResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryLockFeeResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryLockFeeResponse.Merge(m, src)
}
func (m *QueryLockFeeResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryLockFeeResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryLockFeeResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryLockFeeResponse proto.InternalMessageInfo

type QueryHeadBucketExtraRequest struct {
	BucketName string `protobuf:"bytes,1,opt,name=bucket_name,json=bucketName,proto3" json:"bucket_name,omitempty"`
}

func (m *QueryHeadBucketExtraRequest) Reset()         { *m = QueryHeadBucketExtraRequest{} }
func (m *QueryHeadBucketExtraRequest) String() string { return proto.CompactTextString(m) }
func (*QueryHeadBucketExtraRequest) ProtoMessage()    {}
func (*QueryHeadBucketExtraRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b1b80b580af04cb0, []int{37}
}
func (m *QueryHeadBucketExtraRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryHeadBucketExtraRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryHeadBucketExtraRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryHeadBucketExtraRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryHeadBucketExtraRequest.Merge(m, src)
}
func (m *QueryHeadBucketExtraRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryHeadBucketExtraRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryHeadBucketExtraRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryHeadBucketExtraRequest proto.InternalMessageInfo

func (m *QueryHeadBucketExtraRequest) GetBucketName() string {
	if m != nil {
		return m.BucketName
	}
	return ""
}

type QueryHeadBucketExtraResponse struct {
	ExtraInfo *InternalBucketInfo `protobuf:"bytes,1,opt,name=extra_info,json=extraInfo,proto3" json:"extra_info,omitempty"`
}

func (m *QueryHeadBucketExtraResponse) Reset()         { *m = QueryHeadBucketExtraResponse{} }
func (m *QueryHeadBucketExtraResponse) String() string { return proto.CompactTextString(m) }
func (*QueryHeadBucketExtraResponse) ProtoMessage()    {}
func (*QueryHeadBucketExtraResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b1b80b580af04cb0, []int{38}
}
func (m *QueryHeadBucketExtraResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryHeadBucketExtraResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryHeadBucketExtraResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryHeadBucketExtraResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryHeadBucketExtraResponse.Merge(m, src)
}
func (m *QueryHeadBucketExtraResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryHeadBucketExtraResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryHeadBucketExtraResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryHeadBucketExtraResponse proto.InternalMessageInfo

func (m *QueryHeadBucketExtraResponse) GetExtraInfo() *InternalBucketInfo {
	if m != nil {
		return m.ExtraInfo
	}
	return nil
}

type QueryIsPriceChangedRequest struct {
	BucketName string `protobuf:"bytes,1,opt,name=bucket_name,json=bucketName,proto3" json:"bucket_name,omitempty"`
}

func (m *QueryIsPriceChangedRequest) Reset()         { *m = QueryIsPriceChangedRequest{} }
func (m *QueryIsPriceChangedRequest) String() string { return proto.CompactTextString(m) }
func (*QueryIsPriceChangedRequest) ProtoMessage()    {}
func (*QueryIsPriceChangedRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b1b80b580af04cb0, []int{39}
}
func (m *QueryIsPriceChangedRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryIsPriceChangedRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryIsPriceChangedRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryIsPriceChangedRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryIsPriceChangedRequest.Merge(m, src)
}
func (m *QueryIsPriceChangedRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryIsPriceChangedRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryIsPriceChangedRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryIsPriceChangedRequest proto.InternalMessageInfo

func (m *QueryIsPriceChangedRequest) GetBucketName() string {
	if m != nil {
		return m.BucketName
	}
	return ""
}

type QueryIsPriceChangedResponse struct {
	Changed                    bool                                   `protobuf:"varint,1,opt,name=changed,proto3" json:"changed,omitempty"`
	CurrentReadPrice           github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,2,opt,name=current_read_price,json=currentReadPrice,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"current_read_price"`
	CurrentPrimaryStorePrice   github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,3,opt,name=current_primary_store_price,json=currentPrimaryStorePrice,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"current_primary_store_price"`
	CurrentSecondaryStorePrice github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,4,opt,name=current_secondary_store_price,json=currentSecondaryStorePrice,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"current_secondary_store_price"`
	CurrentValidatorTaxRate    github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,5,opt,name=current_validator_tax_rate,json=currentValidatorTaxRate,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"current_validator_tax_rate"`
	NewReadPrice               github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,6,opt,name=new_read_price,json=newReadPrice,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"new_read_price"`
	NewPrimaryStorePrice       github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,7,opt,name=new_primary_store_price,json=newPrimaryStorePrice,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"new_primary_store_price"`
	NewSecondaryStorePrice     github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,8,opt,name=new_secondary_store_price,json=newSecondaryStorePrice,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"new_secondary_store_price"`
	NewValidatorTaxRate        github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,9,opt,name=new_validator_tax_rate,json=newValidatorTaxRate,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"new_validator_tax_rate"`
}

func (m *QueryIsPriceChangedResponse) Reset()         { *m = QueryIsPriceChangedResponse{} }
func (m *QueryIsPriceChangedResponse) String() string { return proto.CompactTextString(m) }
func (*QueryIsPriceChangedResponse) ProtoMessage()    {}
func (*QueryIsPriceChangedResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b1b80b580af04cb0, []int{40}
}
func (m *QueryIsPriceChangedResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryIsPriceChangedResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryIsPriceChangedResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryIsPriceChangedResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryIsPriceChangedResponse.Merge(m, src)
}
func (m *QueryIsPriceChangedResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryIsPriceChangedResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryIsPriceChangedResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryIsPriceChangedResponse proto.InternalMessageInfo

func (m *QueryIsPriceChangedResponse) GetChanged() bool {
	if m != nil {
		return m.Changed
	}
	return false
}

type QueryQuoteUpdateTimeRequest struct {
	BucketName string `protobuf:"bytes,1,opt,name=bucket_name,json=bucketName,proto3" json:"bucket_name,omitempty"`
}

func (m *QueryQuoteUpdateTimeRequest) Reset()         { *m = QueryQuoteUpdateTimeRequest{} }
func (m *QueryQuoteUpdateTimeRequest) String() string { return proto.CompactTextString(m) }
func (*QueryQuoteUpdateTimeRequest) ProtoMessage()    {}
func (*QueryQuoteUpdateTimeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b1b80b580af04cb0, []int{41}
}
func (m *QueryQuoteUpdateTimeRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryQuoteUpdateTimeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryQuoteUpdateTimeRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryQuoteUpdateTimeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryQuoteUpdateTimeRequest.Merge(m, src)
}
func (m *QueryQuoteUpdateTimeRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryQuoteUpdateTimeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryQuoteUpdateTimeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryQuoteUpdateTimeRequest proto.InternalMessageInfo

func (m *QueryQuoteUpdateTimeRequest) GetBucketName() string {
	if m != nil {
		return m.BucketName
	}
	return ""
}

type QueryQuoteUpdateTimeResponse struct {
	UpdateAt int64 `protobuf:"varint,6,opt,name=update_at,json=updateAt,proto3" json:"update_at,omitempty"`
}

func (m *QueryQuoteUpdateTimeResponse) Reset()         { *m = QueryQuoteUpdateTimeResponse{} }
func (m *QueryQuoteUpdateTimeResponse) String() string { return proto.CompactTextString(m) }
func (*QueryQuoteUpdateTimeResponse) ProtoMessage()    {}
func (*QueryQuoteUpdateTimeResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b1b80b580af04cb0, []int{42}
}
func (m *QueryQuoteUpdateTimeResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryQuoteUpdateTimeResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryQuoteUpdateTimeResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryQuoteUpdateTimeResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryQuoteUpdateTimeResponse.Merge(m, src)
}
func (m *QueryQuoteUpdateTimeResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryQuoteUpdateTimeResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryQuoteUpdateTimeResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryQuoteUpdateTimeResponse proto.InternalMessageInfo

func (m *QueryQuoteUpdateTimeResponse) GetUpdateAt() int64 {
	if m != nil {
		return m.UpdateAt
	}
	return 0
}

type QueryGroupMembersExistRequest struct {
	GroupId string   `protobuf:"bytes,1,opt,name=group_id,json=groupId,proto3" json:"group_id,omitempty"`
	Members []string `protobuf:"bytes,2,rep,name=members,proto3" json:"members,omitempty"`
}

func (m *QueryGroupMembersExistRequest) Reset()         { *m = QueryGroupMembersExistRequest{} }
func (m *QueryGroupMembersExistRequest) String() string { return proto.CompactTextString(m) }
func (*QueryGroupMembersExistRequest) ProtoMessage()    {}
func (*QueryGroupMembersExistRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b1b80b580af04cb0, []int{43}
}
func (m *QueryGroupMembersExistRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryGroupMembersExistRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryGroupMembersExistRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryGroupMembersExistRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryGroupMembersExistRequest.Merge(m, src)
}
func (m *QueryGroupMembersExistRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryGroupMembersExistRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryGroupMembersExistRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryGroupMembersExistRequest proto.InternalMessageInfo

func (m *QueryGroupMembersExistRequest) GetGroupId() string {
	if m != nil {
		return m.GroupId
	}
	return ""
}

func (m *QueryGroupMembersExistRequest) GetMembers() []string {
	if m != nil {
		return m.Members
	}
	return nil
}

type QueryGroupMembersExistResponse struct {
	Exists map[string]bool `protobuf:"bytes,1,rep,name=exists,proto3" json:"exists,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
}

func (m *QueryGroupMembersExistResponse) Reset()         { *m = QueryGroupMembersExistResponse{} }
func (m *QueryGroupMembersExistResponse) String() string { return proto.CompactTextString(m) }
func (*QueryGroupMembersExistResponse) ProtoMessage()    {}
func (*QueryGroupMembersExistResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b1b80b580af04cb0, []int{44}
}
func (m *QueryGroupMembersExistResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryGroupMembersExistResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryGroupMembersExistResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryGroupMembersExistResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryGroupMembersExistResponse.Merge(m, src)
}
func (m *QueryGroupMembersExistResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryGroupMembersExistResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryGroupMembersExistResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryGroupMembersExistResponse proto.InternalMessageInfo

func (m *QueryGroupMembersExistResponse) GetExists() map[string]bool {
	if m != nil {
		return m.Exists
	}
	return nil
}

type QueryGroupsExistRequest struct {
	GroupOwner string   `protobuf:"bytes,1,opt,name=group_owner,json=groupOwner,proto3" json:"group_owner,omitempty"`
	GroupNames []string `protobuf:"bytes,2,rep,name=group_names,json=groupNames,proto3" json:"group_names,omitempty"`
}

func (m *QueryGroupsExistRequest) Reset()         { *m = QueryGroupsExistRequest{} }
func (m *QueryGroupsExistRequest) String() string { return proto.CompactTextString(m) }
func (*QueryGroupsExistRequest) ProtoMessage()    {}
func (*QueryGroupsExistRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b1b80b580af04cb0, []int{45}
}
func (m *QueryGroupsExistRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryGroupsExistRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryGroupsExistRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryGroupsExistRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryGroupsExistRequest.Merge(m, src)
}
func (m *QueryGroupsExistRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryGroupsExistRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryGroupsExistRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryGroupsExistRequest proto.InternalMessageInfo

func (m *QueryGroupsExistRequest) GetGroupOwner() string {
	if m != nil {
		return m.GroupOwner
	}
	return ""
}

func (m *QueryGroupsExistRequest) GetGroupNames() []string {
	if m != nil {
		return m.GroupNames
	}
	return nil
}

type QueryGroupsExistByIdRequest struct {
	GroupIds []string `protobuf:"bytes,1,rep,name=group_ids,json=groupIds,proto3" json:"group_ids,omitempty"`
}

func (m *QueryGroupsExistByIdRequest) Reset()         { *m = QueryGroupsExistByIdRequest{} }
func (m *QueryGroupsExistByIdRequest) String() string { return proto.CompactTextString(m) }
func (*QueryGroupsExistByIdRequest) ProtoMessage()    {}
func (*QueryGroupsExistByIdRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b1b80b580af04cb0, []int{46}
}
func (m *QueryGroupsExistByIdRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryGroupsExistByIdRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryGroupsExistByIdRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryGroupsExistByIdRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryGroupsExistByIdRequest.Merge(m, src)
}
func (m *QueryGroupsExistByIdRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryGroupsExistByIdRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryGroupsExistByIdRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryGroupsExistByIdRequest proto.InternalMessageInfo

func (m *QueryGroupsExistByIdRequest) GetGroupIds() []string {
	if m != nil {
		return m.GroupIds
	}
	return nil
}

type QueryGroupsExistResponse struct {
	Exists map[string]bool `protobuf:"bytes,1,rep,name=exists,proto3" json:"exists,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
}

func (m *QueryGroupsExistResponse) Reset()         { *m = QueryGroupsExistResponse{} }
func (m *QueryGroupsExistResponse) String() string { return proto.CompactTextString(m) }
func (*QueryGroupsExistResponse) ProtoMessage()    {}
func (*QueryGroupsExistResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b1b80b580af04cb0, []int{47}
}
func (m *QueryGroupsExistResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryGroupsExistResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryGroupsExistResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryGroupsExistResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryGroupsExistResponse.Merge(m, src)
}
func (m *QueryGroupsExistResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryGroupsExistResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryGroupsExistResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryGroupsExistResponse proto.InternalMessageInfo

func (m *QueryGroupsExistResponse) GetExists() map[string]bool {
	if m != nil {
		return m.Exists
	}
	return nil
}

type QueryPaymentAccountBucketFlowRateLimitRequest struct {
	PaymentAccount string `protobuf:"bytes,1,opt,name=payment_account,json=paymentAccount,proto3" json:"payment_account,omitempty"`
	BucketOwner    string `protobuf:"bytes,2,opt,name=bucket_owner,json=bucketOwner,proto3" json:"bucket_owner,omitempty"`
	BucketName     string `protobuf:"bytes,3,opt,name=bucket_name,json=bucketName,proto3" json:"bucket_name,omitempty"`
}

func (m *QueryPaymentAccountBucketFlowRateLimitRequest) Reset() {
	*m = QueryPaymentAccountBucketFlowRateLimitRequest{}
}
func (m *QueryPaymentAccountBucketFlowRateLimitRequest) String() string {
	return proto.CompactTextString(m)
}
func (*QueryPaymentAccountBucketFlowRateLimitRequest) ProtoMessage() {}
func (*QueryPaymentAccountBucketFlowRateLimitRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b1b80b580af04cb0, []int{48}
}
func (m *QueryPaymentAccountBucketFlowRateLimitRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryPaymentAccountBucketFlowRateLimitRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryPaymentAccountBucketFlowRateLimitRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryPaymentAccountBucketFlowRateLimitRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryPaymentAccountBucketFlowRateLimitRequest.Merge(m, src)
}
func (m *QueryPaymentAccountBucketFlowRateLimitRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryPaymentAccountBucketFlowRateLimitRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryPaymentAccountBucketFlowRateLimitRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryPaymentAccountBucketFlowRateLimitRequest proto.InternalMessageInfo

func (m *QueryPaymentAccountBucketFlowRateLimitRequest) GetPaymentAccount() string {
	if m != nil {
		return m.PaymentAccount
	}
	return ""
}

func (m *QueryPaymentAccountBucketFlowRateLimitRequest) GetBucketOwner() string {
	if m != nil {
		return m.BucketOwner
	}
	return ""
}

func (m *QueryPaymentAccountBucketFlowRateLimitRequest) GetBucketName() string {
	if m != nil {
		return m.BucketName
	}
	return ""
}

type QueryPaymentAccountBucketFlowRateLimitResponse struct {
	IsSet         bool                                   `protobuf:"varint,1,opt,name=is_set,json=isSet,proto3" json:"is_set,omitempty"`
	FlowRateLimit github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,2,opt,name=flow_rate_limit,json=flowRateLimit,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"flow_rate_limit"`
}

func (m *QueryPaymentAccountBucketFlowRateLimitResponse) Reset() {
	*m = QueryPaymentAccountBucketFlowRateLimitResponse{}
}
func (m *QueryPaymentAccountBucketFlowRateLimitResponse) String() string {
	return proto.CompactTextString(m)
}
func (*QueryPaymentAccountBucketFlowRateLimitResponse) ProtoMessage() {}
func (*QueryPaymentAccountBucketFlowRateLimitResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b1b80b580af04cb0, []int{49}
}
func (m *QueryPaymentAccountBucketFlowRateLimitResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryPaymentAccountBucketFlowRateLimitResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryPaymentAccountBucketFlowRateLimitResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryPaymentAccountBucketFlowRateLimitResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryPaymentAccountBucketFlowRateLimitResponse.Merge(m, src)
}
func (m *QueryPaymentAccountBucketFlowRateLimitResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryPaymentAccountBucketFlowRateLimitResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryPaymentAccountBucketFlowRateLimitResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryPaymentAccountBucketFlowRateLimitResponse proto.InternalMessageInfo

func (m *QueryPaymentAccountBucketFlowRateLimitResponse) GetIsSet() bool {
	if m != nil {
		return m.IsSet
	}
	return false
}

func init() {
	proto.RegisterType((*QueryParamsRequest)(nil), "greenfield.storage.QueryParamsRequest")
	proto.RegisterType((*QueryHeadBucketRequest)(nil), "greenfield.storage.QueryHeadBucketRequest")
	proto.RegisterType((*QueryHeadBucketByIdRequest)(nil), "greenfield.storage.QueryHeadBucketByIdRequest")
	proto.RegisterType((*QueryHeadBucketResponse)(nil), "greenfield.storage.QueryHeadBucketResponse")
	proto.RegisterType((*QueryHeadObjectRequest)(nil), "greenfield.storage.QueryHeadObjectRequest")
	proto.RegisterType((*QueryHeadObjectByIdRequest)(nil), "greenfield.storage.QueryHeadObjectByIdRequest")
	proto.RegisterType((*QueryHeadShadowObjectRequest)(nil), "greenfield.storage.QueryHeadShadowObjectRequest")
	proto.RegisterType((*QueryHeadObjectResponse)(nil), "greenfield.storage.QueryHeadObjectResponse")
	proto.RegisterType((*QueryHeadShadowObjectResponse)(nil), "greenfield.storage.QueryHeadShadowObjectResponse")
	proto.RegisterType((*QueryListBucketsRequest)(nil), "greenfield.storage.QueryListBucketsRequest")
	proto.RegisterType((*QueryListBucketsResponse)(nil), "greenfield.storage.QueryListBucketsResponse")
	proto.RegisterType((*QueryListObjectsRequest)(nil), "greenfield.storage.QueryListObjectsRequest")
	proto.RegisterType((*QueryListObjectsByBucketIdRequest)(nil), "greenfield.storage.QueryListObjectsByBucketIdRequest")
	proto.RegisterType((*QueryListObjectsResponse)(nil), "greenfield.storage.QueryListObjectsResponse")
	proto.RegisterType((*QueryNFTRequest)(nil), "greenfield.storage.QueryNFTRequest")
	proto.RegisterType((*QueryBucketNFTResponse)(nil), "greenfield.storage.QueryBucketNFTResponse")
	proto.RegisterType((*QueryObjectNFTResponse)(nil), "greenfield.storage.QueryObjectNFTResponse")
	proto.RegisterType((*QueryGroupNFTResponse)(nil), "greenfield.storage.QueryGroupNFTResponse")
	proto.RegisterType((*QueryPolicyForAccountRequest)(nil), "greenfield.storage.QueryPolicyForAccountRequest")
	proto.RegisterType((*QueryPolicyForAccountResponse)(nil), "greenfield.storage.QueryPolicyForAccountResponse")
	proto.RegisterType((*QueryVerifyPermissionRequest)(nil), "greenfield.storage.QueryVerifyPermissionRequest")
	proto.RegisterType((*QueryVerifyPermissionResponse)(nil), "greenfield.storage.QueryVerifyPermissionResponse")
	proto.RegisterType((*QueryHeadGroupRequest)(nil), "greenfield.storage.QueryHeadGroupRequest")
	proto.RegisterType((*QueryHeadGroupResponse)(nil), "greenfield.storage.QueryHeadGroupResponse")
	proto.RegisterType((*QueryListGroupsRequest)(nil), "greenfield.storage.QueryListGroupsRequest")
	proto.RegisterType((*QueryListGroupsResponse)(nil), "greenfield.storage.QueryListGroupsResponse")
	proto.RegisterType((*QueryHeadGroupMemberRequest)(nil), "greenfield.storage.QueryHeadGroupMemberRequest")
	proto.RegisterType((*QueryHeadGroupMemberResponse)(nil), "greenfield.storage.QueryHeadGroupMemberResponse")
	proto.RegisterType((*QueryPolicyForGroupRequest)(nil), "greenfield.storage.QueryPolicyForGroupRequest")
	proto.RegisterType((*QueryPolicyForGroupResponse)(nil), "greenfield.storage.QueryPolicyForGroupResponse")
	proto.RegisterType((*QueryPolicyByIdRequest)(nil), "greenfield.storage.QueryPolicyByIdRequest")
	proto.RegisterType((*QueryPolicyByIdResponse)(nil), "greenfield.storage.QueryPolicyByIdResponse")
	proto.RegisterType((*QueryLockFeeRequest)(nil), "greenfield.storage.QueryLockFeeRequest")
	proto.RegisterType((*QueryLockFeeResponse)(nil), "greenfield.storage.QueryLockFeeResponse")
	proto.RegisterType((*QueryHeadBucketExtraRequest)(nil), "greenfield.storage.QueryHeadBucketExtraRequest")
	proto.RegisterType((*QueryHeadBucketExtraResponse)(nil), "greenfield.storage.QueryHeadBucketExtraResponse")
	proto.RegisterType((*QueryIsPriceChangedRequest)(nil), "greenfield.storage.QueryIsPriceChangedRequest")
	proto.RegisterType((*QueryIsPriceChangedResponse)(nil), "greenfield.storage.QueryIsPriceChangedResponse")
	proto.RegisterType((*QueryQuoteUpdateTimeRequest)(nil), "greenfield.storage.QueryQuoteUpdateTimeRequest")
	proto.RegisterType((*QueryQuoteUpdateTimeResponse)(nil), "greenfield.storage.QueryQuoteUpdateTimeResponse")
	proto.RegisterType((*QueryGroupMembersExistRequest)(nil), "greenfield.storage.QueryGroupMembersExistRequest")
	proto.RegisterType((*QueryGroupMembersExistResponse)(nil), "greenfield.storage.QueryGroupMembersExistResponse")
	proto.RegisterMapType((map[string]bool)(nil), "greenfield.storage.QueryGroupMembersExistResponse.ExistsEntry")
	proto.RegisterType((*QueryGroupsExistRequest)(nil), "greenfield.storage.QueryGroupsExistRequest")
	proto.RegisterType((*QueryGroupsExistByIdRequest)(nil), "greenfield.storage.QueryGroupsExistByIdRequest")
	proto.RegisterType((*QueryGroupsExistResponse)(nil), "greenfield.storage.QueryGroupsExistResponse")
	proto.RegisterMapType((map[string]bool)(nil), "greenfield.storage.QueryGroupsExistResponse.ExistsEntry")
	proto.RegisterType((*QueryPaymentAccountBucketFlowRateLimitRequest)(nil), "greenfield.storage.QueryPaymentAccountBucketFlowRateLimitRequest")
	proto.RegisterType((*QueryPaymentAccountBucketFlowRateLimitResponse)(nil), "greenfield.storage.QueryPaymentAccountBucketFlowRateLimitResponse")
}

func init() { proto.RegisterFile("greenfield/storage/query.proto", fileDescriptor_b1b80b580af04cb0) }

var fileDescriptor_b1b80b580af04cb0 = []byte{
	// 2855 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x5a, 0xdb, 0x6f, 0x1c, 0x49,
	0xd5, 0x4f, 0xdb, 0x89, 0x63, 0x97, 0xbd, 0x89, 0xbf, 0x5a, 0x67, 0xe3, 0x4c, 0x12, 0x27, 0xe9,
	0xec, 0xe7, 0x64, 0x93, 0x78, 0x26, 0x76, 0x12, 0x14, 0xe7, 0x86, 0xec, 0xb5, 0x1d, 0x8c, 0x72,
	0x71, 0xc6, 0xc6, 0x88, 0x08, 0xd4, 0x94, 0xa7, 0x6b, 0x26, 0xbd, 0x9e, 0xe9, 0x9e, 0x74, 0xf7,
	0xd8, 0x99, 0xb5, 0x46, 0x88, 0x7d, 0x81, 0x47, 0xc4, 0x0a, 0x09, 0x89, 0x8b, 0x10, 0x88, 0xab,
	0x84, 0x10, 0xec, 0x0a, 0x89, 0x27, 0x1e, 0x00, 0x69, 0x25, 0x84, 0x14, 0x96, 0x17, 0xb4, 0x0f,
	0x2b, 0x48, 0xf8, 0x43, 0x50, 0x57, 0x9d, 0xea, 0xae, 0xbe, 0x4c, 0x77, 0x3b, 0x1e, 0x5e, 0x92,
	0x99, 0x9a, 0x3a, 0xe7, 0xfc, 0xce, 0xa5, 0x4e, 0x9d, 0x3a, 0xc7, 0x68, 0xa2, 0x66, 0x53, 0x6a,
	0x56, 0x0d, 0x5a, 0xd7, 0x4b, 0x8e, 0x6b, 0xd9, 0xa4, 0x46, 0x4b, 0x4f, 0x5b, 0xd4, 0x6e, 0x17,
	0x9b, 0xb6, 0xe5, 0x5a, 0x18, 0x07, 0xbf, 0x17, 0xe1, 0xf7, 0xc2, 0x85, 0x8a, 0xe5, 0x34, 0x2c,
	0xa7, 0xb4, 0x41, 0x1c, 0xd8, 0x5c, 0xda, 0x9a, 0xde, 0xa0, 0x2e, 0x99, 0x2e, 0x35, 0x49, 0xcd,
	0x30, 0x89, 0x6b, 0x58, 0x26, 0xa7, 0x2f, 0x1c, 0xe3, 0x7b, 0x35, 0xf6, 0xad, 0xc4, 0xbf, 0xc0,
	0x4f, 0x63, 0x35, 0xab, 0x66, 0xf1, 0x75, 0xef, 0x13, 0xac, 0x9e, 0xa8, 0x59, 0x56, 0xad, 0x4e,
	0x4b, 0xa4, 0x69, 0x94, 0x88, 0x69, 0x5a, 0x2e, 0xe3, 0x26, 0x68, 0x54, 0x09, 0x6e, 0x93, 0xda,
	0x0d, 0xc3, 0x71, 0x0c, 0xcb, 0x2c, 0x55, 0xac, 0x46, 0xc3, 0x17, 0x79, 0x26, 0x79, 0x8f, 0xdb,
	0x6e, 0x52, 0xc1, 0xe6, 0x54, 0x82, 0xd6, 0x4d, 0x62, 0x93, 0x86, 0xd8, 0x90, 0x64, 0x16, 0x99,
	0xc1, 0x59, 0xe9, 0xf7, 0x2d, 0xc3, 0x76, 0x5b, 0xa4, 0x5e, 0xb3, 0xad, 0x56, 0x53, 0xde, 0xa4,
	0x8e, 0x21, 0xfc, 0xc8, 0xb3, 0xce, 0x0a, 0xe3, 0x5c, 0xa6, 0x4f, 0x5b, 0xd4, 0x71, 0xd5, 0x87,
	0xe8, 0xf5, 0xd0, 0xaa, 0xd3, 0xb4, 0x4c, 0x87, 0xe2, 0xeb, 0x68, 0x80, 0x23, 0x18, 0x57, 0x4e,
	0x2b, 0xe7, 0x87, 0x67, 0x0a, 0xc5, 0xb8, 0xe5, 0x8b, 0x9c, 0x66, 0x7e, 0xff, 0x47, 0x9f, 0x9e,
	0xda, 0x57, 0x86, 0xfd, 0xea, 0x6d, 0x74, 0x52, 0x62, 0x38, 0xdf, 0x5e, 0x33, 0x1a, 0xd4, 0x71,
	0x49, 0xa3, 0x09, 0x12, 0xf1, 0x09, 0x34, 0xe4, 0x8a, 0x35, 0xc6, 0xbd, 0xbf, 0x1c, 0x2c, 0xa8,
	0x8f, 0xd1, 0x44, 0x37, 0xf2, 0x3d, 0x43, 0x9b, 0x45, 0x6f, 0x30, 0xde, 0x9f, 0xa3, 0x44, 0x9f,
	0x6f, 0x55, 0x36, 0xa9, 0x2b, 0x30, 0x9d, 0x42, 0xc3, 0x1b, 0x6c, 0x41, 0x33, 0x49, 0x83, 0x32,
	0xc6, 0x43, 0x65, 0xc4, 0x97, 0x1e, 0x90, 0x06, 0x55, 0x67, 0x51, 0x21, 0x42, 0x3a, 0xdf, 0x5e,
	0xd6, 0x05, 0xf9, 0x71, 0x34, 0x04, 0xe4, 0x86, 0x0e, 0xc4, 0x83, 0x7c, 0x61, 0x59, 0x57, 0x7f,
	0xa8, 0xa0, 0xa3, 0x31, 0xb1, 0xa0, 0xcb, 0x67, 0x7d, 0xb9, 0x86, 0x59, 0xb5, 0x40, 0xa1, 0x89,
	0x24, 0x85, 0x38, 0xe1, 0xb2, 0x59, 0xb5, 0x04, 0x2e, 0xef, 0x33, 0x9e, 0x47, 0x88, 0x3e, 0x73,
	0x6d, 0xc2, 0xe9, 0xfb, 0x18, 0xfd, 0xd9, 0xee, 0xf4, 0x8b, 0xde, 0x5e, 0xc6, 0x64, 0x88, 0x8a,
	0x8f, 0xea, 0x63, 0xc9, 0x2c, 0x0f, 0x37, 0xde, 0xa1, 0x95, 0xdc, 0x66, 0xf1, 0x36, 0x58, 0x8c,
	0x82, 0x6f, 0xe8, 0xe3, 0x1b, 0xf8, 0x52, 0xcc, 0x6e, 0x9c, 0x77, 0xc4, 0x6e, 0x40, 0x1e, 0xd8,
	0x8d, 0x2f, 0x2c, 0xeb, 0xea, 0x57, 0xd1, 0x09, 0x9f, 0x74, 0xf5, 0x09, 0xd1, 0xad, 0xed, 0x5e,
	0x83, 0xfb, 0x83, 0xec, 0x19, 0xc1, 0x3c, 0xf0, 0x8c, 0x80, 0x96, 0xe1, 0x19, 0x4e, 0xc8, 0x3d,
	0x63, 0xf9, 0x9f, 0xf1, 0x57, 0xd0, 0x58, 0xad, 0x6e, 0x6d, 0x90, 0xba, 0x06, 0x27, 0x52, 0x63,
	0x47, 0x12, 0x7c, 0x74, 0x51, 0xe6, 0x24, 0x1f, 0xd9, 0xe2, 0x5d, 0x46, 0xb4, 0xce, 0x97, 0xee,
	0x7a, 0x4b, 0x65, 0x5c, 0x8b, 0xad, 0xa9, 0x55, 0x38, 0x66, 0x71, 0xeb, 0x80, 0x02, 0x8b, 0x49,
	0x0a, 0xbc, 0x99, 0xa4, 0x80, 0x4c, 0x1e, 0x55, 0x43, 0x25, 0x60, 0xa2, 0x7b, 0x86, 0xe3, 0xf2,
	0x18, 0x12, 0xa9, 0x03, 0x2f, 0x21, 0x14, 0x24, 0x58, 0x10, 0x30, 0x59, 0x84, 0xa4, 0xea, 0x65,
	0xe3, 0x22, 0x4f, 0xdd, 0x90, 0x8d, 0x8b, 0x2b, 0xa4, 0x46, 0x81, 0xb6, 0x2c, 0x51, 0xaa, 0x3f,
	0x53, 0xd0, 0x78, 0x5c, 0x06, 0xa8, 0x31, 0x87, 0x46, 0xa4, 0x13, 0xe2, 0x9d, 0xf9, 0xfe, 0x1c,
	0x47, 0x64, 0x38, 0x38, 0x22, 0x0e, 0xbe, 0x1b, 0xc2, 0xc9, 0xed, 0x7f, 0x2e, 0x13, 0x27, 0x97,
	0x1f, 0x02, 0xfa, 0x9e, 0x22, 0x19, 0x83, 0xdb, 0xab, 0xd7, 0xc6, 0x88, 0x46, 0x75, 0x5f, 0x2c,
	0x13, 0x7d, 0x53, 0x41, 0x67, 0xa2, 0x20, 0xe6, 0xdb, 0xa0, 0xbb, 0xde, 0x6b, 0x38, 0xa1, 0xcc,
	0xd6, 0x17, 0xc9, 0x6c, 0x21, 0xc7, 0xf9, 0xf6, 0x08, 0x1c, 0x27, 0xc5, 0x5f, 0xaa, 0xe3, 0xa4,
	0xd0, 0x1b, 0x0e, 0x42, 0xaf, 0x87, 0x8e, 0xbb, 0x84, 0x0e, 0x33, 0x9c, 0x0f, 0x96, 0xd6, 0x84,
	0x81, 0x8e, 0xa1, 0x41, 0xd7, 0xda, 0xa4, 0x66, 0x90, 0x79, 0x0e, 0xb2, 0xef, 0xcb, 0xba, 0xfa,
	0x25, 0xc8, 0x87, 0xdc, 0xa6, 0x8c, 0xc6, 0x4f, 0x0a, 0x43, 0x0d, 0xea, 0x12, 0x4d, 0x27, 0x2e,
	0x01, 0xa3, 0xaa, 0xdd, 0x23, 0xf1, 0x3e, 0x75, 0xc9, 0x02, 0x71, 0x49, 0x79, 0xb0, 0x01, 0x9f,
	0x7c, 0xd6, 0x5c, 0xe3, 0x57, 0x61, 0xcd, 0x29, 0x13, 0x58, 0x7f, 0x11, 0x1d, 0x61, 0xac, 0x59,
	0x7a, 0x90, 0x39, 0xdf, 0x89, 0x73, 0x3e, 0x93, 0xc4, 0x99, 0x11, 0x26, 0x30, 0xfe, 0xba, 0x02,
	0x89, 0x78, 0xc5, 0xaa, 0x1b, 0x95, 0xf6, 0x92, 0x65, 0xcf, 0x55, 0x2a, 0x56, 0xcb, 0xf4, 0x13,
	0x71, 0x01, 0x0d, 0xda, 0xd4, 0xb1, 0x5a, 0x76, 0x45, 0x64, 0x61, 0xff, 0x3b, 0x5e, 0x44, 0xff,
	0xd7, 0xb4, 0x0d, 0xb3, 0x62, 0x34, 0x49, 0x5d, 0x23, 0xba, 0x6e, 0x53, 0xc7, 0xe1, 0x71, 0x34,
	0x3f, 0xfe, 0xf1, 0x87, 0x53, 0x63, 0xe0, 0xcc, 0x39, 0xfe, 0xcb, 0xaa, 0x6b, 0x1b, 0x66, 0xad,
	0x3c, 0xea, 0x93, 0xc0, 0xba, 0xba, 0x2e, 0x8a, 0x8a, 0x18, 0x04, 0x50, 0xf2, 0x1a, 0x1a, 0x68,
	0xb2, 0xdf, 0x40, 0xc3, 0x93, 0xb2, 0x86, 0x41, 0xd9, 0x55, 0xe4, 0x0c, 0xca, 0xb0, 0x59, 0xfd,
	0x44, 0xe8, 0xb6, 0x4e, 0x6d, 0xa3, 0xda, 0x5e, 0xf1, 0x37, 0x0a, 0xdd, 0xae, 0xa2, 0x41, 0xab,
	0x49, 0x6d, 0xe2, 0x5a, 0x36, 0xd7, 0x2d, 0x05, 0xb6, 0xbf, 0x33, 0xf3, 0x10, 0x47, 0xaf, 0xa6,
	0xfe, 0xe8, 0xd5, 0x84, 0xe7, 0xd1, 0x30, 0xa9, 0x78, 0xb1, 0xab, 0x79, 0x25, 0xdc, 0xf8, 0xfe,
	0xd3, 0xca, 0xf9, 0x43, 0x61, 0xb7, 0x49, 0x4a, 0xcd, 0xb1, 0x9d, 0x6b, 0xed, 0x26, 0x2d, 0x23,
	0xe2, 0x7f, 0xf6, 0x8d, 0x16, 0xd7, 0x2d, 0x30, 0x1a, 0xad, 0x56, 0x69, 0xc5, 0x65, 0xaa, 0x1d,
	0xea, 0x6a, 0xb4, 0x45, 0xb6, 0xa9, 0x0c, 0x9b, 0xd5, 0xa7, 0x10, 0x69, 0xde, 0xd5, 0xc3, 0x2f,
	0x28, 0x30, 0xd6, 0x2c, 0x1a, 0x66, 0x77, 0x98, 0x66, 0x6d, 0x9b, 0x34, 0xdb, 0x5e, 0x88, 0x6d,
	0x7e, 0xe8, 0xed, 0xc5, 0x27, 0x11, 0xff, 0x26, 0x1b, 0x6c, 0x88, 0xad, 0xb0, 0xa4, 0xb7, 0x2e,
	0x95, 0x28, 0x20, 0x12, 0x74, 0xb8, 0x25, 0x08, 0xa5, 0x5b, 0xee, 0x64, 0xd7, 0xf0, 0xe6, 0xa5,
	0x4f, 0x4d, 0x7c, 0x54, 0xbf, 0xa7, 0x00, 0x63, 0x2f, 0x83, 0xb1, 0x1d, 0x3d, 0x4f, 0xe8, 0x11,
	0xa3, 0xf4, 0xe5, 0x37, 0x8a, 0xfa, 0x63, 0xf9, 0xbe, 0x11, 0xe8, 0x40, 0xef, 0xbb, 0x09, 0xf0,
	0x5e, 0x25, 0x37, 0xe2, 0x3b, 0x02, 0x1f, 0x4f, 0xd3, 0x7d, 0x2c, 0x4d, 0x67, 0x58, 0x10, 0xf9,
	0x16, 0x74, 0xd4, 0x5f, 0x2a, 0xe8, 0x78, 0xd8, 0x37, 0xf7, 0x69, 0x63, 0x83, 0xda, 0xc2, 0x8e,
	0x97, 0xd1, 0x40, 0x83, 0x2d, 0x64, 0xc6, 0x03, 0xec, 0xdb, 0x83, 0xc5, 0x22, 0x61, 0xd4, 0x1f,
	0x0d, 0x23, 0x2a, 0x95, 0x94, 0x21, 0xa8, 0x7e, 0xcd, 0x34, 0xc2, 0xc9, 0x25, 0xc4, 0x91, 0x3c,
	0x2c, 0x1d, 0x0b, 0x99, 0x03, 0x47, 0xcc, 0xbf, 0xa8, 0x55, 0x28, 0x7a, 0xfd, 0x6c, 0x15, 0x3a,
	0x25, 0x69, 0xe9, 0xf2, 0x12, 0xc2, 0x41, 0xba, 0x04, 0xb7, 0x88, 0x7b, 0x37, 0xc8, 0x8a, 0xdc,
	0x11, 0xba, 0xba, 0x06, 0x96, 0x8f, 0xca, 0xd9, 0x5b, 0x4e, 0xbc, 0x06, 0x47, 0x82, 0x2f, 0x47,
	0xca, 0x75, 0xbe, 0x47, 0x2a, 0xd7, 0xf9, 0xc2, 0xb2, 0xae, 0xae, 0x40, 0xac, 0xca, 0x64, 0x7b,
	0x03, 0xf2, 0x03, 0x05, 0xde, 0xa6, 0xf7, 0xac, 0xca, 0xe6, 0x12, 0xa5, 0xc1, 0xc9, 0xf4, 0x8c,
	0xd4, 0x20, 0x76, 0x5b, 0x73, 0x9a, 0xfe, 0xa5, 0xa2, 0xe4, 0xb8, 0x54, 0x3c, 0x9a, 0xd5, 0x26,
	0xac, 0x7b, 0xea, 0x54, 0x6c, 0x4a, 0x5c, 0xaa, 0x11, 0x97, 0xd9, 0xb8, 0xbf, 0x3c, 0xc8, 0x17,
	0xe6, 0x5c, 0x7c, 0x06, 0x8d, 0x34, 0x49, 0xbb, 0x6e, 0x11, 0x5d, 0x73, 0x8c, 0x77, 0x79, 0x2c,
	0xed, 0x2f, 0x0f, 0xc3, 0xda, 0xaa, 0xf1, 0x2e, 0x55, 0xeb, 0x68, 0x2c, 0x0c, 0x0f, 0xd4, 0x5d,
	0x43, 0x03, 0xa4, 0xe1, 0xdd, 0x4e, 0x80, 0xe9, 0x96, 0xf7, 0x08, 0xfd, 0xe4, 0xd3, 0x53, 0x93,
	0x35, 0xc3, 0x7d, 0xd2, 0xda, 0x28, 0x56, 0xac, 0x06, 0xb4, 0x1e, 0xe0, 0xbf, 0x29, 0x47, 0xdf,
	0x84, 0xa7, 0xfa, 0xb2, 0xe9, 0x7e, 0xfc, 0xe1, 0x14, 0x02, 0x0d, 0x96, 0x4d, 0xb7, 0x0c, 0xbc,
	0xd4, 0x3b, 0xd2, 0x31, 0x93, 0x1e, 0x73, 0xb9, 0x5f, 0xb0, 0x72, 0xec, 0x87, 0xe8, 0xfd, 0xd8,
	0x97, 0x5f, 0x92, 0x22, 0xdf, 0x25, 0xa4, 0x81, 0x65, 0xd3, 0xa5, 0xb6, 0x49, 0xea, 0x52, 0xb9,
	0x2d, 0x3d, 0x26, 0x6f, 0x43, 0xec, 0x2f, 0x3b, 0x2b, 0xb6, 0x51, 0xa1, 0x6f, 0x3f, 0x21, 0x66,
	0x8d, 0xea, 0xb9, 0x51, 0xfe, 0xfb, 0x20, 0xa8, 0x19, 0xa5, 0x07, 0x94, 0xe3, 0xe8, 0x60, 0x85,
	0x2f, 0x31, 0xe2, 0xc1, 0xb2, 0xf8, 0x8a, 0xdf, 0x41, 0xb8, 0xd2, 0xb2, 0x6d, 0x6a, 0xba, 0x9a,
	0x4d, 0x89, 0xae, 0x35, 0x3d, 0x72, 0x48, 0x1e, 0xbb, 0xf1, 0xc0, 0x02, 0xad, 0x48, 0x1e, 0x58,
	0xa0, 0x95, 0xf2, 0x28, 0xf0, 0x2d, 0x53, 0xa2, 0x33, 0x50, 0x78, 0x07, 0x1d, 0x17, 0xb2, 0xfc,
	0x48, 0x74, 0x2d, 0x9b, 0x82, 0xd0, 0xfe, 0x1e, 0x08, 0x1d, 0x07, 0x01, 0x2b, 0x10, 0xb5, 0x1e,
	0x7b, 0x2e, 0xfc, 0x6b, 0xe8, 0xa4, 0x10, 0xee, 0xd0, 0x8a, 0x65, 0xea, 0x51, 0xf1, 0xfb, 0x7b,
	0x20, 0xbe, 0x00, 0x22, 0x56, 0x85, 0x04, 0x09, 0x40, 0x1b, 0x89, 0x5f, 0xb5, 0x2d, 0x52, 0x37,
	0x74, 0xaf, 0xe4, 0xd1, 0x5c, 0xf2, 0x4c, 0xb3, 0x89, 0x4b, 0xc7, 0x0f, 0xf4, 0x40, 0xfa, 0x51,
	0xe0, 0xbf, 0x2e, 0xd8, 0xaf, 0x91, 0x67, 0x65, 0xe2, 0x52, 0xbc, 0x81, 0x0e, 0x99, 0x74, 0x5b,
	0x76, 0xf0, 0x40, 0x0f, 0xc4, 0x8d, 0x98, 0x74, 0x3b, 0x70, 0xae, 0x83, 0x8e, 0x7a, 0x32, 0x92,
	0x1c, 0x7b, 0xb0, 0x07, 0xc2, 0xc6, 0x4c, 0xba, 0x1d, 0x77, 0xea, 0x36, 0x3a, 0xe6, 0x09, 0x4d,
	0x76, 0xe8, 0x60, 0x0f, 0xc4, 0xbe, 0x61, 0xd2, 0xed, 0x24, 0x67, 0x3e, 0x45, 0xde, 0x2f, 0x49,
	0x8e, 0x1c, 0xea, 0x81, 0xd4, 0xd7, 0x4d, 0xba, 0x1d, 0x75, 0xa2, 0x9f, 0xc9, 0x1e, 0xb5, 0x2c,
	0x97, 0x7e, 0xa1, 0xa9, 0x13, 0x97, 0xae, 0x19, 0x0d, 0x9a, 0x3b, 0x47, 0xdc, 0x84, 0x4c, 0x16,
	0xa3, 0x87, 0x1c, 0x71, 0x1c, 0x0d, 0xb5, 0xd8, 0xaa, 0x97, 0xd7, 0x07, 0x78, 0x5e, 0xe7, 0x0b,
	0x73, 0xae, 0x6a, 0x42, 0x51, 0x2c, 0x5d, 0xde, 0xce, 0xe2, 0x33, 0xc3, 0x71, 0xa5, 0x87, 0xa1,
	0x7f, 0xf1, 0xc2, 0xc3, 0x90, 0x57, 0x3b, 0x3a, 0x9e, 0x41, 0x07, 0x79, 0x61, 0xc0, 0xcb, 0xa4,
	0xb4, 0xdb, 0x46, 0x6c, 0x54, 0x3f, 0x50, 0xa0, 0xa1, 0x99, 0x20, 0x10, 0xf0, 0xae, 0xa3, 0x01,
	0xea, 0x2d, 0x88, 0x37, 0xf2, 0x9d, 0xa4, 0xac, 0x9b, 0xce, 0xa3, 0xc8, 0xbe, 0x39, 0x8b, 0xa6,
	0x6b, 0xb7, 0xcb, 0xc0, 0xad, 0x30, 0x8b, 0x86, 0xa5, 0x65, 0x3c, 0x8a, 0xfa, 0x37, 0x69, 0x1b,
	0x74, 0xf2, 0x3e, 0xe2, 0x31, 0x74, 0x60, 0x8b, 0xd4, 0x5b, 0x3c, 0x4b, 0x0e, 0x96, 0xf9, 0x97,
	0x1b, 0x7d, 0xd7, 0x15, 0xb5, 0x05, 0x97, 0x39, 0x2f, 0x3a, 0x43, 0xf6, 0xd9, 0x43, 0x91, 0x7f,
	0x4a, 0x90, 0x7a, 0x8e, 0x05, 0x1b, 0xc2, 0x06, 0xcf, 0xb1, 0x8e, 0x7a, 0x03, 0x22, 0x43, 0x12,
	0x1b, 0xa9, 0x3f, 0x84, 0x6b, 0xb8, 0xad, 0x86, 0xca, 0x83, 0xe0, 0x1b, 0x47, 0xfd, 0xb9, 0x68,
	0x46, 0x84, 0x30, 0x83, 0x89, 0x57, 0x22, 0x26, 0xbe, 0x9e, 0x6e, 0xe2, 0xff, 0xad, 0x71, 0x9f,
	0x2b, 0x68, 0x0a, 0x7a, 0xdc, 0xed, 0x06, 0x35, 0x5d, 0x78, 0xcb, 0xf2, 0xfb, 0x74, 0xa9, 0x6e,
	0x6d, 0x7b, 0xa7, 0xe4, 0x9e, 0xd1, 0x30, 0x7c, 0x9b, 0xcf, 0xa1, 0xc3, 0x4d, 0xbe, 0x57, 0x23,
	0x7c, 0x73, 0xa6, 0xdd, 0x0f, 0x35, 0x43, 0xcc, 0xf1, 0x4d, 0xbf, 0x8f, 0x96, 0xaf, 0xaa, 0x86,
	0x33, 0xe8, 0x3b, 0x4e, 0x3e, 0x92, 0xfd, 0xb1, 0x23, 0xf9, 0x6b, 0x05, 0x15, 0xf3, 0xaa, 0x04,
	0x2e, 0x39, 0x82, 0x06, 0x0c, 0x47, 0x73, 0xa8, 0x0b, 0x17, 0xf9, 0x01, 0xc3, 0x59, 0xa5, 0x2e,
	0xd6, 0xd1, 0xe1, 0x6a, 0xdd, 0xda, 0x66, 0x29, 0x48, 0xab, 0x7b, 0x14, 0xaf, 0x70, 0x87, 0xc7,
	0xab, 0xa8, 0xd7, 0xaa, 0x32, 0x88, 0x99, 0xbf, 0x4f, 0xa2, 0x03, 0x0c, 0x2f, 0xee, 0xa0, 0x01,
	0x3e, 0x2b, 0xc0, 0x93, 0x5d, 0x63, 0x22, 0x34, 0x31, 0x29, 0x9c, 0xcb, 0xdc, 0xc7, 0x35, 0x54,
	0xd5, 0xf7, 0xfe, 0xf1, 0x9f, 0xf7, 0xfb, 0x4e, 0xe0, 0x42, 0xa9, 0xeb, 0x7c, 0x07, 0xff, 0x46,
	0x3c, 0x40, 0x63, 0xf3, 0x0e, 0x3c, 0x9d, 0x21, 0x27, 0x3e, 0x5a, 0x29, 0xcc, 0xec, 0x86, 0x04,
	0x50, 0x16, 0x19, 0xca, 0xf3, 0x78, 0xb2, 0x3b, 0xca, 0xd2, 0x8e, 0x3f, 0x9f, 0xe9, 0xe0, 0xef,
	0x2b, 0x08, 0x05, 0x35, 0x24, 0xbe, 0xd0, 0x55, 0x64, 0x6c, 0xca, 0x52, 0xb8, 0x98, 0x6b, 0x2f,
	0xe0, 0xba, 0xc6, 0x70, 0x95, 0xf0, 0x54, 0x12, 0xae, 0x27, 0x5e, 0x01, 0xc0, 0xe3, 0xaf, 0xb4,
	0x23, 0x85, 0x66, 0x07, 0xff, 0x42, 0x41, 0x87, 0xc2, 0x43, 0x1a, 0x5c, 0xcc, 0x21, 0x56, 0x4a,
	0x33, 0xbb, 0x83, 0x39, 0xcb, 0x60, 0x5e, 0xc1, 0xd3, 0x19, 0x30, 0xb5, 0x0d, 0xef, 0xd5, 0xe4,
	0x83, 0x35, 0xf4, 0x0e, 0xfe, 0xae, 0x82, 0x5e, 0x0b, 0x38, 0x3e, 0x58, 0x5a, 0xc3, 0x67, 0xbb,
	0x4a, 0x0e, 0x3a, 0x97, 0x85, 0xee, 0x16, 0x8f, 0x35, 0x2c, 0xd5, 0xcf, 0x30, 0x74, 0x97, 0x71,
	0x31, 0x0b, 0x9d, 0x59, 0x75, 0x4b, 0x3b, 0xa2, 0x21, 0xda, 0xc1, 0xbf, 0x02, 0x27, 0xf3, 0x6e,
	0x63, 0x86, 0x93, 0x43, 0x63, 0x99, 0x0c, 0xeb, 0x85, 0x87, 0x14, 0xea, 0xdb, 0x0c, 0xdf, 0x6d,
	0x7c, 0xb3, 0x2b, 0x3e, 0xde, 0x13, 0x0b, 0x3b, 0xb9, 0xb4, 0x23, 0x35, 0xcf, 0x02, 0x97, 0x07,
	0xf3, 0xa5, 0x0c, 0x97, 0xc7, 0x06, 0x51, 0xbb, 0x03, 0x9d, 0xed, 0x72, 0x80, 0x07, 0x2e, 0xf7,
	0x47, 0x5c, 0x1d, 0xfc, 0x27, 0x05, 0x8d, 0x46, 0x27, 0x36, 0xf8, 0x72, 0xaa, 0xf0, 0x84, 0xd1,
	0x57, 0x61, 0x7a, 0x17, 0x14, 0x00, 0xfa, 0xf3, 0x0c, 0xf4, 0x02, 0x9e, 0xef, 0x0a, 0xda, 0x61,
	0x64, 0x79, 0x0c, 0x2e, 0x02, 0xd7, 0xef, 0x62, 0xef, 0x35, 0x70, 0x63, 0xed, 0xf0, 0x1c, 0x81,
	0x2b, 0x10, 0x85, 0x03, 0xf7, 0xdb, 0x0a, 0x1a, 0x96, 0xc6, 0x48, 0xb8, 0xbb, 0x63, 0xe3, 0x03,
	0xad, 0xc2, 0xa5, 0x7c, 0x9b, 0x01, 0xe2, 0x79, 0x06, 0x51, 0xc5, 0xa7, 0x93, 0x20, 0xd6, 0x0d,
	0xc7, 0x85, 0xb3, 0xe5, 0xe0, 0x1f, 0x01, 0x28, 0x18, 0x91, 0x64, 0x80, 0x0a, 0x0f, 0x96, 0x32,
	0x40, 0x45, 0xa6, 0x2e, 0xe9, 0x76, 0x63, 0xa0, 0xb8, 0xdd, 0x9c, 0x48, 0xda, 0xfc, 0xa3, 0x82,
	0x8e, 0x24, 0x0e, 0x94, 0xf0, 0xb5, 0x3c, 0xf2, 0x63, 0x03, 0xa8, 0x5d, 0xc2, 0x9e, 0x63, 0xb0,
	0x6f, 0xe2, 0xd9, 0x2c, 0xd8, 0xde, 0x99, 0xf2, 0x53, 0x68, 0x28, 0x9b, 0x7e, 0x47, 0x41, 0x23,
	0x7e, 0x5f, 0x2f, 0x77, 0x4c, 0xbe, 0x95, 0x5e, 0x08, 0xca, 0x21, 0x99, 0x7d, 0x21, 0x41, 0x71,
	0x1b, 0x8e, 0xc8, 0xbf, 0x2a, 0xd0, 0x2e, 0x8f, 0xce, 0x2e, 0x52, 0xce, 0x7d, 0x97, 0x49, 0x4b,
	0xca, 0xb9, 0xef, 0x36, 0x18, 0x51, 0xef, 0x33, 0xd4, 0x77, 0xf1, 0x62, 0xe2, 0xf5, 0xce, 0xbb,
	0x79, 0x55, 0xcb, 0x16, 0x75, 0x65, 0x69, 0x47, 0xf4, 0x22, 0x3b, 0xa5, 0x9d, 0xd8, 0xe4, 0xa6,
	0x83, 0xff, 0xa6, 0xa0, 0xd1, 0xe8, 0x3c, 0x21, 0x45, 0x91, 0x2e, 0x63, 0x95, 0x14, 0x45, 0xba,
	0x0d, 0x2b, 0xd4, 0x35, 0xa6, 0xc8, 0x03, 0x7c, 0x2f, 0x49, 0x91, 0x2d, 0x46, 0xa5, 0x49, 0x7f,
	0x5f, 0xb3, 0x23, 0x86, 0x31, 0x9d, 0x68, 0x2a, 0x93, 0xe6, 0x2a, 0x1d, 0xfc, 0x53, 0x05, 0x0d,
	0xf9, 0x51, 0x83, 0xdf, 0x4a, 0xcd, 0xab, 0x72, 0x17, 0xb7, 0x70, 0x21, 0xcf, 0xd6, 0x3c, 0xd1,
	0x1d, 0x44, 0x4e, 0x69, 0x47, 0x7a, 0x58, 0x75, 0xc4, 0x37, 0x7e, 0x3e, 0xbd, 0xaa, 0x2b, 0x98,
	0x02, 0xa4, 0x5c, 0xc8, 0xb1, 0x41, 0x46, 0xe1, 0x62, 0xae, 0xbd, 0x79, 0x82, 0x9c, 0x1d, 0x44,
	0x86, 0xca, 0x09, 0x63, 0xc5, 0x3f, 0x51, 0xd0, 0xe1, 0x48, 0x53, 0x1d, 0x97, 0xb2, 0x2d, 0x14,
	0x9a, 0x14, 0x14, 0x2e, 0xe7, 0x27, 0x00, 0xb4, 0x53, 0x0c, 0xed, 0x39, 0xfc, 0xff, 0x19, 0x47,
	0x12, 0x06, 0x0b, 0x7f, 0x16, 0x0d, 0xe5, 0x70, 0xc3, 0x3c, 0xa5, 0x5a, 0x48, 0xec, 0xe0, 0x17,
	0x4a, 0xb9, 0xf7, 0x03, 0xce, 0x7b, 0x0c, 0xe7, 0x12, 0x5e, 0xc8, 0x38, 0x84, 0x10, 0x06, 0x89,
	0x47, 0x50, 0xbc, 0x7c, 0x3b, 0xde, 0x75, 0x72, 0x38, 0xd2, 0x6a, 0x4f, 0x09, 0x88, 0x58, 0x1b,
	0x3f, 0x25, 0x20, 0xe2, 0xbd, 0x7b, 0xf5, 0x2a, 0x83, 0x5e, 0xc4, 0x97, 0x52, 0xa0, 0x43, 0x9d,
	0xe3, 0xcf, 0x06, 0x3a, 0xf8, 0x1b, 0x0a, 0x1a, 0x91, 0x7b, 0xe3, 0xb8, 0xfb, 0xa3, 0x29, 0xdc,
	0xdc, 0x2f, 0x9c, 0xcf, 0xde, 0x08, 0xc8, 0xde, 0x64, 0xc8, 0x26, 0xf0, 0x89, 0xc4, 0x50, 0xb5,
	0x2a, 0x9b, 0x5a, 0x95, 0x52, 0xfc, 0x5b, 0x88, 0x4c, 0xa9, 0xe5, 0x9d, 0x11, 0x99, 0xf1, 0xe6,
	0x7a, 0x46, 0x64, 0x26, 0x74, 0xd3, 0xd5, 0x9b, 0x0c, 0xdc, 0x35, 0x7c, 0x25, 0xab, 0xf0, 0x66,
	0x9d, 0xf3, 0xc8, 0x65, 0xfc, 0x3b, 0x11, 0xa7, 0xe1, 0x26, 0x78, 0x4a, 0x9c, 0x26, 0x76, 0xdb,
	0x53, 0xe2, 0x34, 0xb9, 0xbb, 0xae, 0xde, 0x60, 0xa8, 0xaf, 0xe2, 0x99, 0x24, 0xd4, 0x86, 0xc3,
	0xdb, 0x91, 0x1a, 0x74, 0xdc, 0x23, 0xa0, 0x7f, 0xaf, 0xc0, 0x38, 0xe4, 0x51, 0xcb, 0x72, 0x49,
	0xd0, 0x96, 0x4b, 0xb1, 0x76, 0x72, 0x03, 0x30, 0xc5, 0xda, 0x5d, 0x3a, 0x7e, 0xe9, 0xd6, 0x7e,
	0xea, 0xe1, 0xd1, 0xa0, 0x23, 0xe8, 0x3d, 0x64, 0x23, 0xc0, 0xff, 0x22, 0x9e, 0xe0, 0xb1, 0xee,
	0x5a, 0xca, 0x13, 0xbc, 0x5b, 0xfb, 0x30, 0xe5, 0x09, 0xde, 0xb5, 0x79, 0xa7, 0x2e, 0x30, 0xf8,
	0x77, 0xf0, 0xad, 0x24, 0xf8, 0x72, 0x06, 0x73, 0x34, 0xd6, 0x7d, 0x12, 0xc9, 0xd7, 0xd0, 0x3b,
	0xa5, 0x1d, 0xf8, 0xa5, 0x83, 0x3f, 0x50, 0xd0, 0x68, 0xb4, 0x85, 0x95, 0x52, 0x6a, 0xc6, 0x5b,
	0x7b, 0x29, 0x35, 0x5b, 0x42, 0x57, 0x2c, 0x07, 0xea, 0x08, 0xdc, 0xf8, 0xbd, 0xe6, 0x74, 0xbc,
	0xf3, 0x39, 0x96, 0xd4, 0xf3, 0x4b, 0x09, 0x9b, 0xe4, 0xee, 0xe0, 0x2e, 0xd1, 0xa7, 0x86, 0xba,
	0x8c, 0x5e, 0x64, 0x37, 0xbf, 0xf3, 0xd8, 0xc1, 0xef, 0xf7, 0xa1, 0xc9, 0x7c, 0xdd, 0x2e, 0x3c,
	0x97, 0xd2, 0x91, 0xc9, 0xd7, 0xfc, 0x2b, 0xcc, 0xef, 0x85, 0x05, 0x68, 0xbb, 0xc1, 0xb4, 0xfd,
	0x32, 0x7e, 0x9c, 0xdc, 0xe4, 0x09, 0xb5, 0x16, 0x45, 0x66, 0x8a, 0xb4, 0xe1, 0x4a, 0x3b, 0x91,
	0x7d, 0x91, 0xc2, 0x6a, 0x7e, 0xe1, 0xa3, 0x17, 0x13, 0xca, 0xf3, 0x17, 0x13, 0xca, 0xbf, 0x5e,
	0x4c, 0x28, 0xdf, 0x7a, 0x39, 0xb1, 0xef, 0xf9, 0xcb, 0x89, 0x7d, 0xff, 0x7c, 0x39, 0xb1, 0xef,
	0xf1, 0x05, 0xa9, 0x65, 0x47, 0xb7, 0x1a, 0x96, 0x03, 0xff, 0x6e, 0x4d, 0xcf, 0x94, 0x9e, 0x85,
	0xff, 0xa0, 0x79, 0x63, 0x80, 0xfd, 0xb1, 0xf2, 0x95, 0xff, 0x06, 0x00, 0x00, 0xff, 0xff, 0xb4,
	0xea, 0x1d, 0x9f, 0x0a, 0x2e, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QueryClient interface {
	// Queries a bucket with specify name.
	HeadBucket(ctx context.Context, in *QueryHeadBucketRequest, opts ...grpc.CallOption) (*QueryHeadBucketResponse, error)
	// Queries a bucket by id
	HeadBucketById(ctx context.Context, in *QueryHeadBucketByIdRequest, opts ...grpc.CallOption) (*QueryHeadBucketResponse, error)
	// Queries a bucket with EIP712 standard metadata info
	HeadBucketNFT(ctx context.Context, in *QueryNFTRequest, opts ...grpc.CallOption) (*QueryBucketNFTResponse, error)
	// Queries a object with specify name.
	HeadObject(ctx context.Context, in *QueryHeadObjectRequest, opts ...grpc.CallOption) (*QueryHeadObjectResponse, error)
	// Queries an object by id
	HeadObjectById(ctx context.Context, in *QueryHeadObjectByIdRequest, opts ...grpc.CallOption) (*QueryHeadObjectResponse, error)
	// Queries a shadow object with specify name.
	HeadShadowObject(ctx context.Context, in *QueryHeadShadowObjectRequest, opts ...grpc.CallOption) (*QueryHeadShadowObjectResponse, error)
	// Queries a object with EIP712 standard metadata info
	HeadObjectNFT(ctx context.Context, in *QueryNFTRequest, opts ...grpc.CallOption) (*QueryObjectNFTResponse, error)
	// Queries a list of bucket items.
	ListBuckets(ctx context.Context, in *QueryListBucketsRequest, opts ...grpc.CallOption) (*QueryListBucketsResponse, error)
	// Queries a list of object items under the bucket.
	ListObjects(ctx context.Context, in *QueryListObjectsRequest, opts ...grpc.CallOption) (*QueryListObjectsResponse, error)
	// Queries a list of object items under the bucket.
	ListObjectsByBucketId(ctx context.Context, in *QueryListObjectsByBucketIdRequest, opts ...grpc.CallOption) (*QueryListObjectsResponse, error)
	// Queries a group with EIP712 standard metadata info
	HeadGroupNFT(ctx context.Context, in *QueryNFTRequest, opts ...grpc.CallOption) (*QueryGroupNFTResponse, error)
	// Queries a policy which grants permission to account
	QueryPolicyForAccount(ctx context.Context, in *QueryPolicyForAccountRequest, opts ...grpc.CallOption) (*QueryPolicyForAccountResponse, error)
	// Queries a list of VerifyPermission items.
	VerifyPermission(ctx context.Context, in *QueryVerifyPermissionRequest, opts ...grpc.CallOption) (*QueryVerifyPermissionResponse, error)
	// Queries a group with specify owner and name .
	HeadGroup(ctx context.Context, in *QueryHeadGroupRequest, opts ...grpc.CallOption) (*QueryHeadGroupResponse, error)
	// Queries a list of ListGroup items.
	ListGroups(ctx context.Context, in *QueryListGroupsRequest, opts ...grpc.CallOption) (*QueryListGroupsResponse, error)
	// Queries a list of HeadGroupMember items.
	HeadGroupMember(ctx context.Context, in *QueryHeadGroupMemberRequest, opts ...grpc.CallOption) (*QueryHeadGroupMemberResponse, error)
	// Queries a policy that grants permission to a group
	QueryPolicyForGroup(ctx context.Context, in *QueryPolicyForGroupRequest, opts ...grpc.CallOption) (*QueryPolicyForGroupResponse, error)
	// Queries a policy by policy id
	QueryPolicyById(ctx context.Context, in *QueryPolicyByIdRequest, opts ...grpc.CallOption) (*QueryPolicyByIdResponse, error)
	// Queries lock fee for storing an object
	QueryLockFee(ctx context.Context, in *QueryLockFeeRequest, opts ...grpc.CallOption) (*QueryLockFeeResponse, error)
	// Queries a bucket extra info (with gvg bindings and price time) with specify name.
	HeadBucketExtra(ctx context.Context, in *QueryHeadBucketExtraRequest, opts ...grpc.CallOption) (*QueryHeadBucketExtraResponse, error)
	// Queries whether read and storage prices changed for the bucket.
	QueryIsPriceChanged(ctx context.Context, in *QueryIsPriceChangedRequest, opts ...grpc.CallOption) (*QueryIsPriceChangedResponse, error)
	// Queries whether read and storage prices changed for the bucket.
	QueryQuotaUpdateTime(ctx context.Context, in *QueryQuoteUpdateTimeRequest, opts ...grpc.CallOption) (*QueryQuoteUpdateTimeResponse, error)
	// Queries whether some members are in the group.
	QueryGroupMembersExist(ctx context.Context, in *QueryGroupMembersExistRequest, opts ...grpc.CallOption) (*QueryGroupMembersExistResponse, error)
	// Queries whether some groups are exist.
	QueryGroupsExist(ctx context.Context, in *QueryGroupsExistRequest, opts ...grpc.CallOption) (*QueryGroupsExistResponse, error)
	// Queries whether some groups are exist by id.
	QueryGroupsExistById(ctx context.Context, in *QueryGroupsExistByIdRequest, opts ...grpc.CallOption) (*QueryGroupsExistResponse, error)
	// Queries the flow rate limit of a bucket for a payment account
	QueryPaymentAccountBucketFlowRateLimit(ctx context.Context, in *QueryPaymentAccountBucketFlowRateLimitRequest, opts ...grpc.CallOption) (*QueryPaymentAccountBucketFlowRateLimitResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) HeadBucket(ctx context.Context, in *QueryHeadBucketRequest, opts ...grpc.CallOption) (*QueryHeadBucketResponse, error) {
	out := new(QueryHeadBucketResponse)
	err := c.cc.Invoke(ctx, "/greenfield.storage.Query/HeadBucket", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) HeadBucketById(ctx context.Context, in *QueryHeadBucketByIdRequest, opts ...grpc.CallOption) (*QueryHeadBucketResponse, error) {
	out := new(QueryHeadBucketResponse)
	err := c.cc.Invoke(ctx, "/greenfield.storage.Query/HeadBucketById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) HeadBucketNFT(ctx context.Context, in *QueryNFTRequest, opts ...grpc.CallOption) (*QueryBucketNFTResponse, error) {
	out := new(QueryBucketNFTResponse)
	err := c.cc.Invoke(ctx, "/greenfield.storage.Query/HeadBucketNFT", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) HeadObject(ctx context.Context, in *QueryHeadObjectRequest, opts ...grpc.CallOption) (*QueryHeadObjectResponse, error) {
	out := new(QueryHeadObjectResponse)
	err := c.cc.Invoke(ctx, "/greenfield.storage.Query/HeadObject", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) HeadObjectById(ctx context.Context, in *QueryHeadObjectByIdRequest, opts ...grpc.CallOption) (*QueryHeadObjectResponse, error) {
	out := new(QueryHeadObjectResponse)
	err := c.cc.Invoke(ctx, "/greenfield.storage.Query/HeadObjectById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) HeadShadowObject(ctx context.Context, in *QueryHeadShadowObjectRequest, opts ...grpc.CallOption) (*QueryHeadShadowObjectResponse, error) {
	out := new(QueryHeadShadowObjectResponse)
	err := c.cc.Invoke(ctx, "/greenfield.storage.Query/HeadShadowObject", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) HeadObjectNFT(ctx context.Context, in *QueryNFTRequest, opts ...grpc.CallOption) (*QueryObjectNFTResponse, error) {
	out := new(QueryObjectNFTResponse)
	err := c.cc.Invoke(ctx, "/greenfield.storage.Query/HeadObjectNFT", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) ListBuckets(ctx context.Context, in *QueryListBucketsRequest, opts ...grpc.CallOption) (*QueryListBucketsResponse, error) {
	out := new(QueryListBucketsResponse)
	err := c.cc.Invoke(ctx, "/greenfield.storage.Query/ListBuckets", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) ListObjects(ctx context.Context, in *QueryListObjectsRequest, opts ...grpc.CallOption) (*QueryListObjectsResponse, error) {
	out := new(QueryListObjectsResponse)
	err := c.cc.Invoke(ctx, "/greenfield.storage.Query/ListObjects", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) ListObjectsByBucketId(ctx context.Context, in *QueryListObjectsByBucketIdRequest, opts ...grpc.CallOption) (*QueryListObjectsResponse, error) {
	out := new(QueryListObjectsResponse)
	err := c.cc.Invoke(ctx, "/greenfield.storage.Query/ListObjectsByBucketId", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) HeadGroupNFT(ctx context.Context, in *QueryNFTRequest, opts ...grpc.CallOption) (*QueryGroupNFTResponse, error) {
	out := new(QueryGroupNFTResponse)
	err := c.cc.Invoke(ctx, "/greenfield.storage.Query/HeadGroupNFT", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) QueryPolicyForAccount(ctx context.Context, in *QueryPolicyForAccountRequest, opts ...grpc.CallOption) (*QueryPolicyForAccountResponse, error) {
	out := new(QueryPolicyForAccountResponse)
	err := c.cc.Invoke(ctx, "/greenfield.storage.Query/QueryPolicyForAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) VerifyPermission(ctx context.Context, in *QueryVerifyPermissionRequest, opts ...grpc.CallOption) (*QueryVerifyPermissionResponse, error) {
	out := new(QueryVerifyPermissionResponse)
	err := c.cc.Invoke(ctx, "/greenfield.storage.Query/VerifyPermission", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) HeadGroup(ctx context.Context, in *QueryHeadGroupRequest, opts ...grpc.CallOption) (*QueryHeadGroupResponse, error) {
	out := new(QueryHeadGroupResponse)
	err := c.cc.Invoke(ctx, "/greenfield.storage.Query/HeadGroup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) ListGroups(ctx context.Context, in *QueryListGroupsRequest, opts ...grpc.CallOption) (*QueryListGroupsResponse, error) {
	out := new(QueryListGroupsResponse)
	err := c.cc.Invoke(ctx, "/greenfield.storage.Query/ListGroups", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) HeadGroupMember(ctx context.Context, in *QueryHeadGroupMemberRequest, opts ...grpc.CallOption) (*QueryHeadGroupMemberResponse, error) {
	out := new(QueryHeadGroupMemberResponse)
	err := c.cc.Invoke(ctx, "/greenfield.storage.Query/HeadGroupMember", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) QueryPolicyForGroup(ctx context.Context, in *QueryPolicyForGroupRequest, opts ...grpc.CallOption) (*QueryPolicyForGroupResponse, error) {
	out := new(QueryPolicyForGroupResponse)
	err := c.cc.Invoke(ctx, "/greenfield.storage.Query/QueryPolicyForGroup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) QueryPolicyById(ctx context.Context, in *QueryPolicyByIdRequest, opts ...grpc.CallOption) (*QueryPolicyByIdResponse, error) {
	out := new(QueryPolicyByIdResponse)
	err := c.cc.Invoke(ctx, "/greenfield.storage.Query/QueryPolicyById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) QueryLockFee(ctx context.Context, in *QueryLockFeeRequest, opts ...grpc.CallOption) (*QueryLockFeeResponse, error) {
	out := new(QueryLockFeeResponse)
	err := c.cc.Invoke(ctx, "/greenfield.storage.Query/QueryLockFee", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) HeadBucketExtra(ctx context.Context, in *QueryHeadBucketExtraRequest, opts ...grpc.CallOption) (*QueryHeadBucketExtraResponse, error) {
	out := new(QueryHeadBucketExtraResponse)
	err := c.cc.Invoke(ctx, "/greenfield.storage.Query/HeadBucketExtra", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) QueryIsPriceChanged(ctx context.Context, in *QueryIsPriceChangedRequest, opts ...grpc.CallOption) (*QueryIsPriceChangedResponse, error) {
	out := new(QueryIsPriceChangedResponse)
	err := c.cc.Invoke(ctx, "/greenfield.storage.Query/QueryIsPriceChanged", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) QueryQuotaUpdateTime(ctx context.Context, in *QueryQuoteUpdateTimeRequest, opts ...grpc.CallOption) (*QueryQuoteUpdateTimeResponse, error) {
	out := new(QueryQuoteUpdateTimeResponse)
	err := c.cc.Invoke(ctx, "/greenfield.storage.Query/QueryQuotaUpdateTime", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) QueryGroupMembersExist(ctx context.Context, in *QueryGroupMembersExistRequest, opts ...grpc.CallOption) (*QueryGroupMembersExistResponse, error) {
	out := new(QueryGroupMembersExistResponse)
	err := c.cc.Invoke(ctx, "/greenfield.storage.Query/QueryGroupMembersExist", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) QueryGroupsExist(ctx context.Context, in *QueryGroupsExistRequest, opts ...grpc.CallOption) (*QueryGroupsExistResponse, error) {
	out := new(QueryGroupsExistResponse)
	err := c.cc.Invoke(ctx, "/greenfield.storage.Query/QueryGroupsExist", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) QueryGroupsExistById(ctx context.Context, in *QueryGroupsExistByIdRequest, opts ...grpc.CallOption) (*QueryGroupsExistResponse, error) {
	out := new(QueryGroupsExistResponse)
	err := c.cc.Invoke(ctx, "/greenfield.storage.Query/QueryGroupsExistById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) QueryPaymentAccountBucketFlowRateLimit(ctx context.Context, in *QueryPaymentAccountBucketFlowRateLimitRequest, opts ...grpc.CallOption) (*QueryPaymentAccountBucketFlowRateLimitResponse, error) {
	out := new(QueryPaymentAccountBucketFlowRateLimitResponse)
	err := c.cc.Invoke(ctx, "/greenfield.storage.Query/QueryPaymentAccountBucketFlowRateLimit", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	// Queries a bucket with specify name.
	HeadBucket(context.Context, *QueryHeadBucketRequest) (*QueryHeadBucketResponse, error)
	// Queries a bucket by id
	HeadBucketById(context.Context, *QueryHeadBucketByIdRequest) (*QueryHeadBucketResponse, error)
	// Queries a bucket with EIP712 standard metadata info
	HeadBucketNFT(context.Context, *QueryNFTRequest) (*QueryBucketNFTResponse, error)
	// Queries a object with specify name.
	HeadObject(context.Context, *QueryHeadObjectRequest) (*QueryHeadObjectResponse, error)
	// Queries an object by id
	HeadObjectById(context.Context, *QueryHeadObjectByIdRequest) (*QueryHeadObjectResponse, error)
	// Queries a shadow object with specify name.
	HeadShadowObject(context.Context, *QueryHeadShadowObjectRequest) (*QueryHeadShadowObjectResponse, error)
	// Queries a object with EIP712 standard metadata info
	HeadObjectNFT(context.Context, *QueryNFTRequest) (*QueryObjectNFTResponse, error)
	// Queries a list of bucket items.
	ListBuckets(context.Context, *QueryListBucketsRequest) (*QueryListBucketsResponse, error)
	// Queries a list of object items under the bucket.
	ListObjects(context.Context, *QueryListObjectsRequest) (*QueryListObjectsResponse, error)
	// Queries a list of object items under the bucket.
	ListObjectsByBucketId(context.Context, *QueryListObjectsByBucketIdRequest) (*QueryListObjectsResponse, error)
	// Queries a group with EIP712 standard metadata info
	HeadGroupNFT(context.Context, *QueryNFTRequest) (*QueryGroupNFTResponse, error)
	// Queries a policy which grants permission to account
	QueryPolicyForAccount(context.Context, *QueryPolicyForAccountRequest) (*QueryPolicyForAccountResponse, error)
	// Queries a list of VerifyPermission items.
	VerifyPermission(context.Context, *QueryVerifyPermissionRequest) (*QueryVerifyPermissionResponse, error)
	// Queries a group with specify owner and name .
	HeadGroup(context.Context, *QueryHeadGroupRequest) (*QueryHeadGroupResponse, error)
	// Queries a list of ListGroup items.
	ListGroups(context.Context, *QueryListGroupsRequest) (*QueryListGroupsResponse, error)
	// Queries a list of HeadGroupMember items.
	HeadGroupMember(context.Context, *QueryHeadGroupMemberRequest) (*QueryHeadGroupMemberResponse, error)
	// Queries a policy that grants permission to a group
	QueryPolicyForGroup(context.Context, *QueryPolicyForGroupRequest) (*QueryPolicyForGroupResponse, error)
	// Queries a policy by policy id
	QueryPolicyById(context.Context, *QueryPolicyByIdRequest) (*QueryPolicyByIdResponse, error)
	// Queries lock fee for storing an object
	QueryLockFee(context.Context, *QueryLockFeeRequest) (*QueryLockFeeResponse, error)
	// Queries a bucket extra info (with gvg bindings and price time) with specify name.
	HeadBucketExtra(context.Context, *QueryHeadBucketExtraRequest) (*QueryHeadBucketExtraResponse, error)
	// Queries whether read and storage prices changed for the bucket.
	QueryIsPriceChanged(context.Context, *QueryIsPriceChangedRequest) (*QueryIsPriceChangedResponse, error)
	// Queries whether read and storage prices changed for the bucket.
	QueryQuotaUpdateTime(context.Context, *QueryQuoteUpdateTimeRequest) (*QueryQuoteUpdateTimeResponse, error)
	// Queries whether some members are in the group.
	QueryGroupMembersExist(context.Context, *QueryGroupMembersExistRequest) (*QueryGroupMembersExistResponse, error)
	// Queries whether some groups are exist.
	QueryGroupsExist(context.Context, *QueryGroupsExistRequest) (*QueryGroupsExistResponse, error)
	// Queries whether some groups are exist by id.
	QueryGroupsExistById(context.Context, *QueryGroupsExistByIdRequest) (*QueryGroupsExistResponse, error)
	// Queries the flow rate limit of a bucket for a payment account
	QueryPaymentAccountBucketFlowRateLimit(context.Context, *QueryPaymentAccountBucketFlowRateLimitRequest) (*QueryPaymentAccountBucketFlowRateLimitResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) HeadBucket(ctx context.Context, req *QueryHeadBucketRequest) (*QueryHeadBucketResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HeadBucket not implemented")
}
func (*UnimplementedQueryServer) HeadBucketById(ctx context.Context, req *QueryHeadBucketByIdRequest) (*QueryHeadBucketResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HeadBucketById not implemented")
}
func (*UnimplementedQueryServer) HeadBucketNFT(ctx context.Context, req *QueryNFTRequest) (*QueryBucketNFTResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HeadBucketNFT not implemented")
}
func (*UnimplementedQueryServer) HeadObject(ctx context.Context, req *QueryHeadObjectRequest) (*QueryHeadObjectResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HeadObject not implemented")
}
func (*UnimplementedQueryServer) HeadObjectById(ctx context.Context, req *QueryHeadObjectByIdRequest) (*QueryHeadObjectResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HeadObjectById not implemented")
}
func (*UnimplementedQueryServer) HeadShadowObject(ctx context.Context, req *QueryHeadShadowObjectRequest) (*QueryHeadShadowObjectResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HeadShadowObject not implemented")
}
func (*UnimplementedQueryServer) HeadObjectNFT(ctx context.Context, req *QueryNFTRequest) (*QueryObjectNFTResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HeadObjectNFT not implemented")
}
func (*UnimplementedQueryServer) ListBuckets(ctx context.Context, req *QueryListBucketsRequest) (*QueryListBucketsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListBuckets not implemented")
}
func (*UnimplementedQueryServer) ListObjects(ctx context.Context, req *QueryListObjectsRequest) (*QueryListObjectsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListObjects not implemented")
}
func (*UnimplementedQueryServer) ListObjectsByBucketId(ctx context.Context, req *QueryListObjectsByBucketIdRequest) (*QueryListObjectsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListObjectsByBucketId not implemented")
}
func (*UnimplementedQueryServer) HeadGroupNFT(ctx context.Context, req *QueryNFTRequest) (*QueryGroupNFTResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HeadGroupNFT not implemented")
}
func (*UnimplementedQueryServer) QueryPolicyForAccount(ctx context.Context, req *QueryPolicyForAccountRequest) (*QueryPolicyForAccountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryPolicyForAccount not implemented")
}
func (*UnimplementedQueryServer) VerifyPermission(ctx context.Context, req *QueryVerifyPermissionRequest) (*QueryVerifyPermissionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VerifyPermission not implemented")
}
func (*UnimplementedQueryServer) HeadGroup(ctx context.Context, req *QueryHeadGroupRequest) (*QueryHeadGroupResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HeadGroup not implemented")
}
func (*UnimplementedQueryServer) ListGroups(ctx context.Context, req *QueryListGroupsRequest) (*QueryListGroupsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListGroups not implemented")
}
func (*UnimplementedQueryServer) HeadGroupMember(ctx context.Context, req *QueryHeadGroupMemberRequest) (*QueryHeadGroupMemberResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HeadGroupMember not implemented")
}
func (*UnimplementedQueryServer) QueryPolicyForGroup(ctx context.Context, req *QueryPolicyForGroupRequest) (*QueryPolicyForGroupResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryPolicyForGroup not implemented")
}
func (*UnimplementedQueryServer) QueryPolicyById(ctx context.Context, req *QueryPolicyByIdRequest) (*QueryPolicyByIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryPolicyById not implemented")
}
func (*UnimplementedQueryServer) QueryLockFee(ctx context.Context, req *QueryLockFeeRequest) (*QueryLockFeeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryLockFee not implemented")
}
func (*UnimplementedQueryServer) HeadBucketExtra(ctx context.Context, req *QueryHeadBucketExtraRequest) (*QueryHeadBucketExtraResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HeadBucketExtra not implemented")
}
func (*UnimplementedQueryServer) QueryIsPriceChanged(ctx context.Context, req *QueryIsPriceChangedRequest) (*QueryIsPriceChangedResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryIsPriceChanged not implemented")
}
func (*UnimplementedQueryServer) QueryQuotaUpdateTime(ctx context.Context, req *QueryQuoteUpdateTimeRequest) (*QueryQuoteUpdateTimeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryQuotaUpdateTime not implemented")
}
func (*UnimplementedQueryServer) QueryGroupMembersExist(ctx context.Context, req *QueryGroupMembersExistRequest) (*QueryGroupMembersExistResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryGroupMembersExist not implemented")
}
func (*UnimplementedQueryServer) QueryGroupsExist(ctx context.Context, req *QueryGroupsExistRequest) (*QueryGroupsExistResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryGroupsExist not implemented")
}
func (*UnimplementedQueryServer) QueryGroupsExistById(ctx context.Context, req *QueryGroupsExistByIdRequest) (*QueryGroupsExistResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryGroupsExistById not implemented")
}
func (*UnimplementedQueryServer) QueryPaymentAccountBucketFlowRateLimit(ctx context.Context, req *QueryPaymentAccountBucketFlowRateLimitRequest) (*QueryPaymentAccountBucketFlowRateLimitResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryPaymentAccountBucketFlowRateLimit not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_HeadBucket_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryHeadBucketRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).HeadBucket(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/greenfield.storage.Query/HeadBucket",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).HeadBucket(ctx, req.(*QueryHeadBucketRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_HeadBucketById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryHeadBucketByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).HeadBucketById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/greenfield.storage.Query/HeadBucketById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).HeadBucketById(ctx, req.(*QueryHeadBucketByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_HeadBucketNFT_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryNFTRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).HeadBucketNFT(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/greenfield.storage.Query/HeadBucketNFT",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).HeadBucketNFT(ctx, req.(*QueryNFTRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_HeadObject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryHeadObjectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).HeadObject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/greenfield.storage.Query/HeadObject",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).HeadObject(ctx, req.(*QueryHeadObjectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_HeadObjectById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryHeadObjectByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).HeadObjectById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/greenfield.storage.Query/HeadObjectById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).HeadObjectById(ctx, req.(*QueryHeadObjectByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_HeadShadowObject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryHeadShadowObjectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).HeadShadowObject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/greenfield.storage.Query/HeadShadowObject",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).HeadShadowObject(ctx, req.(*QueryHeadShadowObjectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_HeadObjectNFT_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryNFTRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).HeadObjectNFT(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/greenfield.storage.Query/HeadObjectNFT",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).HeadObjectNFT(ctx, req.(*QueryNFTRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_ListBuckets_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryListBucketsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).ListBuckets(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/greenfield.storage.Query/ListBuckets",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).ListBuckets(ctx, req.(*QueryListBucketsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_ListObjects_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryListObjectsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).ListObjects(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/greenfield.storage.Query/ListObjects",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).ListObjects(ctx, req.(*QueryListObjectsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_ListObjectsByBucketId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryListObjectsByBucketIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).ListObjectsByBucketId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/greenfield.storage.Query/ListObjectsByBucketId",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).ListObjectsByBucketId(ctx, req.(*QueryListObjectsByBucketIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_HeadGroupNFT_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryNFTRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).HeadGroupNFT(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/greenfield.storage.Query/HeadGroupNFT",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).HeadGroupNFT(ctx, req.(*QueryNFTRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_QueryPolicyForAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryPolicyForAccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).QueryPolicyForAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/greenfield.storage.Query/QueryPolicyForAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).QueryPolicyForAccount(ctx, req.(*QueryPolicyForAccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_VerifyPermission_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryVerifyPermissionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).VerifyPermission(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/greenfield.storage.Query/VerifyPermission",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).VerifyPermission(ctx, req.(*QueryVerifyPermissionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_HeadGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryHeadGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).HeadGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/greenfield.storage.Query/HeadGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).HeadGroup(ctx, req.(*QueryHeadGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_ListGroups_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryListGroupsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).ListGroups(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/greenfield.storage.Query/ListGroups",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).ListGroups(ctx, req.(*QueryListGroupsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_HeadGroupMember_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryHeadGroupMemberRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).HeadGroupMember(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/greenfield.storage.Query/HeadGroupMember",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).HeadGroupMember(ctx, req.(*QueryHeadGroupMemberRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_QueryPolicyForGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryPolicyForGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).QueryPolicyForGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/greenfield.storage.Query/QueryPolicyForGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).QueryPolicyForGroup(ctx, req.(*QueryPolicyForGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_QueryPolicyById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryPolicyByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).QueryPolicyById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/greenfield.storage.Query/QueryPolicyById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).QueryPolicyById(ctx, req.(*QueryPolicyByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_QueryLockFee_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryLockFeeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).QueryLockFee(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/greenfield.storage.Query/QueryLockFee",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).QueryLockFee(ctx, req.(*QueryLockFeeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_HeadBucketExtra_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryHeadBucketExtraRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).HeadBucketExtra(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/greenfield.storage.Query/HeadBucketExtra",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).HeadBucketExtra(ctx, req.(*QueryHeadBucketExtraRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_QueryIsPriceChanged_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryIsPriceChangedRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).QueryIsPriceChanged(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/greenfield.storage.Query/QueryIsPriceChanged",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).QueryIsPriceChanged(ctx, req.(*QueryIsPriceChangedRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_QueryQuotaUpdateTime_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryQuoteUpdateTimeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).QueryQuotaUpdateTime(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/greenfield.storage.Query/QueryQuotaUpdateTime",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).QueryQuotaUpdateTime(ctx, req.(*QueryQuoteUpdateTimeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_QueryGroupMembersExist_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryGroupMembersExistRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).QueryGroupMembersExist(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/greenfield.storage.Query/QueryGroupMembersExist",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).QueryGroupMembersExist(ctx, req.(*QueryGroupMembersExistRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_QueryGroupsExist_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryGroupsExistRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).QueryGroupsExist(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/greenfield.storage.Query/QueryGroupsExist",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).QueryGroupsExist(ctx, req.(*QueryGroupsExistRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_QueryGroupsExistById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryGroupsExistByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).QueryGroupsExistById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/greenfield.storage.Query/QueryGroupsExistById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).QueryGroupsExistById(ctx, req.(*QueryGroupsExistByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_QueryPaymentAccountBucketFlowRateLimit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryPaymentAccountBucketFlowRateLimitRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).QueryPaymentAccountBucketFlowRateLimit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/greenfield.storage.Query/QueryPaymentAccountBucketFlowRateLimit",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).QueryPaymentAccountBucketFlowRateLimit(ctx, req.(*QueryPaymentAccountBucketFlowRateLimitRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "greenfield.storage.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "HeadBucket",
			Handler:    _Query_HeadBucket_Handler,
		},
		{
			MethodName: "HeadBucketById",
			Handler:    _Query_HeadBucketById_Handler,
		},
		{
			MethodName: "HeadBucketNFT",
			Handler:    _Query_HeadBucketNFT_Handler,
		},
		{
			MethodName: "HeadObject",
			Handler:    _Query_HeadObject_Handler,
		},
		{
			MethodName: "HeadObjectById",
			Handler:    _Query_HeadObjectById_Handler,
		},
		{
			MethodName: "HeadShadowObject",
			Handler:    _Query_HeadShadowObject_Handler,
		},
		{
			MethodName: "HeadObjectNFT",
			Handler:    _Query_HeadObjectNFT_Handler,
		},
		{
			MethodName: "ListBuckets",
			Handler:    _Query_ListBuckets_Handler,
		},
		{
			MethodName: "ListObjects",
			Handler:    _Query_ListObjects_Handler,
		},
		{
			MethodName: "ListObjectsByBucketId",
			Handler:    _Query_ListObjectsByBucketId_Handler,
		},
		{
			MethodName: "HeadGroupNFT",
			Handler:    _Query_HeadGroupNFT_Handler,
		},
		{
			MethodName: "QueryPolicyForAccount",
			Handler:    _Query_QueryPolicyForAccount_Handler,
		},
		{
			MethodName: "VerifyPermission",
			Handler:    _Query_VerifyPermission_Handler,
		},
		{
			MethodName: "HeadGroup",
			Handler:    _Query_HeadGroup_Handler,
		},
		{
			MethodName: "ListGroups",
			Handler:    _Query_ListGroups_Handler,
		},
		{
			MethodName: "HeadGroupMember",
			Handler:    _Query_HeadGroupMember_Handler,
		},
		{
			MethodName: "QueryPolicyForGroup",
			Handler:    _Query_QueryPolicyForGroup_Handler,
		},
		{
			MethodName: "QueryPolicyById",
			Handler:    _Query_QueryPolicyById_Handler,
		},
		{
			MethodName: "QueryLockFee",
			Handler:    _Query_QueryLockFee_Handler,
		},
		{
			MethodName: "HeadBucketExtra",
			Handler:    _Query_HeadBucketExtra_Handler,
		},
		{
			MethodName: "QueryIsPriceChanged",
			Handler:    _Query_QueryIsPriceChanged_Handler,
		},
		{
			MethodName: "QueryQuotaUpdateTime",
			Handler:    _Query_QueryQuotaUpdateTime_Handler,
		},
		{
			MethodName: "QueryGroupMembersExist",
			Handler:    _Query_QueryGroupMembersExist_Handler,
		},
		{
			MethodName: "QueryGroupsExist",
			Handler:    _Query_QueryGroupsExist_Handler,
		},
		{
			MethodName: "QueryGroupsExistById",
			Handler:    _Query_QueryGroupsExistById_Handler,
		},
		{
			MethodName: "QueryPaymentAccountBucketFlowRateLimit",
			Handler:    _Query_QueryPaymentAccountBucketFlowRateLimit_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "greenfield/storage/query.proto",
}

func (m *QueryParamsRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryParamsRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryParamsRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *QueryHeadBucketRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryHeadBucketRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryHeadBucketRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.BucketName) > 0 {
		i -= len(m.BucketName)
		copy(dAtA[i:], m.BucketName)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.BucketName)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryHeadBucketByIdRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryHeadBucketByIdRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryHeadBucketByIdRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.BucketId) > 0 {
		i -= len(m.BucketId)
		copy(dAtA[i:], m.BucketId)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.BucketId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryHeadBucketResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryHeadBucketResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryHeadBucketResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.ExtraInfo != nil {
		{
			size, err := m.ExtraInfo.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if m.BucketInfo != nil {
		{
			size, err := m.BucketInfo.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryHeadObjectRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryHeadObjectRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryHeadObjectRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ObjectName) > 0 {
		i -= len(m.ObjectName)
		copy(dAtA[i:], m.ObjectName)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.ObjectName)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.BucketName) > 0 {
		i -= len(m.BucketName)
		copy(dAtA[i:], m.BucketName)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.BucketName)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryHeadObjectByIdRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryHeadObjectByIdRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryHeadObjectByIdRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ObjectId) > 0 {
		i -= len(m.ObjectId)
		copy(dAtA[i:], m.ObjectId)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.ObjectId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryHeadShadowObjectRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryHeadShadowObjectRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryHeadShadowObjectRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ObjectName) > 0 {
		i -= len(m.ObjectName)
		copy(dAtA[i:], m.ObjectName)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.ObjectName)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.BucketName) > 0 {
		i -= len(m.BucketName)
		copy(dAtA[i:], m.BucketName)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.BucketName)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryHeadObjectResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryHeadObjectResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryHeadObjectResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.GlobalVirtualGroup != nil {
		{
			size, err := m.GlobalVirtualGroup.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if m.ObjectInfo != nil {
		{
			size, err := m.ObjectInfo.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryHeadShadowObjectResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryHeadShadowObjectResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryHeadShadowObjectResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.ObjectInfo != nil {
		{
			size, err := m.ObjectInfo.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryListBucketsRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryListBucketsRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryListBucketsRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Pagination != nil {
		{
			size, err := m.Pagination.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryListBucketsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryListBucketsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryListBucketsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Pagination != nil {
		{
			size, err := m.Pagination.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.BucketInfos) > 0 {
		for iNdEx := len(m.BucketInfos) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.BucketInfos[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintQuery(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *QueryListObjectsRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryListObjectsRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryListObjectsRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.BucketName) > 0 {
		i -= len(m.BucketName)
		copy(dAtA[i:], m.BucketName)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.BucketName)))
		i--
		dAtA[i] = 0x12
	}
	if m.Pagination != nil {
		{
			size, err := m.Pagination.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryListObjectsByBucketIdRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryListObjectsByBucketIdRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryListObjectsByBucketIdRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.BucketId) > 0 {
		i -= len(m.BucketId)
		copy(dAtA[i:], m.BucketId)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.BucketId)))
		i--
		dAtA[i] = 0x12
	}
	if m.Pagination != nil {
		{
			size, err := m.Pagination.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryListObjectsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryListObjectsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryListObjectsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Pagination != nil {
		{
			size, err := m.Pagination.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.ObjectInfos) > 0 {
		for iNdEx := len(m.ObjectInfos) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ObjectInfos[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintQuery(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *QueryNFTRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryNFTRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryNFTRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.TokenId) > 0 {
		i -= len(m.TokenId)
		copy(dAtA[i:], m.TokenId)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.TokenId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryBucketNFTResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryBucketNFTResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryBucketNFTResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.MetaData != nil {
		{
			size, err := m.MetaData.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryObjectNFTResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryObjectNFTResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryObjectNFTResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.MetaData != nil {
		{
			size, err := m.MetaData.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryGroupNFTResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryGroupNFTResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryGroupNFTResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.MetaData != nil {
		{
			size, err := m.MetaData.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryPolicyForAccountRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryPolicyForAccountRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryPolicyForAccountRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.PrincipalAddress) > 0 {
		i -= len(m.PrincipalAddress)
		copy(dAtA[i:], m.PrincipalAddress)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.PrincipalAddress)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Resource) > 0 {
		i -= len(m.Resource)
		copy(dAtA[i:], m.Resource)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Resource)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryPolicyForAccountResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryPolicyForAccountResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryPolicyForAccountResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Policy != nil {
		{
			size, err := m.Policy.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryVerifyPermissionRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryVerifyPermissionRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryVerifyPermissionRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.ActionType != 0 {
		i = encodeVarintQuery(dAtA, i, uint64(m.ActionType))
		i--
		dAtA[i] = 0x20
	}
	if len(m.ObjectName) > 0 {
		i -= len(m.ObjectName)
		copy(dAtA[i:], m.ObjectName)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.ObjectName)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.BucketName) > 0 {
		i -= len(m.BucketName)
		copy(dAtA[i:], m.BucketName)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.BucketName)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Operator) > 0 {
		i -= len(m.Operator)
		copy(dAtA[i:], m.Operator)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Operator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryVerifyPermissionResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryVerifyPermissionResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryVerifyPermissionResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Effect != 0 {
		i = encodeVarintQuery(dAtA, i, uint64(m.Effect))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *QueryHeadGroupRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryHeadGroupRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryHeadGroupRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.GroupName) > 0 {
		i -= len(m.GroupName)
		copy(dAtA[i:], m.GroupName)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.GroupName)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.GroupOwner) > 0 {
		i -= len(m.GroupOwner)
		copy(dAtA[i:], m.GroupOwner)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.GroupOwner)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryHeadGroupResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryHeadGroupResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryHeadGroupResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.GroupInfo != nil {
		{
			size, err := m.GroupInfo.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryListGroupsRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryListGroupsRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryListGroupsRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.GroupOwner) > 0 {
		i -= len(m.GroupOwner)
		copy(dAtA[i:], m.GroupOwner)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.GroupOwner)))
		i--
		dAtA[i] = 0x12
	}
	if m.Pagination != nil {
		{
			size, err := m.Pagination.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryListGroupsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryListGroupsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryListGroupsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.GroupInfos) > 0 {
		for iNdEx := len(m.GroupInfos) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.GroupInfos[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintQuery(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if m.Pagination != nil {
		{
			size, err := m.Pagination.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryHeadGroupMemberRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryHeadGroupMemberRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryHeadGroupMemberRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.GroupName) > 0 {
		i -= len(m.GroupName)
		copy(dAtA[i:], m.GroupName)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.GroupName)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.GroupOwner) > 0 {
		i -= len(m.GroupOwner)
		copy(dAtA[i:], m.GroupOwner)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.GroupOwner)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Member) > 0 {
		i -= len(m.Member)
		copy(dAtA[i:], m.Member)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Member)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryHeadGroupMemberResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryHeadGroupMemberResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryHeadGroupMemberResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.GroupMember != nil {
		{
			size, err := m.GroupMember.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryPolicyForGroupRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryPolicyForGroupRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryPolicyForGroupRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.PrincipalGroupId) > 0 {
		i -= len(m.PrincipalGroupId)
		copy(dAtA[i:], m.PrincipalGroupId)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.PrincipalGroupId)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Resource) > 0 {
		i -= len(m.Resource)
		copy(dAtA[i:], m.Resource)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Resource)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryPolicyForGroupResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryPolicyForGroupResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryPolicyForGroupResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Policy != nil {
		{
			size, err := m.Policy.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryPolicyByIdRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryPolicyByIdRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryPolicyByIdRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.PolicyId) > 0 {
		i -= len(m.PolicyId)
		copy(dAtA[i:], m.PolicyId)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.PolicyId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryPolicyByIdResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryPolicyByIdResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryPolicyByIdResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Policy != nil {
		{
			size, err := m.Policy.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryLockFeeRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryLockFeeRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryLockFeeRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.PayloadSize != 0 {
		i = encodeVarintQuery(dAtA, i, uint64(m.PayloadSize))
		i--
		dAtA[i] = 0x18
	}
	if m.CreateAt != 0 {
		i = encodeVarintQuery(dAtA, i, uint64(m.CreateAt))
		i--
		dAtA[i] = 0x10
	}
	if len(m.PrimarySpAddress) > 0 {
		i -= len(m.PrimarySpAddress)
		copy(dAtA[i:], m.PrimarySpAddress)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.PrimarySpAddress)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryLockFeeResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryLockFeeResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryLockFeeResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.Amount.Size()
		i -= size
		if _, err := m.Amount.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintQuery(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *QueryHeadBucketExtraRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryHeadBucketExtraRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryHeadBucketExtraRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.BucketName) > 0 {
		i -= len(m.BucketName)
		copy(dAtA[i:], m.BucketName)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.BucketName)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryHeadBucketExtraResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryHeadBucketExtraResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryHeadBucketExtraResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.ExtraInfo != nil {
		{
			size, err := m.ExtraInfo.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryIsPriceChangedRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryIsPriceChangedRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryIsPriceChangedRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.BucketName) > 0 {
		i -= len(m.BucketName)
		copy(dAtA[i:], m.BucketName)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.BucketName)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryIsPriceChangedResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryIsPriceChangedResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryIsPriceChangedResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.NewValidatorTaxRate.Size()
		i -= size
		if _, err := m.NewValidatorTaxRate.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintQuery(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x4a
	{
		size := m.NewSecondaryStorePrice.Size()
		i -= size
		if _, err := m.NewSecondaryStorePrice.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintQuery(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x42
	{
		size := m.NewPrimaryStorePrice.Size()
		i -= size
		if _, err := m.NewPrimaryStorePrice.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintQuery(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x3a
	{
		size := m.NewReadPrice.Size()
		i -= size
		if _, err := m.NewReadPrice.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintQuery(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x32
	{
		size := m.CurrentValidatorTaxRate.Size()
		i -= size
		if _, err := m.CurrentValidatorTaxRate.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintQuery(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x2a
	{
		size := m.CurrentSecondaryStorePrice.Size()
		i -= size
		if _, err := m.CurrentSecondaryStorePrice.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintQuery(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	{
		size := m.CurrentPrimaryStorePrice.Size()
		i -= size
		if _, err := m.CurrentPrimaryStorePrice.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintQuery(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	{
		size := m.CurrentReadPrice.Size()
		i -= size
		if _, err := m.CurrentReadPrice.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintQuery(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if m.Changed {
		i--
		if m.Changed {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *QueryQuoteUpdateTimeRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryQuoteUpdateTimeRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryQuoteUpdateTimeRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.BucketName) > 0 {
		i -= len(m.BucketName)
		copy(dAtA[i:], m.BucketName)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.BucketName)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryQuoteUpdateTimeResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryQuoteUpdateTimeResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryQuoteUpdateTimeResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.UpdateAt != 0 {
		i = encodeVarintQuery(dAtA, i, uint64(m.UpdateAt))
		i--
		dAtA[i] = 0x30
	}
	return len(dAtA) - i, nil
}

func (m *QueryGroupMembersExistRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryGroupMembersExistRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryGroupMembersExistRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Members) > 0 {
		for iNdEx := len(m.Members) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Members[iNdEx])
			copy(dAtA[i:], m.Members[iNdEx])
			i = encodeVarintQuery(dAtA, i, uint64(len(m.Members[iNdEx])))
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.GroupId) > 0 {
		i -= len(m.GroupId)
		copy(dAtA[i:], m.GroupId)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.GroupId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryGroupMembersExistResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryGroupMembersExistResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryGroupMembersExistResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Exists) > 0 {
		for k := range m.Exists {
			v := m.Exists[k]
			baseI := i
			i--
			if v {
				dAtA[i] = 1
			} else {
				dAtA[i] = 0
			}
			i--
			dAtA[i] = 0x10
			i -= len(k)
			copy(dAtA[i:], k)
			i = encodeVarintQuery(dAtA, i, uint64(len(k)))
			i--
			dAtA[i] = 0xa
			i = encodeVarintQuery(dAtA, i, uint64(baseI-i))
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *QueryGroupsExistRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryGroupsExistRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryGroupsExistRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.GroupNames) > 0 {
		for iNdEx := len(m.GroupNames) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.GroupNames[iNdEx])
			copy(dAtA[i:], m.GroupNames[iNdEx])
			i = encodeVarintQuery(dAtA, i, uint64(len(m.GroupNames[iNdEx])))
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.GroupOwner) > 0 {
		i -= len(m.GroupOwner)
		copy(dAtA[i:], m.GroupOwner)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.GroupOwner)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryGroupsExistByIdRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryGroupsExistByIdRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryGroupsExistByIdRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.GroupIds) > 0 {
		for iNdEx := len(m.GroupIds) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.GroupIds[iNdEx])
			copy(dAtA[i:], m.GroupIds[iNdEx])
			i = encodeVarintQuery(dAtA, i, uint64(len(m.GroupIds[iNdEx])))
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *QueryGroupsExistResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryGroupsExistResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryGroupsExistResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Exists) > 0 {
		for k := range m.Exists {
			v := m.Exists[k]
			baseI := i
			i--
			if v {
				dAtA[i] = 1
			} else {
				dAtA[i] = 0
			}
			i--
			dAtA[i] = 0x10
			i -= len(k)
			copy(dAtA[i:], k)
			i = encodeVarintQuery(dAtA, i, uint64(len(k)))
			i--
			dAtA[i] = 0xa
			i = encodeVarintQuery(dAtA, i, uint64(baseI-i))
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *QueryPaymentAccountBucketFlowRateLimitRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryPaymentAccountBucketFlowRateLimitRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryPaymentAccountBucketFlowRateLimitRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.BucketName) > 0 {
		i -= len(m.BucketName)
		copy(dAtA[i:], m.BucketName)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.BucketName)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.BucketOwner) > 0 {
		i -= len(m.BucketOwner)
		copy(dAtA[i:], m.BucketOwner)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.BucketOwner)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.PaymentAccount) > 0 {
		i -= len(m.PaymentAccount)
		copy(dAtA[i:], m.PaymentAccount)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.PaymentAccount)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryPaymentAccountBucketFlowRateLimitResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryPaymentAccountBucketFlowRateLimitResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryPaymentAccountBucketFlowRateLimitResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.FlowRateLimit.Size()
		i -= size
		if _, err := m.FlowRateLimit.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintQuery(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if m.IsSet {
		i--
		if m.IsSet {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintQuery(dAtA []byte, offset int, v uint64) int {
	offset -= sovQuery(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *QueryParamsRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *QueryHeadBucketRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.BucketName)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryHeadBucketByIdRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.BucketId)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryHeadBucketResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.BucketInfo != nil {
		l = m.BucketInfo.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	if m.ExtraInfo != nil {
		l = m.ExtraInfo.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryHeadObjectRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.BucketName)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	l = len(m.ObjectName)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryHeadObjectByIdRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ObjectId)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryHeadShadowObjectRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.BucketName)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	l = len(m.ObjectName)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryHeadObjectResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ObjectInfo != nil {
		l = m.ObjectInfo.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	if m.GlobalVirtualGroup != nil {
		l = m.GlobalVirtualGroup.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryHeadShadowObjectResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ObjectInfo != nil {
		l = m.ObjectInfo.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryListBucketsRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Pagination != nil {
		l = m.Pagination.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryListBucketsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.BucketInfos) > 0 {
		for _, e := range m.BucketInfos {
			l = e.Size()
			n += 1 + l + sovQuery(uint64(l))
		}
	}
	if m.Pagination != nil {
		l = m.Pagination.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryListObjectsRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Pagination != nil {
		l = m.Pagination.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	l = len(m.BucketName)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryListObjectsByBucketIdRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Pagination != nil {
		l = m.Pagination.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	l = len(m.BucketId)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryListObjectsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.ObjectInfos) > 0 {
		for _, e := range m.ObjectInfos {
			l = e.Size()
			n += 1 + l + sovQuery(uint64(l))
		}
	}
	if m.Pagination != nil {
		l = m.Pagination.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryNFTRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.TokenId)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryBucketNFTResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.MetaData != nil {
		l = m.MetaData.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryObjectNFTResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.MetaData != nil {
		l = m.MetaData.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryGroupNFTResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.MetaData != nil {
		l = m.MetaData.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryPolicyForAccountRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Resource)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	l = len(m.PrincipalAddress)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryPolicyForAccountResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Policy != nil {
		l = m.Policy.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryVerifyPermissionRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Operator)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	l = len(m.BucketName)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	l = len(m.ObjectName)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	if m.ActionType != 0 {
		n += 1 + sovQuery(uint64(m.ActionType))
	}
	return n
}

func (m *QueryVerifyPermissionResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Effect != 0 {
		n += 1 + sovQuery(uint64(m.Effect))
	}
	return n
}

func (m *QueryHeadGroupRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.GroupOwner)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	l = len(m.GroupName)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryHeadGroupResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.GroupInfo != nil {
		l = m.GroupInfo.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryListGroupsRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Pagination != nil {
		l = m.Pagination.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	l = len(m.GroupOwner)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryListGroupsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Pagination != nil {
		l = m.Pagination.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	if len(m.GroupInfos) > 0 {
		for _, e := range m.GroupInfos {
			l = e.Size()
			n += 1 + l + sovQuery(uint64(l))
		}
	}
	return n
}

func (m *QueryHeadGroupMemberRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Member)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	l = len(m.GroupOwner)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	l = len(m.GroupName)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryHeadGroupMemberResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.GroupMember != nil {
		l = m.GroupMember.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryPolicyForGroupRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Resource)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	l = len(m.PrincipalGroupId)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryPolicyForGroupResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Policy != nil {
		l = m.Policy.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryPolicyByIdRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.PolicyId)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryPolicyByIdResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Policy != nil {
		l = m.Policy.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryLockFeeRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.PrimarySpAddress)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	if m.CreateAt != 0 {
		n += 1 + sovQuery(uint64(m.CreateAt))
	}
	if m.PayloadSize != 0 {
		n += 1 + sovQuery(uint64(m.PayloadSize))
	}
	return n
}

func (m *QueryLockFeeResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Amount.Size()
	n += 1 + l + sovQuery(uint64(l))
	return n
}

func (m *QueryHeadBucketExtraRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.BucketName)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryHeadBucketExtraResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ExtraInfo != nil {
		l = m.ExtraInfo.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryIsPriceChangedRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.BucketName)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryIsPriceChangedResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Changed {
		n += 2
	}
	l = m.CurrentReadPrice.Size()
	n += 1 + l + sovQuery(uint64(l))
	l = m.CurrentPrimaryStorePrice.Size()
	n += 1 + l + sovQuery(uint64(l))
	l = m.CurrentSecondaryStorePrice.Size()
	n += 1 + l + sovQuery(uint64(l))
	l = m.CurrentValidatorTaxRate.Size()
	n += 1 + l + sovQuery(uint64(l))
	l = m.NewReadPrice.Size()
	n += 1 + l + sovQuery(uint64(l))
	l = m.NewPrimaryStorePrice.Size()
	n += 1 + l + sovQuery(uint64(l))
	l = m.NewSecondaryStorePrice.Size()
	n += 1 + l + sovQuery(uint64(l))
	l = m.NewValidatorTaxRate.Size()
	n += 1 + l + sovQuery(uint64(l))
	return n
}

func (m *QueryQuoteUpdateTimeRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.BucketName)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryQuoteUpdateTimeResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.UpdateAt != 0 {
		n += 1 + sovQuery(uint64(m.UpdateAt))
	}
	return n
}

func (m *QueryGroupMembersExistRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.GroupId)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	if len(m.Members) > 0 {
		for _, s := range m.Members {
			l = len(s)
			n += 1 + l + sovQuery(uint64(l))
		}
	}
	return n
}

func (m *QueryGroupMembersExistResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Exists) > 0 {
		for k, v := range m.Exists {
			_ = k
			_ = v
			mapEntrySize := 1 + len(k) + sovQuery(uint64(len(k))) + 1 + 1
			n += mapEntrySize + 1 + sovQuery(uint64(mapEntrySize))
		}
	}
	return n
}

func (m *QueryGroupsExistRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.GroupOwner)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	if len(m.GroupNames) > 0 {
		for _, s := range m.GroupNames {
			l = len(s)
			n += 1 + l + sovQuery(uint64(l))
		}
	}
	return n
}

func (m *QueryGroupsExistByIdRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.GroupIds) > 0 {
		for _, s := range m.GroupIds {
			l = len(s)
			n += 1 + l + sovQuery(uint64(l))
		}
	}
	return n
}

func (m *QueryGroupsExistResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Exists) > 0 {
		for k, v := range m.Exists {
			_ = k
			_ = v
			mapEntrySize := 1 + len(k) + sovQuery(uint64(len(k))) + 1 + 1
			n += mapEntrySize + 1 + sovQuery(uint64(mapEntrySize))
		}
	}
	return n
}

func (m *QueryPaymentAccountBucketFlowRateLimitRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.PaymentAccount)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	l = len(m.BucketOwner)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	l = len(m.BucketName)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryPaymentAccountBucketFlowRateLimitResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.IsSet {
		n += 2
	}
	l = m.FlowRateLimit.Size()
	n += 1 + l + sovQuery(uint64(l))
	return n
}

func sovQuery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuery(x uint64) (n int) {
	return sovQuery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *QueryParamsRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryParamsRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryParamsRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}

func (m *QueryHeadBucketRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryHeadBucketRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryHeadBucketRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BucketName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BucketName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryHeadBucketByIdRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryHeadBucketByIdRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryHeadBucketByIdRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BucketId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BucketId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryHeadBucketResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryHeadBucketResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryHeadBucketResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BucketInfo", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.BucketInfo == nil {
				m.BucketInfo = &BucketInfo{}
			}
			if err := m.BucketInfo.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ExtraInfo", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.ExtraInfo == nil {
				m.ExtraInfo = &BucketExtraInfo{}
			}
			if err := m.ExtraInfo.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryHeadObjectRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryHeadObjectRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryHeadObjectRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BucketName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BucketName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ObjectName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ObjectName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryHeadObjectByIdRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryHeadObjectByIdRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryHeadObjectByIdRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ObjectId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ObjectId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryHeadShadowObjectRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryHeadShadowObjectRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryHeadShadowObjectRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BucketName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BucketName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ObjectName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ObjectName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryHeadObjectResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryHeadObjectResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryHeadObjectResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ObjectInfo", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.ObjectInfo == nil {
				m.ObjectInfo = &ObjectInfo{}
			}
			if err := m.ObjectInfo.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GlobalVirtualGroup", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.GlobalVirtualGroup == nil {
				m.GlobalVirtualGroup = &types.GlobalVirtualGroup{}
			}
			if err := m.GlobalVirtualGroup.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryHeadShadowObjectResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryHeadShadowObjectResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryHeadShadowObjectResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ObjectInfo", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.ObjectInfo == nil {
				m.ObjectInfo = &ShadowObjectInfo{}
			}
			if err := m.ObjectInfo.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryListBucketsRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryListBucketsRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryListBucketsRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pagination", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Pagination == nil {
				m.Pagination = &query.PageRequest{}
			}
			if err := m.Pagination.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryListBucketsResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryListBucketsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryListBucketsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BucketInfos", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BucketInfos = append(m.BucketInfos, &BucketInfo{})
			if err := m.BucketInfos[len(m.BucketInfos)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pagination", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Pagination == nil {
				m.Pagination = &query.PageResponse{}
			}
			if err := m.Pagination.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryListObjectsRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryListObjectsRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryListObjectsRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pagination", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Pagination == nil {
				m.Pagination = &query.PageRequest{}
			}
			if err := m.Pagination.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BucketName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BucketName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryListObjectsByBucketIdRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryListObjectsByBucketIdRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryListObjectsByBucketIdRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pagination", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Pagination == nil {
				m.Pagination = &query.PageRequest{}
			}
			if err := m.Pagination.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BucketId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BucketId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryListObjectsResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryListObjectsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryListObjectsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ObjectInfos", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ObjectInfos = append(m.ObjectInfos, &ObjectInfo{})
			if err := m.ObjectInfos[len(m.ObjectInfos)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pagination", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Pagination == nil {
				m.Pagination = &query.PageResponse{}
			}
			if err := m.Pagination.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryNFTRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryNFTRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryNFTRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TokenId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TokenId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryBucketNFTResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryBucketNFTResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryBucketNFTResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MetaData", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.MetaData == nil {
				m.MetaData = &BucketMetaData{}
			}
			if err := m.MetaData.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryObjectNFTResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryObjectNFTResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryObjectNFTResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MetaData", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.MetaData == nil {
				m.MetaData = &ObjectMetaData{}
			}
			if err := m.MetaData.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryGroupNFTResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryGroupNFTResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryGroupNFTResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MetaData", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.MetaData == nil {
				m.MetaData = &GroupMetaData{}
			}
			if err := m.MetaData.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryPolicyForAccountRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryPolicyForAccountRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryPolicyForAccountRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Resource", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Resource = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PrincipalAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PrincipalAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryPolicyForAccountResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryPolicyForAccountResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryPolicyForAccountResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Policy", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Policy == nil {
				m.Policy = &types1.Policy{}
			}
			if err := m.Policy.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryVerifyPermissionRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryVerifyPermissionRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryVerifyPermissionRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Operator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Operator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BucketName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BucketName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ObjectName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ObjectName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ActionType", wireType)
			}
			m.ActionType = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ActionType |= types1.ActionType(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryVerifyPermissionResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryVerifyPermissionResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryVerifyPermissionResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Effect", wireType)
			}
			m.Effect = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Effect |= types1.Effect(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryHeadGroupRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryHeadGroupRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryHeadGroupRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GroupOwner", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.GroupOwner = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GroupName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.GroupName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryHeadGroupResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryHeadGroupResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryHeadGroupResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GroupInfo", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.GroupInfo == nil {
				m.GroupInfo = &GroupInfo{}
			}
			if err := m.GroupInfo.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryListGroupsRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryListGroupsRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryListGroupsRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pagination", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Pagination == nil {
				m.Pagination = &query.PageRequest{}
			}
			if err := m.Pagination.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GroupOwner", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.GroupOwner = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryListGroupsResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryListGroupsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryListGroupsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pagination", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Pagination == nil {
				m.Pagination = &query.PageResponse{}
			}
			if err := m.Pagination.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GroupInfos", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.GroupInfos = append(m.GroupInfos, &GroupInfo{})
			if err := m.GroupInfos[len(m.GroupInfos)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryHeadGroupMemberRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryHeadGroupMemberRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryHeadGroupMemberRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Member", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Member = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GroupOwner", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.GroupOwner = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GroupName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.GroupName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryHeadGroupMemberResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryHeadGroupMemberResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryHeadGroupMemberResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GroupMember", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.GroupMember == nil {
				m.GroupMember = &types1.GroupMember{}
			}
			if err := m.GroupMember.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryPolicyForGroupRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryPolicyForGroupRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryPolicyForGroupRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Resource", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Resource = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PrincipalGroupId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PrincipalGroupId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryPolicyForGroupResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryPolicyForGroupResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryPolicyForGroupResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Policy", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Policy == nil {
				m.Policy = &types1.Policy{}
			}
			if err := m.Policy.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryPolicyByIdRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryPolicyByIdRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryPolicyByIdRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PolicyId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PolicyId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryPolicyByIdResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryPolicyByIdResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryPolicyByIdResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Policy", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Policy == nil {
				m.Policy = &types1.Policy{}
			}
			if err := m.Policy.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryLockFeeRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryLockFeeRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryLockFeeRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PrimarySpAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PrimarySpAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CreateAt", wireType)
			}
			m.CreateAt = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CreateAt |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PayloadSize", wireType)
			}
			m.PayloadSize = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PayloadSize |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryLockFeeResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryLockFeeResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryLockFeeResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Amount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryHeadBucketExtraRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryHeadBucketExtraRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryHeadBucketExtraRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BucketName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BucketName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryHeadBucketExtraResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryHeadBucketExtraResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryHeadBucketExtraResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ExtraInfo", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.ExtraInfo == nil {
				m.ExtraInfo = &InternalBucketInfo{}
			}
			if err := m.ExtraInfo.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryIsPriceChangedRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryIsPriceChangedRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryIsPriceChangedRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BucketName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BucketName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryIsPriceChangedResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryIsPriceChangedResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryIsPriceChangedResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Changed", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Changed = bool(v != 0)
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CurrentReadPrice", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.CurrentReadPrice.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CurrentPrimaryStorePrice", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.CurrentPrimaryStorePrice.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CurrentSecondaryStorePrice", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.CurrentSecondaryStorePrice.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CurrentValidatorTaxRate", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.CurrentValidatorTaxRate.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NewReadPrice", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.NewReadPrice.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NewPrimaryStorePrice", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.NewPrimaryStorePrice.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NewSecondaryStorePrice", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.NewSecondaryStorePrice.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NewValidatorTaxRate", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.NewValidatorTaxRate.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryQuoteUpdateTimeRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryQuoteUpdateTimeRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryQuoteUpdateTimeRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BucketName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BucketName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryQuoteUpdateTimeResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryQuoteUpdateTimeResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryQuoteUpdateTimeResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field UpdateAt", wireType)
			}
			m.UpdateAt = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.UpdateAt |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryGroupMembersExistRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryGroupMembersExistRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryGroupMembersExistRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GroupId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.GroupId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Members", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Members = append(m.Members, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryGroupMembersExistResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryGroupMembersExistResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryGroupMembersExistResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Exists", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Exists == nil {
				m.Exists = make(map[string]bool)
			}
			var mapkey string
			var mapvalue bool
			for iNdEx < postIndex {
				entryPreIndex := iNdEx
				var wire uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowQuery
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					wire |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				fieldNum := int32(wire >> 3)
				if fieldNum == 1 {
					var stringLenmapkey uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowQuery
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						stringLenmapkey |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intStringLenmapkey := int(stringLenmapkey)
					if intStringLenmapkey < 0 {
						return ErrInvalidLengthQuery
					}
					postStringIndexmapkey := iNdEx + intStringLenmapkey
					if postStringIndexmapkey < 0 {
						return ErrInvalidLengthQuery
					}
					if postStringIndexmapkey > l {
						return io.ErrUnexpectedEOF
					}
					mapkey = string(dAtA[iNdEx:postStringIndexmapkey])
					iNdEx = postStringIndexmapkey
				} else if fieldNum == 2 {
					var mapvaluetemp int
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowQuery
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						mapvaluetemp |= int(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					mapvalue = bool(mapvaluetemp != 0)
				} else {
					iNdEx = entryPreIndex
					skippy, err := skipQuery(dAtA[iNdEx:])
					if err != nil {
						return err
					}
					if (skippy < 0) || (iNdEx+skippy) < 0 {
						return ErrInvalidLengthQuery
					}
					if (iNdEx + skippy) > postIndex {
						return io.ErrUnexpectedEOF
					}
					iNdEx += skippy
				}
			}
			m.Exists[mapkey] = mapvalue
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryGroupsExistRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryGroupsExistRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryGroupsExistRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GroupOwner", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.GroupOwner = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GroupNames", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.GroupNames = append(m.GroupNames, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryGroupsExistByIdRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryGroupsExistByIdRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryGroupsExistByIdRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GroupIds", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.GroupIds = append(m.GroupIds, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryGroupsExistResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryGroupsExistResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryGroupsExistResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Exists", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Exists == nil {
				m.Exists = make(map[string]bool)
			}
			var mapkey string
			var mapvalue bool
			for iNdEx < postIndex {
				entryPreIndex := iNdEx
				var wire uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowQuery
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					wire |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				fieldNum := int32(wire >> 3)
				if fieldNum == 1 {
					var stringLenmapkey uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowQuery
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						stringLenmapkey |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intStringLenmapkey := int(stringLenmapkey)
					if intStringLenmapkey < 0 {
						return ErrInvalidLengthQuery
					}
					postStringIndexmapkey := iNdEx + intStringLenmapkey
					if postStringIndexmapkey < 0 {
						return ErrInvalidLengthQuery
					}
					if postStringIndexmapkey > l {
						return io.ErrUnexpectedEOF
					}
					mapkey = string(dAtA[iNdEx:postStringIndexmapkey])
					iNdEx = postStringIndexmapkey
				} else if fieldNum == 2 {
					var mapvaluetemp int
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowQuery
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						mapvaluetemp |= int(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					mapvalue = bool(mapvaluetemp != 0)
				} else {
					iNdEx = entryPreIndex
					skippy, err := skipQuery(dAtA[iNdEx:])
					if err != nil {
						return err
					}
					if (skippy < 0) || (iNdEx+skippy) < 0 {
						return ErrInvalidLengthQuery
					}
					if (iNdEx + skippy) > postIndex {
						return io.ErrUnexpectedEOF
					}
					iNdEx += skippy
				}
			}
			m.Exists[mapkey] = mapvalue
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryPaymentAccountBucketFlowRateLimitRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryPaymentAccountBucketFlowRateLimitRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryPaymentAccountBucketFlowRateLimitRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PaymentAccount", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PaymentAccount = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BucketOwner", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BucketOwner = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BucketName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BucketName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryPaymentAccountBucketFlowRateLimitResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryPaymentAccountBucketFlowRateLimitResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryPaymentAccountBucketFlowRateLimitResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field IsSet", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.IsSet = bool(v != 0)
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FlowRateLimit", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.FlowRateLimit.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipQuery(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthQuery
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupQuery
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthQuery
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthQuery        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowQuery          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupQuery = fmt.Errorf("proto: unexpected end of group")
)