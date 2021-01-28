// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: bid.proto

package rusk

import (
	context "context"
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
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

// Bid representation
type Bid struct {
	EncryptedData        []byte          `protobuf:"bytes,1,opt,name=encrypted_data,json=encryptedData,proto3" json:"encrypted_data,omitempty"`
	HashedSecret         []byte          `protobuf:"bytes,2,opt,name=hashed_secret,json=hashedSecret,proto3" json:"hashed_secret,omitempty"`
	Nonce                []byte          `protobuf:"bytes,3,opt,name=nonce,proto3" json:"nonce,omitempty"`
	StealthAddress       *StealthAddress `protobuf:"bytes,4,opt,name=stealth_address,json=stealthAddress,proto3" json:"stealth_address,omitempty"`
	Commitment           []byte          `protobuf:"bytes,5,opt,name=commitment,proto3" json:"commitment,omitempty"`
	Elegibility          []byte          `protobuf:"bytes,6,opt,name=elegibility,proto3" json:"elegibility,omitempty"`
	Expiration           []byte          `protobuf:"bytes,7,opt,name=expiration,proto3" json:"expiration,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *Bid) Reset()         { *m = Bid{} }
func (m *Bid) String() string { return proto.CompactTextString(m) }
func (*Bid) ProtoMessage()    {}
func (*Bid) Descriptor() ([]byte, []int) {
	return fileDescriptor_25b9ce1844b80bbf, []int{0}
}
func (m *Bid) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Bid.Unmarshal(m, b)
}
func (m *Bid) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Bid.Marshal(b, m, deterministic)
}
func (m *Bid) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Bid.Merge(m, src)
}
func (m *Bid) XXX_Size() int {
	return xxx_messageInfo_Bid.Size(m)
}
func (m *Bid) XXX_DiscardUnknown() {
	xxx_messageInfo_Bid.DiscardUnknown(m)
}

var xxx_messageInfo_Bid proto.InternalMessageInfo

func (m *Bid) GetEncryptedData() []byte {
	if m != nil {
		return m.EncryptedData
	}
	return nil
}

func (m *Bid) GetHashedSecret() []byte {
	if m != nil {
		return m.HashedSecret
	}
	return nil
}

func (m *Bid) GetNonce() []byte {
	if m != nil {
		return m.Nonce
	}
	return nil
}

func (m *Bid) GetStealthAddress() *StealthAddress {
	if m != nil {
		return m.StealthAddress
	}
	return nil
}

func (m *Bid) GetCommitment() []byte {
	if m != nil {
		return m.Commitment
	}
	return nil
}

func (m *Bid) GetElegibility() []byte {
	if m != nil {
		return m.Elegibility
	}
	return nil
}

func (m *Bid) GetExpiration() []byte {
	if m != nil {
		return m.Expiration
	}
	return nil
}

// Used to Request the creation of a Bid
type BidTransactionRequest struct {
	K                    []byte          `protobuf:"bytes,1,opt,name=k,proto3" json:"k,omitempty"`
	Value                uint64          `protobuf:"varint,2,opt,name=value,proto3" json:"value,omitempty"`
	Secret               []byte          `protobuf:"bytes,3,opt,name=secret,proto3" json:"secret,omitempty"`
	StealthAddress       *StealthAddress `protobuf:"bytes,4,opt,name=stealth_address,json=stealthAddress,proto3" json:"stealth_address,omitempty"`
	Seed                 []byte          `protobuf:"bytes,5,opt,name=seed,proto3" json:"seed,omitempty"`
	LatestConsensusRound uint64          `protobuf:"fixed64,6,opt,name=latest_consensus_round,json=latestConsensusRound,proto3" json:"latest_consensus_round,omitempty"`
	LatestConsensusStep  uint64          `protobuf:"fixed64,7,opt,name=latest_consensus_step,json=latestConsensusStep,proto3" json:"latest_consensus_step,omitempty"`
	GasLimit             uint64          `protobuf:"fixed64,8,opt,name=gas_limit,json=gasLimit,proto3" json:"gas_limit,omitempty"`
	GasPrice             uint64          `protobuf:"fixed64,9,opt,name=gas_price,json=gasPrice,proto3" json:"gas_price,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *BidTransactionRequest) Reset()         { *m = BidTransactionRequest{} }
func (m *BidTransactionRequest) String() string { return proto.CompactTextString(m) }
func (*BidTransactionRequest) ProtoMessage()    {}
func (*BidTransactionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_25b9ce1844b80bbf, []int{1}
}
func (m *BidTransactionRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BidTransactionRequest.Unmarshal(m, b)
}
func (m *BidTransactionRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BidTransactionRequest.Marshal(b, m, deterministic)
}
func (m *BidTransactionRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BidTransactionRequest.Merge(m, src)
}
func (m *BidTransactionRequest) XXX_Size() int {
	return xxx_messageInfo_BidTransactionRequest.Size(m)
}
func (m *BidTransactionRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_BidTransactionRequest.DiscardUnknown(m)
}

var xxx_messageInfo_BidTransactionRequest proto.InternalMessageInfo

func (m *BidTransactionRequest) GetK() []byte {
	if m != nil {
		return m.K
	}
	return nil
}

func (m *BidTransactionRequest) GetValue() uint64 {
	if m != nil {
		return m.Value
	}
	return 0
}

func (m *BidTransactionRequest) GetSecret() []byte {
	if m != nil {
		return m.Secret
	}
	return nil
}

func (m *BidTransactionRequest) GetStealthAddress() *StealthAddress {
	if m != nil {
		return m.StealthAddress
	}
	return nil
}

func (m *BidTransactionRequest) GetSeed() []byte {
	if m != nil {
		return m.Seed
	}
	return nil
}

func (m *BidTransactionRequest) GetLatestConsensusRound() uint64 {
	if m != nil {
		return m.LatestConsensusRound
	}
	return 0
}

func (m *BidTransactionRequest) GetLatestConsensusStep() uint64 {
	if m != nil {
		return m.LatestConsensusStep
	}
	return 0
}

func (m *BidTransactionRequest) GetGasLimit() uint64 {
	if m != nil {
		return m.GasLimit
	}
	return 0
}

func (m *BidTransactionRequest) GetGasPrice() uint64 {
	if m != nil {
		return m.GasPrice
	}
	return 0
}

type BidTransaction struct {
	BidTreeStorageIndex  uint64       `protobuf:"varint,1,opt,name=bid_tree_storage_index,json=bidTreeStorageIndex,proto3" json:"bid_tree_storage_index,omitempty"`
	Tx                   *Transaction `protobuf:"bytes,2,opt,name=tx,proto3" json:"tx,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *BidTransaction) Reset()         { *m = BidTransaction{} }
func (m *BidTransaction) String() string { return proto.CompactTextString(m) }
func (*BidTransaction) ProtoMessage()    {}
func (*BidTransaction) Descriptor() ([]byte, []int) {
	return fileDescriptor_25b9ce1844b80bbf, []int{2}
}
func (m *BidTransaction) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BidTransaction.Unmarshal(m, b)
}
func (m *BidTransaction) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BidTransaction.Marshal(b, m, deterministic)
}
func (m *BidTransaction) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BidTransaction.Merge(m, src)
}
func (m *BidTransaction) XXX_Size() int {
	return xxx_messageInfo_BidTransaction.Size(m)
}
func (m *BidTransaction) XXX_DiscardUnknown() {
	xxx_messageInfo_BidTransaction.DiscardUnknown(m)
}

var xxx_messageInfo_BidTransaction proto.InternalMessageInfo

func (m *BidTransaction) GetBidTreeStorageIndex() uint64 {
	if m != nil {
		return m.BidTreeStorageIndex
	}
	return 0
}

func (m *BidTransaction) GetTx() *Transaction {
	if m != nil {
		return m.Tx
	}
	return nil
}

type FindBidRequest struct {
	StealthAddress       *StealthAddress `protobuf:"bytes,1,opt,name=stealth_address,json=stealthAddress,proto3" json:"stealth_address,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *FindBidRequest) Reset()         { *m = FindBidRequest{} }
func (m *FindBidRequest) String() string { return proto.CompactTextString(m) }
func (*FindBidRequest) ProtoMessage()    {}
func (*FindBidRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_25b9ce1844b80bbf, []int{3}
}
func (m *FindBidRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FindBidRequest.Unmarshal(m, b)
}
func (m *FindBidRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FindBidRequest.Marshal(b, m, deterministic)
}
func (m *FindBidRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FindBidRequest.Merge(m, src)
}
func (m *FindBidRequest) XXX_Size() int {
	return xxx_messageInfo_FindBidRequest.Size(m)
}
func (m *FindBidRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_FindBidRequest.DiscardUnknown(m)
}

var xxx_messageInfo_FindBidRequest proto.InternalMessageInfo

func (m *FindBidRequest) GetStealthAddress() *StealthAddress {
	if m != nil {
		return m.StealthAddress
	}
	return nil
}

type BidList struct {
	BidList              []*Bid   `protobuf:"bytes,1,rep,name=bid_list,json=bidList,proto3" json:"bid_list,omitempty"`
	BidHashList          [][]byte `protobuf:"bytes,2,rep,name=bid_hash_list,json=bidHashList,proto3" json:"bid_hash_list,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BidList) Reset()         { *m = BidList{} }
func (m *BidList) String() string { return proto.CompactTextString(m) }
func (*BidList) ProtoMessage()    {}
func (*BidList) Descriptor() ([]byte, []int) {
	return fileDescriptor_25b9ce1844b80bbf, []int{4}
}
func (m *BidList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BidList.Unmarshal(m, b)
}
func (m *BidList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BidList.Marshal(b, m, deterministic)
}
func (m *BidList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BidList.Merge(m, src)
}
func (m *BidList) XXX_Size() int {
	return xxx_messageInfo_BidList.Size(m)
}
func (m *BidList) XXX_DiscardUnknown() {
	xxx_messageInfo_BidList.DiscardUnknown(m)
}

var xxx_messageInfo_BidList proto.InternalMessageInfo

func (m *BidList) GetBidList() []*Bid {
	if m != nil {
		return m.BidList
	}
	return nil
}

func (m *BidList) GetBidHashList() [][]byte {
	if m != nil {
		return m.BidHashList
	}
	return nil
}

func init() {
	proto.RegisterType((*Bid)(nil), "rusk.Bid")
	proto.RegisterType((*BidTransactionRequest)(nil), "rusk.BidTransactionRequest")
	proto.RegisterType((*BidTransaction)(nil), "rusk.BidTransaction")
	proto.RegisterType((*FindBidRequest)(nil), "rusk.FindBidRequest")
	proto.RegisterType((*BidList)(nil), "rusk.BidList")
}

func init() { proto.RegisterFile("bid.proto", fileDescriptor_25b9ce1844b80bbf) }

var fileDescriptor_25b9ce1844b80bbf = []byte{
	// 580 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x54, 0x41, 0x6f, 0xd3, 0x30,
	0x14, 0x5e, 0xda, 0xae, 0x5d, 0x5f, 0xd7, 0xa2, 0x79, 0xdd, 0x14, 0x75, 0x12, 0x2a, 0x01, 0xa4,
	0x0a, 0x69, 0x1d, 0xea, 0xb8, 0x21, 0x0e, 0x14, 0x84, 0x40, 0x9a, 0x00, 0xa5, 0x9c, 0xb8, 0x44,
	0x4e, 0xfc, 0xd4, 0x58, 0x4d, 0x93, 0x62, 0xbf, 0x6c, 0xdd, 0x81, 0x3b, 0xbf, 0x8d, 0x5f, 0x85,
	0x6c, 0x67, 0xdd, 0xca, 0x76, 0x40, 0xe2, 0xe6, 0xf7, 0x7d, 0x9f, 0x5f, 0xfc, 0x7d, 0xf6, 0x0b,
	0xb4, 0x63, 0x29, 0xc6, 0x2b, 0x55, 0x50, 0xc1, 0x1a, 0xaa, 0xd4, 0x8b, 0xc1, 0x01, 0x29, 0x9e,
	0x6b, 0x9e, 0x90, 0x2c, 0x72, 0x47, 0x0c, 0x60, 0x81, 0xd7, 0xda, 0xad, 0x83, 0x5f, 0x35, 0xa8,
	0x4f, 0xa5, 0x60, 0xcf, 0xa1, 0x87, 0x79, 0xa2, 0xae, 0x57, 0x84, 0x22, 0x12, 0x9c, 0xb8, 0xef,
	0x0d, 0xbd, 0xd1, 0x7e, 0xd8, 0xdd, 0xa0, 0xef, 0x39, 0x71, 0xf6, 0x14, 0xba, 0x29, 0xd7, 0x29,
	0x8a, 0x48, 0x63, 0xa2, 0x90, 0xfc, 0x9a, 0x55, 0xed, 0x3b, 0x70, 0x66, 0x31, 0xd6, 0x87, 0xdd,
	0xbc, 0xc8, 0x13, 0xf4, 0xeb, 0x96, 0x74, 0x05, 0x7b, 0x03, 0x8f, 0x34, 0x21, 0xcf, 0x28, 0x8d,
	0xb8, 0x10, 0x0a, 0xb5, 0xf6, 0x1b, 0x43, 0x6f, 0xd4, 0x99, 0xf4, 0xc7, 0xe6, 0xa0, 0xe3, 0x99,
	0x23, 0xdf, 0x3a, 0x2e, 0xec, 0xe9, 0xad, 0x9a, 0x3d, 0x06, 0x48, 0x8a, 0xe5, 0x52, 0xd2, 0x12,
	0x73, 0xf2, 0x77, 0x6d, 0xe7, 0x3b, 0x08, 0x1b, 0x42, 0x07, 0x33, 0x9c, 0xcb, 0x58, 0x66, 0x92,
	0xae, 0xfd, 0xa6, 0x15, 0xdc, 0x85, 0x4c, 0x07, 0x5c, 0xaf, 0xa4, 0xe2, 0x26, 0x0a, 0xbf, 0xe5,
	0x3a, 0xdc, 0x22, 0xc1, 0xef, 0x1a, 0x1c, 0x4d, 0xa5, 0xf8, 0x76, 0x9b, 0x57, 0x88, 0x3f, 0x4a,
	0xd4, 0xc4, 0xf6, 0xc1, 0x5b, 0x54, 0x79, 0x78, 0x0b, 0x63, 0xef, 0x92, 0x67, 0x25, 0x5a, 0xef,
	0x8d, 0xd0, 0x15, 0xec, 0x18, 0x9a, 0x55, 0x24, 0xce, 0x75, 0x55, 0xfd, 0xaf, 0x6d, 0x06, 0x0d,
	0x8d, 0x28, 0x2a, 0xc3, 0x76, 0xcd, 0x5e, 0xc1, 0x71, 0xc6, 0x09, 0x35, 0x45, 0x49, 0x91, 0x6b,
	0xcc, 0x75, 0xa9, 0x23, 0x55, 0x94, 0xb9, 0xb0, 0xae, 0x9b, 0x61, 0xdf, 0xb1, 0xef, 0x6e, 0xc8,
	0xd0, 0x70, 0x6c, 0x02, 0x47, 0xf7, 0x76, 0x69, 0xc2, 0x95, 0x4d, 0xa2, 0x19, 0x1e, 0xfe, 0xb5,
	0x69, 0x46, 0xb8, 0x62, 0x27, 0xd0, 0x9e, 0x73, 0x1d, 0x65, 0x72, 0x29, 0xc9, 0xdf, 0xb3, 0xba,
	0xbd, 0x39, 0xd7, 0x17, 0xa6, 0xbe, 0x21, 0x57, 0x4a, 0x26, 0xe8, 0xb7, 0x37, 0xe4, 0x57, 0x53,
	0x07, 0x29, 0xf4, 0xb6, 0xb3, 0x64, 0xe7, 0x70, 0x1c, 0x4b, 0x11, 0x91, 0x42, 0x8c, 0x34, 0x15,
	0x8a, 0xcf, 0x31, 0x92, 0xb9, 0xc0, 0xb5, 0x4d, 0xb6, 0x11, 0x1e, 0xc6, 0x46, 0x8f, 0x38, 0x73,
	0xdc, 0x27, 0x43, 0xb1, 0x27, 0x50, 0xa3, 0xb5, 0x0d, 0xba, 0x33, 0x39, 0x70, 0x81, 0xdd, 0xbd,
	0x9f, 0x1a, 0xad, 0x83, 0x2f, 0xd0, 0xfb, 0x20, 0x73, 0x31, 0x95, 0xe2, 0xe6, 0xba, 0x1e, 0x88,
	0xdc, 0xfb, 0xf7, 0xc8, 0x83, 0x19, 0xb4, 0xa6, 0x52, 0x5c, 0x48, 0x4d, 0xec, 0x19, 0xec, 0x99,
	0x33, 0x67, 0x52, 0x93, 0xef, 0x0d, 0xeb, 0xa3, 0xce, 0xa4, 0xed, 0x5a, 0x98, 0xaf, 0xb5, 0xe2,
	0x4a, 0x15, 0x40, 0xd7, 0xa8, 0xcc, 0x0c, 0x38, 0x69, 0x6d, 0x58, 0x37, 0x8f, 0x2f, 0x96, 0xe2,
	0x23, 0xd7, 0xa9, 0xd1, 0x4c, 0x7e, 0x02, 0x4c, 0xa5, 0x98, 0xa1, 0xba, 0x94, 0x76, 0x16, 0x9a,
	0x9f, 0xf1, 0xca, 0xcc, 0xdd, 0xc9, 0xa6, 0xdf, 0xfd, 0x77, 0x37, 0xe8, 0x3f, 0x44, 0x06, 0x3b,
	0xec, 0x25, 0xb4, 0x2a, 0xcb, 0xac, 0x92, 0x6c, 0x27, 0x30, 0xe8, 0x6e, 0x36, 0x9a, 0x8f, 0x07,
	0x3b, 0xd3, 0x17, 0xdf, 0x47, 0x73, 0x49, 0x69, 0x19, 0x8f, 0x93, 0x62, 0x79, 0x26, 0x4a, 0xbd,
	0x38, 0xcd, 0x91, 0xae, 0x0a, 0xb5, 0x38, 0x33, 0xca, 0x53, 0x9d, 0xa4, 0xb8, 0xe4, 0xaf, 0xcd,
	0x3a, 0x6e, 0xda, 0x3f, 0xc3, 0xf9, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xa1, 0xb5, 0x31, 0xda,
	0x4b, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// BidServiceClient is the client API for BidService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type BidServiceClient interface {
	// Generate a new Bid and a proof of it's correctness.
	NewBid(ctx context.Context, in *BidTransactionRequest, opts ...grpc.CallOption) (*BidTransaction, error)
	// Look for your owned Bids and return a list of them and it's hash repr
	FindBid(ctx context.Context, in *FindBidRequest, opts ...grpc.CallOption) (*BidList, error)
}

type bidServiceClient struct {
	cc *grpc.ClientConn
}

func NewBidServiceClient(cc *grpc.ClientConn) BidServiceClient {
	return &bidServiceClient{cc}
}

func (c *bidServiceClient) NewBid(ctx context.Context, in *BidTransactionRequest, opts ...grpc.CallOption) (*BidTransaction, error) {
	out := new(BidTransaction)
	err := c.cc.Invoke(ctx, "/rusk.BidService/NewBid", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bidServiceClient) FindBid(ctx context.Context, in *FindBidRequest, opts ...grpc.CallOption) (*BidList, error) {
	out := new(BidList)
	err := c.cc.Invoke(ctx, "/rusk.BidService/FindBid", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BidServiceServer is the server API for BidService service.
type BidServiceServer interface {
	// Generate a new Bid and a proof of it's correctness.
	NewBid(context.Context, *BidTransactionRequest) (*BidTransaction, error)
	// Look for your owned Bids and return a list of them and it's hash repr
	FindBid(context.Context, *FindBidRequest) (*BidList, error)
}

// UnimplementedBidServiceServer can be embedded to have forward compatible implementations.
type UnimplementedBidServiceServer struct {
}

func (*UnimplementedBidServiceServer) NewBid(ctx context.Context, req *BidTransactionRequest) (*BidTransaction, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NewBid not implemented")
}
func (*UnimplementedBidServiceServer) FindBid(ctx context.Context, req *FindBidRequest) (*BidList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindBid not implemented")
}

func RegisterBidServiceServer(s *grpc.Server, srv BidServiceServer) {
	s.RegisterService(&_BidService_serviceDesc, srv)
}

func _BidService_NewBid_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BidTransactionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BidServiceServer).NewBid(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rusk.BidService/NewBid",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BidServiceServer).NewBid(ctx, req.(*BidTransactionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BidService_FindBid_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindBidRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BidServiceServer).FindBid(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rusk.BidService/FindBid",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BidServiceServer).FindBid(ctx, req.(*FindBidRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _BidService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "rusk.BidService",
	HandlerType: (*BidServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "NewBid",
			Handler:    _BidService_NewBid_Handler,
		},
		{
			MethodName: "FindBid",
			Handler:    _BidService_FindBid_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "bid.proto",
}
