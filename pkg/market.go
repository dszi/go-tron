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
	hex "github.com/dszi/go-tron/common/hexutil"
	"github.com/dszi/go-tron/pb/api"
	"github.com/dszi/go-tron/pb/core"
)

// GetMarketOrderByAccount queries market orders for the given account.
func (g *GrpcClient) GetMarketOrderByAccount(addr string) (*core.MarketOrderList, error) {
	req := new(api.BytesMessage)
	var err error

	req.Value, err = base58.DecodeCheck(addr)
	if err != nil {
		return nil, fmt.Errorf("GetMarketOrderByAccount: failed to decode address: %w", err)
	}

	ctx, cancel := g.getContext()
	defer cancel()

	orderList, err := g.Client.GetMarketOrderByAccount(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("GetMarketOrderByAccount: RPC error: %w", err)
	}
	return orderList, nil
}

// GetMarketOrderListByPair queries market order list by sell and buy token IDs.
func (g *GrpcClient) GetMarketOrderListByPair(sellTokenId, buyTokenId string) (*core.MarketOrderList, error) {
	req := new(core.MarketOrderPair)
	req.SellTokenId = []byte(sellTokenId)
	req.BuyTokenId = []byte(buyTokenId)

	ctx, cancel := g.getContext()
	defer cancel()

	orderList, err := g.Client.GetMarketOrderListByPair(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("GetMarketOrderListByPair: %w", err)
	}
	return orderList, nil
}

// GetMarketPriceByPair queries market price information by sell and buy token IDs.
func (g *GrpcClient) GetMarketPriceByPair(sellTokenId, buyTokenId string) (*core.MarketPriceList, error) {
	req := new(core.MarketOrderPair)
	req.SellTokenId = []byte(sellTokenId)
	req.BuyTokenId = []byte(buyTokenId)

	ctx, cancel := g.getContext()
	defer cancel()

	priceList, err := g.Client.GetMarketPriceByPair(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("GetMarketPriceByPair: %w", err)
	}
	return priceList, nil
}

// GetMarketPairList queries the list of all market pairs.
func (g *GrpcClient) GetMarketPairList() (*core.MarketOrderPairList, error) {
	ctx, cancel := g.getContext()
	defer cancel()

	result, err := g.Client.GetMarketPairList(ctx, new(api.EmptyMessage))
	if err != nil {
		return nil, fmt.Errorf("GetMarketPairList: %w", err)
	}
	return result, nil
}

// GetMarketOrderById queries a market order by its ID.
func (g *GrpcClient) GetMarketOrderById(id string) (*core.MarketOrder, error) {
	orderID := new(api.BytesMessage)
	var err error

	orderID.Value, err = hex.FromHex(id)
	if err != nil {
		return nil, fmt.Errorf("GetMarketOrderById: failed to decode id: %w", err)
	}

	ctx, cancel := g.getContext()
	defer cancel()

	order, err := g.Client.GetMarketOrderById(ctx, orderID)
	if err != nil {
		return nil, fmt.Errorf("GetMarketOrderById: RPC error: %w", err)
	}
	return order, nil
}
