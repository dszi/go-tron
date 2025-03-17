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

	hex "github.com/dszi/go-tron/common/hexutil"
	"github.com/dszi/go-tron/pb/api"
	"github.com/dszi/go-tron/pb/core"
	"google.golang.org/grpc"
)

// GetNowBlock retrieves the current block information.
func (g *GrpcClient) GetNowBlock() (*api.BlockExtention, error) {
	ctx, cancel := g.getContext()
	defer cancel()

	result, err := g.Client.GetNowBlock2(ctx, new(api.EmptyMessage))
	if err != nil {
		return nil, fmt.Errorf("GetNowBlock: %w", err)
	}
	return result, nil
}

// GetBlockByNum retrieves a block by its height.
func (g *GrpcClient) GetBlockByNum(num int64) (*api.BlockExtention, error) {
	numMsg := &api.NumberMessage{Num: num}
	ctx, cancel := g.getContext()
	defer cancel()

	maxSizeOption := grpc.MaxCallRecvMsgSize(32 * 10e6)
	result, err := g.Client.GetBlockByNum2(ctx, numMsg, maxSizeOption)
	if err != nil {
		return nil, fmt.Errorf("GetBlockByNum: %w", err)
	}
	return result, nil
}

// GetBlockByID queries block information by block ID.
func (g *GrpcClient) GetBlockByID(id string) (*core.Block, error) {
	blockID := &api.BytesMessage{}
	var err error

	if blockID.Value, err = hex.FromHex(id); err != nil {
		return nil, fmt.Errorf("GetBlockByID: %w", err)
	}

	ctx, cancel := g.getContext()
	defer cancel()

	maxSizeOption := grpc.MaxCallRecvMsgSize(32 * 10e6)
	result, err := g.Client.GetBlockById(ctx, blockID, maxSizeOption)
	if err != nil {
		return nil, fmt.Errorf("GetBlockByID: %w", err)
	}
	return result, nil
}

// GetNextMaintenanceTime queries the next maintenance time.
func (g *GrpcClient) GetNextMaintenanceTime() (*api.NumberMessage, error) {
	ctx, cancel := g.getContext()
	defer cancel()

	nm, err := g.Client.GetNextMaintenanceTime(ctx, new(api.EmptyMessage))
	if err != nil {
		return nil, fmt.Errorf("GetNextMaintenanceTime error: %w", err)
	}
	return nm, nil
}
