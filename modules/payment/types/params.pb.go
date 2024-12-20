// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: mechain/payment/params.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
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

// Params defines the parameters for the module.
type Params struct {
	VersionedParams VersionedParams `protobuf:"bytes,1,opt,name=versioned_params,json=versionedParams,proto3" json:"versioned_params"`
	// The maximum number of payment accounts that can be created by one user
	PaymentAccountCountLimit uint64 `protobuf:"varint,2,opt,name=payment_account_count_limit,json=paymentAccountCountLimit,proto3" json:"payment_account_count_limit,omitempty" yaml:"payment_account_count_limit"`
	// Time duration threshold of forced settlement.
	// If dynamic balance is less than NetOutFlowRate * forcedSettleTime, the account can be forced settled.
	ForcedSettleTime uint64 `protobuf:"varint,3,opt,name=forced_settle_time,json=forcedSettleTime,proto3" json:"forced_settle_time,omitempty" yaml:"forced_settle_time"`
	// the maximum number of flows that will be auto forced settled in one block
	MaxAutoSettleFlowCount uint64 `protobuf:"varint,4,opt,name=max_auto_settle_flow_count,json=maxAutoSettleFlowCount,proto3" json:"max_auto_settle_flow_count,omitempty" yaml:"max_auto_settle_flow_count"`
	// the maximum number of flows that will be auto resumed in one block
	MaxAutoResumeFlowCount uint64 `protobuf:"varint,5,opt,name=max_auto_resume_flow_count,json=maxAutoResumeFlowCount,proto3" json:"max_auto_resume_flow_count,omitempty" yaml:"max_auto_resume_flow_count"`
	// The denom of fee charged in payment module
	FeeDenom string `protobuf:"bytes,6,opt,name=fee_denom,json=feeDenom,proto3" json:"fee_denom,omitempty" yaml:"fee_denom"`
	// The withdrawal amount threshold to trigger time lock
	WithdrawTimeLockThreshold *github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,7,opt,name=withdraw_time_lock_threshold,json=withdrawTimeLockThreshold,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"withdraw_time_lock_threshold,omitempty"`
	// The duration of the time lock for a big amount withdrawal
	WithdrawTimeLockDuration uint64 `protobuf:"varint,8,opt,name=withdraw_time_lock_duration,json=withdrawTimeLockDuration,proto3" json:"withdraw_time_lock_duration,omitempty" yaml:"withdraw_time_lock_duration"`
}

func (m *Params) Reset()         { *m = Params{} }
func (m *Params) String() string { return proto.CompactTextString(m) }
func (*Params) ProtoMessage()    {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_080c11c074cb8da6, []int{0}
}
func (m *Params) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Params) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Params.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Params) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Params.Merge(m, src)
}
func (m *Params) XXX_Size() int {
	return m.Size()
}
func (m *Params) XXX_DiscardUnknown() {
	xxx_messageInfo_Params.DiscardUnknown(m)
}

var xxx_messageInfo_Params proto.InternalMessageInfo

func (m *Params) GetVersionedParams() VersionedParams {
	if m != nil {
		return m.VersionedParams
	}
	return VersionedParams{}
}

func (m *Params) GetPaymentAccountCountLimit() uint64 {
	if m != nil {
		return m.PaymentAccountCountLimit
	}
	return 0
}

func (m *Params) GetForcedSettleTime() uint64 {
	if m != nil {
		return m.ForcedSettleTime
	}
	return 0
}

func (m *Params) GetMaxAutoSettleFlowCount() uint64 {
	if m != nil {
		return m.MaxAutoSettleFlowCount
	}
	return 0
}

func (m *Params) GetMaxAutoResumeFlowCount() uint64 {
	if m != nil {
		return m.MaxAutoResumeFlowCount
	}
	return 0
}

func (m *Params) GetFeeDenom() string {
	if m != nil {
		return m.FeeDenom
	}
	return ""
}

func (m *Params) GetWithdrawTimeLockDuration() uint64 {
	if m != nil {
		return m.WithdrawTimeLockDuration
	}
	return 0
}

// VersionedParams defines the parameters with multiple versions, each version is stored with different timestamp.
type VersionedParams struct {
	// Time duration which the buffer balance need to be reserved for NetOutFlow e.g. 6 month
	ReserveTime uint64 `protobuf:"varint,1,opt,name=reserve_time,json=reserveTime,proto3" json:"reserve_time,omitempty" yaml:"reserve_time"`
	// The tax rate to pay for validators in storage payment. The default value is 1%(0.01)
	ValidatorTaxRate github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,2,opt,name=validator_tax_rate,json=validatorTaxRate,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"validator_tax_rate"`
}

func (m *VersionedParams) Reset()         { *m = VersionedParams{} }
func (m *VersionedParams) String() string { return proto.CompactTextString(m) }
func (*VersionedParams) ProtoMessage()    {}
func (*VersionedParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_080c11c074cb8da6, []int{1}
}
func (m *VersionedParams) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *VersionedParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_VersionedParams.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *VersionedParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VersionedParams.Merge(m, src)
}
func (m *VersionedParams) XXX_Size() int {
	return m.Size()
}
func (m *VersionedParams) XXX_DiscardUnknown() {
	xxx_messageInfo_VersionedParams.DiscardUnknown(m)
}

var xxx_messageInfo_VersionedParams proto.InternalMessageInfo

func (m *VersionedParams) GetReserveTime() uint64 {
	if m != nil {
		return m.ReserveTime
	}
	return 0
}

func init() {
	proto.RegisterType((*Params)(nil), "mechain.payment.Params")
	proto.RegisterType((*VersionedParams)(nil), "mechain.payment.VersionedParams")
}

func init() { proto.RegisterFile("mechain/payment/params.proto", fileDescriptor_080c11c074cb8da6) }

var fileDescriptor_080c11c074cb8da6 = []byte{
	// 581 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x54, 0xc1, 0x6e, 0xd3, 0x4a,
	0x14, 0x8d, 0xdf, 0x6b, 0x4b, 0x3b, 0x45, 0x6a, 0x64, 0x2a, 0x70, 0x4b, 0xb1, 0x83, 0x25, 0xaa,
	0x08, 0xa9, 0x8e, 0x5a, 0x36, 0xa8, 0x62, 0xd3, 0x10, 0x21, 0x55, 0x74, 0x01, 0x26, 0x62, 0xc1,
	0x66, 0x34, 0xb5, 0x6f, 0x12, 0x13, 0x8f, 0x27, 0x1a, 0x8f, 0x9d, 0xe4, 0x2f, 0xf8, 0x18, 0x36,
	0xfc, 0x41, 0x97, 0x15, 0x2b, 0xc4, 0xc2, 0x42, 0xc9, 0x8e, 0x65, 0xbe, 0x00, 0x79, 0xc6, 0x6e,
	0x93, 0x94, 0x56, 0x6c, 0xae, 0x67, 0xee, 0x39, 0x73, 0xce, 0xe8, 0xea, 0x8c, 0xd1, 0x1e, 0x05,
	0xaf, 0x47, 0x82, 0xa8, 0x31, 0x20, 0x63, 0x0a, 0x91, 0x68, 0x0c, 0x08, 0x27, 0x34, 0x76, 0x06,
	0x9c, 0x09, 0xa6, 0x6f, 0x15, 0xa8, 0x53, 0xa0, 0xbb, 0x3b, 0x1e, 0x8b, 0x29, 0x8b, 0xb1, 0x84,
	0x1b, 0x6a, 0xa3, 0xb8, 0xbb, 0xdb, 0x5d, 0xd6, 0x65, 0xaa, 0x9f, 0xaf, 0x54, 0xd7, 0xfe, 0xbd,
	0x8a, 0xd6, 0xde, 0x49, 0x49, 0xfd, 0x3d, 0xaa, 0xa6, 0xc0, 0xe3, 0x80, 0x45, 0xe0, 0x63, 0x65,
	0x63, 0x68, 0x35, 0xad, 0xbe, 0x79, 0x54, 0x73, 0x96, 0x7c, 0x9c, 0x8f, 0x25, 0x51, 0x9d, 0x6d,
	0xae, 0x5c, 0x64, 0x56, 0xc5, 0xdd, 0x4a, 0x17, 0xdb, 0x3a, 0xa0, 0xc7, 0xc5, 0x09, 0x4c, 0x3c,
	0x8f, 0x25, 0x91, 0xc0, 0xaa, 0x86, 0x01, 0x0d, 0x84, 0xf1, 0x5f, 0x4d, 0xab, 0xaf, 0x34, 0xf7,
	0x67, 0x99, 0x65, 0x8f, 0x09, 0x0d, 0x8f, 0xed, 0x3b, 0xc8, 0xb6, 0x6b, 0x14, 0xe8, 0x89, 0x02,
	0x5f, 0xe7, 0xe5, 0x2c, 0x87, 0xf4, 0xb7, 0x48, 0xef, 0x30, 0xee, 0x81, 0x8f, 0x63, 0x10, 0x22,
	0x04, 0x2c, 0x02, 0x0a, 0xc6, 0xff, 0x52, 0xfd, 0xc9, 0x2c, 0xb3, 0x76, 0x94, 0xfa, 0x4d, 0x8e,
	0xed, 0x56, 0x55, 0xf3, 0x83, 0xec, 0xb5, 0x03, 0x0a, 0x3a, 0x41, 0xbb, 0x94, 0x8c, 0x30, 0x49,
	0x04, 0x2b, 0xa9, 0x9d, 0x90, 0x0d, 0xd5, 0x5d, 0x8c, 0x15, 0x29, 0xfa, 0x6c, 0x96, 0x59, 0x4f,
	0x95, 0xe8, 0xed, 0x5c, 0xdb, 0x7d, 0x48, 0xc9, 0xe8, 0x24, 0x11, 0x4c, 0xa9, 0xbf, 0x09, 0xd9,
	0x50, 0x5e, 0x7a, 0xc1, 0x82, 0x43, 0x9c, 0xd0, 0x05, 0x8b, 0xd5, 0x5b, 0x2d, 0x6e, 0x70, 0xaf,
	0x2d, 0x5c, 0x09, 0x5d, 0x5b, 0x1c, 0xa2, 0x8d, 0x0e, 0x00, 0xf6, 0x21, 0x62, 0xd4, 0x58, 0xab,
	0x69, 0xf5, 0x8d, 0xe6, 0xf6, 0x2c, 0xb3, 0xaa, 0xc5, 0x24, 0x4a, 0xc8, 0x76, 0xd7, 0x3b, 0x00,
	0xad, 0x7c, 0xa9, 0x8f, 0xd1, 0xde, 0x30, 0x10, 0x3d, 0x9f, 0x93, 0xa1, 0x1c, 0x0e, 0x0e, 0x99,
	0xd7, 0xc7, 0xa2, 0xc7, 0x21, 0xee, 0xb1, 0xd0, 0x37, 0xee, 0x49, 0x95, 0x97, 0x3f, 0x33, 0x6b,
	0xbf, 0x1b, 0x88, 0x5e, 0x72, 0xee, 0x78, 0x8c, 0x16, 0x19, 0x2b, 0x3e, 0x07, 0xb1, 0xdf, 0x6f,
	0x88, 0xf1, 0x00, 0x62, 0xe7, 0x34, 0x12, 0xdf, 0xbf, 0x1e, 0xa0, 0x22, 0x82, 0xa7, 0x91, 0x70,
	0x77, 0x4a, 0xf5, 0x7c, 0xcc, 0x67, 0xcc, 0xeb, 0xb7, 0x4b, 0xe9, 0x3c, 0x27, 0x7f, 0xb1, 0xf6,
	0x13, 0x4e, 0x44, 0xc0, 0x22, 0x63, 0x7d, 0x39, 0x27, 0x77, 0x90, 0x6d, 0xd7, 0x58, 0xf6, 0x69,
	0x95, 0xd0, 0x37, 0x0d, 0x6d, 0x2d, 0x25, 0x57, 0x3f, 0x46, 0xf7, 0x39, 0xc4, 0xc0, 0xd3, 0x22,
	0x35, 0x9a, 0xf4, 0x7a, 0x34, 0xcb, 0xac, 0x07, 0xca, 0x6b, 0x1e, 0xb5, 0xdd, 0xcd, 0x62, 0x2b,
	0xa3, 0xf2, 0x19, 0xe9, 0x29, 0x09, 0x03, 0x9f, 0x08, 0xc6, 0xb1, 0x20, 0x23, 0xcc, 0x89, 0x00,
	0x99, 0xea, 0x8d, 0xe6, 0xab, 0xfc, 0x45, 0xfc, 0xe3, 0xac, 0x5a, 0xe0, 0xcd, 0xcd, 0xaa, 0x05,
	0x9e, 0x5b, 0xbd, 0xd2, 0x6d, 0x93, 0x91, 0x4b, 0x04, 0x34, 0x5b, 0x17, 0x13, 0x53, 0xbb, 0x9c,
	0x98, 0xda, 0xaf, 0x89, 0xa9, 0x7d, 0x99, 0x9a, 0x95, 0xcb, 0xa9, 0x59, 0xf9, 0x31, 0x35, 0x2b,
	0x9f, 0x9e, 0xcf, 0x39, 0x40, 0x9a, 0x1b, 0xa8, 0x9a, 0x1e, 0x1e, 0x35, 0x46, 0x57, 0xff, 0x0d,
	0xe9, 0x74, 0xbe, 0x26, 0x5f, 0xfd, 0x8b, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x45, 0xfa, 0xb4,
	0x21, 0x57, 0x04, 0x00, 0x00,
}

func (m *Params) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Params) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Params) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.WithdrawTimeLockDuration != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.WithdrawTimeLockDuration))
		i--
		dAtA[i] = 0x40
	}
	if m.WithdrawTimeLockThreshold != nil {
		{
			size := m.WithdrawTimeLockThreshold.Size()
			i -= size
			if _, err := m.WithdrawTimeLockThreshold.MarshalTo(dAtA[i:]); err != nil {
				return 0, err
			}
			i = encodeVarintParams(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x3a
	}
	if len(m.FeeDenom) > 0 {
		i -= len(m.FeeDenom)
		copy(dAtA[i:], m.FeeDenom)
		i = encodeVarintParams(dAtA, i, uint64(len(m.FeeDenom)))
		i--
		dAtA[i] = 0x32
	}
	if m.MaxAutoResumeFlowCount != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.MaxAutoResumeFlowCount))
		i--
		dAtA[i] = 0x28
	}
	if m.MaxAutoSettleFlowCount != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.MaxAutoSettleFlowCount))
		i--
		dAtA[i] = 0x20
	}
	if m.ForcedSettleTime != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.ForcedSettleTime))
		i--
		dAtA[i] = 0x18
	}
	if m.PaymentAccountCountLimit != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.PaymentAccountCountLimit))
		i--
		dAtA[i] = 0x10
	}
	{
		size, err := m.VersionedParams.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *VersionedParams) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *VersionedParams) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *VersionedParams) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.ValidatorTaxRate.Size()
		i -= size
		if _, err := m.ValidatorTaxRate.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if m.ReserveTime != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.ReserveTime))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintParams(dAtA []byte, offset int, v uint64) int {
	offset -= sovParams(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Params) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.VersionedParams.Size()
	n += 1 + l + sovParams(uint64(l))
	if m.PaymentAccountCountLimit != 0 {
		n += 1 + sovParams(uint64(m.PaymentAccountCountLimit))
	}
	if m.ForcedSettleTime != 0 {
		n += 1 + sovParams(uint64(m.ForcedSettleTime))
	}
	if m.MaxAutoSettleFlowCount != 0 {
		n += 1 + sovParams(uint64(m.MaxAutoSettleFlowCount))
	}
	if m.MaxAutoResumeFlowCount != 0 {
		n += 1 + sovParams(uint64(m.MaxAutoResumeFlowCount))
	}
	l = len(m.FeeDenom)
	if l > 0 {
		n += 1 + l + sovParams(uint64(l))
	}
	if m.WithdrawTimeLockThreshold != nil {
		l = m.WithdrawTimeLockThreshold.Size()
		n += 1 + l + sovParams(uint64(l))
	}
	if m.WithdrawTimeLockDuration != 0 {
		n += 1 + sovParams(uint64(m.WithdrawTimeLockDuration))
	}
	return n
}

func (m *VersionedParams) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ReserveTime != 0 {
		n += 1 + sovParams(uint64(m.ReserveTime))
	}
	l = m.ValidatorTaxRate.Size()
	n += 1 + l + sovParams(uint64(l))
	return n
}

func sovParams(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozParams(x uint64) (n int) {
	return sovParams(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowParams
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
			return fmt.Errorf("proto: Params: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Params: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field VersionedParams", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.VersionedParams.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PaymentAccountCountLimit", wireType)
			}
			m.PaymentAccountCountLimit = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PaymentAccountCountLimit |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ForcedSettleTime", wireType)
			}
			m.ForcedSettleTime = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ForcedSettleTime |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxAutoSettleFlowCount", wireType)
			}
			m.MaxAutoSettleFlowCount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MaxAutoSettleFlowCount |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxAutoResumeFlowCount", wireType)
			}
			m.MaxAutoResumeFlowCount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MaxAutoResumeFlowCount |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FeeDenom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.FeeDenom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field WithdrawTimeLockThreshold", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			var v github_com_cosmos_cosmos_sdk_types.Int
			m.WithdrawTimeLockThreshold = &v
			if err := m.WithdrawTimeLockThreshold.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field WithdrawTimeLockDuration", wireType)
			}
			m.WithdrawTimeLockDuration = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.WithdrawTimeLockDuration |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipParams(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthParams
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
func (m *VersionedParams) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowParams
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
			return fmt.Errorf("proto: VersionedParams: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: VersionedParams: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ReserveTime", wireType)
			}
			m.ReserveTime = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ReserveTime |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ValidatorTaxRate", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ValidatorTaxRate.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipParams(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthParams
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
func skipParams(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowParams
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
					return 0, ErrIntOverflowParams
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
					return 0, ErrIntOverflowParams
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
				return 0, ErrInvalidLengthParams
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupParams
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthParams
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthParams        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowParams          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupParams = fmt.Errorf("proto: unexpected end of group")
)
