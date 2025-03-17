//
// Copyright (C) 2024 dszi
//
// This file may be modified and distributed under the terms
// of the MIT license. See the LICENSE file for details.
// Repository: https://github.com/dszi/go-tron
//

package pkg

import (
	"github.com/dszi/go-tron/pb/api"
)

// GetSpendingKey retrieves a spending key.
func (g *GrpcClient) GetSpendingKey() (*api.BytesMessage, error) {
	ctx, cancel := g.getContext()
	defer cancel()

	return g.Client.GetSpendingKey(ctx, new(api.EmptyMessage))
}

// GetExpandedSpendingKey retrieves the expanded spending key.
func (g *GrpcClient) GetExpandedSpendingKey(key string) (*api.ExpandedSpendingKeyMessage, error) {
	ctx, cancel := g.getContext()
	defer cancel()

	return g.Client.GetExpandedSpendingKey(ctx, &api.BytesMessage{Value: []byte(key)})
}

// GetAkFromAsk retrieves `Ak` from `Ask`.
func (g *GrpcClient) GetAkFromAsk(ak string) (*api.BytesMessage, error) {
	ctx, cancel := g.getContext()
	defer cancel()

	return g.Client.GetAkFromAsk(ctx, &api.BytesMessage{Value: []byte(ak)})
}

// GetNkFromNsk retrieves `Nk` from `Nsk`.
func (g *GrpcClient) GetNkFromNsk(nk string) (*api.BytesMessage, error) {
	ctx, cancel := g.getContext()
	defer cancel()

	return g.Client.GetNkFromNsk(ctx, &api.BytesMessage{Value: []byte(nk)})
}

// GetIncomingViewingKey retrieves an incoming viewing key.
func (g *GrpcClient) GetIncomingViewingKey(ak, nk string) (*api.IncomingViewingKeyMessage, error) {
	ctx, cancel := g.getContext()
	defer cancel()

	return g.Client.GetIncomingViewingKey(ctx, &api.ViewingKeyMessage{Ak: []byte(ak), Nk: []byte(nk)})
}

// GetDiversifier retrieves a diversifier message.
func (g *GrpcClient) GetDiversifier() (*api.DiversifierMessage, error) {
	ctx, cancel := g.getContext()
	defer cancel()

	return g.Client.GetDiversifier(ctx, new(api.EmptyMessage))
}

// GetRcm retrieves a random commitment.
func (g *GrpcClient) GetRcm() (*api.BytesMessage, error) {
	ctx, cancel := g.getContext()
	defer cancel()

	return g.Client.GetRcm(ctx, new(api.EmptyMessage))
}

// GetNewShieldedAddress generates a new shielded address.
func (g *GrpcClient) GetNewShieldedAddress() (*api.ShieldedAddressInfo, error) {
	ctx, cancel := g.getContext()
	defer cancel()

	return g.Client.GetNewShieldedAddress(ctx, new(api.EmptyMessage))
}
