package pkg

//
// Copyright (C) 2024 dszi
//
// This file may be modified and distributed under the terms
// of the MIT license. See the LICENSE file for details.
// Repository: https://github.com/dszi/go-tron
//

import (
	"fmt"

	hex "github.com/dszi/go-tron/common/hexutil"
	"github.com/dszi/go-tron/pb/api"
	"github.com/dszi/go-tron/pb/core"
)

// GetBurnTrx retrieves the total burned TRX.
func (g *GrpcClient) GetBurnTrx() (*api.NumberMessage, error) {
	ctx, cancel := g.getContext()
	defer cancel()

	result, err := g.Client.GetBurnTrx(ctx, new(api.EmptyMessage))
	if err != nil {
		return nil, fmt.Errorf("GetBurnTrx: %w", err)
	}
	return result, nil
}

// GetTransactionFromPending retrieves a pending transaction by its ID.
func (g *GrpcClient) GetTransactionFromPending(id string) (*core.Transaction, error) {
	req := new(api.BytesMessage)
	var err error

	req.Value, err = hex.FromHex(id)
	if err != nil {
		return nil, fmt.Errorf("GetTransactionFromPending: failed to decode id: %w", err)
	}

	ctx, cancel := g.getContext()
	defer cancel()

	tx, err := g.Client.GetTransactionFromPending(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("GetTransactionFromPending: %w", err)
	}
	return tx, nil
}

// GetTransactionListFromPending retrieves the list of pending transaction IDs.
func (g *GrpcClient) GetTransactionListFromPending() (*api.TransactionIdList, error) {
	ctx, cancel := g.getContext()
	defer cancel()

	result, err := g.Client.GetTransactionListFromPending(ctx, new(api.EmptyMessage))
	if err != nil {
		return nil, fmt.Errorf("GetTransactionListFromPending: %w", err)
	}
	return result, nil
}

// GetPendingSize queries the size of the pending transaction pool.
func (g *GrpcClient) GetPendingSize() (*api.NumberMessage, error) {
	ctx, cancel := g.getContext()
	defer cancel()

	result, err := g.Client.GetPendingSize(ctx, new(api.EmptyMessage))
	if err != nil {
		return nil, fmt.Errorf("GetPendingSize: %w", err)
	}
	return result, nil
}

// GetBandwidthPrices retrieves the current bandwidth prices.
func (g *GrpcClient) GetBandwidthPrices() (*api.PricesResponseMessage, error) {
	ctx, cancel := g.getContext()
	defer cancel()

	result, err := g.Client.GetBandwidthPrices(ctx, new(api.EmptyMessage))
	if err != nil {
		return nil, fmt.Errorf("GetBandwidthPrices: %w", err)
	}
	return result, nil
}

// GetEnergyPrices retrieves the current energy prices.
func (g *GrpcClient) GetEnergyPrices() (*api.PricesResponseMessage, error) {
	ctx, cancel := g.getContext()
	defer cancel()

	result, err := g.Client.GetEnergyPrices(ctx, new(api.EmptyMessage))
	if err != nil {
		return nil, fmt.Errorf("GetEnergyPrices: %w", err)
	}
	return result, nil
}

// GetMemoFee retrieves the memo fee.
func (g *GrpcClient) GetMemoFee() (*api.PricesResponseMessage, error) {
	ctx, cancel := g.getContext()
	defer cancel()

	result, err := g.Client.GetMemoFee(ctx, new(api.EmptyMessage))
	if err != nil {
		return nil, fmt.Errorf("GetMemoFee: %w", err)
	}
	return result, nil
}
