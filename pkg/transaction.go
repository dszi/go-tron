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
	hex "github.com/dszi/go-tron/common/hexutil"
	"github.com/dszi/go-tron/pb/api"
	"github.com/dszi/go-tron/pb/core"
	"google.golang.org/protobuf/proto"
)

// CreateTransaction creates a TRX transfer transaction.
func (g *GrpcClient) CreateTransaction(from, toAddress string, amount int64) (*api.TransactionExtention, error) {
	contract := &core.TransferContract{}
	var err error

	if contract.OwnerAddress, err = base58.DecodeCheck(from); err != nil {
		return nil, fmt.Errorf("CreateTransaction: failed to decode from address: %w", err)
	}
	if contract.ToAddress, err = base58.DecodeCheck(toAddress); err != nil {
		return nil, fmt.Errorf("CreateTransaction: failed to decode to address: %w", err)
	}
	contract.Amount = amount

	ctx, cancel := g.getContext()
	defer cancel()

	tx, err := g.Client.CreateTransaction2(ctx, contract)
	if err != nil {
		return nil, fmt.Errorf("CreateTransaction RPC error: %w", err)
	}
	if proto.Size(tx) == 0 {
		return nil, fmt.Errorf("CreateTransaction: bad transaction")
	}
	if tx.GetResult().GetCode() != 0 {
		return nil, fmt.Errorf("CreateTransaction: %s", tx.GetResult().GetMessage())
	}
	return tx, nil
}

// BroadcastTransaction broadcasts a signed transaction to the network.
// It returns an error if the broadcast result indicates failure.
func (g *GrpcClient) BroadcastTransaction(tx *core.Transaction) (*api.Return, error) {
	ctx, cancel := g.getContext()
	defer cancel()

	result, err := g.Client.BroadcastTransaction(ctx, tx)
	if err != nil {
		return nil, fmt.Errorf("BroadcastTransaction RPC error: %w", err)
	}
	if !result.GetResult() {
		return result, fmt.Errorf("BroadcastTransaction: result error: %s", result.GetMessage())
	}
	if result.GetCode() != api.Return_SUCCESS {
		return result, fmt.Errorf("BroadcastTransaction: result error (%s): %s", result.GetCode(), result.GetMessage())
	}
	return result, nil
}

// GetTransactionByID retrieves a transaction by its ID.
func (g *GrpcClient) GetTransactionByID(id string) (*core.Transaction, error) {
	req := new(api.BytesMessage)
	var err error

	req.Value, err = hex.FromHex(id)
	if err != nil {
		return nil, fmt.Errorf("GetTransactionByID: failed to decode id: %w", err)
	}

	ctx, cancel := g.getContext()
	defer cancel()

	tx, err := g.Client.GetTransactionById(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("GetTransactionByID RPC error: %w", err)
	}
	if proto.Size(tx) == 0 {
		return nil, fmt.Errorf("GetTransactionByID: transaction info not found")
	}
	return tx, nil
}

// GetTransactionInfoByID queries transaction fee and block information by transaction ID.
func (g *GrpcClient) GetTransactionInfoByID(id string) (*core.TransactionInfo, error) {
	req := new(api.BytesMessage)
	var err error

	req.Value, err = hex.FromHex(id)
	if err != nil {
		return nil, fmt.Errorf("GetTransactionInfoByID: failed to decode id: %w", err)
	}

	ctx, cancel := g.getContext()
	defer cancel()

	txi, err := g.Client.GetTransactionInfoById(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("GetTransactionInfoByID RPC error: %w", err)
	}
	if !bytes.Equal(txi.Id, req.Value) {
		return nil, fmt.Errorf("GetTransactionInfoByID: transaction info not found")
	}
	return txi, nil
}
