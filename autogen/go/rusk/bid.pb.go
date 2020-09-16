// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.24.0
// 	protoc        v3.11.4
// source: bid.proto

package _

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// Bid representation
type Bid struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EncryptedData       *PoseidonCipher   `protobuf:"bytes,1,opt,name=encrypted_data,json=encryptedData,proto3" json:"encrypted_data,omitempty"`
	HashedSecret        *BlsScalar        `protobuf:"bytes,2,opt,name=hashed_secret,json=hashedSecret,proto3" json:"hashed_secret,omitempty"`
	Nonce               *BlsScalar        `protobuf:"bytes,3,opt,name=nonce,proto3" json:"nonce,omitempty"`
	PkR                 *StealthAddress   `protobuf:"bytes,4,opt,name=pk_r,json=pkR,proto3" json:"pk_r,omitempty"`
	Commitment          *JubJubCompressed `protobuf:"bytes,5,opt,name=commitment,proto3" json:"commitment,omitempty"`
	ElegibilityTs       *BlsScalar        `protobuf:"bytes,6,opt,name=elegibility_ts,json=elegibilityTs,proto3" json:"elegibility_ts,omitempty"`
	ExpirationTs        *BlsScalar        `protobuf:"bytes,7,opt,name=expiration_ts,json=expirationTs,proto3" json:"expiration_ts,omitempty"`
	BidTreeStorageIndex uint64            `protobuf:"varint,8,opt,name=bid_tree_storage_index,json=bidTreeStorageIndex,proto3" json:"bid_tree_storage_index,omitempty"`
}

func (x *Bid) Reset() {
	*x = Bid{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bid_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Bid) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Bid) ProtoMessage() {}

func (x *Bid) ProtoReflect() protoreflect.Message {
	mi := &file_bid_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Bid.ProtoReflect.Descriptor instead.
func (*Bid) Descriptor() ([]byte, []int) {
	return file_bid_proto_rawDescGZIP(), []int{0}
}

func (x *Bid) GetEncryptedData() *PoseidonCipher {
	if x != nil {
		return x.EncryptedData
	}
	return nil
}

func (x *Bid) GetHashedSecret() *BlsScalar {
	if x != nil {
		return x.HashedSecret
	}
	return nil
}

func (x *Bid) GetNonce() *BlsScalar {
	if x != nil {
		return x.Nonce
	}
	return nil
}

func (x *Bid) GetPkR() *StealthAddress {
	if x != nil {
		return x.PkR
	}
	return nil
}

func (x *Bid) GetCommitment() *JubJubCompressed {
	if x != nil {
		return x.Commitment
	}
	return nil
}

func (x *Bid) GetElegibilityTs() *BlsScalar {
	if x != nil {
		return x.ElegibilityTs
	}
	return nil
}

func (x *Bid) GetExpirationTs() *BlsScalar {
	if x != nil {
		return x.ExpirationTs
	}
	return nil
}

func (x *Bid) GetBidTreeStorageIndex() uint64 {
	if x != nil {
		return x.BidTreeStorageIndex
	}
	return 0
}

// Used to Request the creation of a Bid
type BidTransactionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	K                    *BlsScalar        `protobuf:"bytes,1,opt,name=k,proto3" json:"k,omitempty"`
	Value                uint64            `protobuf:"varint,2,opt,name=value,proto3" json:"value,omitempty"`
	Secret               *JubJubCompressed `protobuf:"bytes,3,opt,name=secret,proto3" json:"secret,omitempty"`
	PkR                  *StealthAddress   `protobuf:"bytes,4,opt,name=pk_r,json=pkR,proto3" json:"pk_r,omitempty"`
	Seed                 *BlsScalar        `protobuf:"bytes,5,opt,name=seed,proto3" json:"seed,omitempty"`
	LatestConsensusRound uint64            `protobuf:"fixed64,6,opt,name=latest_consensus_round,json=latestConsensusRound,proto3" json:"latest_consensus_round,omitempty"`
	LatestConsensusStep  uint64            `protobuf:"fixed64,7,opt,name=latest_consensus_step,json=latestConsensusStep,proto3" json:"latest_consensus_step,omitempty"` //rusk.Transaction tx = 8;
}

func (x *BidTransactionRequest) Reset() {
	*x = BidTransactionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bid_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BidTransactionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BidTransactionRequest) ProtoMessage() {}

func (x *BidTransactionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_bid_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BidTransactionRequest.ProtoReflect.Descriptor instead.
func (*BidTransactionRequest) Descriptor() ([]byte, []int) {
	return file_bid_proto_rawDescGZIP(), []int{1}
}

func (x *BidTransactionRequest) GetK() *BlsScalar {
	if x != nil {
		return x.K
	}
	return nil
}

func (x *BidTransactionRequest) GetValue() uint64 {
	if x != nil {
		return x.Value
	}
	return 0
}

func (x *BidTransactionRequest) GetSecret() *JubJubCompressed {
	if x != nil {
		return x.Secret
	}
	return nil
}

func (x *BidTransactionRequest) GetPkR() *StealthAddress {
	if x != nil {
		return x.PkR
	}
	return nil
}

func (x *BidTransactionRequest) GetSeed() *BlsScalar {
	if x != nil {
		return x.Seed
	}
	return nil
}

func (x *BidTransactionRequest) GetLatestConsensusRound() uint64 {
	if x != nil {
		return x.LatestConsensusRound
	}
	return 0
}

func (x *BidTransactionRequest) GetLatestConsensusStep() uint64 {
	if x != nil {
		return x.LatestConsensusStep
	}
	return 0
}

type BidTransaction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Bid     *Bid       `protobuf:"bytes,1,opt,name=bid,proto3" json:"bid,omitempty"`
	BidHash *BlsScalar `protobuf:"bytes,2,opt,name=bid_hash,json=bidHash,proto3" json:"bid_hash,omitempty"` // rusk.Transaction tx = 3;
}

func (x *BidTransaction) Reset() {
	*x = BidTransaction{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bid_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BidTransaction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BidTransaction) ProtoMessage() {}

func (x *BidTransaction) ProtoReflect() protoreflect.Message {
	mi := &file_bid_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BidTransaction.ProtoReflect.Descriptor instead.
func (*BidTransaction) Descriptor() ([]byte, []int) {
	return file_bid_proto_rawDescGZIP(), []int{2}
}

func (x *BidTransaction) GetBid() *Bid {
	if x != nil {
		return x.Bid
	}
	return nil
}

func (x *BidTransaction) GetBidHash() *BlsScalar {
	if x != nil {
		return x.BidHash
	}
	return nil
}

type FindBidRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Addr *StealthAddress `protobuf:"bytes,1,opt,name=addr,proto3" json:"addr,omitempty"`
}

func (x *FindBidRequest) Reset() {
	*x = FindBidRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bid_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindBidRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindBidRequest) ProtoMessage() {}

func (x *FindBidRequest) ProtoReflect() protoreflect.Message {
	mi := &file_bid_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindBidRequest.ProtoReflect.Descriptor instead.
func (*FindBidRequest) Descriptor() ([]byte, []int) {
	return file_bid_proto_rawDescGZIP(), []int{3}
}

func (x *FindBidRequest) GetAddr() *StealthAddress {
	if x != nil {
		return x.Addr
	}
	return nil
}

type BidList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BidList     []*Bid       `protobuf:"bytes,1,rep,name=bid_list,json=bidList,proto3" json:"bid_list,omitempty"`
	BidHashList []*BlsScalar `protobuf:"bytes,2,rep,name=bid_hash_list,json=bidHashList,proto3" json:"bid_hash_list,omitempty"`
}

func (x *BidList) Reset() {
	*x = BidList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bid_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BidList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BidList) ProtoMessage() {}

func (x *BidList) ProtoReflect() protoreflect.Message {
	mi := &file_bid_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BidList.ProtoReflect.Descriptor instead.
func (*BidList) Descriptor() ([]byte, []int) {
	return file_bid_proto_rawDescGZIP(), []int{4}
}

func (x *BidList) GetBidList() []*Bid {
	if x != nil {
		return x.BidList
	}
	return nil
}

func (x *BidList) GetBidHashList() []*BlsScalar {
	if x != nil {
		return x.BidHashList
	}
	return nil
}

var File_bid_proto protoreflect.FileDescriptor

var file_bid_proto_rawDesc = []byte{
	0x0a, 0x09, 0x62, 0x69, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x72, 0x75, 0x73,
	0x6b, 0x1a, 0x12, 0x62, 0x61, 0x73, 0x69, 0x63, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0a, 0x6b, 0x65, 0x79, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xa3, 0x03, 0x0a, 0x03, 0x42, 0x69, 0x64, 0x12, 0x3b, 0x0a, 0x0e, 0x65, 0x6e, 0x63,
	0x72, 0x79, 0x70, 0x74, 0x65, 0x64, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x14, 0x2e, 0x72, 0x75, 0x73, 0x6b, 0x2e, 0x50, 0x6f, 0x73, 0x65, 0x69, 0x64, 0x6f,
	0x6e, 0x43, 0x69, 0x70, 0x68, 0x65, 0x72, 0x52, 0x0d, 0x65, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74,
	0x65, 0x64, 0x44, 0x61, 0x74, 0x61, 0x12, 0x34, 0x0a, 0x0d, 0x68, 0x61, 0x73, 0x68, 0x65, 0x64,
	0x5f, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e,
	0x72, 0x75, 0x73, 0x6b, 0x2e, 0x42, 0x6c, 0x73, 0x53, 0x63, 0x61, 0x6c, 0x61, 0x72, 0x52, 0x0c,
	0x68, 0x61, 0x73, 0x68, 0x65, 0x64, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x12, 0x25, 0x0a, 0x05,
	0x6e, 0x6f, 0x6e, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x72, 0x75,
	0x73, 0x6b, 0x2e, 0x42, 0x6c, 0x73, 0x53, 0x63, 0x61, 0x6c, 0x61, 0x72, 0x52, 0x05, 0x6e, 0x6f,
	0x6e, 0x63, 0x65, 0x12, 0x27, 0x0a, 0x04, 0x70, 0x6b, 0x5f, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x14, 0x2e, 0x72, 0x75, 0x73, 0x6b, 0x2e, 0x53, 0x74, 0x65, 0x61, 0x6c, 0x74, 0x68,
	0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x52, 0x03, 0x70, 0x6b, 0x52, 0x12, 0x36, 0x0a, 0x0a,
	0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x16, 0x2e, 0x72, 0x75, 0x73, 0x6b, 0x2e, 0x4a, 0x75, 0x62, 0x4a, 0x75, 0x62, 0x43, 0x6f,
	0x6d, 0x70, 0x72, 0x65, 0x73, 0x73, 0x65, 0x64, 0x52, 0x0a, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74,
	0x6d, 0x65, 0x6e, 0x74, 0x12, 0x36, 0x0a, 0x0e, 0x65, 0x6c, 0x65, 0x67, 0x69, 0x62, 0x69, 0x6c,
	0x69, 0x74, 0x79, 0x5f, 0x74, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x72,
	0x75, 0x73, 0x6b, 0x2e, 0x42, 0x6c, 0x73, 0x53, 0x63, 0x61, 0x6c, 0x61, 0x72, 0x52, 0x0d, 0x65,
	0x6c, 0x65, 0x67, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x54, 0x73, 0x12, 0x34, 0x0a, 0x0d,
	0x65, 0x78, 0x70, 0x69, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x73, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x72, 0x75, 0x73, 0x6b, 0x2e, 0x42, 0x6c, 0x73, 0x53, 0x63,
	0x61, 0x6c, 0x61, 0x72, 0x52, 0x0c, 0x65, 0x78, 0x70, 0x69, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x54, 0x73, 0x12, 0x33, 0x0a, 0x16, 0x62, 0x69, 0x64, 0x5f, 0x74, 0x72, 0x65, 0x65, 0x5f, 0x73,
	0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x13, 0x62, 0x69, 0x64, 0x54, 0x72, 0x65, 0x65, 0x53, 0x74, 0x6f, 0x72, 0x61,
	0x67, 0x65, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x22, 0xb4, 0x02, 0x0a, 0x15, 0x42, 0x69, 0x64, 0x54,
	0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x1d, 0x0a, 0x01, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x72,
	0x75, 0x73, 0x6b, 0x2e, 0x42, 0x6c, 0x73, 0x53, 0x63, 0x61, 0x6c, 0x61, 0x72, 0x52, 0x01, 0x6b,
	0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x2e, 0x0a, 0x06, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x72, 0x75, 0x73, 0x6b, 0x2e, 0x4a, 0x75,
	0x62, 0x4a, 0x75, 0x62, 0x43, 0x6f, 0x6d, 0x70, 0x72, 0x65, 0x73, 0x73, 0x65, 0x64, 0x52, 0x06,
	0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x12, 0x27, 0x0a, 0x04, 0x70, 0x6b, 0x5f, 0x72, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x72, 0x75, 0x73, 0x6b, 0x2e, 0x53, 0x74, 0x65, 0x61,
	0x6c, 0x74, 0x68, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x52, 0x03, 0x70, 0x6b, 0x52, 0x12,
	0x23, 0x0a, 0x04, 0x73, 0x65, 0x65, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e,
	0x72, 0x75, 0x73, 0x6b, 0x2e, 0x42, 0x6c, 0x73, 0x53, 0x63, 0x61, 0x6c, 0x61, 0x72, 0x52, 0x04,
	0x73, 0x65, 0x65, 0x64, 0x12, 0x34, 0x0a, 0x16, 0x6c, 0x61, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x63,
	0x6f, 0x6e, 0x73, 0x65, 0x6e, 0x73, 0x75, 0x73, 0x5f, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x06, 0x52, 0x14, 0x6c, 0x61, 0x74, 0x65, 0x73, 0x74, 0x43, 0x6f, 0x6e, 0x73,
	0x65, 0x6e, 0x73, 0x75, 0x73, 0x52, 0x6f, 0x75, 0x6e, 0x64, 0x12, 0x32, 0x0a, 0x15, 0x6c, 0x61,
	0x74, 0x65, 0x73, 0x74, 0x5f, 0x63, 0x6f, 0x6e, 0x73, 0x65, 0x6e, 0x73, 0x75, 0x73, 0x5f, 0x73,
	0x74, 0x65, 0x70, 0x18, 0x07, 0x20, 0x01, 0x28, 0x06, 0x52, 0x13, 0x6c, 0x61, 0x74, 0x65, 0x73,
	0x74, 0x43, 0x6f, 0x6e, 0x73, 0x65, 0x6e, 0x73, 0x75, 0x73, 0x53, 0x74, 0x65, 0x70, 0x22, 0x59,
	0x0a, 0x0e, 0x42, 0x69, 0x64, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x1b, 0x0a, 0x03, 0x62, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x09, 0x2e,
	0x72, 0x75, 0x73, 0x6b, 0x2e, 0x42, 0x69, 0x64, 0x52, 0x03, 0x62, 0x69, 0x64, 0x12, 0x2a, 0x0a,
	0x08, 0x62, 0x69, 0x64, 0x5f, 0x68, 0x61, 0x73, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0f, 0x2e, 0x72, 0x75, 0x73, 0x6b, 0x2e, 0x42, 0x6c, 0x73, 0x53, 0x63, 0x61, 0x6c, 0x61, 0x72,
	0x52, 0x07, 0x62, 0x69, 0x64, 0x48, 0x61, 0x73, 0x68, 0x22, 0x3a, 0x0a, 0x0e, 0x46, 0x69, 0x6e,
	0x64, 0x42, 0x69, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x28, 0x0a, 0x04, 0x61,
	0x64, 0x64, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x72, 0x75, 0x73, 0x6b,
	0x2e, 0x53, 0x74, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x52,
	0x04, 0x61, 0x64, 0x64, 0x72, 0x22, 0x64, 0x0a, 0x07, 0x42, 0x69, 0x64, 0x4c, 0x69, 0x73, 0x74,
	0x12, 0x24, 0x0a, 0x08, 0x62, 0x69, 0x64, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x09, 0x2e, 0x72, 0x75, 0x73, 0x6b, 0x2e, 0x42, 0x69, 0x64, 0x52, 0x07, 0x62,
	0x69, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x33, 0x0a, 0x0d, 0x62, 0x69, 0x64, 0x5f, 0x68, 0x61,
	0x73, 0x68, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e,
	0x72, 0x75, 0x73, 0x6b, 0x2e, 0x42, 0x6c, 0x73, 0x53, 0x63, 0x61, 0x6c, 0x61, 0x72, 0x52, 0x0b,
	0x62, 0x69, 0x64, 0x48, 0x61, 0x73, 0x68, 0x4c, 0x69, 0x73, 0x74, 0x32, 0x7d, 0x0a, 0x0a, 0x42,
	0x69, 0x64, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3d, 0x0a, 0x06, 0x4e, 0x65, 0x77,
	0x42, 0x69, 0x64, 0x12, 0x1b, 0x2e, 0x72, 0x75, 0x73, 0x6b, 0x2e, 0x42, 0x69, 0x64, 0x54, 0x72,
	0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x14, 0x2e, 0x72, 0x75, 0x73, 0x6b, 0x2e, 0x42, 0x69, 0x64, 0x54, 0x72, 0x61, 0x6e, 0x73,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x00, 0x12, 0x30, 0x0a, 0x07, 0x46, 0x69, 0x6e, 0x64,
	0x42, 0x69, 0x64, 0x12, 0x14, 0x2e, 0x72, 0x75, 0x73, 0x6b, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x42,
	0x69, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0d, 0x2e, 0x72, 0x75, 0x73, 0x6b,
	0x2e, 0x42, 0x69, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x00, 0x42, 0x26, 0x5a, 0x24, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x75, 0x73, 0x6b, 0x2d, 0x6e, 0x65,
	0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2f, 0x72, 0x75, 0x73, 0x6b, 0x2d, 0x73, 0x63, 0x68, 0x65, 0x6d,
	0x61, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_bid_proto_rawDescOnce sync.Once
	file_bid_proto_rawDescData = file_bid_proto_rawDesc
)

func file_bid_proto_rawDescGZIP() []byte {
	file_bid_proto_rawDescOnce.Do(func() {
		file_bid_proto_rawDescData = protoimpl.X.CompressGZIP(file_bid_proto_rawDescData)
	})
	return file_bid_proto_rawDescData
}

var file_bid_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_bid_proto_goTypes = []interface{}{
	(*Bid)(nil),                   // 0: rusk.Bid
	(*BidTransactionRequest)(nil), // 1: rusk.BidTransactionRequest
	(*BidTransaction)(nil),        // 2: rusk.BidTransaction
	(*FindBidRequest)(nil),        // 3: rusk.FindBidRequest
	(*BidList)(nil),               // 4: rusk.BidList
	(*PoseidonCipher)(nil),        // 5: rusk.PoseidonCipher
	(*BlsScalar)(nil),             // 6: rusk.BlsScalar
	(*StealthAddress)(nil),        // 7: rusk.StealthAddress
	(*JubJubCompressed)(nil),      // 8: rusk.JubJubCompressed
}
var file_bid_proto_depIdxs = []int32{
	5,  // 0: rusk.Bid.encrypted_data:type_name -> rusk.PoseidonCipher
	6,  // 1: rusk.Bid.hashed_secret:type_name -> rusk.BlsScalar
	6,  // 2: rusk.Bid.nonce:type_name -> rusk.BlsScalar
	7,  // 3: rusk.Bid.pk_r:type_name -> rusk.StealthAddress
	8,  // 4: rusk.Bid.commitment:type_name -> rusk.JubJubCompressed
	6,  // 5: rusk.Bid.elegibility_ts:type_name -> rusk.BlsScalar
	6,  // 6: rusk.Bid.expiration_ts:type_name -> rusk.BlsScalar
	6,  // 7: rusk.BidTransactionRequest.k:type_name -> rusk.BlsScalar
	8,  // 8: rusk.BidTransactionRequest.secret:type_name -> rusk.JubJubCompressed
	7,  // 9: rusk.BidTransactionRequest.pk_r:type_name -> rusk.StealthAddress
	6,  // 10: rusk.BidTransactionRequest.seed:type_name -> rusk.BlsScalar
	0,  // 11: rusk.BidTransaction.bid:type_name -> rusk.Bid
	6,  // 12: rusk.BidTransaction.bid_hash:type_name -> rusk.BlsScalar
	7,  // 13: rusk.FindBidRequest.addr:type_name -> rusk.StealthAddress
	0,  // 14: rusk.BidList.bid_list:type_name -> rusk.Bid
	6,  // 15: rusk.BidList.bid_hash_list:type_name -> rusk.BlsScalar
	1,  // 16: rusk.BidService.NewBid:input_type -> rusk.BidTransactionRequest
	3,  // 17: rusk.BidService.FindBid:input_type -> rusk.FindBidRequest
	2,  // 18: rusk.BidService.NewBid:output_type -> rusk.BidTransaction
	4,  // 19: rusk.BidService.FindBid:output_type -> rusk.BidList
	18, // [18:20] is the sub-list for method output_type
	16, // [16:18] is the sub-list for method input_type
	16, // [16:16] is the sub-list for extension type_name
	16, // [16:16] is the sub-list for extension extendee
	0,  // [0:16] is the sub-list for field type_name
}

func init() { file_bid_proto_init() }
func file_bid_proto_init() {
	if File_bid_proto != nil {
		return
	}
	file_basic_fields_proto_init()
	file_keys_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_bid_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Bid); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_bid_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BidTransactionRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_bid_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BidTransaction); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_bid_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindBidRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_bid_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BidList); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_bid_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_bid_proto_goTypes,
		DependencyIndexes: file_bid_proto_depIdxs,
		MessageInfos:      file_bid_proto_msgTypes,
	}.Build()
	File_bid_proto = out.File
	file_bid_proto_rawDesc = nil
	file_bid_proto_goTypes = nil
	file_bid_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

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
	cc grpc.ClientConnInterface
}

func NewBidServiceClient(cc grpc.ClientConnInterface) BidServiceClient {
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

func (*UnimplementedBidServiceServer) NewBid(context.Context, *BidTransactionRequest) (*BidTransaction, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NewBid not implemented")
}
func (*UnimplementedBidServiceServer) FindBid(context.Context, *FindBidRequest) (*BidList, error) {
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
