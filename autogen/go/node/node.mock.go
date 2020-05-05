// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: node.proto

package node

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

type WalletMock struct{}

func (m *WalletMock) CreateWallet(ctx context.Context, req *CreateRequest) (*LoadResponse, error) {
	res :=
		&LoadResponse{
			Key: &PubKey{},
		}
	return res, nil
}
func (m *WalletMock) LoadWallet(ctx context.Context, req *LoadRequest) (*LoadResponse, error) {
	res :=
		&LoadResponse{
			Key: &PubKey{},
		}
	return res, nil
}
func (m *WalletMock) CreateFromSeed(ctx context.Context, req *CreateRequest) (*LoadResponse, error) {
	res :=
		&LoadResponse{
			Key: &PubKey{},
		}
	return res, nil
}
func (m *WalletMock) ClearWalletDatabase(ctx context.Context, req *EmptyRequest) (*GenericResponse, error) {
	res :=
		&GenericResponse{
			Response: "numquam",
		}
	return res, nil
}
func (m *WalletMock) GetWalletStatus(ctx context.Context, req *EmptyRequest) (*WalletStatusResponse, error) {
	res :=
		&WalletStatusResponse{
			Loaded: false,
		}
	return res, nil
}
func (m *WalletMock) GetAddress(ctx context.Context, req *EmptyRequest) (*LoadResponse, error) {
	res :=
		&LoadResponse{
			Key: &PubKey{},
		}
	return res, nil
}
func (m *WalletMock) GetBalance(ctx context.Context, req *EmptyRequest) (*BalanceResponse, error) {
	res :=
		&BalanceResponse{}
	return res, nil
}
func (m *WalletMock) GetTxHistory(ctx context.Context, req *EmptyRequest) (*TxHistoryResponse, error) {
	res :=
		&TxHistoryResponse{
			Records: []*TxRecord{
				&TxRecord{
					Direction: 1,
					Timestamp: 494,
					Type:      5,
				},
				&TxRecord{
					Direction: 1,
					Timestamp: 487,
					Type:      3,
				},
				&TxRecord{
					Direction: 1,
					Timestamp: 175,
					Type:      4,
				},
				&TxRecord{
					Direction: 0,
					Timestamp: 755,
					Type:      0,
				},
				&TxRecord{
					Direction: 0,
					Timestamp: 721,
					Type:      0,
				},
				&TxRecord{
					Direction: 1,
					Timestamp: 876,
					Type:      1,
				},
				&TxRecord{
					Direction: 0,
					Timestamp: 932,
					Type:      3,
				},
				&TxRecord{
					Direction: 1,
					Timestamp: 792,
					Type:      0,
				},
			},
		}
	return res, nil
}

type MempoolMock struct{}

func (m *MempoolMock) GetUnconfirmedBalance(ctx context.Context, req *EmptyRequest) (*BalanceResponse, error) {
	res :=
		&BalanceResponse{}
	return res, nil
}
func (m *MempoolMock) SelectTx(ctx context.Context, req *SelectRequest) (*SelectResponse, error) {
	res :=
		&SelectResponse{
			Result: []*Tx{
				&Tx{
					Type: 1,
					Id:   "6179d3ca-f4d2-47cb-930b-a4236a35c8f7",
				},
				&Tx{
					Type: 4,
					Id:   "fe6f1fe8-422e-4668-842e-9ca9e0669d8d",
				},
				&Tx{
					Type: 5,
					Id:   "03610ed5-3740-4d72-a9b3-f330b2488c43",
				},
				&Tx{
					Type: 4,
					Id:   "bbe5c773-0f23-4eff-afd2-f6eda9a6eb34",
				},
				&Tx{
					Type: 1,
					Id:   "153078d8-1490-4341-810f-bb8291ebe9a9",
				},
				&Tx{
					Type: 2,
					Id:   "4220d263-b065-438d-ae9f-ec2d29d8567f",
				},
				&Tx{
					Type: 4,
					Id:   "c4a39a24-e767-46d9-b868-c601d04eb0c0",
				},
				&Tx{
					Type: 5,
					Id:   "84e3c45c-acbb-4c9b-89c6-75b85fdb3e27",
				},
			},
		}
	return res, nil
}

type ChainMock struct{}

func (m *ChainMock) RebuildChain(ctx context.Context, req *EmptyRequest) (*GenericResponse, error) {
	res :=
		&GenericResponse{
			Response: "dolorem",
		}
	return res, nil
}
func (m *ChainMock) GetSyncProgress(ctx context.Context, req *EmptyRequest) (*SyncProgressResponse, error) {
	res :=
		&SyncProgressResponse{
			Progress: 575.4365,
		}
	return res, nil
}

type TransactorMock struct{}

func (m *TransactorMock) CallContract(ctx context.Context, req *CallContractRequest) (*TransactionResponse, error) {
	res :=
		&TransactionResponse{}
	return res, nil
}
func (m *TransactorMock) Transfer(ctx context.Context, req *TransferRequest) (*TransactionResponse, error) {
	res :=
		&TransactionResponse{}
	return res, nil
}
func (m *TransactorMock) Bid(ctx context.Context, req *BidRequest) (*TransactionResponse, error) {
	res :=
		&TransactionResponse{}
	return res, nil
}
func (m *TransactorMock) Stake(ctx context.Context, req *StakeRequest) (*TransactionResponse, error) {
	res :=
		&TransactionResponse{}
	return res, nil
}

type MaintainerMock struct{}

func (m *MaintainerMock) AutomateConsensusTxs(ctx context.Context, req *EmptyRequest) (*GenericResponse, error) {
	res :=
		&GenericResponse{
			Response: "praesentium",
		}
	return res, nil
}
