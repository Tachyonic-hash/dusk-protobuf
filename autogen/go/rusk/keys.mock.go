// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: keys.proto

package rusk

import (
	fmt "fmt"
	math "math"
	proto "github.com/gogo/protobuf/proto"
	context "context"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type KeysMock struct{}

func (m *KeysMock) GenerateKeys(ctx context.Context, req *GenerateKeysRequest) (*GenerateKeysResponse, error) {
	res :=
		&GenerateKeysResponse{
			Sk: &SecretKey{},
			Vk: &ViewKey{},
			Pk: &PublicKey{},
		}
	return res, nil
}
func (m *KeysMock) GenerateStealthAddress(ctx context.Context, req *PublicKey) (*StealthAddress, error) {
	res :=
		&StealthAddress{}
	return res, nil
}
