// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: mempool.proto

package node

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
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

type TxType int32

const (
	// STANDARD indicates either phoenix transaction type or Smart contract calls
	TxType_STANDARD TxType = 0
	// DISTRIBUTE indicates the coinbase and reward distribution contract call
	TxType_DISTRIBUTE TxType = 1
	// WITHDRAWFEES indicates the Provisioners' withdraw contract call
	TxType_WITHDRAWFEES TxType = 2
	// BID transaction propagated by the Block Generator
	TxType_BID TxType = 3
	// STAKE transaction propagated by the Provisioners
	TxType_STAKE TxType = 4
	// SLASH transaction propagated by the consensus to punish the Committee
	// members when they turn byzantine
	TxType_SLASH TxType = 5
	// WITHDRAWSTAKE transaction propagated by the Provisioners to withdraw
	// their stake
	TxType_WITHDRAWSTAKE TxType = 6
	// WITHDRAWBID transaction propagated by the Block Generator to withdraw
	// their bids
	TxType_WITHDRAWBID TxType = 7
)

var TxType_name = map[int32]string{
	0: "STANDARD",
	1: "DISTRIBUTE",
	2: "WITHDRAWFEES",
	3: "BID",
	4: "STAKE",
	5: "SLASH",
	6: "WITHDRAWSTAKE",
	7: "WITHDRAWBID",
}

var TxType_value = map[string]int32{
	"STANDARD":      0,
	"DISTRIBUTE":    1,
	"WITHDRAWFEES":  2,
	"BID":           3,
	"STAKE":         4,
	"SLASH":         5,
	"WITHDRAWSTAKE": 6,
	"WITHDRAWBID":   7,
}

func (x TxType) String() string {
	return proto.EnumName(TxType_name, int32(x))
}

func (TxType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_a84c3667d8c2093a, []int{0}
}

type Tx struct {
	Type                 TxType   `protobuf:"varint,1,opt,name=type,proto3,enum=node.TxType" json:"type,omitempty"`
	Id                   string   `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	LockTime             uint64   `protobuf:"fixed64,3,opt,name=lock_time,json=lockTime,proto3" json:"lock_time,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Tx) Reset()         { *m = Tx{} }
func (m *Tx) String() string { return proto.CompactTextString(m) }
func (*Tx) ProtoMessage()    {}
func (*Tx) Descriptor() ([]byte, []int) {
	return fileDescriptor_a84c3667d8c2093a, []int{0}
}
func (m *Tx) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Tx.Unmarshal(m, b)
}
func (m *Tx) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Tx.Marshal(b, m, deterministic)
}
func (m *Tx) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Tx.Merge(m, src)
}
func (m *Tx) XXX_Size() int {
	return xxx_messageInfo_Tx.Size(m)
}
func (m *Tx) XXX_DiscardUnknown() {
	xxx_messageInfo_Tx.DiscardUnknown(m)
}

var xxx_messageInfo_Tx proto.InternalMessageInfo

func (m *Tx) GetType() TxType {
	if m != nil {
		return m.Type
	}
	return TxType_STANDARD
}

func (m *Tx) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Tx) GetLockTime() uint64 {
	if m != nil {
		return m.LockTime
	}
	return 0
}

// SelectRequest can specify an ID or a transaction type or none
type SelectRequest struct {
	Types                []TxType `protobuf:"varint,1,rep,packed,name=types,proto3,enum=node.TxType" json:"types,omitempty"`
	Id                   string   `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SelectRequest) Reset()         { *m = SelectRequest{} }
func (m *SelectRequest) String() string { return proto.CompactTextString(m) }
func (*SelectRequest) ProtoMessage()    {}
func (*SelectRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a84c3667d8c2093a, []int{1}
}
func (m *SelectRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SelectRequest.Unmarshal(m, b)
}
func (m *SelectRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SelectRequest.Marshal(b, m, deterministic)
}
func (m *SelectRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SelectRequest.Merge(m, src)
}
func (m *SelectRequest) XXX_Size() int {
	return xxx_messageInfo_SelectRequest.Size(m)
}
func (m *SelectRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SelectRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SelectRequest proto.InternalMessageInfo

func (m *SelectRequest) GetTypes() []TxType {
	if m != nil {
		return m.Types
	}
	return nil
}

func (m *SelectRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type SelectResponse struct {
	//Note: the response was a single string describing the transactions.
	//It should actually be responsibility of the caller to format the data or handle it otherwise
	//string msg = 1;
	Result               []*Tx    `protobuf:"bytes,1,rep,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SelectResponse) Reset()         { *m = SelectResponse{} }
func (m *SelectResponse) String() string { return proto.CompactTextString(m) }
func (*SelectResponse) ProtoMessage()    {}
func (*SelectResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a84c3667d8c2093a, []int{2}
}
func (m *SelectResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SelectResponse.Unmarshal(m, b)
}
func (m *SelectResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SelectResponse.Marshal(b, m, deterministic)
}
func (m *SelectResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SelectResponse.Merge(m, src)
}
func (m *SelectResponse) XXX_Size() int {
	return xxx_messageInfo_SelectResponse.Size(m)
}
func (m *SelectResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SelectResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SelectResponse proto.InternalMessageInfo

func (m *SelectResponse) GetResult() []*Tx {
	if m != nil {
		return m.Result
	}
	return nil
}

type GetUnconfirmedBalanceRequest struct {
	Vk                   []byte   `protobuf:"bytes,1,opt,name=vk,proto3" json:"vk,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetUnconfirmedBalanceRequest) Reset()         { *m = GetUnconfirmedBalanceRequest{} }
func (m *GetUnconfirmedBalanceRequest) String() string { return proto.CompactTextString(m) }
func (*GetUnconfirmedBalanceRequest) ProtoMessage()    {}
func (*GetUnconfirmedBalanceRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a84c3667d8c2093a, []int{3}
}
func (m *GetUnconfirmedBalanceRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetUnconfirmedBalanceRequest.Unmarshal(m, b)
}
func (m *GetUnconfirmedBalanceRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetUnconfirmedBalanceRequest.Marshal(b, m, deterministic)
}
func (m *GetUnconfirmedBalanceRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetUnconfirmedBalanceRequest.Merge(m, src)
}
func (m *GetUnconfirmedBalanceRequest) XXX_Size() int {
	return xxx_messageInfo_GetUnconfirmedBalanceRequest.Size(m)
}
func (m *GetUnconfirmedBalanceRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetUnconfirmedBalanceRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetUnconfirmedBalanceRequest proto.InternalMessageInfo

func (m *GetUnconfirmedBalanceRequest) GetVk() []byte {
	if m != nil {
		return m.Vk
	}
	return nil
}

func init() {
	proto.RegisterEnum("node.TxType", TxType_name, TxType_value)
	proto.RegisterType((*Tx)(nil), "node.Tx")
	proto.RegisterType((*SelectRequest)(nil), "node.SelectRequest")
	proto.RegisterType((*SelectResponse)(nil), "node.SelectResponse")
	proto.RegisterType((*GetUnconfirmedBalanceRequest)(nil), "node.GetUnconfirmedBalanceRequest")
}

func init() { proto.RegisterFile("mempool.proto", fileDescriptor_a84c3667d8c2093a) }

var fileDescriptor_a84c3667d8c2093a = []byte{
	// 346 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x91, 0x4d, 0xaf, 0x9a, 0x40,
	0x14, 0x86, 0x0b, 0x28, 0xea, 0x11, 0x29, 0x9d, 0x15, 0x49, 0xbb, 0x20, 0xac, 0x88, 0x49, 0x21,
	0xb1, 0xbf, 0x00, 0x0a, 0xad, 0xa4, 0x4d, 0x17, 0xc3, 0x18, 0x93, 0x6e, 0x1a, 0x85, 0xa3, 0x25,
	0x7c, 0x0c, 0x57, 0x06, 0xa3, 0xb9, 0x7f, 0xfe, 0x06, 0xbc, 0x6c, 0xee, 0xdd, 0x4d, 0xde, 0x27,
	0xe7, 0x39, 0xe7, 0xcd, 0xc0, 0xaa, 0xc2, 0xaa, 0xe1, 0xbc, 0x74, 0x9b, 0x0b, 0x17, 0x9c, 0x4c,
	0x6a, 0x9e, 0xa1, 0x9d, 0x80, 0xcc, 0x6e, 0xc4, 0x82, 0x89, 0xb8, 0x37, 0x68, 0x4a, 0x96, 0xe4,
	0xe8, 0x1b, 0xcd, 0xed, 0x91, 0xcb, 0x6e, 0xec, 0xde, 0x20, 0x1d, 0x08, 0xd1, 0x41, 0xce, 0x33,
	0x53, 0xb6, 0x24, 0x67, 0x41, 0xe5, 0x3c, 0x23, 0x9f, 0x61, 0x51, 0xf2, 0xb4, 0xf8, 0x27, 0xf2,
	0x0a, 0x4d, 0xc5, 0x92, 0x1c, 0x95, 0xce, 0xfb, 0x80, 0xe5, 0x15, 0xda, 0xdf, 0x61, 0x95, 0x60,
	0x89, 0xa9, 0xa0, 0xf8, 0xd4, 0x61, 0x2b, 0x88, 0x0d, 0xd3, 0xde, 0xd2, 0x9a, 0x92, 0xa5, 0xbc,
	0x5b, 0xf0, 0x40, 0x6f, 0x37, 0xd8, 0x1b, 0xd0, 0x47, 0x49, 0xdb, 0xf0, 0xba, 0x45, 0x62, 0x81,
	0x7a, 0xc1, 0xb6, 0x2b, 0xc5, 0xa0, 0x59, 0x6e, 0xe6, 0xa3, 0x86, 0xbe, 0xe6, 0xb6, 0x0b, 0x5f,
	0x7e, 0xa2, 0xd8, 0xd5, 0x29, 0xaf, 0x4f, 0xf9, 0xa5, 0xc2, 0x2c, 0x38, 0x94, 0x87, 0x3a, 0xc5,
	0xf1, 0x0e, 0x1d, 0xe4, 0x6b, 0x31, 0xb4, 0xd4, 0xa8, 0x7c, 0x2d, 0xd6, 0xcf, 0xa0, 0x3e, 0x8e,
	0x20, 0x1a, 0xcc, 0x13, 0xe6, 0xff, 0x09, 0x7d, 0x1a, 0x1a, 0x1f, 0x88, 0x0e, 0x10, 0xc6, 0x09,
	0xa3, 0x71, 0xb0, 0x63, 0x91, 0x21, 0x11, 0x03, 0xb4, 0x7d, 0xcc, 0xb6, 0x21, 0xf5, 0xf7, 0x3f,
	0xa2, 0x28, 0x31, 0x64, 0x32, 0x03, 0x25, 0x88, 0x43, 0x43, 0x21, 0x0b, 0x98, 0x26, 0xcc, 0xff,
	0x15, 0x19, 0x93, 0xe1, 0xf9, 0xdb, 0x4f, 0xb6, 0xc6, 0x94, 0x7c, 0x82, 0xd5, 0x38, 0xf0, 0xa0,
	0x2a, 0xf9, 0x08, 0xcb, 0x31, 0xea, 0x27, 0x67, 0xc1, 0xfa, 0xaf, 0x73, 0xce, 0xc5, 0xff, 0xee,
	0xe8, 0xa6, 0xbc, 0xf2, 0xb2, 0xae, 0x2d, 0xbe, 0x0e, 0x3f, 0x73, 0xec, 0x4e, 0xde, 0xa1, 0x13,
	0xfc, 0x8c, 0xb5, 0x77, 0xe6, 0x5e, 0xdf, 0xf1, 0xa8, 0x0e, 0xe4, 0xdb, 0x4b, 0x00, 0x00, 0x00,
	0xff, 0xff, 0x54, 0x2d, 0xe8, 0xab, 0xc4, 0x01, 0x00, 0x00,
}
