//
// Copyright (C) 2024 dszi
//
// This file may be modified and distributed under the terms
// of the MIT license. See the LICENSE file for details.
// Repository: https://github.com/dszi/go-tron
//

package pkg

import (
	"bytes"
	"fmt"

	"github.com/dszi/go-tron/common/base58"
	"github.com/dszi/go-tron/pb/api"
	"github.com/dszi/go-tron/pb/core"
)

// GetAccount retrieves account information by address.
func (g *GrpcClient) GetAccount(addr string) (*core.Account, error) {
	req := new(core.Account)
	var err error

	req.Address, err = base58.DecodeCheck(addr)
	if err != nil {
		return nil, fmt.Errorf("failed to decode account address: %w", err)
	}

	ctx, cancel := g.getContext()
	defer cancel()

	acc, err := g.Client.GetAccount(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("GetAccount RPC error: %w", err)
	}
	if !bytes.Equal(acc.Address, req.Address) {
		return nil, fmt.Errorf("account not found")
	}
	return acc, nil
}

// GetAccountBalance retrieves the account balance.
func (g *GrpcClient) GetAccountBalance(addr string) (int64, error) {
	req := new(core.Account)
	var err error

	req.Address, err = base58.DecodeCheck(addr)
	if err != nil {
		return 0, fmt.Errorf("failed to decode account address: %w", err)
	}

	ctx, cancel := g.getContext()
	defer cancel()

	acc, err := g.Client.GetAccount(ctx, req)
	if err != nil {
		return 0, fmt.Errorf("GetAccountBalance RPC error: %w", err)
	}
	if !bytes.Equal(acc.Address, req.Address) {
		return 0, fmt.Errorf("account not found")
	}
	return acc.Balance, nil
}

// GetAccountResource retrieves the account resource information.
func (g *GrpcClient) GetAccountResource(addr string) (*api.AccountResourceMessage, error) {
	req := new(core.Account)
	var err error

	req.Address, err = base58.DecodeCheck(addr)
	if err != nil {
		return nil, fmt.Errorf("failed to decode account address: %w", err)
	}

	ctx, cancel := g.getContext()
	defer cancel()

	resource, err := g.Client.GetAccountResource(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("GetAccountResource RPC error: %w", err)
	}
	return resource, nil
}

// CreateAccount creates a new account.
func (g *GrpcClient) CreateAccount(from, addr string) (*api.TransactionExtention, error) {
	contract := new(core.AccountCreateContract)
	var err error

	contract.OwnerAddress, err = base58.DecodeCheck(from)
	if err != nil {
		return nil, fmt.Errorf("failed to decode from address: %w", err)
	}
	contract.AccountAddress, err = base58.DecodeCheck(addr)
	if err != nil {
		return nil, fmt.Errorf("failed to decode target address: %w", err)
	}

	ctx, cancel := g.getContext()
	defer cancel()

	tx, err := g.Client.CreateAccount2(ctx, contract)
	if err != nil {
		return nil, fmt.Errorf("CreateAccount RPC error: %w", err)
	}
	if err := checkTxResult(tx); err != nil {
		return nil, err
	}
	return tx, nil
}

// UpdateAccount updates the account name.
func (g *GrpcClient) UpdateAccount(from, accountName string) (*api.TransactionExtention, error) {
	contract := &core.AccountUpdateContract{
		AccountName: []byte(accountName),
	}
	var err error

	contract.OwnerAddress, err = base58.DecodeCheck(from)
	if err != nil {
		return nil, fmt.Errorf("failed to decode from address: %w", err)
	}

	ctx, cancel := g.getContext()
	defer cancel()

	tx, err := g.Client.UpdateAccount2(ctx, contract)
	if err != nil {
		return nil, fmt.Errorf("UpdateAccount RPC error: %w", err)
	}
	if err := checkTxResult(tx); err != nil {
		return nil, err
	}
	return tx, nil
}

// GetRewardInfo queries the unclaimed reward.
func (g *GrpcClient) GetRewardInfo(addr string) (int64, error) {
	addrBytes, err := base58.DecodeCheck(addr)
	if err != nil {
		return 0, fmt.Errorf("failed to decode address: %w", err)
	}

	ctx, cancel := g.getContext()
	defer cancel()

	rewards, err := g.Client.GetRewardInfo(ctx, GetMessageBytes(addrBytes))
	if err != nil {
		return 0, fmt.Errorf("failed to get reward info: %w", err)
	}
	return rewards.Num, nil
}
