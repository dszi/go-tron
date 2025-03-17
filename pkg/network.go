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

	"github.com/dszi/go-tron/pb/api"
)

// ListNodes queries the list of nodes connected to the API.
func (g *GrpcClient) ListNodes() (*api.NodeList, error) {
	ctx, cancel := g.getContext()
	defer cancel()

	nodeList, err := g.Client.ListNodes(ctx, new(api.EmptyMessage))
	if err != nil {
		return nil, fmt.Errorf("ListNodes error: %w", err)
	}
	return nodeList, nil
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
