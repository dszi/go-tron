//
// Copyright (C) 2024 dszi
//
// This file may be modified and distributed under the terms
// of the MIT license. See the LICENSE file for details.
// Repository: https://github.com/dszi/go-tron
//

package pkg

import (
	"fmt"
	"strconv"
	"time"

	"github.com/dszi/go-tron/common/base58"
	"github.com/dszi/go-tron/pb/api"
	"github.com/dszi/go-tron/pb/core"
)

// CreateAssetIssue issues a token.
func (g *GrpcClient) CreateAssetIssue(from, name, description, abbr, urlStr string,
	precision int32, totalSupply, startTime, endTime, freeAssetNetLimit, publicFreeAssetNetLimit int64,
	trxNum, icoNum, voteScore int32, frozenSupply map[string]string) (*api.TransactionExtention, error) {

	contract := &core.AssetIssueContract{}
	var err error

	if contract.OwnerAddress, err = base58.DecodeCheck(from); err != nil {
		return nil, fmt.Errorf("CreateAssetIssue: failed to decode from address: %w", err)
	}
	contract.Name = []byte(name)
	contract.Abbr = []byte(abbr)
	if precision < 0 || precision > 6 {
		return nil, fmt.Errorf("CreateAssetIssue: precision must be between 0 and 6")
	}
	contract.Precision = precision
	if totalSupply <= 0 {
		return nil, fmt.Errorf("CreateAssetIssue: total supply must be > 0")
	}
	contract.TotalSupply = totalSupply
	if trxNum <= 0 {
		return nil, fmt.Errorf("CreateAssetIssue: trxNum must be > 0")
	}
	contract.TrxNum = trxNum
	if icoNum <= 0 {
		return nil, fmt.Errorf("CreateAssetIssue: icoNum must be > 0")
	}
	contract.Num = icoNum

	now := time.Now().UnixNano() / 1000000
	if startTime <= now {
		return nil, fmt.Errorf("CreateAssetIssue: start time must be greater than current time")
	}
	contract.StartTime = startTime
	if endTime <= startTime {
		return nil, fmt.Errorf("CreateAssetIssue: end time must be greater than start time")
	}
	contract.EndTime = endTime

	if freeAssetNetLimit < 0 {
		return nil, fmt.Errorf("CreateAssetIssue: free asset net limit must be >= 0")
	}
	contract.FreeAssetNetLimit = freeAssetNetLimit
	if publicFreeAssetNetLimit < 0 {
		return nil, fmt.Errorf("CreateAssetIssue: public free asset net limit must be >= 0")
	}
	contract.PublicFreeAssetNetLimit = publicFreeAssetNetLimit

	contract.VoteScore = voteScore
	contract.Description = []byte(description)
	contract.Url = []byte(urlStr)

	// Process frozen supply: key is frozen days, value is frozen amount.
	for key, value := range frozenSupply {
		amount, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return nil, fmt.Errorf("CreateAssetIssue: failed to parse frozen amount for key %s: %w", key, err)
		}
		days, err := strconv.ParseInt(key, 10, 64)
		if err != nil {
			return nil, fmt.Errorf("CreateAssetIssue: failed to parse frozen days for key %s: %w", key, err)
		}
		frozen := &core.AssetIssueContract_FrozenSupply{
			FrozenAmount: amount,
			FrozenDays:   days,
		}
		contract.FrozenSupply = append(contract.FrozenSupply, frozen)
	}

	ctx, cancel := g.getContext()
	defer cancel()

	tx, err := g.Client.CreateAssetIssue2(ctx, contract)
	if err != nil {
		return nil, fmt.Errorf("CreateAssetIssue RPC error: %w", err)
	}
	if err := validateTx(tx); err != nil {
		return nil, err
	}
	return tx, nil
}

// GetAssetIssueList queries the list of all issued tokens.
// If page is -1, returns the full list.
func (g *GrpcClient) GetAssetIssueList(page int64, limit ...int64) (*api.AssetIssueList, error) {
	ctx, cancel := g.getContext()
	defer cancel()

	if page == -1 {
		return g.Client.GetAssetIssueList(ctx, new(api.EmptyMessage))
	}

	useLimit := int64(10)
	if len(limit) == 1 {
		useLimit = limit[0]
	}
	return g.Client.GetPaginatedAssetIssueList(ctx, GetPaginatedMessage(page*useLimit, useLimit))
}

// TransferAsset transfers tokens.
func (g *GrpcClient) TransferAsset(from, toAddress, assetName string, amount int64) (*api.TransactionExtention, error) {
	contract := &core.TransferAssetContract{}
	var err error

	if contract.OwnerAddress, err = base58.DecodeCheck(from); err != nil {
		return nil, fmt.Errorf("TransferAsset: failed to decode from address: %w", err)
	}
	if contract.ToAddress, err = base58.DecodeCheck(toAddress); err != nil {
		return nil, fmt.Errorf("TransferAsset: failed to decode to address: %w", err)
	}

	contract.AssetName = []byte(assetName)
	contract.Amount = amount

	ctx, cancel := g.getContext()
	defer cancel()

	tx, err := g.Client.TransferAsset2(ctx, contract)
	if err != nil {
		return nil, fmt.Errorf("TransferAsset RPC error: %w", err)
	}
	if err := validateTx(tx); err != nil {
		return nil, err
	}
	return tx, nil
}

// ParticipateAssetIssue participates in a token issuance.
func (g *GrpcClient) ParticipateAssetIssue(from, issuerAddress, tokenID string, amount int64) (*api.TransactionExtention, error) {
	contract := &core.ParticipateAssetIssueContract{}
	var err error

	if contract.OwnerAddress, err = base58.DecodeCheck(from); err != nil {
		return nil, fmt.Errorf("ParticipateAssetIssue: failed to decode from address: %w", err)
	}
	if contract.ToAddress, err = base58.DecodeCheck(issuerAddress); err != nil {
		return nil, fmt.Errorf("ParticipateAssetIssue: failed to decode issuer address: %w", err)
	}

	contract.AssetName = []byte(tokenID)
	contract.Amount = amount

	ctx, cancel := g.getContext()
	defer cancel()

	tx, err := g.Client.ParticipateAssetIssue2(ctx, contract)
	if err != nil {
		return nil, fmt.Errorf("ParticipateAssetIssue RPC error: %w", err)
	}
	if err := validateTx(tx); err != nil {
		return nil, err
	}
	return tx, nil
}

// GetPaginatedAssetIssueList queries the list of all issued tokens.
// If page is -1, returns the full list.
func (g *GrpcClient) GetPaginatedAssetIssueList(page int64, limit ...int64) (*api.AssetIssueList, error) {
	ctx, cancel := g.getContext()
	defer cancel()

	if page == -1 {
		return g.Client.GetAssetIssueList(ctx, new(api.EmptyMessage))
	}

	useLimit := int64(10)
	if len(limit) == 1 {
		useLimit = limit[0]
	}
	return g.Client.GetPaginatedAssetIssueList(ctx, GetPaginatedMessage(page*useLimit, useLimit))
}

// GetAssetIssueByAccount queries tokens issued by a given account.
func (g *GrpcClient) GetAssetIssueByAccount(address string) (*api.AssetIssueList, error) {
	req := new(core.Account)
	var err error

	if req.Address, err = base58.DecodeCheck(address); err != nil {
		return nil, fmt.Errorf("GetAssetIssueByAccount: failed to decode address: %w", err)
	}

	ctx, cancel := g.getContext()
	defer cancel()

	return g.Client.GetAssetIssueByAccount(ctx, req)
}

// GetAssetIssueByName queries token information by token name.
func (g *GrpcClient) GetAssetIssueByName(name string) (*core.AssetIssueContract, error) {
	ctx, cancel := g.getContext()
	defer cancel()

	return g.Client.GetAssetIssueByName(ctx, GetMessageBytes([]byte(name)))
}

// GetAssetIssueById queries token information by id.
func (g *GrpcClient) GetAssetIssueById(id string) (*core.AssetIssueContract, error) {
	ctx, cancel := g.getContext()
	defer cancel()

	return g.Client.GetAssetIssueById(ctx, GetMessageBytes([]byte(id)))
}

// UpdateAsset updates asset details such as description and limit.
func (g *GrpcClient) UpdateAsset(from, description, urlStr string, newLimit, newPublicLimit int64) (*api.TransactionExtention, error) {
	addr, err := base58.DecodeCheck(from)
	if err != nil {
		return nil, fmt.Errorf("invalid address: %w", err)
	}

	contract := &core.UpdateAssetContract{
		OwnerAddress:   addr,
		Description:    []byte(description),
		Url:            []byte(urlStr),
		NewLimit:       newLimit,
		NewPublicLimit: newPublicLimit,
	}

	ctx, cancel := g.getContext()
	defer cancel()

	tx, err := g.Client.UpdateAsset2(ctx, contract)
	if err != nil {
		return nil, fmt.Errorf("UpdateAsset failed: %w", err)
	}

	return tx, validateTx(tx)
}
