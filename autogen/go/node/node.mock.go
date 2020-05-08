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
			Response: "dolor",
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
					Direction: 0,
					Timestamp: 843,
					Type:      2,
				},
				&TxRecord{
					Direction: 0,
					Timestamp: 186,
					Type:      0,
				},
				&TxRecord{
					Direction: 0,
					Timestamp: 105,
					Type:      4,
				},
				&TxRecord{
					Direction: 0,
					Timestamp: 83,
					Type:      5,
				},
				&TxRecord{
					Direction: 1,
					Timestamp: 366,
					Type:      0,
				},
				&TxRecord{
					Direction: 1,
					Timestamp: 854,
					Type:      2,
				},
				&TxRecord{
					Direction: 1,
					Timestamp: 412,
					Type:      0,
				},
				&TxRecord{
					Direction: 0,
					Timestamp: 644,
					Type:      1,
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
					Type: 2,
					Id:   "510950e1-f3ed-4f39-b75f-ad46554ecd43",
				},
				&Tx{
					Type: 1,
					Id:   "aae4533e-ece3-4cd3-9cac-25e11f73e8d5",
				},
				&Tx{
					Type: 0,
					Id:   "7ac5cb5d-1998-469c-b5fb-e9aeb459a8e2",
				},
				&Tx{
					Type: 0,
					Id:   "87e4c51a-9ce3-454e-8891-47289fa4c6d0",
				},
				&Tx{
					Type: 1,
					Id:   "b9263f91-c4e8-4734-865a-e3fee70f5d81",
				},
				&Tx{
					Type: 3,
					Id:   "fc974268-889b-477c-aad8-1d9b43a00333",
				},
				&Tx{
					Type: 1,
					Id:   "c0ce763c-f79e-4431-accc-658240bbaa57",
				},
				&Tx{
					Type: 0,
					Id:   "d6ac38e9-1421-49a5-989f-127ea86617e6",
				},
			},
		}
	return res, nil
}

type ChainMock struct{}

func (m *ChainMock) RebuildChain(ctx context.Context, req *EmptyRequest) (*GenericResponse, error) {
	res :=
		&GenericResponse{
			Response: "aut",
		}
	return res, nil
}
func (m *ChainMock) GetSyncProgress(ctx context.Context, req *EmptyRequest) (*SyncProgressResponse, error) {
	res :=
		&SyncProgressResponse{
			Progress: 83.9822,
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
			Response: "aut",
		}
	return res, nil
}
