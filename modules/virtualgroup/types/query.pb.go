package types

import (
	context "context"
	fmt "fmt"
	query "github.com/cosmos/cosmos-sdk/types/query"
	_ "github.com/cosmos/gogoproto/gogoproto"
	grpc1 "github.com/cosmos/gogoproto/grpc"
	proto "github.com/cosmos/gogoproto/proto"
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

type QueryGlobalVirtualGroupRequest struct {
	GlobalVirtualGroupId uint32 `protobuf:"varint,1,opt,name=global_virtual_group_id,json=globalVirtualGroupId,proto3" json:"global_virtual_group_id,omitempty"`
}

func (m *QueryGlobalVirtualGroupRequest) Reset()         { *m = QueryGlobalVirtualGroupRequest{} }
func (m *QueryGlobalVirtualGroupRequest) String() string { return proto.CompactTextString(m) }
func (*QueryGlobalVirtualGroupRequest) ProtoMessage()    {}
func (*QueryGlobalVirtualGroupRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_83cd53fc415e00e7, []int{2}
}
func (m *QueryGlobalVirtualGroupRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryGlobalVirtualGroupRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryGlobalVirtualGroupRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryGlobalVirtualGroupRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryGlobalVirtualGroupRequest.Merge(m, src)
}
func (m *QueryGlobalVirtualGroupRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryGlobalVirtualGroupRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryGlobalVirtualGroupRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryGlobalVirtualGroupRequest proto.InternalMessageInfo

func (m *QueryGlobalVirtualGroupRequest) GetGlobalVirtualGroupId() uint32 {
	if m != nil {
		return m.GlobalVirtualGroupId
	}
	return 0
}

type QueryGlobalVirtualGroupResponse struct {
	GlobalVirtualGroup *GlobalVirtualGroup `protobuf:"bytes,1,opt,name=global_virtual_group,json=globalVirtualGroup,proto3" json:"global_virtual_group,omitempty"`
}

func (m *QueryGlobalVirtualGroupResponse) Reset()         { *m = QueryGlobalVirtualGroupResponse{} }
func (m *QueryGlobalVirtualGroupResponse) String() string { return proto.CompactTextString(m) }
func (*QueryGlobalVirtualGroupResponse) ProtoMessage()    {}
func (*QueryGlobalVirtualGroupResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_83cd53fc415e00e7, []int{3}
}
func (m *QueryGlobalVirtualGroupResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryGlobalVirtualGroupResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryGlobalVirtualGroupResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryGlobalVirtualGroupResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryGlobalVirtualGroupResponse.Merge(m, src)
}
func (m *QueryGlobalVirtualGroupResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryGlobalVirtualGroupResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryGlobalVirtualGroupResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryGlobalVirtualGroupResponse proto.InternalMessageInfo

func (m *QueryGlobalVirtualGroupResponse) GetGlobalVirtualGroup() *GlobalVirtualGroup {
	if m != nil {
		return m.GlobalVirtualGroup
	}
	return nil
}

type QueryGlobalVirtualGroupByFamilyIDRequest struct {
	GlobalVirtualGroupFamilyId uint32 `protobuf:"varint,1,opt,name=global_virtual_group_family_id,json=globalVirtualGroupFamilyId,proto3" json:"global_virtual_group_family_id,omitempty"`
}

func (m *QueryGlobalVirtualGroupByFamilyIDRequest) Reset() {
	*m = QueryGlobalVirtualGroupByFamilyIDRequest{}
}
func (m *QueryGlobalVirtualGroupByFamilyIDRequest) String() string { return proto.CompactTextString(m) }
func (*QueryGlobalVirtualGroupByFamilyIDRequest) ProtoMessage()    {}
func (*QueryGlobalVirtualGroupByFamilyIDRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_83cd53fc415e00e7, []int{4}
}
func (m *QueryGlobalVirtualGroupByFamilyIDRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryGlobalVirtualGroupByFamilyIDRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryGlobalVirtualGroupByFamilyIDRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryGlobalVirtualGroupByFamilyIDRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryGlobalVirtualGroupByFamilyIDRequest.Merge(m, src)
}
func (m *QueryGlobalVirtualGroupByFamilyIDRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryGlobalVirtualGroupByFamilyIDRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryGlobalVirtualGroupByFamilyIDRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryGlobalVirtualGroupByFamilyIDRequest proto.InternalMessageInfo

func (m *QueryGlobalVirtualGroupByFamilyIDRequest) GetGlobalVirtualGroupFamilyId() uint32 {
	if m != nil {
		return m.GlobalVirtualGroupFamilyId
	}
	return 0
}

type QueryGlobalVirtualGroupByFamilyIDResponse struct {
	GlobalVirtualGroups []*GlobalVirtualGroup `protobuf:"bytes,1,rep,name=global_virtual_groups,json=globalVirtualGroups,proto3" json:"global_virtual_groups,omitempty"`
}

func (m *QueryGlobalVirtualGroupByFamilyIDResponse) Reset() {
	*m = QueryGlobalVirtualGroupByFamilyIDResponse{}
}
func (m *QueryGlobalVirtualGroupByFamilyIDResponse) String() string {
	return proto.CompactTextString(m)
}
func (*QueryGlobalVirtualGroupByFamilyIDResponse) ProtoMessage() {}
func (*QueryGlobalVirtualGroupByFamilyIDResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_83cd53fc415e00e7, []int{5}
}
func (m *QueryGlobalVirtualGroupByFamilyIDResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryGlobalVirtualGroupByFamilyIDResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryGlobalVirtualGroupByFamilyIDResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryGlobalVirtualGroupByFamilyIDResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryGlobalVirtualGroupByFamilyIDResponse.Merge(m, src)
}
func (m *QueryGlobalVirtualGroupByFamilyIDResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryGlobalVirtualGroupByFamilyIDResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryGlobalVirtualGroupByFamilyIDResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryGlobalVirtualGroupByFamilyIDResponse proto.InternalMessageInfo

func (m *QueryGlobalVirtualGroupByFamilyIDResponse) GetGlobalVirtualGroups() []*GlobalVirtualGroup {
	if m != nil {
		return m.GlobalVirtualGroups
	}
	return nil
}

type QueryGlobalVirtualGroupFamilyRequest struct {
	FamilyId uint32 `protobuf:"varint,1,opt,name=family_id,json=familyId,proto3" json:"family_id,omitempty"`
}

func (m *QueryGlobalVirtualGroupFamilyRequest) Reset()         { *m = QueryGlobalVirtualGroupFamilyRequest{} }
func (m *QueryGlobalVirtualGroupFamilyRequest) String() string { return proto.CompactTextString(m) }
func (*QueryGlobalVirtualGroupFamilyRequest) ProtoMessage()    {}
func (*QueryGlobalVirtualGroupFamilyRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_83cd53fc415e00e7, []int{6}
}
func (m *QueryGlobalVirtualGroupFamilyRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryGlobalVirtualGroupFamilyRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryGlobalVirtualGroupFamilyRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryGlobalVirtualGroupFamilyRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryGlobalVirtualGroupFamilyRequest.Merge(m, src)
}
func (m *QueryGlobalVirtualGroupFamilyRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryGlobalVirtualGroupFamilyRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryGlobalVirtualGroupFamilyRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryGlobalVirtualGroupFamilyRequest proto.InternalMessageInfo

func (m *QueryGlobalVirtualGroupFamilyRequest) GetFamilyId() uint32 {
	if m != nil {
		return m.FamilyId
	}
	return 0
}

type QueryGlobalVirtualGroupFamilyResponse struct {
	GlobalVirtualGroupFamily *GlobalVirtualGroupFamily `protobuf:"bytes,1,opt,name=global_virtual_group_family,json=globalVirtualGroupFamily,proto3" json:"global_virtual_group_family,omitempty"`
}

func (m *QueryGlobalVirtualGroupFamilyResponse) Reset()         { *m = QueryGlobalVirtualGroupFamilyResponse{} }
func (m *QueryGlobalVirtualGroupFamilyResponse) String() string { return proto.CompactTextString(m) }
func (*QueryGlobalVirtualGroupFamilyResponse) ProtoMessage()    {}
func (*QueryGlobalVirtualGroupFamilyResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_83cd53fc415e00e7, []int{7}
}
func (m *QueryGlobalVirtualGroupFamilyResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryGlobalVirtualGroupFamilyResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryGlobalVirtualGroupFamilyResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryGlobalVirtualGroupFamilyResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryGlobalVirtualGroupFamilyResponse.Merge(m, src)
}
func (m *QueryGlobalVirtualGroupFamilyResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryGlobalVirtualGroupFamilyResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryGlobalVirtualGroupFamilyResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryGlobalVirtualGroupFamilyResponse proto.InternalMessageInfo

func (m *QueryGlobalVirtualGroupFamilyResponse) GetGlobalVirtualGroupFamily() *GlobalVirtualGroupFamily {
	if m != nil {
		return m.GlobalVirtualGroupFamily
	}
	return nil
}

type QueryGlobalVirtualGroupFamiliesRequest struct {
	Pagination *query.PageRequest `protobuf:"bytes,1,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (m *QueryGlobalVirtualGroupFamiliesRequest) Reset() {
	*m = QueryGlobalVirtualGroupFamiliesRequest{}
}
func (m *QueryGlobalVirtualGroupFamiliesRequest) String() string { return proto.CompactTextString(m) }
func (*QueryGlobalVirtualGroupFamiliesRequest) ProtoMessage()    {}
func (*QueryGlobalVirtualGroupFamiliesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_83cd53fc415e00e7, []int{8}
}
func (m *QueryGlobalVirtualGroupFamiliesRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryGlobalVirtualGroupFamiliesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryGlobalVirtualGroupFamiliesRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryGlobalVirtualGroupFamiliesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryGlobalVirtualGroupFamiliesRequest.Merge(m, src)
}
func (m *QueryGlobalVirtualGroupFamiliesRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryGlobalVirtualGroupFamiliesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryGlobalVirtualGroupFamiliesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryGlobalVirtualGroupFamiliesRequest proto.InternalMessageInfo

func (m *QueryGlobalVirtualGroupFamiliesRequest) GetPagination() *query.PageRequest {
	if m != nil {
		return m.Pagination
	}
	return nil
}

type QueryGlobalVirtualGroupFamiliesResponse struct {
	GvgFamilies []*GlobalVirtualGroupFamily `protobuf:"bytes,1,rep,name=gvg_families,json=gvgFamilies,proto3" json:"gvg_families,omitempty"`
	Pagination  *query.PageResponse         `protobuf:"bytes,2,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (m *QueryGlobalVirtualGroupFamiliesResponse) Reset() {
	*m = QueryGlobalVirtualGroupFamiliesResponse{}
}
func (m *QueryGlobalVirtualGroupFamiliesResponse) String() string { return proto.CompactTextString(m) }
func (*QueryGlobalVirtualGroupFamiliesResponse) ProtoMessage()    {}
func (*QueryGlobalVirtualGroupFamiliesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_83cd53fc415e00e7, []int{9}
}
func (m *QueryGlobalVirtualGroupFamiliesResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryGlobalVirtualGroupFamiliesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryGlobalVirtualGroupFamiliesResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryGlobalVirtualGroupFamiliesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryGlobalVirtualGroupFamiliesResponse.Merge(m, src)
}
func (m *QueryGlobalVirtualGroupFamiliesResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryGlobalVirtualGroupFamiliesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryGlobalVirtualGroupFamiliesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryGlobalVirtualGroupFamiliesResponse proto.InternalMessageInfo

func (m *QueryGlobalVirtualGroupFamiliesResponse) GetGvgFamilies() []*GlobalVirtualGroupFamily {
	if m != nil {
		return m.GvgFamilies
	}
	return nil
}

func (m *QueryGlobalVirtualGroupFamiliesResponse) GetPagination() *query.PageResponse {
	if m != nil {
		return m.Pagination
	}
	return nil
}

type AvailableGlobalVirtualGroupFamiliesRequest struct {
	GlobalVirtualGroupFamilyIds []uint32 `protobuf:"varint,1,rep,packed,name=global_virtual_group_family_ids,json=globalVirtualGroupFamilyIds,proto3" json:"global_virtual_group_family_ids,omitempty"`
}

func (m *AvailableGlobalVirtualGroupFamiliesRequest) Reset() {
	*m = AvailableGlobalVirtualGroupFamiliesRequest{}
}
func (m *AvailableGlobalVirtualGroupFamiliesRequest) String() string {
	return proto.CompactTextString(m)
}
func (*AvailableGlobalVirtualGroupFamiliesRequest) ProtoMessage() {}
func (*AvailableGlobalVirtualGroupFamiliesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_83cd53fc415e00e7, []int{10}
}
func (m *AvailableGlobalVirtualGroupFamiliesRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *AvailableGlobalVirtualGroupFamiliesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AvailableGlobalVirtualGroupFamiliesRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *AvailableGlobalVirtualGroupFamiliesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AvailableGlobalVirtualGroupFamiliesRequest.Merge(m, src)
}
func (m *AvailableGlobalVirtualGroupFamiliesRequest) XXX_Size() int {
	return m.Size()
}
func (m *AvailableGlobalVirtualGroupFamiliesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AvailableGlobalVirtualGroupFamiliesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AvailableGlobalVirtualGroupFamiliesRequest proto.InternalMessageInfo

func (m *AvailableGlobalVirtualGroupFamiliesRequest) GetGlobalVirtualGroupFamilyIds() []uint32 {
	if m != nil {
		return m.GlobalVirtualGroupFamilyIds
	}
	return nil
}

type AvailableGlobalVirtualGroupFamiliesResponse struct {
	GlobalVirtualGroupFamilyIds []uint32 `protobuf:"varint,1,rep,packed,name=global_virtual_group_family_ids,json=globalVirtualGroupFamilyIds,proto3" json:"global_virtual_group_family_ids,omitempty"`
}

func (m *AvailableGlobalVirtualGroupFamiliesResponse) Reset() {
	*m = AvailableGlobalVirtualGroupFamiliesResponse{}
}
func (m *AvailableGlobalVirtualGroupFamiliesResponse) String() string {
	return proto.CompactTextString(m)
}
func (*AvailableGlobalVirtualGroupFamiliesResponse) ProtoMessage() {}
func (*AvailableGlobalVirtualGroupFamiliesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_83cd53fc415e00e7, []int{11}
}
func (m *AvailableGlobalVirtualGroupFamiliesResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *AvailableGlobalVirtualGroupFamiliesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AvailableGlobalVirtualGroupFamiliesResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *AvailableGlobalVirtualGroupFamiliesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AvailableGlobalVirtualGroupFamiliesResponse.Merge(m, src)
}
func (m *AvailableGlobalVirtualGroupFamiliesResponse) XXX_Size() int {
	return m.Size()
}
func (m *AvailableGlobalVirtualGroupFamiliesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AvailableGlobalVirtualGroupFamiliesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AvailableGlobalVirtualGroupFamiliesResponse proto.InternalMessageInfo

func (m *AvailableGlobalVirtualGroupFamiliesResponse) GetGlobalVirtualGroupFamilyIds() []uint32 {
	if m != nil {
		return m.GlobalVirtualGroupFamilyIds
	}
	return nil
}

type QuerySwapInInfoRequest struct {
	GlobalVirtualGroupFamilyId uint32 `protobuf:"varint,1,opt,name=global_virtual_group_family_id,json=globalVirtualGroupFamilyId,proto3" json:"global_virtual_group_family_id,omitempty"`
	GlobalVirtualGroupId       uint32 `protobuf:"varint,2,opt,name=global_virtual_group_id,json=globalVirtualGroupId,proto3" json:"global_virtual_group_id,omitempty"`
}

func (m *QuerySwapInInfoRequest) Reset()         { *m = QuerySwapInInfoRequest{} }
func (m *QuerySwapInInfoRequest) String() string { return proto.CompactTextString(m) }
func (*QuerySwapInInfoRequest) ProtoMessage()    {}
func (*QuerySwapInInfoRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_83cd53fc415e00e7, []int{12}
}
func (m *QuerySwapInInfoRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QuerySwapInInfoRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QuerySwapInInfoRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QuerySwapInInfoRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QuerySwapInInfoRequest.Merge(m, src)
}
func (m *QuerySwapInInfoRequest) XXX_Size() int {
	return m.Size()
}
func (m *QuerySwapInInfoRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QuerySwapInInfoRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QuerySwapInInfoRequest proto.InternalMessageInfo

func (m *QuerySwapInInfoRequest) GetGlobalVirtualGroupFamilyId() uint32 {
	if m != nil {
		return m.GlobalVirtualGroupFamilyId
	}
	return 0
}

func (m *QuerySwapInInfoRequest) GetGlobalVirtualGroupId() uint32 {
	if m != nil {
		return m.GlobalVirtualGroupId
	}
	return 0
}

type QuerySwapInInfoResponse struct {
	SwapInInfo *SwapInInfo `protobuf:"bytes,1,opt,name=swap_in_info,json=swapInInfo,proto3" json:"swap_in_info,omitempty"`
}

func (m *QuerySwapInInfoResponse) Reset()         { *m = QuerySwapInInfoResponse{} }
func (m *QuerySwapInInfoResponse) String() string { return proto.CompactTextString(m) }
func (*QuerySwapInInfoResponse) ProtoMessage()    {}
func (*QuerySwapInInfoResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_83cd53fc415e00e7, []int{13}
}
func (m *QuerySwapInInfoResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QuerySwapInInfoResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QuerySwapInInfoResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QuerySwapInInfoResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QuerySwapInInfoResponse.Merge(m, src)
}
func (m *QuerySwapInInfoResponse) XXX_Size() int {
	return m.Size()
}
func (m *QuerySwapInInfoResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QuerySwapInInfoResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QuerySwapInInfoResponse proto.InternalMessageInfo

func (m *QuerySwapInInfoResponse) GetSwapInInfo() *SwapInInfo {
	if m != nil {
		return m.SwapInInfo
	}
	return nil
}

type QuerySPGVGStatisticsRequest struct {
	SpId uint32 `protobuf:"varint,1,opt,name=sp_id,json=spId,proto3" json:"sp_id,omitempty"`
}

func (m *QuerySPGVGStatisticsRequest) Reset()         { *m = QuerySPGVGStatisticsRequest{} }
func (m *QuerySPGVGStatisticsRequest) String() string { return proto.CompactTextString(m) }
func (*QuerySPGVGStatisticsRequest) ProtoMessage()    {}
func (*QuerySPGVGStatisticsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_83cd53fc415e00e7, []int{14}
}
func (m *QuerySPGVGStatisticsRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QuerySPGVGStatisticsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QuerySPGVGStatisticsRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QuerySPGVGStatisticsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QuerySPGVGStatisticsRequest.Merge(m, src)
}
func (m *QuerySPGVGStatisticsRequest) XXX_Size() int {
	return m.Size()
}
func (m *QuerySPGVGStatisticsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QuerySPGVGStatisticsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QuerySPGVGStatisticsRequest proto.InternalMessageInfo

func (m *QuerySPGVGStatisticsRequest) GetSpId() uint32 {
	if m != nil {
		return m.SpId
	}
	return 0
}

type QuerySPGVGStatisticsResponse struct {
	GvgStats *GVGStatisticsWithinSP `protobuf:"bytes,1,opt,name=gvg_stats,json=gvgStats,proto3" json:"gvg_stats,omitempty"`
}

func (m *QuerySPGVGStatisticsResponse) Reset()         { *m = QuerySPGVGStatisticsResponse{} }
func (m *QuerySPGVGStatisticsResponse) String() string { return proto.CompactTextString(m) }
func (*QuerySPGVGStatisticsResponse) ProtoMessage()    {}
func (*QuerySPGVGStatisticsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_83cd53fc415e00e7, []int{15}
}
func (m *QuerySPGVGStatisticsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QuerySPGVGStatisticsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QuerySPGVGStatisticsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QuerySPGVGStatisticsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QuerySPGVGStatisticsResponse.Merge(m, src)
}
func (m *QuerySPGVGStatisticsResponse) XXX_Size() int {
	return m.Size()
}
func (m *QuerySPGVGStatisticsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QuerySPGVGStatisticsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QuerySPGVGStatisticsResponse proto.InternalMessageInfo

func (m *QuerySPGVGStatisticsResponse) GetGvgStats() *GVGStatisticsWithinSP {
	if m != nil {
		return m.GvgStats
	}
	return nil
}

type QuerySPAvailableGlobalVirtualGroupFamiliesRequest struct {
	SpId uint32 `protobuf:"varint,1,opt,name=sp_id,json=spId,proto3" json:"sp_id,omitempty"`
}

func (m *QuerySPAvailableGlobalVirtualGroupFamiliesRequest) Reset() {
	*m = QuerySPAvailableGlobalVirtualGroupFamiliesRequest{}
}
func (m *QuerySPAvailableGlobalVirtualGroupFamiliesRequest) String() string {
	return proto.CompactTextString(m)
}
func (*QuerySPAvailableGlobalVirtualGroupFamiliesRequest) ProtoMessage() {}
func (*QuerySPAvailableGlobalVirtualGroupFamiliesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_83cd53fc415e00e7, []int{16}
}
func (m *QuerySPAvailableGlobalVirtualGroupFamiliesRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QuerySPAvailableGlobalVirtualGroupFamiliesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QuerySPAvailableGlobalVirtualGroupFamiliesRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QuerySPAvailableGlobalVirtualGroupFamiliesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QuerySPAvailableGlobalVirtualGroupFamiliesRequest.Merge(m, src)
}
func (m *QuerySPAvailableGlobalVirtualGroupFamiliesRequest) XXX_Size() int {
	return m.Size()
}
func (m *QuerySPAvailableGlobalVirtualGroupFamiliesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QuerySPAvailableGlobalVirtualGroupFamiliesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QuerySPAvailableGlobalVirtualGroupFamiliesRequest proto.InternalMessageInfo

func (m *QuerySPAvailableGlobalVirtualGroupFamiliesRequest) GetSpId() uint32 {
	if m != nil {
		return m.SpId
	}
	return 0
}

type QuerySPAvailableGlobalVirtualGroupFamiliesResponse struct {
	GlobalVirtualGroupFamilyIds []uint32 `protobuf:"varint,1,rep,packed,name=global_virtual_group_family_ids,json=globalVirtualGroupFamilyIds,proto3" json:"global_virtual_group_family_ids,omitempty"`
}

func (m *QuerySPAvailableGlobalVirtualGroupFamiliesResponse) Reset() {
	*m = QuerySPAvailableGlobalVirtualGroupFamiliesResponse{}
}
func (m *QuerySPAvailableGlobalVirtualGroupFamiliesResponse) String() string {
	return proto.CompactTextString(m)
}
func (*QuerySPAvailableGlobalVirtualGroupFamiliesResponse) ProtoMessage() {}
func (*QuerySPAvailableGlobalVirtualGroupFamiliesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_83cd53fc415e00e7, []int{17}
}
func (m *QuerySPAvailableGlobalVirtualGroupFamiliesResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QuerySPAvailableGlobalVirtualGroupFamiliesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QuerySPAvailableGlobalVirtualGroupFamiliesResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QuerySPAvailableGlobalVirtualGroupFamiliesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QuerySPAvailableGlobalVirtualGroupFamiliesResponse.Merge(m, src)
}
func (m *QuerySPAvailableGlobalVirtualGroupFamiliesResponse) XXX_Size() int {
	return m.Size()
}
func (m *QuerySPAvailableGlobalVirtualGroupFamiliesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QuerySPAvailableGlobalVirtualGroupFamiliesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QuerySPAvailableGlobalVirtualGroupFamiliesResponse proto.InternalMessageInfo

func (m *QuerySPAvailableGlobalVirtualGroupFamiliesResponse) GetGlobalVirtualGroupFamilyIds() []uint32 {
	if m != nil {
		return m.GlobalVirtualGroupFamilyIds
	}
	return nil
}

type QuerySpOptimalGlobalVirtualGroupFamilyRequest struct {
	SpId            uint32          `protobuf:"varint,1,opt,name=sp_id,json=spId,proto3" json:"sp_id,omitempty"`
	PickVgfStrategy PickVGFStrategy `protobuf:"varint,2,opt,name=pick_vgf_strategy,json=pickVgfStrategy,proto3,enum=greenfield.virtualgroup.PickVGFStrategy" json:"pick_vgf_strategy,omitempty"`
}

func (m *QuerySpOptimalGlobalVirtualGroupFamilyRequest) Reset() {
	*m = QuerySpOptimalGlobalVirtualGroupFamilyRequest{}
}
func (m *QuerySpOptimalGlobalVirtualGroupFamilyRequest) String() string {
	return proto.CompactTextString(m)
}
func (*QuerySpOptimalGlobalVirtualGroupFamilyRequest) ProtoMessage() {}
func (*QuerySpOptimalGlobalVirtualGroupFamilyRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_83cd53fc415e00e7, []int{18}
}
func (m *QuerySpOptimalGlobalVirtualGroupFamilyRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QuerySpOptimalGlobalVirtualGroupFamilyRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QuerySpOptimalGlobalVirtualGroupFamilyRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QuerySpOptimalGlobalVirtualGroupFamilyRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QuerySpOptimalGlobalVirtualGroupFamilyRequest.Merge(m, src)
}
func (m *QuerySpOptimalGlobalVirtualGroupFamilyRequest) XXX_Size() int {
	return m.Size()
}
func (m *QuerySpOptimalGlobalVirtualGroupFamilyRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QuerySpOptimalGlobalVirtualGroupFamilyRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QuerySpOptimalGlobalVirtualGroupFamilyRequest proto.InternalMessageInfo

func (m *QuerySpOptimalGlobalVirtualGroupFamilyRequest) GetSpId() uint32 {
	if m != nil {
		return m.SpId
	}
	return 0
}

func (m *QuerySpOptimalGlobalVirtualGroupFamilyRequest) GetPickVgfStrategy() PickVGFStrategy {
	if m != nil {
		return m.PickVgfStrategy
	}
	return Strategy_Maximize_Free_Store_Size
}

type QuerySpOptimalGlobalVirtualGroupFamilyResponse struct {
	GlobalVirtualGroupFamilyId uint32 `protobuf:"varint,1,opt,name=global_virtual_group_family_id,json=globalVirtualGroupFamilyId,proto3" json:"global_virtual_group_family_id,omitempty"`
}

func (m *QuerySpOptimalGlobalVirtualGroupFamilyResponse) Reset() {
	*m = QuerySpOptimalGlobalVirtualGroupFamilyResponse{}
}
func (m *QuerySpOptimalGlobalVirtualGroupFamilyResponse) String() string {
	return proto.CompactTextString(m)
}
func (*QuerySpOptimalGlobalVirtualGroupFamilyResponse) ProtoMessage() {}
func (*QuerySpOptimalGlobalVirtualGroupFamilyResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_83cd53fc415e00e7, []int{19}
}
func (m *QuerySpOptimalGlobalVirtualGroupFamilyResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QuerySpOptimalGlobalVirtualGroupFamilyResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QuerySpOptimalGlobalVirtualGroupFamilyResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QuerySpOptimalGlobalVirtualGroupFamilyResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QuerySpOptimalGlobalVirtualGroupFamilyResponse.Merge(m, src)
}
func (m *QuerySpOptimalGlobalVirtualGroupFamilyResponse) XXX_Size() int {
	return m.Size()
}
func (m *QuerySpOptimalGlobalVirtualGroupFamilyResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QuerySpOptimalGlobalVirtualGroupFamilyResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QuerySpOptimalGlobalVirtualGroupFamilyResponse proto.InternalMessageInfo

func (m *QuerySpOptimalGlobalVirtualGroupFamilyResponse) GetGlobalVirtualGroupFamilyId() uint32 {
	if m != nil {
		return m.GlobalVirtualGroupFamilyId
	}
	return 0
}

func init() {
	proto.RegisterType((*QueryGlobalVirtualGroupRequest)(nil), "greenfield.virtualgroup.QueryGlobalVirtualGroupRequest")
	proto.RegisterType((*QueryGlobalVirtualGroupResponse)(nil), "greenfield.virtualgroup.QueryGlobalVirtualGroupResponse")
	proto.RegisterType((*QueryGlobalVirtualGroupByFamilyIDRequest)(nil), "greenfield.virtualgroup.QueryGlobalVirtualGroupByFamilyIDRequest")
	proto.RegisterType((*QueryGlobalVirtualGroupByFamilyIDResponse)(nil), "greenfield.virtualgroup.QueryGlobalVirtualGroupByFamilyIDResponse")
	proto.RegisterType((*QueryGlobalVirtualGroupFamilyRequest)(nil), "greenfield.virtualgroup.QueryGlobalVirtualGroupFamilyRequest")
	proto.RegisterType((*QueryGlobalVirtualGroupFamilyResponse)(nil), "greenfield.virtualgroup.QueryGlobalVirtualGroupFamilyResponse")
	proto.RegisterType((*QueryGlobalVirtualGroupFamiliesRequest)(nil), "greenfield.virtualgroup.QueryGlobalVirtualGroupFamiliesRequest")
	proto.RegisterType((*QueryGlobalVirtualGroupFamiliesResponse)(nil), "greenfield.virtualgroup.QueryGlobalVirtualGroupFamiliesResponse")
	proto.RegisterType((*AvailableGlobalVirtualGroupFamiliesRequest)(nil), "greenfield.virtualgroup.AvailableGlobalVirtualGroupFamiliesRequest")
	proto.RegisterType((*AvailableGlobalVirtualGroupFamiliesResponse)(nil), "greenfield.virtualgroup.AvailableGlobalVirtualGroupFamiliesResponse")
	proto.RegisterType((*QuerySwapInInfoRequest)(nil), "greenfield.virtualgroup.QuerySwapInInfoRequest")
	proto.RegisterType((*QuerySwapInInfoResponse)(nil), "greenfield.virtualgroup.QuerySwapInInfoResponse")
	proto.RegisterType((*QuerySPGVGStatisticsRequest)(nil), "greenfield.virtualgroup.QuerySPGVGStatisticsRequest")
	proto.RegisterType((*QuerySPGVGStatisticsResponse)(nil), "greenfield.virtualgroup.QuerySPGVGStatisticsResponse")
	proto.RegisterType((*QuerySPAvailableGlobalVirtualGroupFamiliesRequest)(nil), "greenfield.virtualgroup.QuerySPAvailableGlobalVirtualGroupFamiliesRequest")
	proto.RegisterType((*QuerySPAvailableGlobalVirtualGroupFamiliesResponse)(nil), "greenfield.virtualgroup.QuerySPAvailableGlobalVirtualGroupFamiliesResponse")
	proto.RegisterType((*QuerySpOptimalGlobalVirtualGroupFamilyRequest)(nil), "greenfield.virtualgroup.QuerySpOptimalGlobalVirtualGroupFamilyRequest")
	proto.RegisterType((*QuerySpOptimalGlobalVirtualGroupFamilyResponse)(nil), "greenfield.virtualgroup.QuerySpOptimalGlobalVirtualGroupFamilyResponse")
}

func init() {
	proto.RegisterFile("greenfield/virtualgroup/query.proto", fileDescriptor_83cd53fc415e00e7)
}

var fileDescriptor_83cd53fc415e00e7 = []byte{
	// 1119 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x98, 0xcf, 0x73, 0xdb, 0x44,
	0x14, 0xc7, 0xa3, 0xd0, 0x66, 0xda, 0xd7, 0x86, 0x1f, 0x9b, 0x40, 0x32, 0x4a, 0xc6, 0x01, 0xa5,
	0x6d, 0x42, 0xda, 0x48, 0xd8, 0x34, 0xa5, 0xc3, 0xf4, 0x57, 0x9c, 0x34, 0xae, 0xc9, 0x81, 0xe0,
	0x64, 0xd2, 0x19, 0x66, 0x18, 0xb1, 0x76, 0xe4, 0xad, 0x26, 0xb6, 0x56, 0xf5, 0xca, 0x2e, 0xe6,
	0xc4, 0xf4, 0xdc, 0x03, 0x43, 0x4f, 0x70, 0xe0, 0x6f, 0xe0, 0x2f, 0xe0, 0x9c, 0x63, 0x67, 0x38,
	0xc0, 0x09, 0x98, 0x84, 0x0b, 0xfc, 0x09, 0x9c, 0x18, 0xad, 0x56, 0xb1, 0x8d, 0xb5, 0xb2, 0xe4,
	0xf8, 0x92, 0xf1, 0xc8, 0xfb, 0xbe, 0xef, 0xfb, 0x79, 0xfb, 0xf4, 0x76, 0x63, 0x58, 0x24, 0x0d,
	0xcb, 0x72, 0xaa, 0xb6, 0x55, 0x3b, 0x30, 0x5a, 0x76, 0xc3, 0x6b, 0xe2, 0x1a, 0x69, 0xd0, 0xa6,
	0x6b, 0x3c, 0x6d, 0x5a, 0x8d, 0xb6, 0xee, 0x36, 0xa8, 0x47, 0xd1, 0x4c, 0x67, 0x91, 0xde, 0xbd,
	0x48, 0x5d, 0xa9, 0x50, 0x56, 0xa7, 0xcc, 0x28, 0x63, 0x66, 0x05, 0x11, 0x46, 0x2b, 0x5b, 0xb6,
	0x3c, 0x9c, 0x35, 0x5c, 0x4c, 0x6c, 0x07, 0x7b, 0x36, 0x75, 0x02, 0x11, 0x75, 0x9a, 0x50, 0x42,
	0xf9, 0x47, 0xc3, 0xff, 0x24, 0x9e, 0xce, 0x13, 0x4a, 0x49, 0xcd, 0x32, 0xb0, 0x6b, 0x1b, 0xd8,
	0x71, 0xa8, 0xc7, 0x43, 0x98, 0xf8, 0xf6, 0x8a, 0xcc, 0x5d, 0x85, 0xd6, 0xeb, 0xa7, 0xca, 0xd2,
	0x55, 0x2e, 0x6e, 0xe0, 0x7a, 0xa8, 0x25, 0x25, 0xf5, 0xda, 0xae, 0x25, 0x16, 0x69, 0xd3, 0x80,
	0x3e, 0xf3, 0x31, 0x76, 0x78, 0x64, 0xc9, 0x7a, 0xda, 0xb4, 0x98, 0xa7, 0xed, 0xc1, 0x54, 0xcf,
	0x53, 0xe6, 0x52, 0x87, 0x59, 0xe8, 0x2e, 0x4c, 0x04, 0x19, 0x66, 0x95, 0x77, 0x95, 0xe5, 0x4b,
	0xb9, 0x05, 0x5d, 0x52, 0x27, 0x3d, 0x08, 0xcc, 0x9f, 0x3b, 0xfa, 0x7d, 0x61, 0xac, 0x24, 0x82,
	0xb4, 0xc7, 0x90, 0xe1, 0xaa, 0x85, 0x1a, 0x2d, 0xe3, 0xda, 0x7e, 0xb0, 0xbe, 0xe0, 0xaf, 0x17,
	0x79, 0xd1, 0x1a, 0xcc, 0x10, 0xfe, 0xa5, 0x29, 0xd4, 0x4c, 0x2e, 0x67, 0xda, 0x07, 0x3c, 0xe3,
	0x64, 0x69, 0x9a, 0xf4, 0xc5, 0x16, 0x0f, 0xb4, 0x6f, 0x14, 0x58, 0x90, 0x2a, 0x0b, 0xef, 0x5f,
	0xc0, 0x74, 0x94, 0xb4, 0x20, 0xb9, 0x2e, 0x25, 0x89, 0x90, 0x44, 0xfd, 0x26, 0x34, 0x07, 0x96,
	0x25, 0x0e, 0xf2, 0xed, 0x2d, 0x5c, 0xb7, 0x6b, 0xed, 0xe2, 0x66, 0x48, 0x99, 0x87, 0x4c, 0x24,
	0x65, 0x95, 0xaf, 0xeb, 0xc0, 0xaa, 0xfd, 0x79, 0x84, 0xd4, 0x81, 0xf6, 0x42, 0x81, 0xf7, 0x13,
	0x24, 0x14, 0xf0, 0x26, 0xbc, 0x1d, 0x95, 0xd1, 0xdf, 0xc7, 0xd7, 0xd2, 0xd2, 0x4f, 0xf5, 0xbb,
	0x62, 0xda, 0x06, 0x5c, 0x91, 0xb8, 0x09, 0xbc, 0x84, 0xe8, 0x73, 0x70, 0xf1, 0xff, 0x94, 0x17,
	0xaa, 0x21, 0xd3, 0xf7, 0x0a, 0x5c, 0x1d, 0xa0, 0x22, 0x78, 0x5c, 0x98, 0x8b, 0xa9, 0xa0, 0xd8,
	0xd3, 0x6c, 0x0a, 0x2a, 0xa1, 0x3f, 0x2b, 0xab, 0xb8, 0xe6, 0xc2, 0xb5, 0x38, 0x6b, 0xb6, 0x15,
	0xbe, 0x3b, 0x68, 0x0b, 0xa0, 0x33, 0x0a, 0x84, 0x95, 0x6b, 0x7a, 0x30, 0x37, 0x74, 0x7f, 0x6e,
	0xe8, 0xc1, 0xa4, 0x11, 0x73, 0x43, 0xdf, 0xc1, 0xc4, 0x12, 0xb1, 0xa5, 0xae, 0x48, 0xed, 0x48,
	0x81, 0xa5, 0x81, 0x29, 0x45, 0x3d, 0xf6, 0xe0, 0x32, 0x69, 0x91, 0x00, 0xdf, 0xb6, 0xc2, 0x6d,
	0x1d, 0xa2, 0x00, 0x97, 0x48, 0x8b, 0x84, 0xea, 0xa8, 0xd0, 0x43, 0x32, 0xce, 0x49, 0x96, 0x06,
	0x92, 0x04, 0x96, 0x7a, 0x50, 0x1a, 0xb0, 0xb2, 0xde, 0xc2, 0x76, 0x0d, 0x97, 0x6b, 0xd6, 0xe0,
	0x02, 0x6e, 0xc2, 0x42, 0xfc, 0xeb, 0x11, 0xf0, 0x4d, 0x96, 0xe6, 0xe4, 0xef, 0x07, 0xd3, 0x18,
	0x5c, 0x4f, 0x94, 0x53, 0x54, 0x70, 0x34, 0x49, 0x5f, 0x2a, 0xf0, 0x0e, 0xdf, 0xb3, 0xdd, 0x67,
	0xd8, 0x2d, 0x3a, 0x45, 0xa7, 0x4a, 0x47, 0xf8, 0xd2, 0xc7, 0x8d, 0xc7, 0xf1, 0x98, 0xf1, 0xf8,
	0x25, 0xcc, 0xf4, 0x99, 0x12, 0xd8, 0x0f, 0xe1, 0x32, 0x7b, 0x86, 0x5d, 0xd3, 0x76, 0x4c, 0xdb,
	0xa9, 0x52, 0xd1, 0xae, 0x8b, 0xd2, 0xc6, 0xe9, 0x92, 0x00, 0x76, 0xfa, 0x59, 0xcb, 0xc1, 0x5c,
	0x90, 0x61, 0xa7, 0xb0, 0x5f, 0xd8, 0xf5, 0x8f, 0x34, 0xe6, 0xd9, 0x95, 0xd3, 0x1d, 0x9d, 0x82,
	0xf3, 0xac, 0x6b, 0x88, 0x9f, 0x63, 0xbe, 0xab, 0x43, 0x98, 0x8f, 0x8e, 0x11, 0xd6, 0xb6, 0xe1,
	0xa2, 0xdf, 0xd3, 0xcc, 0xc3, 0x5e, 0x78, 0xde, 0xe8, 0xf2, 0x86, 0xee, 0x96, 0x78, 0x6c, 0x7b,
	0x4f, 0x6c, 0x67, 0x77, 0xa7, 0x74, 0x81, 0xb4, 0x88, 0xff, 0x98, 0x69, 0x8f, 0x20, 0x2b, 0x92,
	0xa5, 0x68, 0xc4, 0x48, 0xdb, 0x5f, 0x43, 0x2e, 0x8d, 0xd2, 0x48, 0xdb, 0xeb, 0x07, 0x05, 0x56,
	0x83, 0xe4, 0xee, 0xa7, 0xae, 0x67, 0xd7, 0x71, 0x6d, 0xd0, 0xbc, 0x8d, 0x42, 0x40, 0x7b, 0xf0,
	0x96, 0x6b, 0x57, 0x0e, 0xcd, 0x16, 0xa9, 0x9a, 0xcc, 0x6b, 0x60, 0xcf, 0x22, 0x6d, 0xde, 0x40,
	0xaf, 0xe7, 0x96, 0xe5, 0x27, 0xba, 0x5d, 0x39, 0xdc, 0x2f, 0x6c, 0xed, 0x8a, 0xf5, 0xa5, 0x37,
	0x7c, 0x89, 0x7d, 0x52, 0x0d, 0x1f, 0x68, 0x1e, 0xe8, 0x49, 0xbd, 0x89, 0xa2, 0x8c, 0xe0, 0x95,
	0xc8, 0xfd, 0xf3, 0x26, 0x9c, 0xe7, 0x69, 0xd1, 0x0b, 0x05, 0x26, 0x82, 0x6b, 0x07, 0x92, 0x9f,
	0x67, 0xfd, 0x77, 0x1d, 0xf5, 0x46, 0xb2, 0xc5, 0x81, 0x67, 0x6d, 0xe9, 0xf9, 0x2f, 0x7f, 0xbd,
	0x1c, 0x7f, 0x0f, 0x2d, 0x18, 0xf1, 0x77, 0x30, 0xf4, 0xb3, 0x02, 0xa8, 0xbf, 0x02, 0xe8, 0xa3,
	0xf8, 0x6c, 0xd2, 0xab, 0x91, 0x7a, 0x3b, 0x7d, 0xa0, 0xb0, 0xbc, 0xc6, 0x2d, 0x1b, 0x68, 0x55,
	0x6a, 0x39, 0x6a, 0x17, 0xd0, 0xdf, 0x0a, 0xcc, 0xc7, 0x5d, 0x2e, 0xd0, 0x7a, 0x5a, 0x47, 0x7d,
	0x37, 0x21, 0x35, 0x7f, 0x16, 0x09, 0x81, 0x97, 0xe7, 0x78, 0x77, 0xd0, 0xc7, 0xa9, 0xf0, 0xcc,
	0x72, 0xbb, 0xd3, 0x67, 0xe8, 0x57, 0x05, 0x66, 0x65, 0xed, 0x8a, 0xee, 0xa6, 0x35, 0xd9, 0xf3,
	0x0a, 0xaa, 0xf7, 0x86, 0x0d, 0x17, 0x7c, 0x77, 0x38, 0xdf, 0x2d, 0x74, 0x33, 0x1d, 0x5f, 0x00,
	0x87, 0xfe, 0x50, 0x40, 0x95, 0xcf, 0x27, 0x74, 0x7f, 0x28, 0x73, 0x9d, 0x19, 0xa9, 0x3e, 0x18,
	0x5e, 0x40, 0xf0, 0xdd, 0xe3, 0x7c, 0xb7, 0xd1, 0xad, 0x21, 0xf8, 0x7c, 0x84, 0x7f, 0x15, 0x58,
	0x4c, 0x30, 0x8a, 0xd1, 0x86, 0xd4, 0x69, 0xf2, 0x23, 0x41, 0xdd, 0x3c, 0x9b, 0x88, 0x40, 0x7e,
	0xc4, 0x91, 0xf3, 0xe8, 0x81, 0x14, 0x19, 0x87, 0x6a, 0x66, 0x3c, 0xfc, 0x8f, 0x0a, 0x40, 0xe7,
	0x4c, 0x46, 0x46, 0xfc, 0x6e, 0xf4, 0xdd, 0x4a, 0xd4, 0x0f, 0x92, 0x07, 0x08, 0xef, 0xab, 0xdc,
	0xfb, 0x12, 0xba, 0x2a, 0xf5, 0xde, 0x7d, 0xa1, 0x40, 0x3f, 0x29, 0x30, 0xd9, 0x73, 0x38, 0xa3,
	0x9b, 0x03, 0x52, 0x46, 0x5e, 0x21, 0xd4, 0xb5, 0x94, 0x51, 0xc2, 0x6d, 0x8e, 0xbb, 0xbd, 0x81,
	0x56, 0xe4, 0x6e, 0x5d, 0x33, 0xbc, 0x66, 0x08, 0x83, 0xdf, 0x8d, 0xc3, 0x8a, 0x38, 0xc9, 0x92,
	0xf4, 0xd5, 0x27, 0x83, 0x9c, 0xa5, 0x68, 0xaf, 0xed, 0x91, 0x68, 0x09, 0xf6, 0x6d, 0xce, 0xfe,
	0x10, 0x6d, 0xc4, 0xb1, 0x27, 0x6d, 0xb4, 0xe7, 0xe3, 0xe2, 0x1f, 0xa0, 0x81, 0xc7, 0x3b, 0xda,
	0x1a, 0x00, 0x91, 0xf0, 0xee, 0xa2, 0x16, 0xce, 0xac, 0x23, 0x0a, 0x51, 0xe0, 0x85, 0x58, 0x47,
	0xf7, 0xe3, 0x0a, 0x41, 0x03, 0xb1, 0x98, 0x32, 0xb4, 0xf3, 0xc5, 0xa3, 0xe3, 0x8c, 0xf2, 0xea,
	0x38, 0xa3, 0xfc, 0x79, 0x9c, 0x51, 0xbe, 0x3d, 0xc9, 0x8c, 0xbd, 0x3a, 0xc9, 0x8c, 0xfd, 0x76,
	0x92, 0x19, 0xfb, 0xdc, 0x20, 0xb6, 0xf7, 0xa4, 0x59, 0xd6, 0x2b, 0xb4, 0x6e, 0x58, 0xad, 0x3a,
	0x65, 0xe2, 0x6f, 0x2b, 0x9b, 0x33, 0xbe, 0x8a, 0xf8, 0xf5, 0xa5, 0x3c, 0xc1, 0x7f, 0x7e, 0xf9,
	0xf0, 0xbf, 0x00, 0x00, 0x00, 0xff, 0xff, 0xe9, 0x65, 0x17, 0xec, 0x8f, 0x12, 0x00, 0x00,
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
	// Queries a global virtual group by its id.
	GlobalVirtualGroup(ctx context.Context, in *QueryGlobalVirtualGroupRequest, opts ...grpc.CallOption) (*QueryGlobalVirtualGroupResponse, error)
	// Queries a list of global virtual groups by family id.
	GlobalVirtualGroupByFamilyID(ctx context.Context, in *QueryGlobalVirtualGroupByFamilyIDRequest, opts ...grpc.CallOption) (*QueryGlobalVirtualGroupByFamilyIDResponse, error)
	// Queries a global virtual group family by its id.
	GlobalVirtualGroupFamily(ctx context.Context, in *QueryGlobalVirtualGroupFamilyRequest, opts ...grpc.CallOption) (*QueryGlobalVirtualGroupFamilyResponse, error)
	// Queries a list of GlobalVirtualGroupFamilies items.
	GlobalVirtualGroupFamilies(ctx context.Context, in *QueryGlobalVirtualGroupFamiliesRequest, opts ...grpc.CallOption) (*QueryGlobalVirtualGroupFamiliesResponse, error)
	// AvailableGlobalVirtualGroupFamilies filters a list of GlobalVirtualGroupFamilies ID which are qualified to create bucket on
	AvailableGlobalVirtualGroupFamilies(ctx context.Context, in *AvailableGlobalVirtualGroupFamiliesRequest, opts ...grpc.CallOption) (*AvailableGlobalVirtualGroupFamiliesResponse, error)
	// SwapInInfo gets reserved swapIn info for a specific global virtual group family or global virtual group
	SwapInInfo(ctx context.Context, in *QuerySwapInInfoRequest, opts ...grpc.CallOption) (*QuerySwapInInfoResponse, error)
	// GVGStatistics gets gvg statistics for a SP
	GVGStatistics(ctx context.Context, in *QuerySPGVGStatisticsRequest, opts ...grpc.CallOption) (*QuerySPGVGStatisticsResponse, error)
	// QuerySpAvailableGlobalVirtualGroupFamilies filters a list of GlobalVirtualGroupFamilies IDs under a certain SP that are qualified to create a bucket on
	QuerySpAvailableGlobalVirtualGroupFamilies(ctx context.Context, in *QuerySPAvailableGlobalVirtualGroupFamiliesRequest, opts ...grpc.CallOption) (*QuerySPAvailableGlobalVirtualGroupFamiliesResponse, error)
	// QuerySpOptimalGlobalVirtualGroupFamily filters the optimal GlobalVirtualGroupFamily under a certain SP that is qualified to create a bucket on
	QuerySpOptimalGlobalVirtualGroupFamily(ctx context.Context, in *QuerySpOptimalGlobalVirtualGroupFamilyRequest, opts ...grpc.CallOption) (*QuerySpOptimalGlobalVirtualGroupFamilyResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) GlobalVirtualGroup(ctx context.Context, in *QueryGlobalVirtualGroupRequest, opts ...grpc.CallOption) (*QueryGlobalVirtualGroupResponse, error) {
	out := new(QueryGlobalVirtualGroupResponse)
	err := c.cc.Invoke(ctx, "/greenfield.virtualgroup.Query/GlobalVirtualGroup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) GlobalVirtualGroupByFamilyID(ctx context.Context, in *QueryGlobalVirtualGroupByFamilyIDRequest, opts ...grpc.CallOption) (*QueryGlobalVirtualGroupByFamilyIDResponse, error) {
	out := new(QueryGlobalVirtualGroupByFamilyIDResponse)
	err := c.cc.Invoke(ctx, "/greenfield.virtualgroup.Query/GlobalVirtualGroupByFamilyID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) GlobalVirtualGroupFamily(ctx context.Context, in *QueryGlobalVirtualGroupFamilyRequest, opts ...grpc.CallOption) (*QueryGlobalVirtualGroupFamilyResponse, error) {
	out := new(QueryGlobalVirtualGroupFamilyResponse)
	err := c.cc.Invoke(ctx, "/greenfield.virtualgroup.Query/GlobalVirtualGroupFamily", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) GlobalVirtualGroupFamilies(ctx context.Context, in *QueryGlobalVirtualGroupFamiliesRequest, opts ...grpc.CallOption) (*QueryGlobalVirtualGroupFamiliesResponse, error) {
	out := new(QueryGlobalVirtualGroupFamiliesResponse)
	err := c.cc.Invoke(ctx, "/greenfield.virtualgroup.Query/GlobalVirtualGroupFamilies", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) AvailableGlobalVirtualGroupFamilies(ctx context.Context, in *AvailableGlobalVirtualGroupFamiliesRequest, opts ...grpc.CallOption) (*AvailableGlobalVirtualGroupFamiliesResponse, error) {
	out := new(AvailableGlobalVirtualGroupFamiliesResponse)
	err := c.cc.Invoke(ctx, "/greenfield.virtualgroup.Query/AvailableGlobalVirtualGroupFamilies", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) SwapInInfo(ctx context.Context, in *QuerySwapInInfoRequest, opts ...grpc.CallOption) (*QuerySwapInInfoResponse, error) {
	out := new(QuerySwapInInfoResponse)
	err := c.cc.Invoke(ctx, "/greenfield.virtualgroup.Query/SwapInInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) GVGStatistics(ctx context.Context, in *QuerySPGVGStatisticsRequest, opts ...grpc.CallOption) (*QuerySPGVGStatisticsResponse, error) {
	out := new(QuerySPGVGStatisticsResponse)
	err := c.cc.Invoke(ctx, "/greenfield.virtualgroup.Query/GVGStatistics", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) QuerySpAvailableGlobalVirtualGroupFamilies(ctx context.Context, in *QuerySPAvailableGlobalVirtualGroupFamiliesRequest, opts ...grpc.CallOption) (*QuerySPAvailableGlobalVirtualGroupFamiliesResponse, error) {
	out := new(QuerySPAvailableGlobalVirtualGroupFamiliesResponse)
	err := c.cc.Invoke(ctx, "/greenfield.virtualgroup.Query/QuerySpAvailableGlobalVirtualGroupFamilies", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) QuerySpOptimalGlobalVirtualGroupFamily(ctx context.Context, in *QuerySpOptimalGlobalVirtualGroupFamilyRequest, opts ...grpc.CallOption) (*QuerySpOptimalGlobalVirtualGroupFamilyResponse, error) {
	out := new(QuerySpOptimalGlobalVirtualGroupFamilyResponse)
	err := c.cc.Invoke(ctx, "/greenfield.virtualgroup.Query/QuerySpOptimalGlobalVirtualGroupFamily", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	// Queries a global virtual group by its id.
	GlobalVirtualGroup(context.Context, *QueryGlobalVirtualGroupRequest) (*QueryGlobalVirtualGroupResponse, error)
	// Queries a list of global virtual groups by family id.
	GlobalVirtualGroupByFamilyID(context.Context, *QueryGlobalVirtualGroupByFamilyIDRequest) (*QueryGlobalVirtualGroupByFamilyIDResponse, error)
	// Queries a global virtual group family by its id.
	GlobalVirtualGroupFamily(context.Context, *QueryGlobalVirtualGroupFamilyRequest) (*QueryGlobalVirtualGroupFamilyResponse, error)
	// Queries a list of GlobalVirtualGroupFamilies items.
	GlobalVirtualGroupFamilies(context.Context, *QueryGlobalVirtualGroupFamiliesRequest) (*QueryGlobalVirtualGroupFamiliesResponse, error)
	// AvailableGlobalVirtualGroupFamilies filters a list of GlobalVirtualGroupFamilies ID which are qualified to create bucket on
	AvailableGlobalVirtualGroupFamilies(context.Context, *AvailableGlobalVirtualGroupFamiliesRequest) (*AvailableGlobalVirtualGroupFamiliesResponse, error)
	// SwapInInfo gets reserved swapIn info for a specific global virtual group family or global virtual group
	SwapInInfo(context.Context, *QuerySwapInInfoRequest) (*QuerySwapInInfoResponse, error)
	// GVGStatistics gets gvg statistics for a SP
	GVGStatistics(context.Context, *QuerySPGVGStatisticsRequest) (*QuerySPGVGStatisticsResponse, error)
	// QuerySpAvailableGlobalVirtualGroupFamilies filters a list of GlobalVirtualGroupFamilies IDs under a certain SP that are qualified to create a bucket on
	QuerySpAvailableGlobalVirtualGroupFamilies(context.Context, *QuerySPAvailableGlobalVirtualGroupFamiliesRequest) (*QuerySPAvailableGlobalVirtualGroupFamiliesResponse, error)
	// QuerySpOptimalGlobalVirtualGroupFamily filters the optimal GlobalVirtualGroupFamily under a certain SP that is qualified to create a bucket on
	QuerySpOptimalGlobalVirtualGroupFamily(context.Context, *QuerySpOptimalGlobalVirtualGroupFamilyRequest) (*QuerySpOptimalGlobalVirtualGroupFamilyResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) GlobalVirtualGroup(ctx context.Context, req *QueryGlobalVirtualGroupRequest) (*QueryGlobalVirtualGroupResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GlobalVirtualGroup not implemented")
}
func (*UnimplementedQueryServer) GlobalVirtualGroupByFamilyID(ctx context.Context, req *QueryGlobalVirtualGroupByFamilyIDRequest) (*QueryGlobalVirtualGroupByFamilyIDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GlobalVirtualGroupByFamilyID not implemented")
}
func (*UnimplementedQueryServer) GlobalVirtualGroupFamily(ctx context.Context, req *QueryGlobalVirtualGroupFamilyRequest) (*QueryGlobalVirtualGroupFamilyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GlobalVirtualGroupFamily not implemented")
}
func (*UnimplementedQueryServer) GlobalVirtualGroupFamilies(ctx context.Context, req *QueryGlobalVirtualGroupFamiliesRequest) (*QueryGlobalVirtualGroupFamiliesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GlobalVirtualGroupFamilies not implemented")
}
func (*UnimplementedQueryServer) AvailableGlobalVirtualGroupFamilies(ctx context.Context, req *AvailableGlobalVirtualGroupFamiliesRequest) (*AvailableGlobalVirtualGroupFamiliesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AvailableGlobalVirtualGroupFamilies not implemented")
}
func (*UnimplementedQueryServer) SwapInInfo(ctx context.Context, req *QuerySwapInInfoRequest) (*QuerySwapInInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SwapInInfo not implemented")
}
func (*UnimplementedQueryServer) GVGStatistics(ctx context.Context, req *QuerySPGVGStatisticsRequest) (*QuerySPGVGStatisticsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GVGStatistics not implemented")
}
func (*UnimplementedQueryServer) QuerySpAvailableGlobalVirtualGroupFamilies(ctx context.Context, req *QuerySPAvailableGlobalVirtualGroupFamiliesRequest) (*QuerySPAvailableGlobalVirtualGroupFamiliesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QuerySpAvailableGlobalVirtualGroupFamilies not implemented")
}
func (*UnimplementedQueryServer) QuerySpOptimalGlobalVirtualGroupFamily(ctx context.Context, req *QuerySpOptimalGlobalVirtualGroupFamilyRequest) (*QuerySpOptimalGlobalVirtualGroupFamilyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QuerySpOptimalGlobalVirtualGroupFamily not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_GlobalVirtualGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryGlobalVirtualGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).GlobalVirtualGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/greenfield.virtualgroup.Query/GlobalVirtualGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).GlobalVirtualGroup(ctx, req.(*QueryGlobalVirtualGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_GlobalVirtualGroupByFamilyID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryGlobalVirtualGroupByFamilyIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).GlobalVirtualGroupByFamilyID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/greenfield.virtualgroup.Query/GlobalVirtualGroupByFamilyID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).GlobalVirtualGroupByFamilyID(ctx, req.(*QueryGlobalVirtualGroupByFamilyIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_GlobalVirtualGroupFamily_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryGlobalVirtualGroupFamilyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).GlobalVirtualGroupFamily(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/greenfield.virtualgroup.Query/GlobalVirtualGroupFamily",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).GlobalVirtualGroupFamily(ctx, req.(*QueryGlobalVirtualGroupFamilyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_GlobalVirtualGroupFamilies_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryGlobalVirtualGroupFamiliesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).GlobalVirtualGroupFamilies(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/greenfield.virtualgroup.Query/GlobalVirtualGroupFamilies",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).GlobalVirtualGroupFamilies(ctx, req.(*QueryGlobalVirtualGroupFamiliesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_AvailableGlobalVirtualGroupFamilies_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AvailableGlobalVirtualGroupFamiliesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).AvailableGlobalVirtualGroupFamilies(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/greenfield.virtualgroup.Query/AvailableGlobalVirtualGroupFamilies",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).AvailableGlobalVirtualGroupFamilies(ctx, req.(*AvailableGlobalVirtualGroupFamiliesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_SwapInInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QuerySwapInInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).SwapInInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/greenfield.virtualgroup.Query/SwapInInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).SwapInInfo(ctx, req.(*QuerySwapInInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_GVGStatistics_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QuerySPGVGStatisticsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).GVGStatistics(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/greenfield.virtualgroup.Query/GVGStatistics",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).GVGStatistics(ctx, req.(*QuerySPGVGStatisticsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_QuerySpAvailableGlobalVirtualGroupFamilies_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QuerySPAvailableGlobalVirtualGroupFamiliesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).QuerySpAvailableGlobalVirtualGroupFamilies(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/greenfield.virtualgroup.Query/QuerySpAvailableGlobalVirtualGroupFamilies",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).QuerySpAvailableGlobalVirtualGroupFamilies(ctx, req.(*QuerySPAvailableGlobalVirtualGroupFamiliesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_QuerySpOptimalGlobalVirtualGroupFamily_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QuerySpOptimalGlobalVirtualGroupFamilyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).QuerySpOptimalGlobalVirtualGroupFamily(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/greenfield.virtualgroup.Query/QuerySpOptimalGlobalVirtualGroupFamily",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).QuerySpOptimalGlobalVirtualGroupFamily(ctx, req.(*QuerySpOptimalGlobalVirtualGroupFamilyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "greenfield.virtualgroup.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GlobalVirtualGroup",
			Handler:    _Query_GlobalVirtualGroup_Handler,
		},
		{
			MethodName: "GlobalVirtualGroupByFamilyID",
			Handler:    _Query_GlobalVirtualGroupByFamilyID_Handler,
		},
		{
			MethodName: "GlobalVirtualGroupFamily",
			Handler:    _Query_GlobalVirtualGroupFamily_Handler,
		},
		{
			MethodName: "GlobalVirtualGroupFamilies",
			Handler:    _Query_GlobalVirtualGroupFamilies_Handler,
		},
		{
			MethodName: "AvailableGlobalVirtualGroupFamilies",
			Handler:    _Query_AvailableGlobalVirtualGroupFamilies_Handler,
		},
		{
			MethodName: "SwapInInfo",
			Handler:    _Query_SwapInInfo_Handler,
		},
		{
			MethodName: "GVGStatistics",
			Handler:    _Query_GVGStatistics_Handler,
		},
		{
			MethodName: "QuerySpAvailableGlobalVirtualGroupFamilies",
			Handler:    _Query_QuerySpAvailableGlobalVirtualGroupFamilies_Handler,
		},
		{
			MethodName: "QuerySpOptimalGlobalVirtualGroupFamily",
			Handler:    _Query_QuerySpOptimalGlobalVirtualGroupFamily_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "greenfield/virtualgroup/query.proto",
}

func (m *QueryGlobalVirtualGroupRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryGlobalVirtualGroupRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryGlobalVirtualGroupRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.GlobalVirtualGroupId != 0 {
		i = encodeVarintQuery(dAtA, i, uint64(m.GlobalVirtualGroupId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *QueryGlobalVirtualGroupResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryGlobalVirtualGroupResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryGlobalVirtualGroupResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryGlobalVirtualGroupByFamilyIDRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryGlobalVirtualGroupByFamilyIDRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryGlobalVirtualGroupByFamilyIDRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.GlobalVirtualGroupFamilyId != 0 {
		i = encodeVarintQuery(dAtA, i, uint64(m.GlobalVirtualGroupFamilyId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *QueryGlobalVirtualGroupByFamilyIDResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryGlobalVirtualGroupByFamilyIDResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryGlobalVirtualGroupByFamilyIDResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.GlobalVirtualGroups) > 0 {
		for iNdEx := len(m.GlobalVirtualGroups) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.GlobalVirtualGroups[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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

func (m *QueryGlobalVirtualGroupFamilyRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryGlobalVirtualGroupFamilyRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryGlobalVirtualGroupFamilyRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.FamilyId != 0 {
		i = encodeVarintQuery(dAtA, i, uint64(m.FamilyId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *QueryGlobalVirtualGroupFamilyResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryGlobalVirtualGroupFamilyResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryGlobalVirtualGroupFamilyResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.GlobalVirtualGroupFamily != nil {
		{
			size, err := m.GlobalVirtualGroupFamily.MarshalToSizedBuffer(dAtA[:i])
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

func (m *QueryGlobalVirtualGroupFamiliesRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryGlobalVirtualGroupFamiliesRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryGlobalVirtualGroupFamiliesRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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

func (m *QueryGlobalVirtualGroupFamiliesResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryGlobalVirtualGroupFamiliesResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryGlobalVirtualGroupFamiliesResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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
	if len(m.GvgFamilies) > 0 {
		for iNdEx := len(m.GvgFamilies) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.GvgFamilies[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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

func (m *AvailableGlobalVirtualGroupFamiliesRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AvailableGlobalVirtualGroupFamiliesRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *AvailableGlobalVirtualGroupFamiliesRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.GlobalVirtualGroupFamilyIds) > 0 {
		dAtA7 := make([]byte, len(m.GlobalVirtualGroupFamilyIds)*10)
		var j6 int
		for _, num := range m.GlobalVirtualGroupFamilyIds {
			for num >= 1<<7 {
				dAtA7[j6] = uint8(uint64(num)&0x7f | 0x80)
				num >>= 7
				j6++
			}
			dAtA7[j6] = uint8(num)
			j6++
		}
		i -= j6
		copy(dAtA[i:], dAtA7[:j6])
		i = encodeVarintQuery(dAtA, i, uint64(j6))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *AvailableGlobalVirtualGroupFamiliesResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AvailableGlobalVirtualGroupFamiliesResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *AvailableGlobalVirtualGroupFamiliesResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.GlobalVirtualGroupFamilyIds) > 0 {
		dAtA9 := make([]byte, len(m.GlobalVirtualGroupFamilyIds)*10)
		var j8 int
		for _, num := range m.GlobalVirtualGroupFamilyIds {
			for num >= 1<<7 {
				dAtA9[j8] = uint8(uint64(num)&0x7f | 0x80)
				num >>= 7
				j8++
			}
			dAtA9[j8] = uint8(num)
			j8++
		}
		i -= j8
		copy(dAtA[i:], dAtA9[:j8])
		i = encodeVarintQuery(dAtA, i, uint64(j8))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QuerySwapInInfoRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QuerySwapInInfoRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QuerySwapInInfoRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.GlobalVirtualGroupId != 0 {
		i = encodeVarintQuery(dAtA, i, uint64(m.GlobalVirtualGroupId))
		i--
		dAtA[i] = 0x10
	}
	if m.GlobalVirtualGroupFamilyId != 0 {
		i = encodeVarintQuery(dAtA, i, uint64(m.GlobalVirtualGroupFamilyId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *QuerySwapInInfoResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QuerySwapInInfoResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QuerySwapInInfoResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.SwapInInfo != nil {
		{
			size, err := m.SwapInInfo.MarshalToSizedBuffer(dAtA[:i])
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

func (m *QuerySPGVGStatisticsRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QuerySPGVGStatisticsRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QuerySPGVGStatisticsRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.SpId != 0 {
		i = encodeVarintQuery(dAtA, i, uint64(m.SpId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *QuerySPGVGStatisticsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QuerySPGVGStatisticsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QuerySPGVGStatisticsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.GvgStats != nil {
		{
			size, err := m.GvgStats.MarshalToSizedBuffer(dAtA[:i])
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

func (m *QuerySPAvailableGlobalVirtualGroupFamiliesRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QuerySPAvailableGlobalVirtualGroupFamiliesRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QuerySPAvailableGlobalVirtualGroupFamiliesRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.SpId != 0 {
		i = encodeVarintQuery(dAtA, i, uint64(m.SpId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *QuerySPAvailableGlobalVirtualGroupFamiliesResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QuerySPAvailableGlobalVirtualGroupFamiliesResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QuerySPAvailableGlobalVirtualGroupFamiliesResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.GlobalVirtualGroupFamilyIds) > 0 {
		dAtA13 := make([]byte, len(m.GlobalVirtualGroupFamilyIds)*10)
		var j12 int
		for _, num := range m.GlobalVirtualGroupFamilyIds {
			for num >= 1<<7 {
				dAtA13[j12] = uint8(uint64(num)&0x7f | 0x80)
				num >>= 7
				j12++
			}
			dAtA13[j12] = uint8(num)
			j12++
		}
		i -= j12
		copy(dAtA[i:], dAtA13[:j12])
		i = encodeVarintQuery(dAtA, i, uint64(j12))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QuerySpOptimalGlobalVirtualGroupFamilyRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QuerySpOptimalGlobalVirtualGroupFamilyRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QuerySpOptimalGlobalVirtualGroupFamilyRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.PickVgfStrategy != 0 {
		i = encodeVarintQuery(dAtA, i, uint64(m.PickVgfStrategy))
		i--
		dAtA[i] = 0x10
	}
	if m.SpId != 0 {
		i = encodeVarintQuery(dAtA, i, uint64(m.SpId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *QuerySpOptimalGlobalVirtualGroupFamilyResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QuerySpOptimalGlobalVirtualGroupFamilyResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QuerySpOptimalGlobalVirtualGroupFamilyResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.GlobalVirtualGroupFamilyId != 0 {
		i = encodeVarintQuery(dAtA, i, uint64(m.GlobalVirtualGroupFamilyId))
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

func (m *QueryGlobalVirtualGroupRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.GlobalVirtualGroupId != 0 {
		n += 1 + sovQuery(uint64(m.GlobalVirtualGroupId))
	}
	return n
}

func (m *QueryGlobalVirtualGroupResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.GlobalVirtualGroup != nil {
		l = m.GlobalVirtualGroup.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryGlobalVirtualGroupByFamilyIDRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.GlobalVirtualGroupFamilyId != 0 {
		n += 1 + sovQuery(uint64(m.GlobalVirtualGroupFamilyId))
	}
	return n
}

func (m *QueryGlobalVirtualGroupByFamilyIDResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.GlobalVirtualGroups) > 0 {
		for _, e := range m.GlobalVirtualGroups {
			l = e.Size()
			n += 1 + l + sovQuery(uint64(l))
		}
	}
	return n
}

func (m *QueryGlobalVirtualGroupFamilyRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.FamilyId != 0 {
		n += 1 + sovQuery(uint64(m.FamilyId))
	}
	return n
}

func (m *QueryGlobalVirtualGroupFamilyResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.GlobalVirtualGroupFamily != nil {
		l = m.GlobalVirtualGroupFamily.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryGlobalVirtualGroupFamiliesRequest) Size() (n int) {
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

func (m *QueryGlobalVirtualGroupFamiliesResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.GvgFamilies) > 0 {
		for _, e := range m.GvgFamilies {
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

func (m *AvailableGlobalVirtualGroupFamiliesRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.GlobalVirtualGroupFamilyIds) > 0 {
		l = 0
		for _, e := range m.GlobalVirtualGroupFamilyIds {
			l += sovQuery(uint64(e))
		}
		n += 1 + sovQuery(uint64(l)) + l
	}
	return n
}

func (m *AvailableGlobalVirtualGroupFamiliesResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.GlobalVirtualGroupFamilyIds) > 0 {
		l = 0
		for _, e := range m.GlobalVirtualGroupFamilyIds {
			l += sovQuery(uint64(e))
		}
		n += 1 + sovQuery(uint64(l)) + l
	}
	return n
}

func (m *QuerySwapInInfoRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.GlobalVirtualGroupFamilyId != 0 {
		n += 1 + sovQuery(uint64(m.GlobalVirtualGroupFamilyId))
	}
	if m.GlobalVirtualGroupId != 0 {
		n += 1 + sovQuery(uint64(m.GlobalVirtualGroupId))
	}
	return n
}

func (m *QuerySwapInInfoResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.SwapInInfo != nil {
		l = m.SwapInInfo.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QuerySPGVGStatisticsRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.SpId != 0 {
		n += 1 + sovQuery(uint64(m.SpId))
	}
	return n
}

func (m *QuerySPGVGStatisticsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.GvgStats != nil {
		l = m.GvgStats.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QuerySPAvailableGlobalVirtualGroupFamiliesRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.SpId != 0 {
		n += 1 + sovQuery(uint64(m.SpId))
	}
	return n
}

func (m *QuerySPAvailableGlobalVirtualGroupFamiliesResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.GlobalVirtualGroupFamilyIds) > 0 {
		l = 0
		for _, e := range m.GlobalVirtualGroupFamilyIds {
			l += sovQuery(uint64(e))
		}
		n += 1 + sovQuery(uint64(l)) + l
	}
	return n
}

func (m *QuerySpOptimalGlobalVirtualGroupFamilyRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.SpId != 0 {
		n += 1 + sovQuery(uint64(m.SpId))
	}
	if m.PickVgfStrategy != 0 {
		n += 1 + sovQuery(uint64(m.PickVgfStrategy))
	}
	return n
}

func (m *QuerySpOptimalGlobalVirtualGroupFamilyResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.GlobalVirtualGroupFamilyId != 0 {
		n += 1 + sovQuery(uint64(m.GlobalVirtualGroupFamilyId))
	}
	return n
}

func sovQuery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuery(x uint64) (n int) {
	return sovQuery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *QueryGlobalVirtualGroupRequest) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryGlobalVirtualGroupRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryGlobalVirtualGroupRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field GlobalVirtualGroupId", wireType)
			}
			m.GlobalVirtualGroupId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.GlobalVirtualGroupId |= uint32(b&0x7F) << shift
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
func (m *QueryGlobalVirtualGroupResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryGlobalVirtualGroupResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryGlobalVirtualGroupResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
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
				m.GlobalVirtualGroup = &GlobalVirtualGroup{}
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
func (m *QueryGlobalVirtualGroupByFamilyIDRequest) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryGlobalVirtualGroupByFamilyIDRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryGlobalVirtualGroupByFamilyIDRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field GlobalVirtualGroupFamilyId", wireType)
			}
			m.GlobalVirtualGroupFamilyId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.GlobalVirtualGroupFamilyId |= uint32(b&0x7F) << shift
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
func (m *QueryGlobalVirtualGroupByFamilyIDResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryGlobalVirtualGroupByFamilyIDResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryGlobalVirtualGroupByFamilyIDResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GlobalVirtualGroups", wireType)
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
			m.GlobalVirtualGroups = append(m.GlobalVirtualGroups, &GlobalVirtualGroup{})
			if err := m.GlobalVirtualGroups[len(m.GlobalVirtualGroups)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
func (m *QueryGlobalVirtualGroupFamilyRequest) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryGlobalVirtualGroupFamilyRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryGlobalVirtualGroupFamilyRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field FamilyId", wireType)
			}
			m.FamilyId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.FamilyId |= uint32(b&0x7F) << shift
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
func (m *QueryGlobalVirtualGroupFamilyResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryGlobalVirtualGroupFamilyResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryGlobalVirtualGroupFamilyResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GlobalVirtualGroupFamily", wireType)
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
			if m.GlobalVirtualGroupFamily == nil {
				m.GlobalVirtualGroupFamily = &GlobalVirtualGroupFamily{}
			}
			if err := m.GlobalVirtualGroupFamily.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
func (m *QueryGlobalVirtualGroupFamiliesRequest) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryGlobalVirtualGroupFamiliesRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryGlobalVirtualGroupFamiliesRequest: illegal tag %d (wire type %d)", fieldNum, wire)
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
func (m *QueryGlobalVirtualGroupFamiliesResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryGlobalVirtualGroupFamiliesResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryGlobalVirtualGroupFamiliesResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GvgFamilies", wireType)
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
			m.GvgFamilies = append(m.GvgFamilies, &GlobalVirtualGroupFamily{})
			if err := m.GvgFamilies[len(m.GvgFamilies)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
func (m *AvailableGlobalVirtualGroupFamiliesRequest) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: AvailableGlobalVirtualGroupFamiliesRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AvailableGlobalVirtualGroupFamiliesRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType == 0 {
				var v uint32
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowQuery
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= uint32(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				m.GlobalVirtualGroupFamilyIds = append(m.GlobalVirtualGroupFamilyIds, v)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowQuery
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					packedLen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if packedLen < 0 {
					return ErrInvalidLengthQuery
				}
				postIndex := iNdEx + packedLen
				if postIndex < 0 {
					return ErrInvalidLengthQuery
				}
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				var elementCount int
				var count int
				for _, integer := range dAtA[iNdEx:postIndex] {
					if integer < 128 {
						count++
					}
				}
				elementCount = count
				if elementCount != 0 && len(m.GlobalVirtualGroupFamilyIds) == 0 {
					m.GlobalVirtualGroupFamilyIds = make([]uint32, 0, elementCount)
				}
				for iNdEx < postIndex {
					var v uint32
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowQuery
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						v |= uint32(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					m.GlobalVirtualGroupFamilyIds = append(m.GlobalVirtualGroupFamilyIds, v)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field GlobalVirtualGroupFamilyIds", wireType)
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
func (m *AvailableGlobalVirtualGroupFamiliesResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: AvailableGlobalVirtualGroupFamiliesResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AvailableGlobalVirtualGroupFamiliesResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType == 0 {
				var v uint32
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowQuery
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= uint32(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				m.GlobalVirtualGroupFamilyIds = append(m.GlobalVirtualGroupFamilyIds, v)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowQuery
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					packedLen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if packedLen < 0 {
					return ErrInvalidLengthQuery
				}
				postIndex := iNdEx + packedLen
				if postIndex < 0 {
					return ErrInvalidLengthQuery
				}
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				var elementCount int
				var count int
				for _, integer := range dAtA[iNdEx:postIndex] {
					if integer < 128 {
						count++
					}
				}
				elementCount = count
				if elementCount != 0 && len(m.GlobalVirtualGroupFamilyIds) == 0 {
					m.GlobalVirtualGroupFamilyIds = make([]uint32, 0, elementCount)
				}
				for iNdEx < postIndex {
					var v uint32
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowQuery
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						v |= uint32(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					m.GlobalVirtualGroupFamilyIds = append(m.GlobalVirtualGroupFamilyIds, v)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field GlobalVirtualGroupFamilyIds", wireType)
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
func (m *QuerySwapInInfoRequest) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QuerySwapInInfoRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QuerySwapInInfoRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field GlobalVirtualGroupFamilyId", wireType)
			}
			m.GlobalVirtualGroupFamilyId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.GlobalVirtualGroupFamilyId |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field GlobalVirtualGroupId", wireType)
			}
			m.GlobalVirtualGroupId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.GlobalVirtualGroupId |= uint32(b&0x7F) << shift
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
func (m *QuerySwapInInfoResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QuerySwapInInfoResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QuerySwapInInfoResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SwapInInfo", wireType)
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
			if m.SwapInInfo == nil {
				m.SwapInInfo = &SwapInInfo{}
			}
			if err := m.SwapInInfo.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
func (m *QuerySPGVGStatisticsRequest) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QuerySPGVGStatisticsRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QuerySPGVGStatisticsRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SpId", wireType)
			}
			m.SpId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SpId |= uint32(b&0x7F) << shift
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
func (m *QuerySPGVGStatisticsResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QuerySPGVGStatisticsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QuerySPGVGStatisticsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GvgStats", wireType)
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
			if m.GvgStats == nil {
				m.GvgStats = &GVGStatisticsWithinSP{}
			}
			if err := m.GvgStats.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
func (m *QuerySPAvailableGlobalVirtualGroupFamiliesRequest) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QuerySPAvailableGlobalVirtualGroupFamiliesRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QuerySPAvailableGlobalVirtualGroupFamiliesRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SpId", wireType)
			}
			m.SpId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SpId |= uint32(b&0x7F) << shift
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
func (m *QuerySPAvailableGlobalVirtualGroupFamiliesResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QuerySPAvailableGlobalVirtualGroupFamiliesResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QuerySPAvailableGlobalVirtualGroupFamiliesResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType == 0 {
				var v uint32
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowQuery
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= uint32(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				m.GlobalVirtualGroupFamilyIds = append(m.GlobalVirtualGroupFamilyIds, v)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowQuery
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					packedLen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if packedLen < 0 {
					return ErrInvalidLengthQuery
				}
				postIndex := iNdEx + packedLen
				if postIndex < 0 {
					return ErrInvalidLengthQuery
				}
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				var elementCount int
				var count int
				for _, integer := range dAtA[iNdEx:postIndex] {
					if integer < 128 {
						count++
					}
				}
				elementCount = count
				if elementCount != 0 && len(m.GlobalVirtualGroupFamilyIds) == 0 {
					m.GlobalVirtualGroupFamilyIds = make([]uint32, 0, elementCount)
				}
				for iNdEx < postIndex {
					var v uint32
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowQuery
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						v |= uint32(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					m.GlobalVirtualGroupFamilyIds = append(m.GlobalVirtualGroupFamilyIds, v)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field GlobalVirtualGroupFamilyIds", wireType)
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
func (m *QuerySpOptimalGlobalVirtualGroupFamilyRequest) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QuerySpOptimalGlobalVirtualGroupFamilyRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QuerySpOptimalGlobalVirtualGroupFamilyRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SpId", wireType)
			}
			m.SpId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SpId |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PickVgfStrategy", wireType)
			}
			m.PickVgfStrategy = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PickVgfStrategy |= PickVGFStrategy(b&0x7F) << shift
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
func (m *QuerySpOptimalGlobalVirtualGroupFamilyResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QuerySpOptimalGlobalVirtualGroupFamilyResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QuerySpOptimalGlobalVirtualGroupFamilyResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field GlobalVirtualGroupFamilyId", wireType)
			}
			m.GlobalVirtualGroupFamilyId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.GlobalVirtualGroupFamilyId |= uint32(b&0x7F) << shift
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
