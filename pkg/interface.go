//
// Copyright (C) 2024 dszi
//
// This file may be modified and distributed under the terms
// of the MIT license. See the LICENSE file for details.
// Repository: https://github.com/dszi/go-tron
//

package pkg

import (
	"context"

	"github.com/dszi/go-tron/pb/api"
	"github.com/dszi/go-tron/pb/core"
)

// TronClient provides an interface for interacting with the TRON blockchain via gRPC.
// It includes methods for managing accounts, transactions, assets, resources, smart contracts, and network state.
//
// This interface abstracts the core functionality needed to interact with the TRON network.
type TronClient interface {
	// Account Management
	GetAccount(addr string) (*core.Account, error)
	GetAccountBalance(addr string) (int64, error)
	GetAccountResource(addr string) (*api.AccountResourceMessage, error)
	CreateAccount(from, addr string) (*api.TransactionExtention, error)
	UpdateAccount(from, accountName string) (*api.TransactionExtention, error)
	GetRewardInfo(addr string) (int64, error)

	// Transactions
	CreateTransaction(from, toAddress string, amount int64) (*api.TransactionExtention, error)
	BroadcastTransaction(tx *core.Transaction) (*api.Return, error)
	GetTransactionByID(id string) (*core.Transaction, error)
	GetTransactionInfoByID(id string) (*core.TransactionInfo, error)
	GetTransactionFromPending(id string) (*core.Transaction, error)
	GetTransactionListFromPending() (*api.TransactionIdList, error)
	TotalTransaction() (*api.NumberMessage, error)

	// Resource Management
	FreezeBalance(from, delegateTo string, resource core.ResourceCode, frozenBalance int64) (*api.TransactionExtention, error)
	UnfreezeBalance(from, delegateTo string, resource core.ResourceCode) (*api.TransactionExtention, error)
	WithdrawBalance(from string) (*api.TransactionExtention, error)
	UnfreezeAsset(from string) (*api.TransactionExtention, error)
	UnfreezeBalanceV2(from string, resource core.ResourceCode, unfreezeBalance int64) (*api.TransactionExtention, error)
	WithdrawExpireUnfreeze(from string, timestamp int64) (*api.TransactionExtention, error)
	DelegateResource(from, to string, resource core.ResourceCode, delegateBalance int64, lock bool, lockPeriod int64) (*api.TransactionExtention, error)
	UnDelegateResource(owner, receiver string, resource core.ResourceCode, delegateBalance int64, lock bool) (*api.TransactionExtention, error)
	CancelAllUnfreezeV2(from string) (*api.TransactionExtention, error)

	// Witnesses Management
	VoteWitnessAccount(from string, witnessMap map[string]int64) (*api.TransactionExtention, error)
	ListWitnesses() (*api.WitnessList, error)
	CreateWitness(from, urlStr string) (*api.TransactionExtention, error)
	UpdateWitness(from, urlStr string) (*api.TransactionExtention, error)
	GetBrokerageInfo(witness string) (float64, error)
	UpdateBrokerage(from string, brokerage int32) (*api.TransactionExtention, error)

	// Asset Management
	CreateAssetIssue(from, name, description, abbr, urlStr string, precision int32, totalSupply, startTime, endTime, FreeAssetNetLimit, PublicFreeAssetNetLimit int64, trxNum, icoNum, voteScore int32, frozenSupply map[string]string) (*api.TransactionExtention, error)
	GetAssetIssueList(page int64, limit ...int64) (*api.AssetIssueList, error)
	GetPaginatedAssetIssueList(page int64, limit ...int64) (*api.AssetIssueList, error)
	GetAssetIssueByAccount(address string) (*api.AssetIssueList, error)
	GetAssetIssueByName(name string) (*core.AssetIssueContract, error)
	GetAssetIssueById(id string) (*core.AssetIssueContract, error)
	TransferAsset(from, toAddress, assetName string, amount int64) (*api.TransactionExtention, error)
	ParticipateAssetIssue(from, issuerAddress, tokenID string, amount int64) (*api.TransactionExtention, error)
	UpdateAsset(from, description, urlStr string, newLimit, newPublicLimit int64) (*api.TransactionExtention, error)

	// Block Management
	GetNowBlock() (*api.BlockExtention, error)
	GetBlockByNum(num int64) (*api.BlockExtention, error)
	GetBlockByID(id string) (*core.Block, error)
	GetNextMaintenanceTime() (*api.NumberMessage, error)

	// Market Management
	GetMarketOrderByAccount(addr string) (*core.MarketOrderList, error)
	GetMarketPairList() (*core.MarketOrderPairList, error)
	GetMarketOrderListByPair(sellTokenId, buyTokenId string) (*core.MarketOrderList, error)
	GetMarketPriceByPair(sellTokenId, buyTokenId string) (*core.MarketPriceList, error)
	GetMarketOrderById(id string) (*core.MarketOrder, error)
	GetBurnTrx() (*api.NumberMessage, error)

	// Contracts
	DeployContract(from, contractName string, abi *core.SmartContract_ABI, codeStr string, feeLimit, curPercent, oeLimit int64) (*api.TransactionExtention, error)
	TriggerContract(from, contractAddress, method, jsonString string, feeLimit, tAmount int64, tTokenID string, tTokenAmount int64) (*api.TransactionExtention, error)

	// Shielded & Privacy
	GetSpendingKey() (*api.BytesMessage, error)
	GetExpandedSpendingKey(key string) (*api.ExpandedSpendingKeyMessage, error)
	GetAkFromAsk(ak string) (*api.BytesMessage, error)
	GetNkFromNsk(nk string) (*api.BytesMessage, error)
	GetIncomingViewingKey(ak, nk string) (*api.IncomingViewingKeyMessage, error)
	GetDiversifier() (*api.DiversifierMessage, error)
	GetRcm() (*api.BytesMessage, error)
	GetNewShieldedAddress() (*api.ShieldedAddressInfo, error)

	// Network
	ListNodes() (*api.NodeList, error)
	GetPendingSize() (*api.NumberMessage, error)
	GetBandwidthPrices() (*api.PricesResponseMessage, error)
	GetEnergyPrices() (*api.PricesResponseMessage, error)
	GetMemoFee() (*api.PricesResponseMessage, error)

	ConnectionManager
}

// ConnectionManager defines methods for managing the gRPC connection.
type ConnectionManager interface {
	Start() error
	Stop()
	Reconnect(url string) error
	getContext() (context.Context, context.CancelFunc)
}
