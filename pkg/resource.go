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
	"google.golang.org/protobuf/proto"
)

// FreezeBalance stakes TRX (deprecated, use FreezeBalanceV2 instead).
func (g *GrpcClient) FreezeBalance(from, delegateTo string, resource core.ResourceCode, frozenBalance int64) (*api.TransactionExtention, error) {
	contract := &core.FreezeBalanceContract{}
	var err error

	if contract.OwnerAddress, err = base58.DecodeCheck(from); err != nil {
		return nil, fmt.Errorf("FreezeBalance: failed to decode from address: %w", err)
	}
	contract.FrozenBalance = frozenBalance
	contract.FrozenDuration = 3 // Tron only allows 3 days freeze

	if delegateTo != "" {
		if contract.ReceiverAddress, err = base58.DecodeCheck(delegateTo); err != nil {
			return nil, fmt.Errorf("FreezeBalance: failed to decode delegateTo address: %w", err)
		}
	}
	contract.Resource = resource

	ctx, cancel := g.getContext()
	defer cancel()

	tx, err := g.Client.FreezeBalance2(ctx, contract)
	if err != nil {
		return nil, fmt.Errorf("FreezeBalance RPC error: %w", err)
	}
	if err := validateTx(tx); err != nil {
		return nil, err
	}
	return tx, nil
}

// UnfreezeBalance unstakes TRX staked during Stake1.0.
func (g *GrpcClient) UnfreezeBalance(from, delegateTo string, resource core.ResourceCode) (*api.TransactionExtention, error) {
	contract := &core.UnfreezeBalanceContract{}
	var err error

	if contract.OwnerAddress, err = base58.DecodeCheck(from); err != nil {
		return nil, fmt.Errorf("UnfreezeBalance: failed to decode from address: %w", err)
	}
	if delegateTo != "" {
		if contract.ReceiverAddress, err = base58.DecodeCheck(delegateTo); err != nil {
			return nil, fmt.Errorf("UnfreezeBalance: failed to decode delegateTo address: %w", err)
		}
	}
	contract.Resource = resource

	ctx, cancel := g.getContext()
	defer cancel()

	tx, err := g.Client.UnfreezeBalance2(ctx, contract)
	if err != nil {
		return nil, fmt.Errorf("UnfreezeBalance RPC error: %w", err)
	}
	if err := validateTx(tx); err != nil {
		return nil, err
	}
	return tx, nil
}

// WithdrawBalance redeems block producing reward.
func (g *GrpcClient) WithdrawBalance(from string) (*api.TransactionExtention, error) {
	contract := &core.WithdrawBalanceContract{}
	var err error

	if contract.OwnerAddress, err = base58.DecodeCheck(from); err != nil {
		return nil, fmt.Errorf("WithdrawBalance: failed to decode from address: %w", err)
	}

	ctx, cancel := g.getContext()
	defer cancel()

	tx, err := g.Client.WithdrawBalance2(ctx, contract)
	if err != nil {
		return nil, fmt.Errorf("WithdrawBalance RPC error: %w", err)
	}
	if err := validateTx(tx); err != nil {
		return nil, err
	}
	return tx, nil
}

// UnfreezeAsset unstakes token balance.
func (g *GrpcClient) UnfreezeAsset(from string) (*api.TransactionExtention, error) {
	contract := &core.UnfreezeAssetContract{}
	var err error

	if contract.OwnerAddress, err = base58.DecodeCheck(from); err != nil {
		return nil, fmt.Errorf("UnfreezeAsset: failed to decode from address: %w", err)
	}

	ctx, cancel := g.getContext()
	defer cancel()

	tx, err := g.Client.UnfreezeAsset2(ctx, contract)
	if err != nil {
		return nil, fmt.Errorf("UnfreezeAsset RPC error: %w", err)
	}
	if err := validateTx(tx); err != nil {
		return nil, err
	}
	return tx, nil
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

// UnfreezeBalanceV2 unfreezes TRX (new version).
func (g *GrpcClient) UnfreezeBalanceV2(from string, resource core.ResourceCode, unfreezeBalance int64) (*api.TransactionExtention, error) {
	contract := &core.UnfreezeBalanceV2Contract{}
	var err error

	if contract.OwnerAddress, err = base58.DecodeCheck(from); err != nil {
		return nil, fmt.Errorf("UnfreezeBalanceV2: failed to decode from address: %w", err)
	}
	contract.UnfreezeBalance = unfreezeBalance
	contract.Resource = resource

	ctx, cancel := g.getContext()
	defer cancel()

	tx, err := g.Client.UnfreezeBalanceV2(ctx, contract)
	if err != nil {
		return nil, fmt.Errorf("UnfreezeBalanceV2 RPC error: %w", err)
	}
	if err := validateTx(tx); err != nil {
		return nil, err
	}
	return tx, nil
}

// WithdrawExpireUnfreeze withdraws staked TRX after expiration.
func (g *GrpcClient) WithdrawExpireUnfreeze(from string, timestamp int64) (*api.TransactionExtention, error) {
	contract := &core.WithdrawExpireUnfreezeContract{}
	var err error

	if contract.OwnerAddress, err = base58.DecodeCheck(from); err != nil {
		return nil, fmt.Errorf("WithdrawExpireUnfreeze: failed to decode from address: %w", err)
	}

	ctx, cancel := g.getContext()
	defer cancel()

	tx, err := g.Client.WithdrawExpireUnfreeze(ctx, contract)
	if err != nil {
		return nil, fmt.Errorf("WithdrawExpireUnfreeze RPC error: %w", err)
	}
	if proto.Size(tx) == 0 {
		return nil, fmt.Errorf("WithdrawExpireUnfreeze: bad transaction")
	}
	return tx, nil
}

// DelegateResource delegates resources for staking.
func (g *GrpcClient) DelegateResource(from, to string, resource core.ResourceCode, delegateBalance int64, lock bool, lockPeriod int64) (*api.TransactionExtention, error) {
	addrFrom, err := base58.DecodeCheck(from)
	if err != nil {
		return nil, fmt.Errorf("DelegateResource: failed to decode from address: %w", err)
	}
	addrTo, err := base58.DecodeCheck(to)
	if err != nil {
		return nil, fmt.Errorf("DelegateResource: failed to decode to address: %w", err)
	}

	contract := &core.DelegateResourceContract{
		OwnerAddress:    addrFrom,
		ReceiverAddress: addrTo,
		Resource:        resource,
		Balance:         delegateBalance,
		Lock:            lock,
		LockPeriod:      lockPeriod,
	}

	ctx, cancel := g.getContext()
	defer cancel()

	resp, err := g.Client.DelegateResource(ctx, contract)
	if err != nil {
		return nil, fmt.Errorf("DelegateResource RPC error: %w", err)
	}
	return resp, nil
}

// UnDelegateResource revokes delegated resources.
func (g *GrpcClient) UnDelegateResource(owner, receiver string, resource core.ResourceCode, delegateBalance int64, lock bool) (*api.TransactionExtention, error) {
	addrOwner, err := base58.DecodeCheck(owner)
	if err != nil {
		return nil, fmt.Errorf("UnDelegateResource: failed to decode owner address: %w", err)
	}
	addrReceiver, err := base58.DecodeCheck(receiver)
	if err != nil {
		return nil, fmt.Errorf("UnDelegateResource: failed to decode receiver address: %w", err)
	}

	contract := &core.UnDelegateResourceContract{
		OwnerAddress:    addrOwner,
		ReceiverAddress: addrReceiver,
		Resource:        resource,
		Balance:         delegateBalance,
	}

	ctx, cancel := g.getContext()
	defer cancel()

	resp, err := g.Client.UnDelegateResource(ctx, contract)
	if err != nil {
		return nil, fmt.Errorf("UnDelegateResource RPC error: %w", err)
	}
	return resp, nil
}

// CancelAllUnfreezeV2 cancels all pending unfreeze operations.
func (g *GrpcClient) CancelAllUnfreezeV2(from string) (*api.TransactionExtention, error) {
	contract := &core.CancelAllUnfreezeV2Contract{}
	var err error

	if contract.OwnerAddress, err = base58.DecodeCheck(from); err != nil {
		return nil, fmt.Errorf("CancelAllUnfreezeV2: failed to decode from address: %w", err)
	}

	ctx, cancel := g.getContext()
	defer cancel()

	tx, err := g.Client.CancelAllUnfreezeV2(ctx, contract)
	if err != nil {
		return nil, fmt.Errorf("CancelAllUnfreezeV2 RPC error: %w", err)
	}
	if proto.Size(tx) == 0 {
		return nil, fmt.Errorf("CancelAllUnfreezeV2: bad transaction")
	}
	return tx, nil
}
