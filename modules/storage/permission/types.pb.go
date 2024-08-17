package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
	github_com_cosmos_gogoproto_types "github.com/cosmos/gogoproto/types"
	resource "github.com/forbole/bdjuno/v4/types/resource"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	io "io"
	math "math"
	math_bits "math/bits"
	time "time"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type Policy struct {
	// id is an unique u256 sequence for each policy. It also be used as NFT tokenID
	Id Uint `protobuf:"bytes,1,opt,name=id,proto3,customtype=Uint" json:"id"`
	// principal defines the accounts/group which the permission grants to
	Principal *Principal `protobuf:"bytes,2,opt,name=principal,proto3" json:"principal,omitempty"`
	// resource_type defines the type of resource that grants permission for
	ResourceType resource.ResourceType `protobuf:"varint,3,opt,name=resource_type,json=resourceType,proto3,enum=greenfield.resource.ResourceType" json:"resource_type,omitempty"`
	// resource_id defines the bucket/object/group id of the resource that grants permission for
	ResourceId Uint `protobuf:"bytes,4,opt,name=resource_id,json=resourceId,proto3,customtype=Uint" json:"resource_id"`
	// statements defines the details content of the permission, including effect/actions/sub-resources
	Statements []*Statement `protobuf:"bytes,5,rep,name=statements,proto3" json:"statements,omitempty"`
	// expiration_time defines the whole expiration time of all the statements.
	// Notices: Its priority is higher than the expiration time inside the Statement
	ExpirationTime *time.Time `protobuf:"bytes,6,opt,name=expiration_time,json=expirationTime,proto3,stdtime" json:"expiration_time,omitempty"`
}

func (m *Policy) Reset()         { *m = Policy{} }
func (m *Policy) String() string { return proto.CompactTextString(m) }
func (*Policy) ProtoMessage()    {}
func (*Policy) Descriptor() ([]byte, []int) {
	return fileDescriptor_0d2afeea9f743f03, []int{0}
}
func (m *Policy) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Policy) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Policy.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Policy) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Policy.Merge(m, src)
}
func (m *Policy) XXX_Size() int {
	return m.Size()
}
func (m *Policy) XXX_DiscardUnknown() {
	xxx_messageInfo_Policy.DiscardUnknown(m)
}

var xxx_messageInfo_Policy proto.InternalMessageInfo

func (m *Policy) GetPrincipal() *Principal {
	if m != nil {
		return m.Principal
	}
	return nil
}

func (m *Policy) GetResourceType() resource.ResourceType {
	if m != nil {
		return m.ResourceType
	}
	return resource.RESOURCE_TYPE_UNSPECIFIED
}

func (m *Policy) GetStatements() []*Statement {
	if m != nil {
		return m.Statements
	}
	return nil
}

func (m *Policy) GetExpirationTime() *time.Time {
	if m != nil {
		return m.ExpirationTime
	}
	return nil
}

// PolicyGroup refers to a group of policies which grant permission to Group, which is limited to MaxGroupNum (default 10).
// This means that a single resource can only grant permission to 10 groups. The reason for
// this is to enable on-chain determination of whether an operator has permission within a limited time.
type PolicyGroup struct {
	// items define a pair of policy_id and group_id. Each resource can only grant its own permissions to a limited number of groups
	Items []*PolicyGroup_Item `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
}

func (m *PolicyGroup) Reset()         { *m = PolicyGroup{} }
func (m *PolicyGroup) String() string { return proto.CompactTextString(m) }
func (*PolicyGroup) ProtoMessage()    {}
func (*PolicyGroup) Descriptor() ([]byte, []int) {
	return fileDescriptor_0d2afeea9f743f03, []int{1}
}
func (m *PolicyGroup) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PolicyGroup) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PolicyGroup.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PolicyGroup) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PolicyGroup.Merge(m, src)
}
func (m *PolicyGroup) XXX_Size() int {
	return m.Size()
}
func (m *PolicyGroup) XXX_DiscardUnknown() {
	xxx_messageInfo_PolicyGroup.DiscardUnknown(m)
}

var xxx_messageInfo_PolicyGroup proto.InternalMessageInfo

func (m *PolicyGroup) GetItems() []*PolicyGroup_Item {
	if m != nil {
		return m.Items
	}
	return nil
}

type PolicyGroup_Item struct {
	PolicyId Uint `protobuf:"bytes,1,opt,name=policy_id,json=policyId,proto3,customtype=Uint" json:"policy_id"`
	GroupId  Uint `protobuf:"bytes,2,opt,name=group_id,json=groupId,proto3,customtype=Uint" json:"group_id"`
}

func (m *PolicyGroup_Item) Reset()         { *m = PolicyGroup_Item{} }
func (m *PolicyGroup_Item) String() string { return proto.CompactTextString(m) }
func (*PolicyGroup_Item) ProtoMessage()    {}
func (*PolicyGroup_Item) Descriptor() ([]byte, []int) {
	return fileDescriptor_0d2afeea9f743f03, []int{1, 0}
}
func (m *PolicyGroup_Item) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PolicyGroup_Item) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PolicyGroup_Item.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PolicyGroup_Item) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PolicyGroup_Item.Merge(m, src)
}
func (m *PolicyGroup_Item) XXX_Size() int {
	return m.Size()
}
func (m *PolicyGroup_Item) XXX_DiscardUnknown() {
	xxx_messageInfo_PolicyGroup_Item.DiscardUnknown(m)
}

var xxx_messageInfo_PolicyGroup_Item proto.InternalMessageInfo

type GroupMember struct {
	// id is an unique u256 sequence for each group member. It also be used as NFT tokenID
	Id Uint `protobuf:"bytes,1,opt,name=id,proto3,customtype=Uint" json:"id"`
	// group_id is the unique id of the group
	GroupId Uint `protobuf:"bytes,2,opt,name=group_id,json=groupId,proto3,customtype=Uint" json:"group_id"`
	// member is the account address of the member
	Member string `protobuf:"bytes,3,opt,name=member,proto3" json:"member,omitempty"`
	// expiration_time defines the expiration time of the group member
	ExpirationTime *time.Time `protobuf:"bytes,4,opt,name=expiration_time,json=expirationTime,proto3,stdtime" json:"expiration_time,omitempty"`
}

func (m *GroupMember) Reset()         { *m = GroupMember{} }
func (m *GroupMember) String() string { return proto.CompactTextString(m) }
func (*GroupMember) ProtoMessage()    {}
func (*GroupMember) Descriptor() ([]byte, []int) {
	return fileDescriptor_0d2afeea9f743f03, []int{2}
}
func (m *GroupMember) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GroupMember) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GroupMember.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GroupMember) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GroupMember.Merge(m, src)
}
func (m *GroupMember) XXX_Size() int {
	return m.Size()
}
func (m *GroupMember) XXX_DiscardUnknown() {
	xxx_messageInfo_GroupMember.DiscardUnknown(m)
}

var xxx_messageInfo_GroupMember proto.InternalMessageInfo

func (m *GroupMember) GetMember() string {
	if m != nil {
		return m.Member
	}
	return ""
}

func (m *GroupMember) GetExpirationTime() *time.Time {
	if m != nil {
		return m.ExpirationTime
	}
	return nil
}

func init() {
	proto.RegisterType((*Policy)(nil), "greenfield.permission.Policy")
	proto.RegisterType((*PolicyGroup)(nil), "greenfield.permission.PolicyGroup")
	proto.RegisterType((*PolicyGroup_Item)(nil), "greenfield.permission.PolicyGroup.Item")
	proto.RegisterType((*GroupMember)(nil), "greenfield.permission.GroupMember")
}

func init() { proto.RegisterFile("greenfield/permission/types.proto", fileDescriptor_0d2afeea9f743f03) }

var fileDescriptor_0d2afeea9f743f03 = []byte{
	// 531 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x94, 0xc1, 0x6e, 0xd3, 0x30,
	0x18, 0xc7, 0xeb, 0xae, 0x2b, 0xad, 0x0b, 0x43, 0xb2, 0x86, 0x14, 0x8a, 0x94, 0x66, 0xbd, 0x50,
	0x09, 0xd5, 0x81, 0x22, 0x21, 0x0e, 0x80, 0xa0, 0x07, 0xa6, 0x1c, 0x26, 0x4d, 0xd9, 0xb8, 0x70,
	0xa9, 0xda, 0xc4, 0x0b, 0x96, 0xea, 0xd8, 0xb2, 0xdd, 0xa9, 0x7d, 0x07, 0x0e, 0x7b, 0x98, 0x3d,
	0x03, 0xea, 0x71, 0xda, 0x09, 0x71, 0x28, 0xa8, 0x7d, 0x00, 0x5e, 0x01, 0xc5, 0x49, 0x96, 0x4a,
	0x2b, 0x14, 0xb8, 0x54, 0xfe, 0xdc, 0xdf, 0xff, 0xfb, 0xfe, 0xf6, 0xdf, 0x0a, 0x3c, 0x88, 0x24,
	0x21, 0xf1, 0x19, 0x25, 0xe3, 0xd0, 0x15, 0x44, 0x32, 0xaa, 0x14, 0xe5, 0xb1, 0xab, 0x67, 0x82,
	0x28, 0x2c, 0x24, 0xd7, 0x1c, 0x3d, 0x28, 0x10, 0x5c, 0x20, 0xcd, 0x87, 0x01, 0x57, 0x8c, 0xab,
	0x81, 0x81, 0xdc, 0xb4, 0x48, 0x15, 0xcd, 0xfd, 0x88, 0x47, 0x3c, 0xdd, 0x4f, 0x56, 0xd9, 0x6e,
	0x2b, 0xe2, 0x3c, 0x1a, 0x13, 0xd7, 0x54, 0xa3, 0xc9, 0x99, 0xab, 0x29, 0x23, 0x4a, 0x0f, 0x99,
	0xc8, 0x80, 0xf6, 0x66, 0x2f, 0x01, 0x67, 0x8c, 0xc7, 0x37, 0x4d, 0x0a, 0x46, 0x12, 0xc5, 0x27,
	0x32, 0x20, 0xeb, 0x6e, 0xdb, 0x9f, 0x77, 0x60, 0xf5, 0x98, 0x8f, 0x69, 0x30, 0x43, 0x4f, 0x60,
	0x99, 0x86, 0x16, 0x70, 0x40, 0xa7, 0xde, 0x7f, 0x34, 0x5f, 0xb4, 0x4a, 0xdf, 0x16, 0xad, 0xca,
	0x07, 0x1a, 0xeb, 0xeb, 0xcb, 0x6e, 0x23, 0x33, 0x9c, 0x94, 0x7e, 0x99, 0x86, 0xe8, 0x0d, 0xac,
	0x0b, 0x49, 0xe3, 0x80, 0x8a, 0xe1, 0xd8, 0x2a, 0x3b, 0xa0, 0xd3, 0xe8, 0x39, 0x78, 0xe3, 0xc9,
	0xf1, 0x71, 0xce, 0xf9, 0x85, 0x04, 0xbd, 0x87, 0xf7, 0x72, 0x3f, 0x83, 0xc4, 0x8f, 0xb5, 0xe3,
	0x80, 0xce, 0x5e, 0xef, 0x60, 0xbd, 0x47, 0x0e, 0x60, 0x3f, 0x5b, 0x9c, 0xce, 0x04, 0xf1, 0xef,
	0xca, 0xb5, 0x0a, 0xbd, 0x82, 0x8d, 0x9b, 0x3e, 0x34, 0xb4, 0x2a, 0xdb, 0xdd, 0xc3, 0x9c, 0xf7,
	0x42, 0xf4, 0x16, 0x42, 0xa5, 0x87, 0x9a, 0x30, 0x12, 0x6b, 0x65, 0xed, 0x3a, 0x3b, 0x7f, 0x38,
	0xc6, 0x49, 0x0e, 0xfa, 0x6b, 0x1a, 0x74, 0x04, 0xef, 0x93, 0xa9, 0xa0, 0x72, 0xa8, 0x29, 0x8f,
	0x07, 0x49, 0x44, 0x56, 0xd5, 0xdc, 0x46, 0x13, 0xa7, 0xf9, 0xe1, 0x3c, 0x3f, 0x7c, 0x9a, 0xe7,
	0xd7, 0xaf, 0xcd, 0x17, 0x2d, 0x70, 0xf1, 0xbd, 0x05, 0xfc, 0xbd, 0x42, 0x9c, 0xfc, 0xdd, 0xfe,
	0x02, 0x60, 0x23, 0x8d, 0xe3, 0x50, 0xf2, 0x89, 0x40, 0xaf, 0xe1, 0x2e, 0xd5, 0x84, 0x29, 0x0b,
	0x18, 0x6f, 0x8f, 0x7f, 0x77, 0xc5, 0x85, 0x04, 0x7b, 0x9a, 0x30, 0x3f, 0x55, 0x35, 0xa7, 0xb0,
	0x92, 0x94, 0xe8, 0x25, 0xac, 0x0b, 0x83, 0x0c, 0xfe, 0x2e, 0xe1, 0x5a, 0x4a, 0x7b, 0x21, 0x7a,
	0x01, 0x6b, 0x51, 0xd2, 0x36, 0x11, 0x96, 0xb7, 0x0b, 0xef, 0x18, 0xd8, 0x0b, 0xdb, 0x3f, 0x01,
	0x6c, 0x18, 0x3f, 0x47, 0x84, 0x8d, 0x88, 0xfc, 0xb7, 0xc7, 0xf5, 0x9f, 0x43, 0xd1, 0x53, 0x58,
	0x65, 0x66, 0x9c, 0x79, 0x4d, 0xf5, 0xbe, 0x75, 0x7d, 0xd9, 0xdd, 0xcf, 0xc8, 0x77, 0x61, 0x28,
	0x89, 0x52, 0x27, 0x5a, 0xd2, 0x38, 0xf2, 0x33, 0x0e, 0x79, 0xb7, 0xe3, 0xab, 0x6c, 0x8d, 0xaf,
	0xb2, 0x29, 0xba, 0xfe, 0xe1, 0x7c, 0x69, 0x83, 0xab, 0xa5, 0x0d, 0x7e, 0x2c, 0x6d, 0x70, 0xb1,
	0xb2, 0x4b, 0x57, 0x2b, 0xbb, 0xf4, 0x75, 0x65, 0x97, 0x3e, 0x76, 0x23, 0xaa, 0x3f, 0x4d, 0x46,
	0x38, 0xe0, 0xcc, 0x25, 0xe7, 0x8c, 0xab, 0xec, 0xf7, 0xfc, 0x59, 0xcf, 0x9d, 0xde, 0xfa, 0x8c,
	0x8c, 0xaa, 0x66, 0xe4, 0xf3, 0x5f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x3a, 0xae, 0x8e, 0x35, 0x6c,
	0x04, 0x00, 0x00,
}

func (m *Policy) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Policy) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Policy) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.ExpirationTime != nil {
		n1, err1 := github_com_cosmos_gogoproto_types.StdTimeMarshalTo(*m.ExpirationTime, dAtA[i-github_com_cosmos_gogoproto_types.SizeOfStdTime(*m.ExpirationTime):])
		if err1 != nil {
			return 0, err1
		}
		i -= n1
		i = encodeVarintTypes(dAtA, i, uint64(n1))
		i--
		dAtA[i] = 0x32
	}
	if len(m.Statements) > 0 {
		for iNdEx := len(m.Statements) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Statements[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTypes(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x2a
		}
	}
	{
		size := m.ResourceId.Size()
		i -= size
		if _, err := m.ResourceId.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintTypes(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	if m.ResourceType != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.ResourceType))
		i--
		dAtA[i] = 0x18
	}
	if m.Principal != nil {
		{
			size, err := m.Principal.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTypes(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	{
		size := m.Id.Size()
		i -= size
		if _, err := m.Id.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintTypes(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *PolicyGroup) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PolicyGroup) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PolicyGroup) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Items) > 0 {
		for iNdEx := len(m.Items) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Items[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTypes(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *PolicyGroup_Item) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PolicyGroup_Item) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PolicyGroup_Item) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.GroupId.Size()
		i -= size
		if _, err := m.GroupId.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintTypes(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	{
		size := m.PolicyId.Size()
		i -= size
		if _, err := m.PolicyId.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintTypes(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *GroupMember) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GroupMember) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GroupMember) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.ExpirationTime != nil {
		n3, err3 := github_com_cosmos_gogoproto_types.StdTimeMarshalTo(*m.ExpirationTime, dAtA[i-github_com_cosmos_gogoproto_types.SizeOfStdTime(*m.ExpirationTime):])
		if err3 != nil {
			return 0, err3
		}
		i -= n3
		i = encodeVarintTypes(dAtA, i, uint64(n3))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Member) > 0 {
		i -= len(m.Member)
		copy(dAtA[i:], m.Member)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.Member)))
		i--
		dAtA[i] = 0x1a
	}
	{
		size := m.GroupId.Size()
		i -= size
		if _, err := m.GroupId.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintTypes(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	{
		size := m.Id.Size()
		i -= size
		if _, err := m.Id.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintTypes(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintTypes(dAtA []byte, offset int, v uint64) int {
	offset -= sovTypes(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Policy) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Id.Size()
	n += 1 + l + sovTypes(uint64(l))
	if m.Principal != nil {
		l = m.Principal.Size()
		n += 1 + l + sovTypes(uint64(l))
	}
	if m.ResourceType != 0 {
		n += 1 + sovTypes(uint64(m.ResourceType))
	}
	l = m.ResourceId.Size()
	n += 1 + l + sovTypes(uint64(l))
	if len(m.Statements) > 0 {
		for _, e := range m.Statements {
			l = e.Size()
			n += 1 + l + sovTypes(uint64(l))
		}
	}
	if m.ExpirationTime != nil {
		l = github_com_cosmos_gogoproto_types.SizeOfStdTime(*m.ExpirationTime)
		n += 1 + l + sovTypes(uint64(l))
	}
	return n
}

func (m *PolicyGroup) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Items) > 0 {
		for _, e := range m.Items {
			l = e.Size()
			n += 1 + l + sovTypes(uint64(l))
		}
	}
	return n
}

func (m *PolicyGroup_Item) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.PolicyId.Size()
	n += 1 + l + sovTypes(uint64(l))
	l = m.GroupId.Size()
	n += 1 + l + sovTypes(uint64(l))
	return n
}

func (m *GroupMember) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Id.Size()
	n += 1 + l + sovTypes(uint64(l))
	l = m.GroupId.Size()
	n += 1 + l + sovTypes(uint64(l))
	l = len(m.Member)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	if m.ExpirationTime != nil {
		l = github_com_cosmos_gogoproto_types.SizeOfStdTime(*m.ExpirationTime)
		n += 1 + l + sovTypes(uint64(l))
	}
	return n
}

func sovTypes(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTypes(x uint64) (n int) {
	return sovTypes(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Policy) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
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
			return fmt.Errorf("proto: Policy: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Policy: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Id.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Principal", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Principal == nil {
				m.Principal = &Principal{}
			}
			if err := m.Principal.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ResourceType", wireType)
			}
			m.ResourceType = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ResourceType |= resource.ResourceType(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ResourceId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ResourceId.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Statements", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Statements = append(m.Statements, &Statement{})
			if err := m.Statements[len(m.Statements)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ExpirationTime", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.ExpirationTime == nil {
				m.ExpirationTime = new(time.Time)
			}
			if err := github_com_cosmos_gogoproto_types.StdTimeUnmarshal(m.ExpirationTime, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTypes
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
func (m *PolicyGroup) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
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
			return fmt.Errorf("proto: PolicyGroup: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PolicyGroup: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Items", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Items = append(m.Items, &PolicyGroup_Item{})
			if err := m.Items[len(m.Items)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTypes
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
func (m *PolicyGroup_Item) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
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
			return fmt.Errorf("proto: Item: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Item: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PolicyId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.PolicyId.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GroupId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.GroupId.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTypes
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
func (m *GroupMember) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
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
			return fmt.Errorf("proto: GroupMember: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GroupMember: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Id.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GroupId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.GroupId.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Member", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Member = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ExpirationTime", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.ExpirationTime == nil {
				m.ExpirationTime = new(time.Time)
			}
			if err := github_com_cosmos_gogoproto_types.StdTimeUnmarshal(m.ExpirationTime, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTypes
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
func skipTypes(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTypes
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
					return 0, ErrIntOverflowTypes
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
					return 0, ErrIntOverflowTypes
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
				return 0, ErrInvalidLengthTypes
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTypes
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTypes
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTypes        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTypes          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTypes = fmt.Errorf("proto: unexpected end of group")
)