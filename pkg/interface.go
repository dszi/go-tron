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

type TronClient interface {
	GetAccount(addr string) (*core.Account, error)
	GetAccountBalance(addr string) (int64, error)
	GetAccountResource(addr string) (*api.AccountResourceMessage, error)
	CreateTransaction(from, toAddress string, amount int64) (*api.TransactionExtention, error)
	BroadcastTransaction(tx *core.Transaction) (*api.Return, error)
	CreateAccount(from, addr string) (*api.TransactionExtention, error)
	UpdateAccount(from, accountName string) (*api.TransactionExtention, error)
	VoteWitnessAccount(from string, witnessMap map[string]int64) (*api.TransactionExtention, error)
	GetBrokerageInfo(witness string) (float64, error)
	GetRewardInfo(addr string) (int64, error)
	UpdateBrokerage(from string, brokerage int32) (*api.TransactionExtention, error)
	CreateAssetIssue(from, name, description, abbr, urlStr string, precision int32, totalSupply, startTime, endTime, FreeAssetNetLimit, PublicFreeAssetNetLimit int64, trxNum, icoNum, voteScore int32, frozenSupply map[string]string) (*api.TransactionExtention, error)
	ListWitnesses() (*api.WitnessList, error)
	CreateWitness(from, urlStr string) (*api.TransactionExtention, error)
	UpdateWitness(from, urlStr string) (*api.TransactionExtention, error)
	TransferAsset(from, toAddress, assetName string, amount int64) (*api.TransactionExtention, error)
	ParticipateAssetIssue(from, issuerAddress, tokenID string, amount int64) (*api.TransactionExtention, error)
	ListNodes() (*api.NodeList, error)
	GetAssetIssueList(page int64, limit ...int) (*api.AssetIssueList, error)
	GetAssetIssueByAccount(address string) (*api.AssetIssueList, error)
	GetAssetIssueByName(name string) (*core.AssetIssueContract, error)
	GetAssetIssueById(id string) (*core.AssetIssueContract, error)
	GetNowBlock() (*api.BlockExtention, error)
	GetBlockByNum(num int64) (*api.BlockExtention, error)
	TotalTransaction() (*api.NumberMessage, error)
	GetTransactionByID(id string) (*core.Transaction, error)
	FreezeBalance(from, delegateTo string, resource core.ResourceCode, frozenBalance int64) (*api.TransactionExtention, error)
	UnfreezeBalance(from, delegateTo string, resource core.ResourceCode) (*api.TransactionExtention, error)
	WithdrawBalance(from string) (*api.TransactionExtention, error)
	UnfreezeAsset(from string) (*api.TransactionExtention, error)
	GetNextMaintenanceTime() (*api.NumberMessage, error)
	GetTransactionInfoByID(id string) (*core.TransactionInfo, error)
	GetBlockByID(id string) (*core.Block, error)
	GetMarketOrderByAccount(addr string) (*core.MarketOrderList, error)
	GetMarketPairList() (*core.MarketOrderPairList, error)
	GetMarketOrderListByPair(sellTokenId, buyTokenId string) (*core.MarketOrderList, error)
	GetMarketPriceByPair(sellTokenId, buyTokenId string) (*core.MarketPriceList, error)
	GetMarketOrderById(id string) (*core.MarketOrder, error)
	GetBurnTrx() (*api.NumberMessage, error)
	UnfreezeBalanceV2(from string, resource core.ResourceCode, unfreezeBalance int64) (*api.TransactionExtention, error)
	WithdrawExpireUnfreeze(from string, timestamp int64) (*api.TransactionExtention, error)
	DelegateResource(from, to string, resource core.ResourceCode, delegateBalance int64, lock bool, lockPeriod int64) (*api.TransactionExtention, error)
	UnDelegateResource(owner, receiver string, resource core.ResourceCode, delegateBalance int64, lock bool) (*api.TransactionExtention, error)
	GetTransactionFromPending(id string) (*core.Transaction, error)
	GetTransactionListFromPending() (*api.TransactionIdList, error)
	GetPendingSize() (*api.NumberMessage, error)
	CancelAllUnfreezeV2(from string) (*api.TransactionExtention, error)
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
