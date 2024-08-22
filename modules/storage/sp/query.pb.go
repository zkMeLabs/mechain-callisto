package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
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

type QueryStorageProvidersRequest struct {
	// pagination defines an optional pagination for the request.
	Pagination *query.PageRequest `protobuf:"bytes,1,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (m *QueryStorageProvidersRequest) Reset()         { *m = QueryStorageProvidersRequest{} }
func (m *QueryStorageProvidersRequest) String() string { return proto.CompactTextString(m) }
func (*QueryStorageProvidersRequest) ProtoMessage()    {}
func (*QueryStorageProvidersRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_48dd9c8aad3b7a6d, []int{2}
}
func (m *QueryStorageProvidersRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryStorageProvidersRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryStorageProvidersRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryStorageProvidersRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryStorageProvidersRequest.Merge(m, src)
}
func (m *QueryStorageProvidersRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryStorageProvidersRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryStorageProvidersRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryStorageProvidersRequest proto.InternalMessageInfo

func (m *QueryStorageProvidersRequest) GetPagination() *query.PageRequest {
	if m != nil {
		return m.Pagination
	}
	return nil
}

type QueryStorageProvidersResponse struct {
	Sps []*StorageProvider `protobuf:"bytes,1,rep,name=sps,proto3" json:"sps,omitempty"`
	// pagination defines the pagination in the response.
	Pagination *query.PageResponse `protobuf:"bytes,2,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (m *QueryStorageProvidersResponse) Reset()         { *m = QueryStorageProvidersResponse{} }
func (m *QueryStorageProvidersResponse) String() string { return proto.CompactTextString(m) }
func (*QueryStorageProvidersResponse) ProtoMessage()    {}
func (*QueryStorageProvidersResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_48dd9c8aad3b7a6d, []int{3}
}
func (m *QueryStorageProvidersResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryStorageProvidersResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryStorageProvidersResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryStorageProvidersResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryStorageProvidersResponse.Merge(m, src)
}
func (m *QueryStorageProvidersResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryStorageProvidersResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryStorageProvidersResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryStorageProvidersResponse proto.InternalMessageInfo

func (m *QueryStorageProvidersResponse) GetSps() []*StorageProvider {
	if m != nil {
		return m.Sps
	}
	return nil
}

func (m *QueryStorageProvidersResponse) GetPagination() *query.PageResponse {
	if m != nil {
		return m.Pagination
	}
	return nil
}

type QuerySpStoragePriceRequest struct {
	// operator address of sp
	SpAddr string `protobuf:"bytes,1,opt,name=sp_addr,json=spAddr,proto3" json:"sp_addr,omitempty"`
}

func (m *QuerySpStoragePriceRequest) Reset()         { *m = QuerySpStoragePriceRequest{} }
func (m *QuerySpStoragePriceRequest) String() string { return proto.CompactTextString(m) }
func (*QuerySpStoragePriceRequest) ProtoMessage()    {}
func (*QuerySpStoragePriceRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_48dd9c8aad3b7a6d, []int{4}
}
func (m *QuerySpStoragePriceRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QuerySpStoragePriceRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QuerySpStoragePriceRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QuerySpStoragePriceRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QuerySpStoragePriceRequest.Merge(m, src)
}
func (m *QuerySpStoragePriceRequest) XXX_Size() int {
	return m.Size()
}
func (m *QuerySpStoragePriceRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QuerySpStoragePriceRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QuerySpStoragePriceRequest proto.InternalMessageInfo

func (m *QuerySpStoragePriceRequest) GetSpAddr() string {
	if m != nil {
		return m.SpAddr
	}
	return ""
}

type QuerySpStoragePriceResponse struct {
	SpStoragePrice SpStoragePrice `protobuf:"bytes,1,opt,name=sp_storage_price,json=spStoragePrice,proto3" json:"sp_storage_price"`
}

func (m *QuerySpStoragePriceResponse) Reset()         { *m = QuerySpStoragePriceResponse{} }
func (m *QuerySpStoragePriceResponse) String() string { return proto.CompactTextString(m) }
func (*QuerySpStoragePriceResponse) ProtoMessage()    {}
func (*QuerySpStoragePriceResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_48dd9c8aad3b7a6d, []int{5}
}
func (m *QuerySpStoragePriceResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QuerySpStoragePriceResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QuerySpStoragePriceResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QuerySpStoragePriceResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QuerySpStoragePriceResponse.Merge(m, src)
}
func (m *QuerySpStoragePriceResponse) XXX_Size() int {
	return m.Size()
}
func (m *QuerySpStoragePriceResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QuerySpStoragePriceResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QuerySpStoragePriceResponse proto.InternalMessageInfo

func (m *QuerySpStoragePriceResponse) GetSpStoragePrice() SpStoragePrice {
	if m != nil {
		return m.SpStoragePrice
	}
	return SpStoragePrice{}
}

type QueryGlobalSpStorePriceByTimeRequest struct {
	// unix timestamp in seconds. If it's 0, it will return the latest price.
	Timestamp int64 `protobuf:"varint,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
}

func (m *QueryGlobalSpStorePriceByTimeRequest) Reset()         { *m = QueryGlobalSpStorePriceByTimeRequest{} }
func (m *QueryGlobalSpStorePriceByTimeRequest) String() string { return proto.CompactTextString(m) }
func (*QueryGlobalSpStorePriceByTimeRequest) ProtoMessage()    {}
func (*QueryGlobalSpStorePriceByTimeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_48dd9c8aad3b7a6d, []int{6}
}
func (m *QueryGlobalSpStorePriceByTimeRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryGlobalSpStorePriceByTimeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryGlobalSpStorePriceByTimeRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryGlobalSpStorePriceByTimeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryGlobalSpStorePriceByTimeRequest.Merge(m, src)
}
func (m *QueryGlobalSpStorePriceByTimeRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryGlobalSpStorePriceByTimeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryGlobalSpStorePriceByTimeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryGlobalSpStorePriceByTimeRequest proto.InternalMessageInfo

func (m *QueryGlobalSpStorePriceByTimeRequest) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

type QueryGlobalSpStorePriceByTimeResponse struct {
	GlobalSpStorePrice GlobalSpStorePrice `protobuf:"bytes,1,opt,name=global_sp_store_price,json=globalSpStorePrice,proto3" json:"global_sp_store_price"`
}

func (m *QueryGlobalSpStorePriceByTimeResponse) Reset()         { *m = QueryGlobalSpStorePriceByTimeResponse{} }
func (m *QueryGlobalSpStorePriceByTimeResponse) String() string { return proto.CompactTextString(m) }
func (*QueryGlobalSpStorePriceByTimeResponse) ProtoMessage()    {}
func (*QueryGlobalSpStorePriceByTimeResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_48dd9c8aad3b7a6d, []int{7}
}
func (m *QueryGlobalSpStorePriceByTimeResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryGlobalSpStorePriceByTimeResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryGlobalSpStorePriceByTimeResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryGlobalSpStorePriceByTimeResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryGlobalSpStorePriceByTimeResponse.Merge(m, src)
}
func (m *QueryGlobalSpStorePriceByTimeResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryGlobalSpStorePriceByTimeResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryGlobalSpStorePriceByTimeResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryGlobalSpStorePriceByTimeResponse proto.InternalMessageInfo

func (m *QueryGlobalSpStorePriceByTimeResponse) GetGlobalSpStorePrice() GlobalSpStorePrice {
	if m != nil {
		return m.GlobalSpStorePrice
	}
	return GlobalSpStorePrice{}
}

type QueryStorageProviderRequest struct {
	Id uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (m *QueryStorageProviderRequest) Reset()         { *m = QueryStorageProviderRequest{} }
func (m *QueryStorageProviderRequest) String() string { return proto.CompactTextString(m) }
func (*QueryStorageProviderRequest) ProtoMessage()    {}
func (*QueryStorageProviderRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_48dd9c8aad3b7a6d, []int{8}
}
func (m *QueryStorageProviderRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryStorageProviderRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryStorageProviderRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryStorageProviderRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryStorageProviderRequest.Merge(m, src)
}
func (m *QueryStorageProviderRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryStorageProviderRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryStorageProviderRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryStorageProviderRequest proto.InternalMessageInfo

func (m *QueryStorageProviderRequest) GetId() uint32 {
	if m != nil {
		return m.Id
	}
	return 0
}

type QueryStorageProviderResponse struct {
	StorageProvider *StorageProvider `protobuf:"bytes,1,opt,name=storageProvider,proto3" json:"storageProvider,omitempty"`
}

func (m *QueryStorageProviderResponse) Reset()         { *m = QueryStorageProviderResponse{} }
func (m *QueryStorageProviderResponse) String() string { return proto.CompactTextString(m) }
func (*QueryStorageProviderResponse) ProtoMessage()    {}
func (*QueryStorageProviderResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_48dd9c8aad3b7a6d, []int{9}
}
func (m *QueryStorageProviderResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryStorageProviderResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryStorageProviderResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryStorageProviderResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryStorageProviderResponse.Merge(m, src)
}
func (m *QueryStorageProviderResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryStorageProviderResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryStorageProviderResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryStorageProviderResponse proto.InternalMessageInfo

func (m *QueryStorageProviderResponse) GetStorageProvider() *StorageProvider {
	if m != nil {
		return m.StorageProvider
	}
	return nil
}

type QueryStorageProviderByOperatorAddressRequest struct {
	OperatorAddress string `protobuf:"bytes,1,opt,name=operator_address,json=operatorAddress,proto3" json:"operator_address,omitempty"`
}

func (m *QueryStorageProviderByOperatorAddressRequest) Reset() {
	*m = QueryStorageProviderByOperatorAddressRequest{}
}
func (m *QueryStorageProviderByOperatorAddressRequest) String() string {
	return proto.CompactTextString(m)
}
func (*QueryStorageProviderByOperatorAddressRequest) ProtoMessage() {}
func (*QueryStorageProviderByOperatorAddressRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_48dd9c8aad3b7a6d, []int{10}
}
func (m *QueryStorageProviderByOperatorAddressRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryStorageProviderByOperatorAddressRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryStorageProviderByOperatorAddressRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryStorageProviderByOperatorAddressRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryStorageProviderByOperatorAddressRequest.Merge(m, src)
}
func (m *QueryStorageProviderByOperatorAddressRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryStorageProviderByOperatorAddressRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryStorageProviderByOperatorAddressRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryStorageProviderByOperatorAddressRequest proto.InternalMessageInfo

func (m *QueryStorageProviderByOperatorAddressRequest) GetOperatorAddress() string {
	if m != nil {
		return m.OperatorAddress
	}
	return ""
}

type QueryStorageProviderByOperatorAddressResponse struct {
	StorageProvider *StorageProvider `protobuf:"bytes,1,opt,name=storageProvider,proto3" json:"storageProvider,omitempty"`
}

func (m *QueryStorageProviderByOperatorAddressResponse) Reset() {
	*m = QueryStorageProviderByOperatorAddressResponse{}
}
func (m *QueryStorageProviderByOperatorAddressResponse) String() string {
	return proto.CompactTextString(m)
}
func (*QueryStorageProviderByOperatorAddressResponse) ProtoMessage() {}
func (*QueryStorageProviderByOperatorAddressResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_48dd9c8aad3b7a6d, []int{11}
}
func (m *QueryStorageProviderByOperatorAddressResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryStorageProviderByOperatorAddressResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryStorageProviderByOperatorAddressResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryStorageProviderByOperatorAddressResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryStorageProviderByOperatorAddressResponse.Merge(m, src)
}
func (m *QueryStorageProviderByOperatorAddressResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryStorageProviderByOperatorAddressResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryStorageProviderByOperatorAddressResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryStorageProviderByOperatorAddressResponse proto.InternalMessageInfo

func (m *QueryStorageProviderByOperatorAddressResponse) GetStorageProvider() *StorageProvider {
	if m != nil {
		return m.StorageProvider
	}
	return nil
}

type QueryStorageProviderMaintenanceRecordsRequest struct {
	OperatorAddress string `protobuf:"bytes,1,opt,name=operator_address,json=operatorAddress,proto3" json:"operator_address,omitempty"`
}

func (m *QueryStorageProviderMaintenanceRecordsRequest) Reset() {
	*m = QueryStorageProviderMaintenanceRecordsRequest{}
}
func (m *QueryStorageProviderMaintenanceRecordsRequest) String() string {
	return proto.CompactTextString(m)
}
func (*QueryStorageProviderMaintenanceRecordsRequest) ProtoMessage() {}
func (*QueryStorageProviderMaintenanceRecordsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_48dd9c8aad3b7a6d, []int{12}
}
func (m *QueryStorageProviderMaintenanceRecordsRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryStorageProviderMaintenanceRecordsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryStorageProviderMaintenanceRecordsRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryStorageProviderMaintenanceRecordsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryStorageProviderMaintenanceRecordsRequest.Merge(m, src)
}
func (m *QueryStorageProviderMaintenanceRecordsRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryStorageProviderMaintenanceRecordsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryStorageProviderMaintenanceRecordsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryStorageProviderMaintenanceRecordsRequest proto.InternalMessageInfo

func (m *QueryStorageProviderMaintenanceRecordsRequest) GetOperatorAddress() string {
	if m != nil {
		return m.OperatorAddress
	}
	return ""
}

type QueryStorageProviderMaintenanceRecordsResponse struct {
	Records []*MaintenanceRecord `protobuf:"bytes,1,rep,name=records,proto3" json:"records,omitempty"`
}

func (m *QueryStorageProviderMaintenanceRecordsResponse) Reset() {
	*m = QueryStorageProviderMaintenanceRecordsResponse{}
}
func (m *QueryStorageProviderMaintenanceRecordsResponse) String() string {
	return proto.CompactTextString(m)
}
func (*QueryStorageProviderMaintenanceRecordsResponse) ProtoMessage() {}
func (*QueryStorageProviderMaintenanceRecordsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_48dd9c8aad3b7a6d, []int{13}
}
func (m *QueryStorageProviderMaintenanceRecordsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryStorageProviderMaintenanceRecordsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryStorageProviderMaintenanceRecordsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryStorageProviderMaintenanceRecordsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryStorageProviderMaintenanceRecordsResponse.Merge(m, src)
}
func (m *QueryStorageProviderMaintenanceRecordsResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryStorageProviderMaintenanceRecordsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryStorageProviderMaintenanceRecordsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryStorageProviderMaintenanceRecordsResponse proto.InternalMessageInfo

func (m *QueryStorageProviderMaintenanceRecordsResponse) GetRecords() []*MaintenanceRecord {
	if m != nil {
		return m.Records
	}
	return nil
}

func init() {
	proto.RegisterType((*QueryStorageProvidersRequest)(nil), "greenfield.sp.QueryStorageProvidersRequest")
	proto.RegisterType((*QueryStorageProvidersResponse)(nil), "greenfield.sp.QueryStorageProvidersResponse")
	proto.RegisterType((*QuerySpStoragePriceRequest)(nil), "greenfield.sp.QuerySpStoragePriceRequest")
	proto.RegisterType((*QuerySpStoragePriceResponse)(nil), "greenfield.sp.QuerySpStoragePriceResponse")
	proto.RegisterType((*QueryGlobalSpStorePriceByTimeRequest)(nil), "greenfield.sp.QueryGlobalSpStorePriceByTimeRequest")
	proto.RegisterType((*QueryGlobalSpStorePriceByTimeResponse)(nil), "greenfield.sp.QueryGlobalSpStorePriceByTimeResponse")
	proto.RegisterType((*QueryStorageProviderRequest)(nil), "greenfield.sp.QueryStorageProviderRequest")
	proto.RegisterType((*QueryStorageProviderResponse)(nil), "greenfield.sp.QueryStorageProviderResponse")
	proto.RegisterType((*QueryStorageProviderByOperatorAddressRequest)(nil), "greenfield.sp.QueryStorageProviderByOperatorAddressRequest")
	proto.RegisterType((*QueryStorageProviderByOperatorAddressResponse)(nil), "greenfield.sp.QueryStorageProviderByOperatorAddressResponse")
	proto.RegisterType((*QueryStorageProviderMaintenanceRecordsRequest)(nil), "greenfield.sp.QueryStorageProviderMaintenanceRecordsRequest")
	proto.RegisterType((*QueryStorageProviderMaintenanceRecordsResponse)(nil), "greenfield.sp.QueryStorageProviderMaintenanceRecordsResponse")
}

func init() { proto.RegisterFile("greenfield/sp/query.proto", fileDescriptor_48dd9c8aad3b7a6d) }

var fileDescriptor_48dd9c8aad3b7a6d = []byte{
	// 874 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x56, 0x4f, 0x4f, 0x3b, 0x45,
	0x18, 0xee, 0x16, 0x2d, 0x61, 0x08, 0x7f, 0x32, 0x42, 0x90, 0x95, 0x16, 0x58, 0x45, 0x29, 0x7f,
	0x76, 0x6d, 0x2b, 0x89, 0x01, 0x8d, 0x4a, 0x8c, 0xa8, 0x09, 0x01, 0x57, 0x2f, 0x72, 0xd9, 0x4c,
	0xbb, 0xc3, 0xb2, 0x49, 0xbb, 0x33, 0xec, 0x2c, 0x8d, 0x0d, 0xe1, 0x82, 0x5f, 0xc0, 0x44, 0xbd,
	0x78, 0xf1, 0xeb, 0x70, 0xf0, 0x40, 0xe2, 0xc5, 0x78, 0x30, 0x06, 0xfc, 0x12, 0xde, 0xcc, 0xce,
	0xcc, 0xb6, 0xee, 0x74, 0x4b, 0xf7, 0x47, 0x7e, 0x17, 0xc2, 0xce, 0xfb, 0x3e, 0xef, 0xfb, 0x3c,
	0xef, 0xbc, 0xfb, 0x74, 0xc1, 0xb2, 0x17, 0x62, 0x1c, 0x9c, 0xfb, 0xb8, 0xed, 0x5a, 0x8c, 0x5a,
	0x97, 0x57, 0x38, 0xec, 0x99, 0x34, 0x24, 0x11, 0x81, 0x33, 0x83, 0x90, 0xc9, 0xa8, 0xbe, 0xd5,
	0x22, 0xac, 0x43, 0x98, 0xd5, 0x44, 0x0c, 0x8b, 0x3c, 0xab, 0x5b, 0x6b, 0xe2, 0x08, 0xd5, 0x2c,
	0x8a, 0x3c, 0x3f, 0x40, 0x91, 0x4f, 0x02, 0x01, 0xd5, 0x97, 0x45, 0xae, 0xc3, 0x9f, 0x2c, 0xf1,
	0x20, 0x43, 0x0b, 0x1e, 0xf1, 0x88, 0x38, 0x8f, 0xff, 0x93, 0xa7, 0x2b, 0x1e, 0x21, 0x5e, 0x1b,
	0x5b, 0x88, 0xfa, 0x16, 0x0a, 0x02, 0x12, 0xf1, 0x6a, 0x09, 0x46, 0x4f, 0x93, 0xa4, 0x28, 0x44,
	0x9d, 0x24, 0xa6, 0x08, 0x88, 0x7a, 0x14, 0xcb, 0x90, 0xb1, 0x00, 0xe0, 0x57, 0x31, 0xcf, 0x53,
	0x9e, 0x6f, 0xe3, 0xcb, 0x2b, 0xcc, 0x22, 0xe3, 0x4b, 0xf0, 0x5a, 0xea, 0x94, 0x51, 0x12, 0x30,
	0x0c, 0x1b, 0xa0, 0x24, 0xea, 0xbe, 0xae, 0xad, 0x69, 0x9b, 0xd3, 0xf5, 0x45, 0x33, 0x25, 0xdf,
	0x14, 0xe9, 0x87, 0xaf, 0xdc, 0xfd, 0xb5, 0x5a, 0xb0, 0x65, 0xaa, 0x71, 0x0e, 0x56, 0x78, 0xad,
	0xaf, 0x23, 0x12, 0x22, 0x0f, 0x9f, 0x86, 0xa4, 0xeb, 0xbb, 0x38, 0x4c, 0x7a, 0xc1, 0xcf, 0x00,
	0x18, 0xcc, 0x46, 0x16, 0x7e, 0xdb, 0x94, 0xf3, 0x88, 0x07, 0x69, 0x8a, 0x81, 0xcb, 0x41, 0x9a,
	0xa7, 0xc8, 0xc3, 0x12, 0x6b, 0xff, 0x0f, 0x69, 0xfc, 0xa2, 0x81, 0xf2, 0x88, 0x46, 0x92, 0xfe,
	0xbb, 0x60, 0x82, 0xd1, 0x98, 0xfb, 0xc4, 0xe6, 0x74, 0xbd, 0xa2, 0x70, 0x57, 0x50, 0x76, 0x9c,
	0x0a, 0x8f, 0x52, 0xdc, 0x8a, 0x9c, 0xdb, 0x3b, 0x63, 0xb9, 0x89, 0x76, 0x29, 0x72, 0x7b, 0x40,
	0x17, 0xdc, 0x68, 0xbf, 0x8f, 0xdf, 0x4a, 0x64, 0xc0, 0x25, 0x30, 0xc9, 0xa8, 0x83, 0x5c, 0x37,
	0xe4, 0xfa, 0xa7, 0xec, 0x12, 0xa3, 0x9f, 0xb8, 0x6e, 0x68, 0xb4, 0xc1, 0x1b, 0x99, 0x30, 0x29,
	0xe8, 0x18, 0xcc, 0x33, 0xea, 0x30, 0x11, 0x72, 0x68, 0x1c, 0x93, 0x03, 0x2c, 0xab, 0xea, 0x52,
	0x05, 0xe4, 0x0d, 0xcd, 0xb2, 0xd4, 0xa9, 0xf1, 0x29, 0x78, 0x8b, 0x77, 0x3b, 0x6a, 0x93, 0x26,
	0x6a, 0x0b, 0x88, 0x04, 0xf4, 0xbe, 0xf1, 0x3b, 0x7d, 0xba, 0x2b, 0x60, 0x2a, 0xf2, 0x3b, 0x98,
	0x45, 0xa8, 0x43, 0x79, 0xbf, 0x09, 0x7b, 0x70, 0x60, 0x7c, 0xaf, 0x81, 0x8d, 0x31, 0x65, 0x24,
	0xfd, 0x33, 0xb0, 0xe8, 0xf1, 0x1c, 0x47, 0xaa, 0x48, 0x6b, 0x58, 0x57, 0x34, 0x64, 0xd4, 0x13,
	0x3a, 0xa0, 0x37, 0x14, 0x31, 0x76, 0x93, 0xc9, 0x29, 0xd7, 0x2a, 0x25, 0xcc, 0x82, 0xa2, 0xef,
	0xf2, 0x3e, 0x33, 0x76, 0xd1, 0x77, 0x8d, 0x8b, 0xec, 0x25, 0xed, 0x53, 0xfd, 0x1c, 0xcc, 0xb1,
	0x74, 0x48, 0x92, 0x1c, 0xb7, 0x46, 0x2a, 0xcc, 0xf8, 0x16, 0xec, 0x64, 0x75, 0x3a, 0xec, 0x9d,
	0x50, 0x1c, 0xa2, 0x88, 0x84, 0xf1, 0xc5, 0x63, 0xd6, 0x7f, 0x3d, 0xaa, 0x60, 0x9e, 0xc8, 0x08,
	0xdf, 0x10, 0xcc, 0x98, 0x5c, 0x92, 0x39, 0x92, 0x46, 0x18, 0x3d, 0xb0, 0x9b, 0xb3, 0xf4, 0x4b,
	0x57, 0x75, 0x96, 0xdd, 0xfa, 0x18, 0xf9, 0x41, 0x84, 0x03, 0x14, 0xc4, 0x4b, 0xdb, 0x22, 0xa1,
	0xfb, 0x1c, 0x59, 0x6d, 0x60, 0xe6, 0xad, 0x2d, 0x75, 0xed, 0x83, 0xc9, 0x50, 0x1c, 0xc9, 0x97,
	0x7d, 0x4d, 0xd1, 0x33, 0x84, 0xb5, 0x13, 0x40, 0xfd, 0xdf, 0x29, 0xf0, 0x2a, 0x6f, 0x07, 0x03,
	0x50, 0x12, 0x86, 0x06, 0xd5, 0x4d, 0x1c, 0x76, 0x4c, 0xdd, 0x78, 0x2a, 0x45, 0xd0, 0x32, 0xca,
	0xb7, 0xbf, 0xff, 0xf3, 0x63, 0x71, 0x09, 0x2e, 0x5a, 0x59, 0x5e, 0x0d, 0x7f, 0xd2, 0xc0, 0xbc,
	0xea, 0x5d, 0x70, 0x3b, 0xab, 0xee, 0x08, 0x2b, 0xd5, 0x77, 0xf2, 0x25, 0x4b, 0x3a, 0x1b, 0x9c,
	0xce, 0x2a, 0x2c, 0xa7, 0xe8, 0xf4, 0xcd, 0x24, 0x61, 0xf0, 0xab, 0x26, 0x7f, 0x0c, 0xd2, 0x1e,
	0x02, 0xab, 0x99, 0xcd, 0xb2, 0xfc, 0x4d, 0xdf, 0xca, 0x93, 0x2a, 0x59, 0xd5, 0x38, 0xab, 0x6d,
	0x58, 0x55, 0x86, 0xa4, 0x1a, 0x9d, 0x75, 0x2d, 0x2d, 0xf3, 0x06, 0xfe, 0x96, 0x38, 0xff, 0x28,
	0xc7, 0x81, 0x8d, 0x2c, 0x02, 0x63, 0x6c, 0x4e, 0x7f, 0xef, 0xc5, 0x40, 0x92, 0xff, 0xc7, 0x9c,
	0xff, 0x3e, 0x7c, 0x5f, 0xe1, 0x9f, 0xe9, 0x74, 0x4e, 0xb3, 0xe7, 0xc4, 0xce, 0x69, 0x5d, 0xf7,
	0xfd, 0xf3, 0x06, 0xfe, 0xac, 0x81, 0x39, 0xe5, 0xd2, 0xe0, 0x56, 0x8e, 0x9b, 0x4d, 0x78, 0x6f,
	0xe7, 0xca, 0x95, 0x74, 0xab, 0x9c, 0xee, 0x9b, 0x70, 0xfd, 0xa9, 0x25, 0xb0, 0xae, 0x7d, 0xf7,
	0x06, 0xfe, 0xa9, 0x81, 0xb5, 0x71, 0xd6, 0x02, 0x0f, 0x72, 0x34, 0x1f, 0xe5, 0x75, 0xfa, 0x07,
	0xcf, 0x03, 0x4b, 0x29, 0x07, 0x5c, 0xca, 0x1e, 0x6c, 0xa8, 0x9b, 0xa3, 0xa8, 0x89, 0x87, 0xae,
	0x7a, 0x0f, 0xbc, 0x2d, 0x82, 0xfa, 0x58, 0x83, 0x19, 0x96, 0x9b, 0x87, 0xf1, 0x48, 0x13, 0xd4,
	0x3f, 0x7c, 0x26, 0x5a, 0x0a, 0x3e, 0xe1, 0x82, 0xbf, 0x80, 0x47, 0xe3, 0x04, 0x77, 0x06, 0x35,
	0x1c, 0xe9, 0x73, 0x59, 0x43, 0x38, 0xfc, 0xe8, 0xee, 0xa1, 0xa2, 0xdd, 0x3f, 0x54, 0xb4, 0xbf,
	0x1f, 0x2a, 0xda, 0x0f, 0x8f, 0x95, 0xc2, 0xfd, 0x63, 0xa5, 0xf0, 0xc7, 0x63, 0xa5, 0x70, 0xb6,
	0xe1, 0xf9, 0xd1, 0xc5, 0x55, 0xd3, 0x6c, 0x91, 0x8e, 0x85, 0xbb, 0xf1, 0x27, 0xae, 0xf8, 0xdb,
	0xad, 0xd5, 0xad, 0xef, 0xfa, 0xdf, 0x94, 0xcd, 0x12, 0xff, 0xa8, 0x6c, 0xfc, 0x17, 0x00, 0x00,
	0xff, 0xff, 0xa5, 0xa7, 0x1e, 0x82, 0x32, 0x0b, 0x00, 0x00,
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
	// Queries a list of GetStorageProviders items.
	StorageProviders(ctx context.Context, in *QueryStorageProvidersRequest, opts ...grpc.CallOption) (*QueryStorageProvidersResponse, error)
	// get the latest storage price of specific sp
	QuerySpStoragePrice(ctx context.Context, in *QuerySpStoragePriceRequest, opts ...grpc.CallOption) (*QuerySpStoragePriceResponse, error)
	// get global store price by time
	QueryGlobalSpStorePriceByTime(ctx context.Context, in *QueryGlobalSpStorePriceByTimeRequest, opts ...grpc.CallOption) (*QueryGlobalSpStorePriceByTimeResponse, error)
	// Queries a storage provider with specify id
	StorageProvider(ctx context.Context, in *QueryStorageProviderRequest, opts ...grpc.CallOption) (*QueryStorageProviderResponse, error)
	// Queries a StorageProvider by specify operator address.
	StorageProviderByOperatorAddress(ctx context.Context, in *QueryStorageProviderByOperatorAddressRequest, opts ...grpc.CallOption) (*QueryStorageProviderByOperatorAddressResponse, error)
	// Queries a StorageProvider by specify operator address.
	StorageProviderMaintenanceRecordsByOperatorAddress(ctx context.Context, in *QueryStorageProviderMaintenanceRecordsRequest, opts ...grpc.CallOption) (*QueryStorageProviderMaintenanceRecordsResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) StorageProviders(ctx context.Context, in *QueryStorageProvidersRequest, opts ...grpc.CallOption) (*QueryStorageProvidersResponse, error) {
	out := new(QueryStorageProvidersResponse)
	err := c.cc.Invoke(ctx, "/greenfield.sp.Query/StorageProviders", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) QuerySpStoragePrice(ctx context.Context, in *QuerySpStoragePriceRequest, opts ...grpc.CallOption) (*QuerySpStoragePriceResponse, error) {
	out := new(QuerySpStoragePriceResponse)
	err := c.cc.Invoke(ctx, "/greenfield.sp.Query/QuerySpStoragePrice", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) QueryGlobalSpStorePriceByTime(ctx context.Context, in *QueryGlobalSpStorePriceByTimeRequest, opts ...grpc.CallOption) (*QueryGlobalSpStorePriceByTimeResponse, error) {
	out := new(QueryGlobalSpStorePriceByTimeResponse)
	err := c.cc.Invoke(ctx, "/greenfield.sp.Query/QueryGlobalSpStorePriceByTime", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) StorageProvider(ctx context.Context, in *QueryStorageProviderRequest, opts ...grpc.CallOption) (*QueryStorageProviderResponse, error) {
	out := new(QueryStorageProviderResponse)
	err := c.cc.Invoke(ctx, "/greenfield.sp.Query/StorageProvider", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) StorageProviderByOperatorAddress(ctx context.Context, in *QueryStorageProviderByOperatorAddressRequest, opts ...grpc.CallOption) (*QueryStorageProviderByOperatorAddressResponse, error) {
	out := new(QueryStorageProviderByOperatorAddressResponse)
	err := c.cc.Invoke(ctx, "/greenfield.sp.Query/StorageProviderByOperatorAddress", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) StorageProviderMaintenanceRecordsByOperatorAddress(ctx context.Context, in *QueryStorageProviderMaintenanceRecordsRequest, opts ...grpc.CallOption) (*QueryStorageProviderMaintenanceRecordsResponse, error) {
	out := new(QueryStorageProviderMaintenanceRecordsResponse)
	err := c.cc.Invoke(ctx, "/greenfield.sp.Query/StorageProviderMaintenanceRecordsByOperatorAddress", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	// Queries a list of GetStorageProviders items.
	StorageProviders(context.Context, *QueryStorageProvidersRequest) (*QueryStorageProvidersResponse, error)
	// get the latest storage price of specific sp
	QuerySpStoragePrice(context.Context, *QuerySpStoragePriceRequest) (*QuerySpStoragePriceResponse, error)
	// get global store price by time
	QueryGlobalSpStorePriceByTime(context.Context, *QueryGlobalSpStorePriceByTimeRequest) (*QueryGlobalSpStorePriceByTimeResponse, error)
	// Queries a storage provider with specify id
	StorageProvider(context.Context, *QueryStorageProviderRequest) (*QueryStorageProviderResponse, error)
	// Queries a StorageProvider by specify operator address.
	StorageProviderByOperatorAddress(context.Context, *QueryStorageProviderByOperatorAddressRequest) (*QueryStorageProviderByOperatorAddressResponse, error)
	// Queries a StorageProvider by specify operator address.
	StorageProviderMaintenanceRecordsByOperatorAddress(context.Context, *QueryStorageProviderMaintenanceRecordsRequest) (*QueryStorageProviderMaintenanceRecordsResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) StorageProviders(ctx context.Context, req *QueryStorageProvidersRequest) (*QueryStorageProvidersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StorageProviders not implemented")
}
func (*UnimplementedQueryServer) QuerySpStoragePrice(ctx context.Context, req *QuerySpStoragePriceRequest) (*QuerySpStoragePriceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QuerySpStoragePrice not implemented")
}
func (*UnimplementedQueryServer) QueryGlobalSpStorePriceByTime(ctx context.Context, req *QueryGlobalSpStorePriceByTimeRequest) (*QueryGlobalSpStorePriceByTimeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryGlobalSpStorePriceByTime not implemented")
}
func (*UnimplementedQueryServer) StorageProvider(ctx context.Context, req *QueryStorageProviderRequest) (*QueryStorageProviderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StorageProvider not implemented")
}
func (*UnimplementedQueryServer) StorageProviderByOperatorAddress(ctx context.Context, req *QueryStorageProviderByOperatorAddressRequest) (*QueryStorageProviderByOperatorAddressResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StorageProviderByOperatorAddress not implemented")
}
func (*UnimplementedQueryServer) StorageProviderMaintenanceRecordsByOperatorAddress(ctx context.Context, req *QueryStorageProviderMaintenanceRecordsRequest) (*QueryStorageProviderMaintenanceRecordsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StorageProviderMaintenanceRecordsByOperatorAddress not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_StorageProviders_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryStorageProvidersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).StorageProviders(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/greenfield.sp.Query/StorageProviders",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).StorageProviders(ctx, req.(*QueryStorageProvidersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_QuerySpStoragePrice_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QuerySpStoragePriceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).QuerySpStoragePrice(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/greenfield.sp.Query/QuerySpStoragePrice",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).QuerySpStoragePrice(ctx, req.(*QuerySpStoragePriceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_QueryGlobalSpStorePriceByTime_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryGlobalSpStorePriceByTimeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).QueryGlobalSpStorePriceByTime(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/greenfield.sp.Query/QueryGlobalSpStorePriceByTime",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).QueryGlobalSpStorePriceByTime(ctx, req.(*QueryGlobalSpStorePriceByTimeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_StorageProvider_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryStorageProviderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).StorageProvider(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/greenfield.sp.Query/StorageProvider",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).StorageProvider(ctx, req.(*QueryStorageProviderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_StorageProviderByOperatorAddress_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryStorageProviderByOperatorAddressRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).StorageProviderByOperatorAddress(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/greenfield.sp.Query/StorageProviderByOperatorAddress",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).StorageProviderByOperatorAddress(ctx, req.(*QueryStorageProviderByOperatorAddressRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_StorageProviderMaintenanceRecordsByOperatorAddress_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryStorageProviderMaintenanceRecordsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).StorageProviderMaintenanceRecordsByOperatorAddress(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/greenfield.sp.Query/StorageProviderMaintenanceRecordsByOperatorAddress",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).StorageProviderMaintenanceRecordsByOperatorAddress(ctx, req.(*QueryStorageProviderMaintenanceRecordsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "greenfield.sp.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "StorageProviders",
			Handler:    _Query_StorageProviders_Handler,
		},
		{
			MethodName: "QuerySpStoragePrice",
			Handler:    _Query_QuerySpStoragePrice_Handler,
		},
		{
			MethodName: "QueryGlobalSpStorePriceByTime",
			Handler:    _Query_QueryGlobalSpStorePriceByTime_Handler,
		},
		{
			MethodName: "StorageProvider",
			Handler:    _Query_StorageProvider_Handler,
		},
		{
			MethodName: "StorageProviderByOperatorAddress",
			Handler:    _Query_StorageProviderByOperatorAddress_Handler,
		},
		{
			MethodName: "StorageProviderMaintenanceRecordsByOperatorAddress",
			Handler:    _Query_StorageProviderMaintenanceRecordsByOperatorAddress_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "greenfield/sp/query.proto",
}

func (m *QueryStorageProvidersRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryStorageProvidersRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryStorageProvidersRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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

func (m *QueryStorageProvidersResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryStorageProvidersResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryStorageProvidersResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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
	if len(m.Sps) > 0 {
		for iNdEx := len(m.Sps) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Sps[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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

func (m *QuerySpStoragePriceRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QuerySpStoragePriceRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QuerySpStoragePriceRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.SpAddr) > 0 {
		i -= len(m.SpAddr)
		copy(dAtA[i:], m.SpAddr)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.SpAddr)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QuerySpStoragePriceResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QuerySpStoragePriceResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QuerySpStoragePriceResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.SpStoragePrice.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintQuery(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *QueryGlobalSpStorePriceByTimeRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryGlobalSpStorePriceByTimeRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryGlobalSpStorePriceByTimeRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Timestamp != 0 {
		i = encodeVarintQuery(dAtA, i, uint64(m.Timestamp))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *QueryGlobalSpStorePriceByTimeResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryGlobalSpStorePriceByTimeResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryGlobalSpStorePriceByTimeResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.GlobalSpStorePrice.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintQuery(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *QueryStorageProviderRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryStorageProviderRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryStorageProviderRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Id != 0 {
		i = encodeVarintQuery(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *QueryStorageProviderResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryStorageProviderResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryStorageProviderResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.StorageProvider != nil {
		{
			size, err := m.StorageProvider.MarshalToSizedBuffer(dAtA[:i])
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

func (m *QueryStorageProviderByOperatorAddressRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryStorageProviderByOperatorAddressRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryStorageProviderByOperatorAddressRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.OperatorAddress) > 0 {
		i -= len(m.OperatorAddress)
		copy(dAtA[i:], m.OperatorAddress)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.OperatorAddress)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryStorageProviderByOperatorAddressResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryStorageProviderByOperatorAddressResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryStorageProviderByOperatorAddressResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.StorageProvider != nil {
		{
			size, err := m.StorageProvider.MarshalToSizedBuffer(dAtA[:i])
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

func (m *QueryStorageProviderMaintenanceRecordsRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryStorageProviderMaintenanceRecordsRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryStorageProviderMaintenanceRecordsRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.OperatorAddress) > 0 {
		i -= len(m.OperatorAddress)
		copy(dAtA[i:], m.OperatorAddress)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.OperatorAddress)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryStorageProviderMaintenanceRecordsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryStorageProviderMaintenanceRecordsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryStorageProviderMaintenanceRecordsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Records) > 0 {
		for iNdEx := len(m.Records) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Records[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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

func (m *QueryStorageProvidersRequest) Size() (n int) {
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

func (m *QueryStorageProvidersResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Sps) > 0 {
		for _, e := range m.Sps {
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

func (m *QuerySpStoragePriceRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.SpAddr)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QuerySpStoragePriceResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.SpStoragePrice.Size()
	n += 1 + l + sovQuery(uint64(l))
	return n
}

func (m *QueryGlobalSpStorePriceByTimeRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Timestamp != 0 {
		n += 1 + sovQuery(uint64(m.Timestamp))
	}
	return n
}

func (m *QueryGlobalSpStorePriceByTimeResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.GlobalSpStorePrice.Size()
	n += 1 + l + sovQuery(uint64(l))
	return n
}

func (m *QueryStorageProviderRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovQuery(uint64(m.Id))
	}
	return n
}

func (m *QueryStorageProviderResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.StorageProvider != nil {
		l = m.StorageProvider.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryStorageProviderByOperatorAddressRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.OperatorAddress)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryStorageProviderByOperatorAddressResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.StorageProvider != nil {
		l = m.StorageProvider.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryStorageProviderMaintenanceRecordsRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.OperatorAddress)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryStorageProviderMaintenanceRecordsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Records) > 0 {
		for _, e := range m.Records {
			l = e.Size()
			n += 1 + l + sovQuery(uint64(l))
		}
	}
	return n
}

func sovQuery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuery(x uint64) (n int) {
	return sovQuery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}

func (m *QueryStorageProvidersRequest) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryStorageProvidersRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryStorageProvidersRequest: illegal tag %d (wire type %d)", fieldNum, wire)
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
func (m *QueryStorageProvidersResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryStorageProvidersResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryStorageProvidersResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sps", wireType)
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
			m.Sps = append(m.Sps, &StorageProvider{})
			if err := m.Sps[len(m.Sps)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
func (m *QuerySpStoragePriceRequest) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QuerySpStoragePriceRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QuerySpStoragePriceRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SpAddr", wireType)
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
			m.SpAddr = string(dAtA[iNdEx:postIndex])
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
func (m *QuerySpStoragePriceResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QuerySpStoragePriceResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QuerySpStoragePriceResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SpStoragePrice", wireType)
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
			if err := m.SpStoragePrice.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
func (m *QueryGlobalSpStorePriceByTimeRequest) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryGlobalSpStorePriceByTimeRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryGlobalSpStorePriceByTimeRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Timestamp", wireType)
			}
			m.Timestamp = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
func (m *QueryGlobalSpStorePriceByTimeResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryGlobalSpStorePriceByTimeResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryGlobalSpStorePriceByTimeResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GlobalSpStorePrice", wireType)
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
			if err := m.GlobalSpStorePrice.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
func (m *QueryStorageProviderRequest) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryStorageProviderRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryStorageProviderRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= uint32(b&0x7F) << shift
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
func (m *QueryStorageProviderResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryStorageProviderResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryStorageProviderResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StorageProvider", wireType)
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
			if m.StorageProvider == nil {
				m.StorageProvider = &StorageProvider{}
			}
			if err := m.StorageProvider.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
func (m *QueryStorageProviderByOperatorAddressRequest) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryStorageProviderByOperatorAddressRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryStorageProviderByOperatorAddressRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OperatorAddress", wireType)
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
			m.OperatorAddress = string(dAtA[iNdEx:postIndex])
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
func (m *QueryStorageProviderByOperatorAddressResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryStorageProviderByOperatorAddressResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryStorageProviderByOperatorAddressResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StorageProvider", wireType)
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
			if m.StorageProvider == nil {
				m.StorageProvider = &StorageProvider{}
			}
			if err := m.StorageProvider.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
func (m *QueryStorageProviderMaintenanceRecordsRequest) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryStorageProviderMaintenanceRecordsRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryStorageProviderMaintenanceRecordsRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OperatorAddress", wireType)
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
			m.OperatorAddress = string(dAtA[iNdEx:postIndex])
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
func (m *QueryStorageProviderMaintenanceRecordsResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryStorageProviderMaintenanceRecordsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryStorageProviderMaintenanceRecordsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Records", wireType)
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
			m.Records = append(m.Records, &MaintenanceRecord{})
			if err := m.Records[len(m.Records)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
