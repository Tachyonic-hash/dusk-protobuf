// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: bid.proto

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

type BidServiceMock struct{}

func (m *BidServiceMock) NewBid(ctx context.Context, req *BidTransactionRequest) (*BidTransaction, error) {
	res :=
		&BidTransaction{
			Bid: &Bid{
				EncryptedData: &PoseidonCipher{},
				HashedSecret:  &BlsScalar{},
				Nonce:         &BlsScalar{},
				PkR: &StealthAddress{
					RG:  &JubJubCompressed{},
					PkR: &JubJubCompressed{},
				},
				Commitment:          &JubJubCompressed{},
				ElegibilityTs:       &BlsScalar{},
				ExpirationTs:        &BlsScalar{},
				BidTreeStorageIndex: 654,
			},
			BidHash: &BlsScalar{},
		}
	return res, nil
}
func (m *BidServiceMock) FindBid(ctx context.Context, req *FindBidRequest) (*BidList, error) {
	res :=
		&BidList{
			BidList: []*Bid{
				&Bid{
					EncryptedData: &PoseidonCipher{},
					HashedSecret:  &BlsScalar{},
					Nonce:         &BlsScalar{},
					PkR: &StealthAddress{
						RG:  &JubJubCompressed{},
						PkR: &JubJubCompressed{},
					},
					Commitment:          &JubJubCompressed{},
					ElegibilityTs:       &BlsScalar{},
					ExpirationTs:        &BlsScalar{},
					BidTreeStorageIndex: 961,
				},
				&Bid{
					EncryptedData: &PoseidonCipher{},
					HashedSecret:  &BlsScalar{},
					Nonce:         &BlsScalar{},
					PkR: &StealthAddress{
						RG:  &JubJubCompressed{},
						PkR: &JubJubCompressed{},
					},
					Commitment:          &JubJubCompressed{},
					ElegibilityTs:       &BlsScalar{},
					ExpirationTs:        &BlsScalar{},
					BidTreeStorageIndex: 450,
				},
				&Bid{
					EncryptedData: &PoseidonCipher{},
					HashedSecret:  &BlsScalar{},
					Nonce:         &BlsScalar{},
					PkR: &StealthAddress{
						RG:  &JubJubCompressed{},
						PkR: &JubJubCompressed{},
					},
					Commitment:          &JubJubCompressed{},
					ElegibilityTs:       &BlsScalar{},
					ExpirationTs:        &BlsScalar{},
					BidTreeStorageIndex: 231,
				},
				&Bid{
					EncryptedData: &PoseidonCipher{},
					HashedSecret:  &BlsScalar{},
					Nonce:         &BlsScalar{},
					PkR: &StealthAddress{
						RG:  &JubJubCompressed{},
						PkR: &JubJubCompressed{},
					},
					Commitment:          &JubJubCompressed{},
					ElegibilityTs:       &BlsScalar{},
					ExpirationTs:        &BlsScalar{},
					BidTreeStorageIndex: 226,
				},
				&Bid{
					EncryptedData: &PoseidonCipher{},
					HashedSecret:  &BlsScalar{},
					Nonce:         &BlsScalar{},
					PkR: &StealthAddress{
						RG:  &JubJubCompressed{},
						PkR: &JubJubCompressed{},
					},
					Commitment:          &JubJubCompressed{},
					ElegibilityTs:       &BlsScalar{},
					ExpirationTs:        &BlsScalar{},
					BidTreeStorageIndex: 598,
				},
				&Bid{
					EncryptedData: &PoseidonCipher{},
					HashedSecret:  &BlsScalar{},
					Nonce:         &BlsScalar{},
					PkR: &StealthAddress{
						RG:  &JubJubCompressed{},
						PkR: &JubJubCompressed{},
					},
					Commitment:          &JubJubCompressed{},
					ElegibilityTs:       &BlsScalar{},
					ExpirationTs:        &BlsScalar{},
					BidTreeStorageIndex: 637,
				},
				&Bid{
					EncryptedData: &PoseidonCipher{},
					HashedSecret:  &BlsScalar{},
					Nonce:         &BlsScalar{},
					PkR: &StealthAddress{
						RG:  &JubJubCompressed{},
						PkR: &JubJubCompressed{},
					},
					Commitment:          &JubJubCompressed{},
					ElegibilityTs:       &BlsScalar{},
					ExpirationTs:        &BlsScalar{},
					BidTreeStorageIndex: 516,
				},
				&Bid{
					EncryptedData: &PoseidonCipher{},
					HashedSecret:  &BlsScalar{},
					Nonce:         &BlsScalar{},
					PkR: &StealthAddress{
						RG:  &JubJubCompressed{},
						PkR: &JubJubCompressed{},
					},
					Commitment:          &JubJubCompressed{},
					ElegibilityTs:       &BlsScalar{},
					ExpirationTs:        &BlsScalar{},
					BidTreeStorageIndex: 139,
				},
			},
			BidHashList: []*BlsScalar{
				&BlsScalar{},
				&BlsScalar{},
				&BlsScalar{},
				&BlsScalar{},
				&BlsScalar{},
				&BlsScalar{},
				&BlsScalar{},
				&BlsScalar{},
			},
		}
	return res, nil
}
