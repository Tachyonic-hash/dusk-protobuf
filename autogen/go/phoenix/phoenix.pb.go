// Code generated by protoc-gen-go. DO NOT EDIT.
// source: phoenix.proto

package phoenix

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type EchoMethod struct {
	M                    string   `protobuf:"bytes,1,opt,name=m,proto3" json:"m,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EchoMethod) Reset()         { *m = EchoMethod{} }
func (m *EchoMethod) String() string { return proto.CompactTextString(m) }
func (*EchoMethod) ProtoMessage()    {}
func (*EchoMethod) Descriptor() ([]byte, []int) {
	return fileDescriptor_a1c51cd95b21b6dd, []int{0}
}

func (m *EchoMethod) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EchoMethod.Unmarshal(m, b)
}
func (m *EchoMethod) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EchoMethod.Marshal(b, m, deterministic)
}
func (m *EchoMethod) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EchoMethod.Merge(m, src)
}
func (m *EchoMethod) XXX_Size() int {
	return xxx_messageInfo_EchoMethod.Size(m)
}
func (m *EchoMethod) XXX_DiscardUnknown() {
	xxx_messageInfo_EchoMethod.DiscardUnknown(m)
}

var xxx_messageInfo_EchoMethod proto.InternalMessageInfo

func (m *EchoMethod) GetM() string {
	if m != nil {
		return m.M
	}
	return ""
}

type GenerateSecretKeyRequest struct {
	B                    []byte   `protobuf:"bytes,1,opt,name=b,proto3" json:"b,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GenerateSecretKeyRequest) Reset()         { *m = GenerateSecretKeyRequest{} }
func (m *GenerateSecretKeyRequest) String() string { return proto.CompactTextString(m) }
func (*GenerateSecretKeyRequest) ProtoMessage()    {}
func (*GenerateSecretKeyRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a1c51cd95b21b6dd, []int{1}
}

func (m *GenerateSecretKeyRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GenerateSecretKeyRequest.Unmarshal(m, b)
}
func (m *GenerateSecretKeyRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GenerateSecretKeyRequest.Marshal(b, m, deterministic)
}
func (m *GenerateSecretKeyRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenerateSecretKeyRequest.Merge(m, src)
}
func (m *GenerateSecretKeyRequest) XXX_Size() int {
	return xxx_messageInfo_GenerateSecretKeyRequest.Size(m)
}
func (m *GenerateSecretKeyRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GenerateSecretKeyRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GenerateSecretKeyRequest proto.InternalMessageInfo

func (m *GenerateSecretKeyRequest) GetB() []byte {
	if m != nil {
		return m.B
	}
	return nil
}

type NullifierRequest struct {
	Note                 *Note      `protobuf:"bytes,1,opt,name=note,proto3" json:"note,omitempty"`
	Sk                   *SecretKey `protobuf:"bytes,2,opt,name=sk,proto3" json:"sk,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *NullifierRequest) Reset()         { *m = NullifierRequest{} }
func (m *NullifierRequest) String() string { return proto.CompactTextString(m) }
func (*NullifierRequest) ProtoMessage()    {}
func (*NullifierRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a1c51cd95b21b6dd, []int{2}
}

func (m *NullifierRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NullifierRequest.Unmarshal(m, b)
}
func (m *NullifierRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NullifierRequest.Marshal(b, m, deterministic)
}
func (m *NullifierRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NullifierRequest.Merge(m, src)
}
func (m *NullifierRequest) XXX_Size() int {
	return xxx_messageInfo_NullifierRequest.Size(m)
}
func (m *NullifierRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_NullifierRequest.DiscardUnknown(m)
}

var xxx_messageInfo_NullifierRequest proto.InternalMessageInfo

func (m *NullifierRequest) GetNote() *Note {
	if m != nil {
		return m.Note
	}
	return nil
}

func (m *NullifierRequest) GetSk() *SecretKey {
	if m != nil {
		return m.Sk
	}
	return nil
}

type NullifierResponse struct {
	Nullifier            *Nullifier `protobuf:"bytes,1,opt,name=nullifier,proto3" json:"nullifier,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *NullifierResponse) Reset()         { *m = NullifierResponse{} }
func (m *NullifierResponse) String() string { return proto.CompactTextString(m) }
func (*NullifierResponse) ProtoMessage()    {}
func (*NullifierResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a1c51cd95b21b6dd, []int{3}
}

func (m *NullifierResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NullifierResponse.Unmarshal(m, b)
}
func (m *NullifierResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NullifierResponse.Marshal(b, m, deterministic)
}
func (m *NullifierResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NullifierResponse.Merge(m, src)
}
func (m *NullifierResponse) XXX_Size() int {
	return xxx_messageInfo_NullifierResponse.Size(m)
}
func (m *NullifierResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_NullifierResponse.DiscardUnknown(m)
}

var xxx_messageInfo_NullifierResponse proto.InternalMessageInfo

func (m *NullifierResponse) GetNullifier() *Nullifier {
	if m != nil {
		return m.Nullifier
	}
	return nil
}

type NullifierStatusRequest struct {
	Nullifier            *Nullifier `protobuf:"bytes,1,opt,name=nullifier,proto3" json:"nullifier,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *NullifierStatusRequest) Reset()         { *m = NullifierStatusRequest{} }
func (m *NullifierStatusRequest) String() string { return proto.CompactTextString(m) }
func (*NullifierStatusRequest) ProtoMessage()    {}
func (*NullifierStatusRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a1c51cd95b21b6dd, []int{4}
}

func (m *NullifierStatusRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NullifierStatusRequest.Unmarshal(m, b)
}
func (m *NullifierStatusRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NullifierStatusRequest.Marshal(b, m, deterministic)
}
func (m *NullifierStatusRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NullifierStatusRequest.Merge(m, src)
}
func (m *NullifierStatusRequest) XXX_Size() int {
	return xxx_messageInfo_NullifierStatusRequest.Size(m)
}
func (m *NullifierStatusRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_NullifierStatusRequest.DiscardUnknown(m)
}

var xxx_messageInfo_NullifierStatusRequest proto.InternalMessageInfo

func (m *NullifierStatusRequest) GetNullifier() *Nullifier {
	if m != nil {
		return m.Nullifier
	}
	return nil
}

type NullifierStatusResponse struct {
	Unspent              bool     `protobuf:"varint,1,opt,name=unspent,proto3" json:"unspent,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NullifierStatusResponse) Reset()         { *m = NullifierStatusResponse{} }
func (m *NullifierStatusResponse) String() string { return proto.CompactTextString(m) }
func (*NullifierStatusResponse) ProtoMessage()    {}
func (*NullifierStatusResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a1c51cd95b21b6dd, []int{5}
}

func (m *NullifierStatusResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NullifierStatusResponse.Unmarshal(m, b)
}
func (m *NullifierStatusResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NullifierStatusResponse.Marshal(b, m, deterministic)
}
func (m *NullifierStatusResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NullifierStatusResponse.Merge(m, src)
}
func (m *NullifierStatusResponse) XXX_Size() int {
	return xxx_messageInfo_NullifierStatusResponse.Size(m)
}
func (m *NullifierStatusResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_NullifierStatusResponse.DiscardUnknown(m)
}

var xxx_messageInfo_NullifierStatusResponse proto.InternalMessageInfo

func (m *NullifierStatusResponse) GetUnspent() bool {
	if m != nil {
		return m.Unspent
	}
	return false
}

type DecryptNoteRequest struct {
	Note                 *Note    `protobuf:"bytes,1,opt,name=note,proto3" json:"note,omitempty"`
	Vk                   *ViewKey `protobuf:"bytes,2,opt,name=vk,proto3" json:"vk,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DecryptNoteRequest) Reset()         { *m = DecryptNoteRequest{} }
func (m *DecryptNoteRequest) String() string { return proto.CompactTextString(m) }
func (*DecryptNoteRequest) ProtoMessage()    {}
func (*DecryptNoteRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a1c51cd95b21b6dd, []int{6}
}

func (m *DecryptNoteRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DecryptNoteRequest.Unmarshal(m, b)
}
func (m *DecryptNoteRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DecryptNoteRequest.Marshal(b, m, deterministic)
}
func (m *DecryptNoteRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DecryptNoteRequest.Merge(m, src)
}
func (m *DecryptNoteRequest) XXX_Size() int {
	return xxx_messageInfo_DecryptNoteRequest.Size(m)
}
func (m *DecryptNoteRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DecryptNoteRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DecryptNoteRequest proto.InternalMessageInfo

func (m *DecryptNoteRequest) GetNote() *Note {
	if m != nil {
		return m.Note
	}
	return nil
}

func (m *DecryptNoteRequest) GetVk() *ViewKey {
	if m != nil {
		return m.Vk
	}
	return nil
}

type OwnedNotesRequest struct {
	Vk                   *ViewKey `protobuf:"bytes,1,opt,name=vk,proto3" json:"vk,omitempty"`
	Notes                []*Note  `protobuf:"bytes,2,rep,name=notes,proto3" json:"notes,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OwnedNotesRequest) Reset()         { *m = OwnedNotesRequest{} }
func (m *OwnedNotesRequest) String() string { return proto.CompactTextString(m) }
func (*OwnedNotesRequest) ProtoMessage()    {}
func (*OwnedNotesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a1c51cd95b21b6dd, []int{7}
}

func (m *OwnedNotesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OwnedNotesRequest.Unmarshal(m, b)
}
func (m *OwnedNotesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OwnedNotesRequest.Marshal(b, m, deterministic)
}
func (m *OwnedNotesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OwnedNotesRequest.Merge(m, src)
}
func (m *OwnedNotesRequest) XXX_Size() int {
	return xxx_messageInfo_OwnedNotesRequest.Size(m)
}
func (m *OwnedNotesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_OwnedNotesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_OwnedNotesRequest proto.InternalMessageInfo

func (m *OwnedNotesRequest) GetVk() *ViewKey {
	if m != nil {
		return m.Vk
	}
	return nil
}

func (m *OwnedNotesRequest) GetNotes() []*Note {
	if m != nil {
		return m.Notes
	}
	return nil
}

type OwnedNotesResponse struct {
	Notes                []*DecryptedNote `protobuf:"bytes,1,rep,name=notes,proto3" json:"notes,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *OwnedNotesResponse) Reset()         { *m = OwnedNotesResponse{} }
func (m *OwnedNotesResponse) String() string { return proto.CompactTextString(m) }
func (*OwnedNotesResponse) ProtoMessage()    {}
func (*OwnedNotesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a1c51cd95b21b6dd, []int{8}
}

func (m *OwnedNotesResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OwnedNotesResponse.Unmarshal(m, b)
}
func (m *OwnedNotesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OwnedNotesResponse.Marshal(b, m, deterministic)
}
func (m *OwnedNotesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OwnedNotesResponse.Merge(m, src)
}
func (m *OwnedNotesResponse) XXX_Size() int {
	return xxx_messageInfo_OwnedNotesResponse.Size(m)
}
func (m *OwnedNotesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_OwnedNotesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_OwnedNotesResponse proto.InternalMessageInfo

func (m *OwnedNotesResponse) GetNotes() []*DecryptedNote {
	if m != nil {
		return m.Notes
	}
	return nil
}

type NewTransactionInputRequest struct {
	Pos                  *Idx       `protobuf:"bytes,1,opt,name=pos,proto3" json:"pos,omitempty"`
	Sk                   *SecretKey `protobuf:"bytes,2,opt,name=sk,proto3" json:"sk,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *NewTransactionInputRequest) Reset()         { *m = NewTransactionInputRequest{} }
func (m *NewTransactionInputRequest) String() string { return proto.CompactTextString(m) }
func (*NewTransactionInputRequest) ProtoMessage()    {}
func (*NewTransactionInputRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a1c51cd95b21b6dd, []int{9}
}

func (m *NewTransactionInputRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NewTransactionInputRequest.Unmarshal(m, b)
}
func (m *NewTransactionInputRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NewTransactionInputRequest.Marshal(b, m, deterministic)
}
func (m *NewTransactionInputRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NewTransactionInputRequest.Merge(m, src)
}
func (m *NewTransactionInputRequest) XXX_Size() int {
	return xxx_messageInfo_NewTransactionInputRequest.Size(m)
}
func (m *NewTransactionInputRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_NewTransactionInputRequest.DiscardUnknown(m)
}

var xxx_messageInfo_NewTransactionInputRequest proto.InternalMessageInfo

func (m *NewTransactionInputRequest) GetPos() *Idx {
	if m != nil {
		return m.Pos
	}
	return nil
}

func (m *NewTransactionInputRequest) GetSk() *SecretKey {
	if m != nil {
		return m.Sk
	}
	return nil
}

type NewTransactionOutputRequest struct {
	NoteType             NoteType   `protobuf:"varint,1,opt,name=note_type,json=noteType,proto3,enum=phoenix.NoteType" json:"note_type,omitempty"`
	Pk                   *PublicKey `protobuf:"bytes,2,opt,name=pk,proto3" json:"pk,omitempty"`
	Value                uint64     `protobuf:"varint,3,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *NewTransactionOutputRequest) Reset()         { *m = NewTransactionOutputRequest{} }
func (m *NewTransactionOutputRequest) String() string { return proto.CompactTextString(m) }
func (*NewTransactionOutputRequest) ProtoMessage()    {}
func (*NewTransactionOutputRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a1c51cd95b21b6dd, []int{10}
}

func (m *NewTransactionOutputRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NewTransactionOutputRequest.Unmarshal(m, b)
}
func (m *NewTransactionOutputRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NewTransactionOutputRequest.Marshal(b, m, deterministic)
}
func (m *NewTransactionOutputRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NewTransactionOutputRequest.Merge(m, src)
}
func (m *NewTransactionOutputRequest) XXX_Size() int {
	return xxx_messageInfo_NewTransactionOutputRequest.Size(m)
}
func (m *NewTransactionOutputRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_NewTransactionOutputRequest.DiscardUnknown(m)
}

var xxx_messageInfo_NewTransactionOutputRequest proto.InternalMessageInfo

func (m *NewTransactionOutputRequest) GetNoteType() NoteType {
	if m != nil {
		return m.NoteType
	}
	return NoteType_TRANSPARENT
}

func (m *NewTransactionOutputRequest) GetPk() *PublicKey {
	if m != nil {
		return m.Pk
	}
	return nil
}

func (m *NewTransactionOutputRequest) GetValue() uint64 {
	if m != nil {
		return m.Value
	}
	return 0
}

type NewTransactionRequest struct {
	Inputs               []*TransactionInput  `protobuf:"bytes,1,rep,name=inputs,proto3" json:"inputs,omitempty"`
	Outputs              []*TransactionOutput `protobuf:"bytes,2,rep,name=outputs,proto3" json:"outputs,omitempty"`
	Fee                  uint64               `protobuf:"varint,3,opt,name=fee,proto3" json:"fee,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *NewTransactionRequest) Reset()         { *m = NewTransactionRequest{} }
func (m *NewTransactionRequest) String() string { return proto.CompactTextString(m) }
func (*NewTransactionRequest) ProtoMessage()    {}
func (*NewTransactionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a1c51cd95b21b6dd, []int{11}
}

func (m *NewTransactionRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NewTransactionRequest.Unmarshal(m, b)
}
func (m *NewTransactionRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NewTransactionRequest.Marshal(b, m, deterministic)
}
func (m *NewTransactionRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NewTransactionRequest.Merge(m, src)
}
func (m *NewTransactionRequest) XXX_Size() int {
	return xxx_messageInfo_NewTransactionRequest.Size(m)
}
func (m *NewTransactionRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_NewTransactionRequest.DiscardUnknown(m)
}

var xxx_messageInfo_NewTransactionRequest proto.InternalMessageInfo

func (m *NewTransactionRequest) GetInputs() []*TransactionInput {
	if m != nil {
		return m.Inputs
	}
	return nil
}

func (m *NewTransactionRequest) GetOutputs() []*TransactionOutput {
	if m != nil {
		return m.Outputs
	}
	return nil
}

func (m *NewTransactionRequest) GetFee() uint64 {
	if m != nil {
		return m.Fee
	}
	return 0
}

type VerifyTransactionResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VerifyTransactionResponse) Reset()         { *m = VerifyTransactionResponse{} }
func (m *VerifyTransactionResponse) String() string { return proto.CompactTextString(m) }
func (*VerifyTransactionResponse) ProtoMessage()    {}
func (*VerifyTransactionResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a1c51cd95b21b6dd, []int{12}
}

func (m *VerifyTransactionResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VerifyTransactionResponse.Unmarshal(m, b)
}
func (m *VerifyTransactionResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VerifyTransactionResponse.Marshal(b, m, deterministic)
}
func (m *VerifyTransactionResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VerifyTransactionResponse.Merge(m, src)
}
func (m *VerifyTransactionResponse) XXX_Size() int {
	return xxx_messageInfo_VerifyTransactionResponse.Size(m)
}
func (m *VerifyTransactionResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_VerifyTransactionResponse.DiscardUnknown(m)
}

var xxx_messageInfo_VerifyTransactionResponse proto.InternalMessageInfo

type VerifyTransactionRootRequest struct {
	Transaction          *Transaction `protobuf:"bytes,1,opt,name=transaction,proto3" json:"transaction,omitempty"`
	Root                 *Scalar      `protobuf:"bytes,2,opt,name=root,proto3" json:"root,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *VerifyTransactionRootRequest) Reset()         { *m = VerifyTransactionRootRequest{} }
func (m *VerifyTransactionRootRequest) String() string { return proto.CompactTextString(m) }
func (*VerifyTransactionRootRequest) ProtoMessage()    {}
func (*VerifyTransactionRootRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a1c51cd95b21b6dd, []int{13}
}

func (m *VerifyTransactionRootRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VerifyTransactionRootRequest.Unmarshal(m, b)
}
func (m *VerifyTransactionRootRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VerifyTransactionRootRequest.Marshal(b, m, deterministic)
}
func (m *VerifyTransactionRootRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VerifyTransactionRootRequest.Merge(m, src)
}
func (m *VerifyTransactionRootRequest) XXX_Size() int {
	return xxx_messageInfo_VerifyTransactionRootRequest.Size(m)
}
func (m *VerifyTransactionRootRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_VerifyTransactionRootRequest.DiscardUnknown(m)
}

var xxx_messageInfo_VerifyTransactionRootRequest proto.InternalMessageInfo

func (m *VerifyTransactionRootRequest) GetTransaction() *Transaction {
	if m != nil {
		return m.Transaction
	}
	return nil
}

func (m *VerifyTransactionRootRequest) GetRoot() *Scalar {
	if m != nil {
		return m.Root
	}
	return nil
}

type VerifyTransactionRootResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VerifyTransactionRootResponse) Reset()         { *m = VerifyTransactionRootResponse{} }
func (m *VerifyTransactionRootResponse) String() string { return proto.CompactTextString(m) }
func (*VerifyTransactionRootResponse) ProtoMessage()    {}
func (*VerifyTransactionRootResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a1c51cd95b21b6dd, []int{14}
}

func (m *VerifyTransactionRootResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VerifyTransactionRootResponse.Unmarshal(m, b)
}
func (m *VerifyTransactionRootResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VerifyTransactionRootResponse.Marshal(b, m, deterministic)
}
func (m *VerifyTransactionRootResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VerifyTransactionRootResponse.Merge(m, src)
}
func (m *VerifyTransactionRootResponse) XXX_Size() int {
	return xxx_messageInfo_VerifyTransactionRootResponse.Size(m)
}
func (m *VerifyTransactionRootResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_VerifyTransactionRootResponse.DiscardUnknown(m)
}

var xxx_messageInfo_VerifyTransactionRootResponse proto.InternalMessageInfo

type StoreTransactionsRequest struct {
	Transactions         []*Transaction `protobuf:"bytes,1,rep,name=transactions,proto3" json:"transactions,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *StoreTransactionsRequest) Reset()         { *m = StoreTransactionsRequest{} }
func (m *StoreTransactionsRequest) String() string { return proto.CompactTextString(m) }
func (*StoreTransactionsRequest) ProtoMessage()    {}
func (*StoreTransactionsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a1c51cd95b21b6dd, []int{15}
}

func (m *StoreTransactionsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StoreTransactionsRequest.Unmarshal(m, b)
}
func (m *StoreTransactionsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StoreTransactionsRequest.Marshal(b, m, deterministic)
}
func (m *StoreTransactionsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StoreTransactionsRequest.Merge(m, src)
}
func (m *StoreTransactionsRequest) XXX_Size() int {
	return xxx_messageInfo_StoreTransactionsRequest.Size(m)
}
func (m *StoreTransactionsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_StoreTransactionsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_StoreTransactionsRequest proto.InternalMessageInfo

func (m *StoreTransactionsRequest) GetTransactions() []*Transaction {
	if m != nil {
		return m.Transactions
	}
	return nil
}

type StoreTransactionsResponse struct {
	Notes                []*Note  `protobuf:"bytes,1,rep,name=notes,proto3" json:"notes,omitempty"`
	Root                 *Scalar  `protobuf:"bytes,2,opt,name=root,proto3" json:"root,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StoreTransactionsResponse) Reset()         { *m = StoreTransactionsResponse{} }
func (m *StoreTransactionsResponse) String() string { return proto.CompactTextString(m) }
func (*StoreTransactionsResponse) ProtoMessage()    {}
func (*StoreTransactionsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a1c51cd95b21b6dd, []int{16}
}

func (m *StoreTransactionsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StoreTransactionsResponse.Unmarshal(m, b)
}
func (m *StoreTransactionsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StoreTransactionsResponse.Marshal(b, m, deterministic)
}
func (m *StoreTransactionsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StoreTransactionsResponse.Merge(m, src)
}
func (m *StoreTransactionsResponse) XXX_Size() int {
	return xxx_messageInfo_StoreTransactionsResponse.Size(m)
}
func (m *StoreTransactionsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_StoreTransactionsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_StoreTransactionsResponse proto.InternalMessageInfo

func (m *StoreTransactionsResponse) GetNotes() []*Note {
	if m != nil {
		return m.Notes
	}
	return nil
}

func (m *StoreTransactionsResponse) GetRoot() *Scalar {
	if m != nil {
		return m.Root
	}
	return nil
}

type SetFeePkRequest struct {
	Transaction          *Transaction `protobuf:"bytes,1,opt,name=transaction,proto3" json:"transaction,omitempty"`
	Pk                   *PublicKey   `protobuf:"bytes,2,opt,name=pk,proto3" json:"pk,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *SetFeePkRequest) Reset()         { *m = SetFeePkRequest{} }
func (m *SetFeePkRequest) String() string { return proto.CompactTextString(m) }
func (*SetFeePkRequest) ProtoMessage()    {}
func (*SetFeePkRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a1c51cd95b21b6dd, []int{17}
}

func (m *SetFeePkRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SetFeePkRequest.Unmarshal(m, b)
}
func (m *SetFeePkRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SetFeePkRequest.Marshal(b, m, deterministic)
}
func (m *SetFeePkRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetFeePkRequest.Merge(m, src)
}
func (m *SetFeePkRequest) XXX_Size() int {
	return xxx_messageInfo_SetFeePkRequest.Size(m)
}
func (m *SetFeePkRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SetFeePkRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SetFeePkRequest proto.InternalMessageInfo

func (m *SetFeePkRequest) GetTransaction() *Transaction {
	if m != nil {
		return m.Transaction
	}
	return nil
}

func (m *SetFeePkRequest) GetPk() *PublicKey {
	if m != nil {
		return m.Pk
	}
	return nil
}

type KeysResponse struct {
	Vk                   *ViewKey   `protobuf:"bytes,1,opt,name=vk,proto3" json:"vk,omitempty"`
	Pk                   *PublicKey `protobuf:"bytes,2,opt,name=pk,proto3" json:"pk,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *KeysResponse) Reset()         { *m = KeysResponse{} }
func (m *KeysResponse) String() string { return proto.CompactTextString(m) }
func (*KeysResponse) ProtoMessage()    {}
func (*KeysResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a1c51cd95b21b6dd, []int{18}
}

func (m *KeysResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KeysResponse.Unmarshal(m, b)
}
func (m *KeysResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KeysResponse.Marshal(b, m, deterministic)
}
func (m *KeysResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KeysResponse.Merge(m, src)
}
func (m *KeysResponse) XXX_Size() int {
	return xxx_messageInfo_KeysResponse.Size(m)
}
func (m *KeysResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_KeysResponse.DiscardUnknown(m)
}

var xxx_messageInfo_KeysResponse proto.InternalMessageInfo

func (m *KeysResponse) GetVk() *ViewKey {
	if m != nil {
		return m.Vk
	}
	return nil
}

func (m *KeysResponse) GetPk() *PublicKey {
	if m != nil {
		return m.Pk
	}
	return nil
}

func init() {
	proto.RegisterType((*EchoMethod)(nil), "phoenix.EchoMethod")
	proto.RegisterType((*GenerateSecretKeyRequest)(nil), "phoenix.GenerateSecretKeyRequest")
	proto.RegisterType((*NullifierRequest)(nil), "phoenix.NullifierRequest")
	proto.RegisterType((*NullifierResponse)(nil), "phoenix.NullifierResponse")
	proto.RegisterType((*NullifierStatusRequest)(nil), "phoenix.NullifierStatusRequest")
	proto.RegisterType((*NullifierStatusResponse)(nil), "phoenix.NullifierStatusResponse")
	proto.RegisterType((*DecryptNoteRequest)(nil), "phoenix.DecryptNoteRequest")
	proto.RegisterType((*OwnedNotesRequest)(nil), "phoenix.OwnedNotesRequest")
	proto.RegisterType((*OwnedNotesResponse)(nil), "phoenix.OwnedNotesResponse")
	proto.RegisterType((*NewTransactionInputRequest)(nil), "phoenix.NewTransactionInputRequest")
	proto.RegisterType((*NewTransactionOutputRequest)(nil), "phoenix.NewTransactionOutputRequest")
	proto.RegisterType((*NewTransactionRequest)(nil), "phoenix.NewTransactionRequest")
	proto.RegisterType((*VerifyTransactionResponse)(nil), "phoenix.VerifyTransactionResponse")
	proto.RegisterType((*VerifyTransactionRootRequest)(nil), "phoenix.VerifyTransactionRootRequest")
	proto.RegisterType((*VerifyTransactionRootResponse)(nil), "phoenix.VerifyTransactionRootResponse")
	proto.RegisterType((*StoreTransactionsRequest)(nil), "phoenix.StoreTransactionsRequest")
	proto.RegisterType((*StoreTransactionsResponse)(nil), "phoenix.StoreTransactionsResponse")
	proto.RegisterType((*SetFeePkRequest)(nil), "phoenix.SetFeePkRequest")
	proto.RegisterType((*KeysResponse)(nil), "phoenix.KeysResponse")
}

func init() {
	proto.RegisterFile("phoenix.proto", fileDescriptor_a1c51cd95b21b6dd)
}

var fileDescriptor_a1c51cd95b21b6dd = []byte{
	// 882 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x56, 0x6d, 0x6f, 0xdc, 0x44,
	0x10, 0xb6, 0x2f, 0xd7, 0x26, 0x37, 0xb9, 0x36, 0xb9, 0x69, 0x52, 0x2e, 0x1b, 0x68, 0xaf, 0x1b,
	0x40, 0x11, 0x42, 0x11, 0x24, 0x08, 0xf1, 0x81, 0x4f, 0xa8, 0x49, 0x29, 0x85, 0x24, 0xf2, 0x45,
	0x41, 0xad, 0x90, 0xe0, 0xe2, 0x4c, 0x74, 0xd6, 0x39, 0x5e, 0x63, 0xaf, 0x93, 0x9e, 0xf8, 0x01,
	0x7c, 0xe7, 0xc7, 0xf0, 0xfb, 0x90, 0x5f, 0x76, 0x6d, 0xd7, 0xeb, 0xe6, 0x50, 0xbf, 0xed, 0xed,
	0x3c, 0xf3, 0xcc, 0xcb, 0x8e, 0x9f, 0x39, 0x78, 0x10, 0x4e, 0x05, 0x05, 0xde, 0xdb, 0xbd, 0x30,
	0x12, 0x52, 0xe0, 0x72, 0xf1, 0x93, 0xad, 0x5e, 0x79, 0xe4, 0x5f, 0xe6, 0xb7, 0x0c, 0x66, 0x34,
	0x8f, 0xd5, 0x39, 0x10, 0x92, 0x8a, 0xf3, 0x40, 0x46, 0x93, 0x20, 0x9e, 0xb8, 0xd2, 0x13, 0x41,
	0x7e, 0xc5, 0x19, 0xc0, 0xa1, 0x3b, 0x15, 0xbf, 0x90, 0x9c, 0x8a, 0x4b, 0xec, 0x83, 0x7d, 0x3d,
	0xb4, 0x47, 0xf6, 0x6e, 0xcf, 0xb1, 0xaf, 0xf9, 0x2e, 0x0c, 0x5f, 0x50, 0x40, 0xd1, 0x44, 0xd2,
	0x98, 0xdc, 0x88, 0xe4, 0x2b, 0x9a, 0x3b, 0xf4, 0x67, 0x42, 0xb1, 0x4c, 0x91, 0x17, 0x19, 0xb2,
	0xef, 0xd8, 0x17, 0xfc, 0x35, 0xac, 0x1f, 0x27, 0xbe, 0xef, 0x5d, 0x79, 0x14, 0x29, 0xc4, 0x33,
	0xe8, 0xa6, 0xa1, 0x33, 0xd0, 0xea, 0xfe, 0x83, 0x3d, 0x95, 0xf8, 0xb1, 0x90, 0xe4, 0x64, 0x26,
	0xe4, 0xd0, 0x89, 0x67, 0xc3, 0x4e, 0x06, 0x40, 0x0d, 0x28, 0x63, 0x75, 0xe2, 0x19, 0x3f, 0x84,
	0x41, 0x85, 0x3a, 0x0e, 0x45, 0x10, 0x13, 0x7e, 0x05, 0xbd, 0x40, 0x5d, 0x16, 0x01, 0x4a, 0xff,
	0x12, 0x5e, 0x82, 0xf8, 0x4f, 0xf0, 0x58, 0xdf, 0x8f, 0xe5, 0x44, 0x26, 0xb1, 0xca, 0xf3, 0xff,
	0x73, 0x1d, 0xc0, 0x47, 0x0d, 0xae, 0x22, 0xb1, 0x21, 0x2c, 0x27, 0x41, 0x1c, 0x52, 0x20, 0x33,
	0xaa, 0x15, 0x47, 0xfd, 0xe4, 0xaf, 0x01, 0x9f, 0x93, 0x1b, 0xcd, 0x43, 0x99, 0x35, 0x60, 0xf1,
	0x26, 0x8d, 0xa0, 0x73, 0xa3, 0x9a, 0xb4, 0xae, 0x01, 0xe7, 0x1e, 0xdd, 0x66, 0x2d, 0xba, 0x99,
	0xf1, 0x37, 0x30, 0x38, 0xb9, 0x0d, 0xe8, 0x32, 0x75, 0xd2, 0x65, 0xe5, 0x6e, 0x76, 0xbb, 0x1b,
	0xee, 0xc0, 0xbd, 0x34, 0x40, 0x3c, 0xec, 0x8c, 0x96, 0x9a, 0xc1, 0x73, 0x1b, 0xff, 0x01, 0xb0,
	0xca, 0x5d, 0x94, 0xf9, 0xa5, 0x72, 0xb5, 0x33, 0xd7, 0xc7, 0xda, 0xb5, 0x28, 0x31, 0xc7, 0x2b,
	0x8e, 0x3f, 0x80, 0x1d, 0xd3, 0xed, 0x59, 0x39, 0x7b, 0x2f, 0x83, 0x30, 0x91, 0x2a, 0xd1, 0x27,
	0xb0, 0x14, 0x8a, 0xb8, 0xc8, 0xb4, 0xaf, 0x99, 0x5e, 0x5e, 0xbe, 0x75, 0x52, 0xc3, 0x42, 0x43,
	0xf2, 0xb7, 0x0d, 0xdb, 0xf5, 0x10, 0x27, 0x89, 0xac, 0xc4, 0xd8, 0x83, 0x5e, 0x9a, 0xca, 0xef,
	0x72, 0x1e, 0xe6, 0xbd, 0x7e, 0xb8, 0x3f, 0xa8, 0x95, 0x7b, 0x36, 0x0f, 0xc9, 0x59, 0x09, 0x8a,
	0x53, 0x1a, 0x33, 0x6c, 0xc6, 0x3c, 0x4d, 0x2e, 0x7c, 0xcf, 0xcd, 0x62, 0x86, 0x33, 0xdc, 0x80,
	0x7b, 0x37, 0x13, 0x3f, 0xa1, 0xe1, 0xd2, 0xc8, 0xde, 0xed, 0x3a, 0xf9, 0x0f, 0xfe, 0x8f, 0x0d,
	0x9b, 0xf5, 0x4c, 0x54, 0x0e, 0x5f, 0xc3, 0x7d, 0x2f, 0xad, 0x5b, 0x35, 0x6d, 0x4b, 0xf3, 0x36,
	0x3a, 0x53, 0x00, 0xf1, 0x1b, 0x58, 0x16, 0x59, 0x1d, 0xea, 0x8d, 0x98, 0xc9, 0xa7, 0x28, 0x55,
	0x41, 0x71, 0x1d, 0x96, 0xae, 0x48, 0xa5, 0x95, 0x1e, 0xf9, 0x36, 0x6c, 0x9d, 0x53, 0xe4, 0x5d,
	0xcd, 0x6b, 0x69, 0xe5, 0x6f, 0xc9, 0xff, 0x82, 0x8f, 0x9b, 0x46, 0x21, 0x74, 0xef, 0xbe, 0x85,
	0xd5, 0x8a, 0x6c, 0x14, 0xef, 0xb4, 0x61, 0x4a, 0xc4, 0xa9, 0x02, 0x71, 0x07, 0xba, 0x91, 0x10,
	0xb2, 0xe8, 0xe2, 0x5a, 0xf9, 0x72, 0xee, 0xc4, 0x9f, 0x44, 0x4e, 0x66, 0xe4, 0x4f, 0xe1, 0x93,
	0x96, 0xe0, 0x45, 0x76, 0x67, 0x30, 0x1c, 0x4b, 0x11, 0x51, 0xc5, 0xae, 0x47, 0xfc, 0x3b, 0xe8,
	0x57, 0x02, 0xaa, 0xbe, 0x9a, 0x53, 0xab, 0x21, 0x39, 0xc1, 0x96, 0x81, 0xb5, 0x18, 0xee, 0x9d,
	0xfa, 0x70, 0x1b, 0xbf, 0x8b, 0xc5, 0xaa, 0xbb, 0x86, 0xb5, 0x31, 0xc9, 0x23, 0xa2, 0xd3, 0xd9,
	0x87, 0x76, 0x73, 0x81, 0x89, 0xe4, 0x67, 0xd0, 0x7f, 0x45, 0xf3, 0xb2, 0x90, 0xbb, 0x25, 0x60,
	0x01, 0xd6, 0xfd, 0x7f, 0x7b, 0xb0, 0x7c, 0x9a, 0x5b, 0x70, 0x1f, 0xba, 0xe9, 0xb6, 0xc0, 0x47,
	0x1a, 0x5b, 0x2e, 0x0f, 0x66, 0xba, 0xe4, 0x16, 0xfe, 0x0c, 0x83, 0xc6, 0x16, 0xc1, 0x67, 0x1a,
	0xdb, 0xb6, 0x61, 0x98, 0xe1, 0x5b, 0xe7, 0x16, 0x1e, 0x40, 0x37, 0xad, 0x11, 0x0d, 0x56, 0xb6,
	0xa9, 0xef, 0xaa, 0x6d, 0xe0, 0x16, 0x3e, 0x87, 0x9e, 0x16, 0x6c, 0xdc, 0x32, 0x88, 0x7b, 0x11,
	0x92, 0x99, 0x4c, 0x9a, 0xe5, 0x1c, 0xd6, 0xde, 0x91, 0x7d, 0x7c, 0xda, 0x74, 0xa8, 0x2d, 0x17,
	0x36, 0x6a, 0x07, 0x68, 0xde, 0x2f, 0xa0, 0x77, 0x44, 0xd2, 0x9d, 0xa6, 0xe3, 0x85, 0x35, 0x01,
	0x64, 0xf5, 0xd9, 0xcb, 0x2a, 0x59, 0xad, 0x6c, 0x11, 0xdc, 0x7e, 0x57, 0x78, 0x2b, 0xbb, 0x85,
	0xb5, 0xa8, 0x32, 0xb7, 0xf0, 0x05, 0x40, 0x29, 0xea, 0x58, 0x56, 0xdd, 0xd8, 0x22, 0x6c, 0xdb,
	0x68, 0xd3, 0xa9, 0x1f, 0x02, 0x1e, 0x25, 0xbe, 0x3f, 0x76, 0x27, 0x41, 0x85, 0xb0, 0x31, 0x6b,
	0x77, 0xd1, 0xfc, 0x0a, 0x8f, 0x0c, 0x0b, 0x02, 0x77, 0xca, 0xea, 0x5b, 0xd7, 0x07, 0x6b, 0x97,
	0x51, 0x6e, 0xe1, 0x1b, 0xd8, 0x30, 0xad, 0x05, 0xfc, 0xb4, 0x85, 0xb9, 0xb6, 0x35, 0xd8, 0x7b,
	0xd4, 0x96, 0x5b, 0xf8, 0x23, 0x3c, 0xac, 0x3b, 0xe3, 0x93, 0x16, 0x56, 0xc5, 0x67, 0xfc, 0xcc,
	0xb9, 0x85, 0xdf, 0xc3, 0x8a, 0x92, 0x09, 0x1c, 0x56, 0xe6, 0xba, 0xa6, 0x1c, 0xad, 0xde, 0x27,
	0x30, 0x68, 0x48, 0x28, 0x1a, 0xc1, 0x8c, 0x97, 0x0f, 0xd3, 0xba, 0x0e, 0x2c, 0x9c, 0xc2, 0xa6,
	0x51, 0x93, 0xf1, 0xb3, 0xf7, 0xb8, 0x97, 0x0b, 0x83, 0x7d, 0x7e, 0x17, 0x4c, 0x47, 0xfa, 0x0d,
	0x06, 0x0d, 0x19, 0xae, 0x48, 0x43, 0x9b, 0xf0, 0x57, 0xea, 0x68, 0x55, 0x71, 0x6e, 0x5d, 0xdc,
	0xcf, 0xfe, 0xe1, 0x1e, 0xfc, 0x17, 0x00, 0x00, 0xff, 0xff, 0x0b, 0x2c, 0x15, 0xed, 0x33, 0x0b,
	0x00, 0x00,
}
