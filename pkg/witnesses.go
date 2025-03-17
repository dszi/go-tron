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

	"github.com/dszi/go-tron/common/base58"
	"github.com/dszi/go-tron/pb/api"
	"github.com/dszi/go-tron/pb/core"
)

// VoteWitnessAccount submits a vote for super representative candidates.
func (g *GrpcClient) VoteWitnessAccount(from string, witnessMap map[string]int64) (*api.TransactionExtention, error) {
	contract := &core.VoteWitnessContract{}
	var err error

	contract.OwnerAddress, err = base58.DecodeCheck(from)
	if err != nil {
		return nil, fmt.Errorf("VoteWitnessAccount: failed to decode from address: %w", err)
	}

	// Construct votes from witnessMap.
	for addr, count := range witnessMap {
		witnessAddress, err := base58.DecodeCheck(addr)
		if err != nil {
			return nil, fmt.Errorf("VoteWitnessAccount: failed to decode witness address (%s): %w", addr, err)
		}
		contract.Votes = append(contract.Votes, &core.VoteWitnessContract_Vote{
			VoteAddress: witnessAddress,
			VoteCount:   count,
		})
	}

	ctx, cancel := g.getContext()
	defer cancel()

	tx, err := g.Client.VoteWitnessAccount2(ctx, contract)
	if err != nil {
		return nil, fmt.Errorf("VoteWitnessAccount RPC error: %w", err)
	}
	if err := validateTx(tx); err != nil {
		return nil, err
	}
	return tx, nil
}

// ListWitnesses queries the list of super representative candidates.
func (g *GrpcClient) ListWitnesses() (*api.WitnessList, error) {
	ctx, cancel := g.getContext()
	defer cancel()

	witnessList, err := g.Client.ListWitnesses(ctx, new(api.EmptyMessage))
	if err != nil {
		return nil, fmt.Errorf("ListWitnesses RPC error: %w", err)
	}
	return witnessList, nil
}

// CreateWitness applies to become a super representative candidate.
func (g *GrpcClient) CreateWitness(from, urlStr string) (*api.TransactionExtention, error) {
	contract := &core.WitnessCreateContract{
		Url: []byte(urlStr),
	}
	var err error
	contract.OwnerAddress, err = base58.DecodeCheck(from)
	if err != nil {
		return nil, fmt.Errorf("CreateWitness: failed to decode from address: %w", err)
	}

	ctx, cancel := g.getContext()
	defer cancel()

	tx, err := g.Client.CreateWitness2(ctx, contract)
	if err != nil {
		return nil, fmt.Errorf("CreateWitness RPC error: %w", err)
	}
	if err := validateTx(tx); err != nil {
		return nil, err
	}
	return tx, nil
}

// UpdateWitness updates the website URL of a super representative candidate.
func (g *GrpcClient) UpdateWitness(from, urlStr string) (*api.TransactionExtention, error) {
	contract := &core.WitnessUpdateContract{
		UpdateUrl: []byte(urlStr),
	}
	var err error
	contract.OwnerAddress, err = base58.DecodeCheck(from)
	if err != nil {
		return nil, fmt.Errorf("UpdateWitness: failed to decode from address: %w", err)
	}

	ctx, cancel := g.getContext()
	defer cancel()

	tx, err := g.Client.UpdateWitness2(ctx, contract)
	if err != nil {
		return nil, fmt.Errorf("UpdateWitness RPC error: %w", err)
	}
	if err := validateTx(tx); err != nil {
		return nil, err
	}
	return tx, nil
}

// GetBrokerageInfo queries the unclaimed brokerage reward.
func (g *GrpcClient) GetBrokerageInfo(witness string) (float64, error) {
	addr, err := base58.DecodeCheck(witness)
	if err != nil {
		return 0, fmt.Errorf("GetBrokerageInfo: failed to decode witness address: %w", err)
	}

	ctx, cancel := g.getContext()
	defer cancel()

	result, err := g.Client.GetBrokerageInfo(ctx, GetMessageBytes(addr))
	if err != nil {
		return 0, fmt.Errorf("GetBrokerageInfo RPC error: %w", err)
	}
	return float64(result.Num), nil
}

// UpdateBrokerage updates the brokerage ratio.
func (g *GrpcClient) UpdateBrokerage(from string, brokerage int32) (*api.TransactionExtention, error) {
	contract := &core.UpdateBrokerageContract{
		Brokerage: brokerage,
	}
	var err error
	contract.OwnerAddress, err = base58.DecodeCheck(from)
	if err != nil {
		return nil, fmt.Errorf("UpdateBrokerage: failed to decode from address: %w", err)
	}

	ctx, cancel := g.getContext()
	defer cancel()

	tx, err := g.Client.UpdateBrokerage(ctx, contract)
	if err != nil {
		return nil, fmt.Errorf("UpdateBrokerage RPC error: %w", err)
	}
	if err := validateTx(tx); err != nil {
		return nil, err
	}
	return tx, nil
}
