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

// TotalTransaction retrieves the total number of transactions.
func (g *GrpcClient) TotalTransaction() (*api.NumberMessage, error) {
	ctx, cancel := g.getContext()
	defer cancel()

	totalTx, err := g.Client.TotalTransaction(ctx, new(api.EmptyMessage))
	if err != nil {
		return nil, fmt.Errorf("TotalTransaction error: %w", err)
	}
	return totalTx, nil
}
