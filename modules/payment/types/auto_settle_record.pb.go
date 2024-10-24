// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: mechain/payment/auto_settle_record.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
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

// AutoSettleRecord is the record keeps the auto settle information.
// The EndBlocker of payment module will scan the list of AutoSettleRecord
// and settle the stream account if the timestamp is less than the current time.
type AutoSettleRecord struct {
	// timestamp is the unix timestamp when the stream account will be settled.
	Timestamp int64 `protobuf:"varint,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	// A stream account address
	Addr string `protobuf:"bytes,2,opt,name=addr,proto3" json:"addr,omitempty"`
}

func (m *AutoSettleRecord) Reset()         { *m = AutoSettleRecord{} }
func (m *AutoSettleRecord) String() string { return proto.CompactTextString(m) }
func (*AutoSettleRecord) ProtoMessage()    {}
func (*AutoSettleRecord) Descriptor() ([]byte, []int) {
	return fileDescriptor_e3fb6b8de54e2985, []int{0}
}
func (m *AutoSettleRecord) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *AutoSettleRecord) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AutoSettleRecord.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *AutoSettleRecord) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AutoSettleRecord.Merge(m, src)
}
func (m *AutoSettleRecord) XXX_Size() int {
	return m.Size()
}
func (m *AutoSettleRecord) XXX_DiscardUnknown() {
	xxx_messageInfo_AutoSettleRecord.DiscardUnknown(m)
}

var xxx_messageInfo_AutoSettleRecord proto.InternalMessageInfo

func (m *AutoSettleRecord) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *AutoSettleRecord) GetAddr() string {
	if m != nil {
		return m.Addr
	}
	return ""
}

func init() {
	proto.RegisterType((*AutoSettleRecord)(nil), "mechain.payment.AutoSettleRecord")
}

func init() {
	proto.RegisterFile("mechain/payment/auto_settle_record.proto", fileDescriptor_e3fb6b8de54e2985)
}

var fileDescriptor_e3fb6b8de54e2985 = []byte{
	// 228 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xd2, 0xc8, 0x4d, 0x4d, 0xce,
	0x48, 0xcc, 0xcc, 0xd3, 0x2f, 0x48, 0xac, 0xcc, 0x4d, 0xcd, 0x2b, 0xd1, 0x4f, 0x2c, 0x2d, 0xc9,
	0x8f, 0x2f, 0x4e, 0x2d, 0x29, 0xc9, 0x49, 0x8d, 0x2f, 0x4a, 0x4d, 0xce, 0x2f, 0x4a, 0xd1, 0x2b,
	0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x87, 0xaa, 0xd4, 0x83, 0xaa, 0x94, 0x92, 0x4c, 0xce, 0x2f,
	0xce, 0xcd, 0x2f, 0x8e, 0x07, 0x4b, 0xeb, 0x43, 0x38, 0x10, 0xb5, 0x4a, 0x71, 0x5c, 0x02, 0x8e,
	0xa5, 0x25, 0xf9, 0xc1, 0x60, 0x63, 0x82, 0xc0, 0xa6, 0x08, 0xc9, 0x70, 0x71, 0x96, 0x64, 0xe6,
	0xa6, 0x16, 0x97, 0x24, 0xe6, 0x16, 0x48, 0x30, 0x2a, 0x30, 0x6a, 0x30, 0x07, 0x21, 0x04, 0x84,
	0x74, 0xb8, 0x58, 0x12, 0x53, 0x52, 0x8a, 0x24, 0x98, 0x14, 0x18, 0x35, 0x38, 0x9d, 0x24, 0x2e,
	0x6d, 0xd1, 0x15, 0x81, 0x9a, 0xe8, 0x98, 0x92, 0x52, 0x94, 0x5a, 0x5c, 0x1c, 0x5c, 0x52, 0x94,
	0x99, 0x97, 0x1e, 0x04, 0x56, 0xe5, 0xe4, 0x72, 0xe2, 0x91, 0x1c, 0xe3, 0x85, 0x47, 0x72, 0x8c,
	0x0f, 0x1e, 0xc9, 0x31, 0x4e, 0x78, 0x2c, 0xc7, 0x70, 0xe1, 0xb1, 0x1c, 0xc3, 0x8d, 0xc7, 0x72,
	0x0c, 0x51, 0x5a, 0xe9, 0x99, 0x25, 0x19, 0xa5, 0x49, 0x7a, 0xc9, 0xf9, 0xb9, 0xfa, 0xa9, 0x65,
	0xb9, 0xf9, 0xc5, 0x50, 0xb2, 0xcc, 0xd0, 0x48, 0xbf, 0x02, 0xee, 0xc9, 0x92, 0xca, 0x82, 0xd4,
	0xe2, 0x24, 0x36, 0xb0, 0x63, 0x8d, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0x4c, 0x15, 0x56, 0x86,
	0x04, 0x01, 0x00, 0x00,
}

func (m *AutoSettleRecord) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AutoSettleRecord) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *AutoSettleRecord) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Addr) > 0 {
		i -= len(m.Addr)
		copy(dAtA[i:], m.Addr)
		i = encodeVarintAutoSettleRecord(dAtA, i, uint64(len(m.Addr)))
		i--
		dAtA[i] = 0x12
	}
	if m.Timestamp != 0 {
		i = encodeVarintAutoSettleRecord(dAtA, i, uint64(m.Timestamp))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintAutoSettleRecord(dAtA []byte, offset int, v uint64) int {
	offset -= sovAutoSettleRecord(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *AutoSettleRecord) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Timestamp != 0 {
		n += 1 + sovAutoSettleRecord(uint64(m.Timestamp))
	}
	l = len(m.Addr)
	if l > 0 {
		n += 1 + l + sovAutoSettleRecord(uint64(l))
	}
	return n
}

func sovAutoSettleRecord(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozAutoSettleRecord(x uint64) (n int) {
	return sovAutoSettleRecord(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *AutoSettleRecord) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAutoSettleRecord
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
			return fmt.Errorf("proto: AutoSettleRecord: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AutoSettleRecord: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Timestamp", wireType)
			}
			m.Timestamp = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAutoSettleRecord
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Timestamp |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Addr", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAutoSettleRecord
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
				return ErrInvalidLengthAutoSettleRecord
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAutoSettleRecord
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Addr = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipAutoSettleRecord(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthAutoSettleRecord
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
func skipAutoSettleRecord(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowAutoSettleRecord
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
					return 0, ErrIntOverflowAutoSettleRecord
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
					return 0, ErrIntOverflowAutoSettleRecord
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
				return 0, ErrInvalidLengthAutoSettleRecord
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupAutoSettleRecord
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthAutoSettleRecord
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthAutoSettleRecord        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowAutoSettleRecord          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupAutoSettleRecord = fmt.Errorf("proto: unexpected end of group")
)