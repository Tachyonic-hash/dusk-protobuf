// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: note.proto

package rusk

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

type NoteType int32

const (
	NoteType_TRANSPARENT NoteType = 0
	NoteType_OBFUSCATED  NoteType = 1
)

var NoteType_name = map[int32]string{
	0: "TRANSPARENT",
	1: "OBFUSCATED",
}

var NoteType_value = map[string]int32{
	"TRANSPARENT": 0,
	"OBFUSCATED":  1,
}

func (x NoteType) String() string {
	return proto.EnumName(NoteType_name, int32(x))
}

func (NoteType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_640dafe07df50d4e, []int{0}
}

type Nullifier struct {
	H                    *Scalar  `protobuf:"bytes,1,opt,name=h,proto3" json:"h,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Nullifier) Reset()         { *m = Nullifier{} }
func (m *Nullifier) String() string { return proto.CompactTextString(m) }
func (*Nullifier) ProtoMessage()    {}
func (*Nullifier) Descriptor() ([]byte, []int) {
	return fileDescriptor_640dafe07df50d4e, []int{0}
}
func (m *Nullifier) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Nullifier.Unmarshal(m, b)
}
func (m *Nullifier) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Nullifier.Marshal(b, m, deterministic)
}
func (m *Nullifier) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Nullifier.Merge(m, src)
}
func (m *Nullifier) XXX_Size() int {
	return xxx_messageInfo_Nullifier.Size(m)
}
func (m *Nullifier) XXX_DiscardUnknown() {
	xxx_messageInfo_Nullifier.DiscardUnknown(m)
}

var xxx_messageInfo_Nullifier proto.InternalMessageInfo

func (m *Nullifier) GetH() *Scalar {
	if m != nil {
		return m.H
	}
	return nil
}

type Note struct {
	NoteType        NoteType         `protobuf:"varint,1,opt,name=note_type,json=noteType,proto3,enum=rusk.NoteType" json:"note_type,omitempty"`
	Pos             uint64           `protobuf:"fixed64,2,opt,name=pos,proto3" json:"pos,omitempty"`
	Nonce           *Nonce           `protobuf:"bytes,3,opt,name=nonce,proto3" json:"nonce,omitempty"`
	RG              *CompressedPoint `protobuf:"bytes,4,opt,name=r_g,json=rG,proto3" json:"r_g,omitempty"`
	PkR             *CompressedPoint `protobuf:"bytes,5,opt,name=pk_r,json=pkR,proto3" json:"pk_r,omitempty"`
	ValueCommitment *Scalar          `protobuf:"bytes,6,opt,name=value_commitment,json=valueCommitment,proto3" json:"value_commitment,omitempty"`
	// Types that are valid to be assigned to BlindingFactor:
	//	*Note_TransparentBlindingFactor
	//	*Note_EncryptedBlindingFactor
	BlindingFactor isNote_BlindingFactor `protobuf_oneof:"blinding_factor"`
	// Types that are valid to be assigned to Value:
	//	*Note_TransparentValue
	//	*Note_EncryptedValue
	Value                isNote_Value `protobuf_oneof:"value"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *Note) Reset()         { *m = Note{} }
func (m *Note) String() string { return proto.CompactTextString(m) }
func (*Note) ProtoMessage()    {}
func (*Note) Descriptor() ([]byte, []int) {
	return fileDescriptor_640dafe07df50d4e, []int{1}
}
func (m *Note) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Note.Unmarshal(m, b)
}
func (m *Note) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Note.Marshal(b, m, deterministic)
}
func (m *Note) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Note.Merge(m, src)
}
func (m *Note) XXX_Size() int {
	return xxx_messageInfo_Note.Size(m)
}
func (m *Note) XXX_DiscardUnknown() {
	xxx_messageInfo_Note.DiscardUnknown(m)
}

var xxx_messageInfo_Note proto.InternalMessageInfo

type isNote_BlindingFactor interface {
	isNote_BlindingFactor()
}
type isNote_Value interface {
	isNote_Value()
}

type Note_TransparentBlindingFactor struct {
	TransparentBlindingFactor *Scalar `protobuf:"bytes,7,opt,name=transparent_blinding_factor,json=transparentBlindingFactor,proto3,oneof" json:"transparent_blinding_factor,omitempty"`
}
type Note_EncryptedBlindingFactor struct {
	EncryptedBlindingFactor []byte `protobuf:"bytes,8,opt,name=encrypted_blinding_factor,json=encryptedBlindingFactor,proto3,oneof" json:"encrypted_blinding_factor,omitempty"`
}
type Note_TransparentValue struct {
	TransparentValue uint64 `protobuf:"fixed64,9,opt,name=transparent_value,json=transparentValue,proto3,oneof" json:"transparent_value,omitempty"`
}
type Note_EncryptedValue struct {
	EncryptedValue []byte `protobuf:"bytes,10,opt,name=encrypted_value,json=encryptedValue,proto3,oneof" json:"encrypted_value,omitempty"`
}

func (*Note_TransparentBlindingFactor) isNote_BlindingFactor() {}
func (*Note_EncryptedBlindingFactor) isNote_BlindingFactor()   {}
func (*Note_TransparentValue) isNote_Value()                   {}
func (*Note_EncryptedValue) isNote_Value()                     {}

func (m *Note) GetBlindingFactor() isNote_BlindingFactor {
	if m != nil {
		return m.BlindingFactor
	}
	return nil
}
func (m *Note) GetValue() isNote_Value {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *Note) GetNoteType() NoteType {
	if m != nil {
		return m.NoteType
	}
	return NoteType_TRANSPARENT
}

func (m *Note) GetPos() uint64 {
	if m != nil {
		return m.Pos
	}
	return 0
}

func (m *Note) GetNonce() *Nonce {
	if m != nil {
		return m.Nonce
	}
	return nil
}

func (m *Note) GetRG() *CompressedPoint {
	if m != nil {
		return m.RG
	}
	return nil
}

func (m *Note) GetPkR() *CompressedPoint {
	if m != nil {
		return m.PkR
	}
	return nil
}

func (m *Note) GetValueCommitment() *Scalar {
	if m != nil {
		return m.ValueCommitment
	}
	return nil
}

func (m *Note) GetTransparentBlindingFactor() *Scalar {
	if x, ok := m.GetBlindingFactor().(*Note_TransparentBlindingFactor); ok {
		return x.TransparentBlindingFactor
	}
	return nil
}

func (m *Note) GetEncryptedBlindingFactor() []byte {
	if x, ok := m.GetBlindingFactor().(*Note_EncryptedBlindingFactor); ok {
		return x.EncryptedBlindingFactor
	}
	return nil
}

func (m *Note) GetTransparentValue() uint64 {
	if x, ok := m.GetValue().(*Note_TransparentValue); ok {
		return x.TransparentValue
	}
	return 0
}

func (m *Note) GetEncryptedValue() []byte {
	if x, ok := m.GetValue().(*Note_EncryptedValue); ok {
		return x.EncryptedValue
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Note) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Note_TransparentBlindingFactor)(nil),
		(*Note_EncryptedBlindingFactor)(nil),
		(*Note_TransparentValue)(nil),
		(*Note_EncryptedValue)(nil),
	}
}

type DecryptedNote struct {
	NoteType        NoteType         `protobuf:"varint,1,opt,name=note_type,json=noteType,proto3,enum=rusk.NoteType" json:"note_type,omitempty"`
	Pos             uint64           `protobuf:"fixed64,2,opt,name=pos,proto3" json:"pos,omitempty"`
	Value           uint64           `protobuf:"fixed64,3,opt,name=value,proto3" json:"value,omitempty"`
	Nonce           *Nonce           `protobuf:"bytes,4,opt,name=nonce,proto3" json:"nonce,omitempty"`
	RG              *CompressedPoint `protobuf:"bytes,5,opt,name=r_g,json=rG,proto3" json:"r_g,omitempty"`
	PkR             *CompressedPoint `protobuf:"bytes,6,opt,name=pk_r,json=pkR,proto3" json:"pk_r,omitempty"`
	ValueCommitment *Scalar          `protobuf:"bytes,7,opt,name=value_commitment,json=valueCommitment,proto3" json:"value_commitment,omitempty"`
	BlindingFactor  *Scalar          `protobuf:"bytes,8,opt,name=blinding_factor,json=blindingFactor,proto3" json:"blinding_factor,omitempty"`
	// Types that are valid to be assigned to RawBlindingFactor:
	//	*DecryptedNote_TransparentBlindingFactor
	//	*DecryptedNote_EncryptedBlindingFactor
	RawBlindingFactor isDecryptedNote_RawBlindingFactor `protobuf_oneof:"rawBlindingFactor"`
	// Types that are valid to be assigned to RawValue:
	//	*DecryptedNote_TransparentValue
	//	*DecryptedNote_EncryptedValue
	RawValue             isDecryptedNote_RawValue `protobuf_oneof:"rawValue"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *DecryptedNote) Reset()         { *m = DecryptedNote{} }
func (m *DecryptedNote) String() string { return proto.CompactTextString(m) }
func (*DecryptedNote) ProtoMessage()    {}
func (*DecryptedNote) Descriptor() ([]byte, []int) {
	return fileDescriptor_640dafe07df50d4e, []int{2}
}
func (m *DecryptedNote) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DecryptedNote.Unmarshal(m, b)
}
func (m *DecryptedNote) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DecryptedNote.Marshal(b, m, deterministic)
}
func (m *DecryptedNote) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DecryptedNote.Merge(m, src)
}
func (m *DecryptedNote) XXX_Size() int {
	return xxx_messageInfo_DecryptedNote.Size(m)
}
func (m *DecryptedNote) XXX_DiscardUnknown() {
	xxx_messageInfo_DecryptedNote.DiscardUnknown(m)
}

var xxx_messageInfo_DecryptedNote proto.InternalMessageInfo

type isDecryptedNote_RawBlindingFactor interface {
	isDecryptedNote_RawBlindingFactor()
}
type isDecryptedNote_RawValue interface {
	isDecryptedNote_RawValue()
}

type DecryptedNote_TransparentBlindingFactor struct {
	TransparentBlindingFactor *Scalar `protobuf:"bytes,9,opt,name=transparent_blinding_factor,json=transparentBlindingFactor,proto3,oneof" json:"transparent_blinding_factor,omitempty"`
}
type DecryptedNote_EncryptedBlindingFactor struct {
	EncryptedBlindingFactor []byte `protobuf:"bytes,10,opt,name=encrypted_blinding_factor,json=encryptedBlindingFactor,proto3,oneof" json:"encrypted_blinding_factor,omitempty"`
}
type DecryptedNote_TransparentValue struct {
	TransparentValue uint64 `protobuf:"fixed64,11,opt,name=transparent_value,json=transparentValue,proto3,oneof" json:"transparent_value,omitempty"`
}
type DecryptedNote_EncryptedValue struct {
	EncryptedValue []byte `protobuf:"bytes,12,opt,name=encrypted_value,json=encryptedValue,proto3,oneof" json:"encrypted_value,omitempty"`
}

func (*DecryptedNote_TransparentBlindingFactor) isDecryptedNote_RawBlindingFactor() {}
func (*DecryptedNote_EncryptedBlindingFactor) isDecryptedNote_RawBlindingFactor()   {}
func (*DecryptedNote_TransparentValue) isDecryptedNote_RawValue()                   {}
func (*DecryptedNote_EncryptedValue) isDecryptedNote_RawValue()                     {}

func (m *DecryptedNote) GetRawBlindingFactor() isDecryptedNote_RawBlindingFactor {
	if m != nil {
		return m.RawBlindingFactor
	}
	return nil
}
func (m *DecryptedNote) GetRawValue() isDecryptedNote_RawValue {
	if m != nil {
		return m.RawValue
	}
	return nil
}

func (m *DecryptedNote) GetNoteType() NoteType {
	if m != nil {
		return m.NoteType
	}
	return NoteType_TRANSPARENT
}

func (m *DecryptedNote) GetPos() uint64 {
	if m != nil {
		return m.Pos
	}
	return 0
}

func (m *DecryptedNote) GetValue() uint64 {
	if m != nil {
		return m.Value
	}
	return 0
}

func (m *DecryptedNote) GetNonce() *Nonce {
	if m != nil {
		return m.Nonce
	}
	return nil
}

func (m *DecryptedNote) GetRG() *CompressedPoint {
	if m != nil {
		return m.RG
	}
	return nil
}

func (m *DecryptedNote) GetPkR() *CompressedPoint {
	if m != nil {
		return m.PkR
	}
	return nil
}

func (m *DecryptedNote) GetValueCommitment() *Scalar {
	if m != nil {
		return m.ValueCommitment
	}
	return nil
}

func (m *DecryptedNote) GetBlindingFactor() *Scalar {
	if m != nil {
		return m.BlindingFactor
	}
	return nil
}

func (m *DecryptedNote) GetTransparentBlindingFactor() *Scalar {
	if x, ok := m.GetRawBlindingFactor().(*DecryptedNote_TransparentBlindingFactor); ok {
		return x.TransparentBlindingFactor
	}
	return nil
}

func (m *DecryptedNote) GetEncryptedBlindingFactor() []byte {
	if x, ok := m.GetRawBlindingFactor().(*DecryptedNote_EncryptedBlindingFactor); ok {
		return x.EncryptedBlindingFactor
	}
	return nil
}

func (m *DecryptedNote) GetTransparentValue() uint64 {
	if x, ok := m.GetRawValue().(*DecryptedNote_TransparentValue); ok {
		return x.TransparentValue
	}
	return 0
}

func (m *DecryptedNote) GetEncryptedValue() []byte {
	if x, ok := m.GetRawValue().(*DecryptedNote_EncryptedValue); ok {
		return x.EncryptedValue
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*DecryptedNote) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*DecryptedNote_TransparentBlindingFactor)(nil),
		(*DecryptedNote_EncryptedBlindingFactor)(nil),
		(*DecryptedNote_TransparentValue)(nil),
		(*DecryptedNote_EncryptedValue)(nil),
	}
}

func init() {
	proto.RegisterEnum("rusk.NoteType", NoteType_name, NoteType_value)
	proto.RegisterType((*Nullifier)(nil), "rusk.Nullifier")
	proto.RegisterType((*Note)(nil), "rusk.Note")
	proto.RegisterType((*DecryptedNote)(nil), "rusk.DecryptedNote")
}

func init() { proto.RegisterFile("note.proto", fileDescriptor_640dafe07df50d4e) }

var fileDescriptor_640dafe07df50d4e = []byte{
	// 510 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x94, 0xcd, 0x6e, 0xda, 0x4c,
	0x14, 0x86, 0x71, 0x30, 0x7f, 0x07, 0x3e, 0x30, 0xf3, 0xb5, 0xaa, 0x93, 0x6e, 0x28, 0x8b, 0xd6,
	0x4d, 0x14, 0x90, 0x52, 0x55, 0xdd, 0x74, 0x83, 0x49, 0x52, 0x56, 0x6e, 0x34, 0xd0, 0x2e, 0xba,
	0xb1, 0xfc, 0x33, 0x18, 0x0b, 0x7b, 0xc6, 0x1a, 0x8f, 0x1b, 0x71, 0x9b, 0xbd, 0x99, 0x6e, 0x2b,
	0x8f, 0x09, 0x21, 0xd0, 0x36, 0x69, 0x94, 0xdd, 0xd8, 0xe7, 0x79, 0xcf, 0x19, 0xcf, 0x3c, 0x32,
	0x00, 0x65, 0x82, 0x0c, 0x12, 0xce, 0x04, 0x43, 0x2a, 0xcf, 0xd2, 0xe5, 0x51, 0x73, 0x1e, 0x92,
	0xc8, 0x2f, 0x5e, 0xf5, 0xdf, 0x40, 0xc3, 0xca, 0xa2, 0x28, 0x9c, 0x87, 0x84, 0xa3, 0x23, 0x50,
	0x16, 0xba, 0xd2, 0x53, 0x8c, 0xe6, 0x59, 0x6b, 0x90, 0xb3, 0x83, 0xa9, 0xe7, 0x44, 0x0e, 0xc7,
	0xca, 0xa2, 0xff, 0xb3, 0x0c, 0xaa, 0xc5, 0x04, 0x41, 0x27, 0xd0, 0xc8, 0x5b, 0xda, 0x62, 0x95,
	0x10, 0x09, 0xb7, 0xcf, 0xda, 0x05, 0x9c, 0x97, 0x67, 0xab, 0x84, 0xe0, 0x3a, 0x5d, 0xaf, 0x90,
	0x06, 0xe5, 0x84, 0xa5, 0xfa, 0x41, 0x4f, 0x31, 0xaa, 0x38, 0x5f, 0xa2, 0x57, 0x50, 0xa1, 0x8c,
	0x7a, 0x44, 0x2f, 0xcb, 0x39, 0xcd, 0x9b, 0x28, 0xf5, 0x08, 0x2e, 0x2a, 0xe8, 0x35, 0x94, 0xb9,
	0x1d, 0xe8, 0xaa, 0x04, 0x9e, 0x17, 0xc0, 0x98, 0xc5, 0x09, 0x27, 0x69, 0x4a, 0xfc, 0x2b, 0x16,
	0x52, 0x81, 0x0f, 0xf8, 0x27, 0x64, 0x80, 0x9a, 0x2c, 0x6d, 0xae, 0x57, 0xfe, 0x06, 0x96, 0x93,
	0x25, 0x46, 0x1f, 0x40, 0xfb, 0xee, 0x44, 0x19, 0xb1, 0x3d, 0x16, 0xc7, 0xa1, 0x88, 0x09, 0x15,
	0x7a, 0xf5, 0x37, 0xdf, 0xd9, 0x91, 0xd4, 0x78, 0x03, 0x21, 0x0b, 0x5e, 0x0a, 0xee, 0xd0, 0x34,
	0x71, 0x38, 0xa1, 0xc2, 0x76, 0xa3, 0x90, 0xfa, 0x21, 0x0d, 0xec, 0xb9, 0xe3, 0x09, 0xc6, 0xf5,
	0xda, 0x7e, 0x8f, 0x49, 0x09, 0x1f, 0x6e, 0x45, 0xcc, 0x75, 0xe2, 0x52, 0x06, 0xd0, 0x47, 0x38,
	0x24, 0xd4, 0xe3, 0xab, 0x44, 0x10, 0x7f, 0xaf, 0x5b, 0xbd, 0xa7, 0x18, 0xad, 0x49, 0x09, 0xbf,
	0xd8, 0x20, 0x3b, 0xe9, 0x53, 0xe8, 0x6e, 0xef, 0x46, 0x6e, 0x56, 0x6f, 0xe4, 0x67, 0x3b, 0x51,
	0xb0, 0xb6, 0x55, 0xfa, 0x9a, 0x57, 0xd0, 0x5b, 0xe8, 0xdc, 0x0e, 0x2b, 0x60, 0x90, 0x23, 0x14,
	0xdc, 0xde, 0x14, 0x24, 0x6a, 0x76, 0xa1, 0xb3, 0xb3, 0x1b, 0xb3, 0x06, 0x15, 0x99, 0xe9, 0xff,
	0x50, 0xe1, 0xbf, 0x73, 0xb2, 0xc6, 0x9f, 0x42, 0x81, 0x67, 0xeb, 0xce, 0x52, 0x81, 0x2a, 0x2e,
	0x1e, 0x6e, 0xc5, 0x50, 0xef, 0x13, 0xa3, 0xf2, 0x50, 0x31, 0xaa, 0x8f, 0x12, 0xa3, 0xf6, 0x10,
	0x31, 0xde, 0xef, 0x1d, 0x98, 0xbc, 0xbe, 0xdd, 0x5c, 0xdb, 0xbd, 0x7b, 0x83, 0xf7, 0xf8, 0xd4,
	0x78, 0x52, 0x9f, 0xe0, 0x51, 0x3e, 0x35, 0xff, 0xc5, 0xa7, 0xd6, 0x1f, 0x7c, 0xfa, 0x1f, 0xba,
	0xdc, 0xb9, 0xbe, 0x3b, 0xce, 0x04, 0xa8, 0x73, 0xe7, 0x5a, 0x02, 0xc7, 0x27, 0x50, 0xbf, 0x71,
	0x05, 0x75, 0xa0, 0x39, 0xc3, 0x23, 0x6b, 0x7a, 0x35, 0xc2, 0x17, 0xd6, 0x4c, 0x2b, 0xa1, 0x36,
	0xc0, 0x67, 0xf3, 0xf2, 0xcb, 0x74, 0x3c, 0x9a, 0x5d, 0x9c, 0x6b, 0x8a, 0x79, 0xfc, 0xcd, 0x08,
	0x42, 0xb1, 0xc8, 0xdc, 0x81, 0xc7, 0xe2, 0xa1, 0x9f, 0xa5, 0xcb, 0x53, 0xf9, 0xf7, 0x72, 0xb3,
	0xf9, 0xd0, 0xc9, 0x04, 0x0b, 0x08, 0x1d, 0x06, 0x6c, 0x98, 0x9f, 0x9a, 0x5b, 0x95, 0x95, 0x77,
	0xbf, 0x02, 0x00, 0x00, 0xff, 0xff, 0x92, 0x04, 0xc7, 0x6c, 0xf8, 0x04, 0x00, 0x00,
}
