//
// Copyright (C) 2024 dszi
//
// This file may be modified and distributed under the terms
// of the MIT license. See the LICENSE file for details.
// Repository: https://github.com/dszi/go-tron
//

package pkg

import (
	"crypto/sha256"
	"fmt"
	"strconv"

	"github.com/dszi/go-tron/common/base58"
	hex "github.com/dszi/go-tron/common/hexutil"
	"github.com/dszi/go-tron/pb/api"
	"github.com/dszi/go-tron/pb/core"
	"github.com/dszi/go-tron/pkg/abi"
	"google.golang.org/protobuf/proto"
)

// DeployContract deploys a contract and returns the transaction result.
func (g *GrpcClient) DeployContract(from, contractName string, abi *core.SmartContract_ABI, codeStr string, feeLimit, curPercent, oeLimit int64) (*api.TransactionExtention, error) {
	var err error

	fromDesc, err := base58.DecodeCheck(from)
	if err != nil {
		return nil, fmt.Errorf("invalid sender address: %w", err)
	}
	if curPercent < 0 || curPercent > 100 {
		return nil, fmt.Errorf("consume_user_resource_percent should be between 0 and 100")
	}
	if oeLimit <= 0 {
		return nil, fmt.Errorf("origin_energy_limit must be greater than 0")
	}

	bc, err := hex.FromHex(codeStr)
	if err != nil {
		return nil, fmt.Errorf("invalid bytecode: %w", err)
	}

	ct := &core.CreateSmartContract{
		OwnerAddress: fromDesc,
		NewContract: &core.SmartContract{
			OriginAddress:              fromDesc,
			Abi:                        abi,
			Name:                       contractName,
			ConsumeUserResourcePercent: curPercent,
			OriginEnergyLimit:          oeLimit,
			Bytecode:                   bc,
		},
	}

	ctx, cancel := g.getContext()
	defer cancel()

	tx, err := g.Client.DeployContract(ctx, ct)
	if err != nil {
		return nil, fmt.Errorf("failed to deploy contract: %w", err)
	}
	if feeLimit > 0 {
		tx.Transaction.RawData.FeeLimit = feeLimit
		g.UpdateHash(tx)
	}
	return tx, err
}

// TriggerContract executes a contract function.
func (g *GrpcClient) TriggerContract(from, contractAddress, method, jsonString string, feeLimit, tAmount int64, tTokenID string, tTokenAmount int64) (*api.TransactionExtention, error) {
	fromDesc, err := base58.DecodeCheck(from)
	if err != nil {
		return nil, fmt.Errorf("invalid sender address: %w", err)
	}

	contractDesc, err := base58.DecodeCheck(contractAddress)
	if err != nil {
		return nil, fmt.Errorf("invalid contract address: %w", err)
	}

	param, err := abi.LoadFromJSON(jsonString)
	if err != nil {
		return nil, fmt.Errorf("failed to load JSON parameters: %w", err)
	}

	dataBytes, err := abi.Pack(method, param)
	if err != nil {
		return nil, fmt.Errorf("failed to encode method parameters: %w", err)
	}

	ct := &core.TriggerSmartContract{
		OwnerAddress:    fromDesc,
		ContractAddress: contractDesc,
		Data:            dataBytes,
		CallValue:       tAmount,
	}

	if len(tTokenID) > 0 && tTokenAmount > 0 {
		ct.CallTokenValue = tTokenAmount
		ct.TokenId, err = strconv.ParseInt(tTokenID, 10, 64)
		if err != nil {
			return nil, fmt.Errorf("invalid token ID: %w", err)
		}
	}

	return g.triggerContract(ct, feeLimit)
}

// triggerContract sends a smart contract execution transaction.
func (g *GrpcClient) triggerContract(ct *core.TriggerSmartContract, feeLimit int64) (*api.TransactionExtention, error) {
	ctx, cancel := g.getContext()
	defer cancel()

	tx, err := g.Client.TriggerContract(ctx, ct)
	if err != nil {
		return nil, fmt.Errorf("failed to trigger contract: %w", err)
	}

	if tx.Result.Code > 0 {
		return nil, fmt.Errorf("contract execution failed: %s", string(tx.Result.Message))
	}
	if feeLimit > 0 {
		tx.Transaction.RawData.FeeLimit = feeLimit
		g.UpdateHash(tx)
	}
	return tx, err
}

// UpdateHash updates the transaction hash after local modifications.
func (g *GrpcClient) UpdateHash(tx *api.TransactionExtention) error {
	rawData, err := proto.Marshal(tx.Transaction.GetRawData())
	if err != nil {
		return err
	}

	h256h := sha256.New()
	h256h.Write(rawData)
	hash := h256h.Sum(nil)
	tx.Txid = hash
	return nil
}

// GetContractABI retrieves the ABI of a deployed contract.
func (g *GrpcClient) GetContractABI(contractAddress string) (*core.SmartContract_ABI, error) {
	contractDesc, err := base58.DecodeCheck(contractAddress)
	if err != nil {
		return nil, fmt.Errorf("invalid contract address: %w", err)
	}

	ctx, cancel := g.getContext()
	defer cancel()

	sm, err := g.Client.GetContract(ctx, GetMessageBytes(contractDesc))
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve contract: %w", err)
	}

	if sm == nil {
		return nil, fmt.Errorf("contract ABI not found")
	}

	return sm.Abi, nil
}
