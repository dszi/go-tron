//
// Copyright (C) 2024 dszi
//
// This file may be modified and distributed under the terms
// of the MIT license. See the LICENSE file for details.
// Repository: https://github.com/dszi/go-tron
//

package abi

import (
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"golang.org/x/crypto/sha3"
	"math/big"
	"reflect"
	"strconv"
	"strings"

	"github.com/dszi/go-tron/common/base58"
	eABI "github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
)

// Param defines a key-value structure for smart contract parameters.
type Param map[string]any

// LoadFromJSON parses a JSON string into an array of Params.
// It expects a JSON array format and returns parsed parameters or an error.
func LoadFromJSON(jString string) ([]Param, error) {
	if len(jString) == 0 {
		return nil, nil
	}
	var data []Param
	err := json.Unmarshal([]byte(jString), &data)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON: %w", err)
	}
	return data, nil
}

// convertToAddress converts a TRON base58-encoded address into an Ethereum-style address (20 bytes).
// It extracts the last 20 bytes from the decoded address.
func convertToAddress(v any) (common.Address, error) {
	str, ok := v.(string)
	if !ok {
		return common.Address{}, errors.New("invalid address type")
	}

	addr, err := base58.DecodeCheck(str)
	if err != nil {
		return common.Address{}, fmt.Errorf("invalid base58 address %s: %w", str, err)
	}

	if len(addr) < 20 {
		return common.Address{}, fmt.Errorf("invalid address length: expected 20, got %d", len(addr))
	}

	return common.BytesToAddress(addr[len(addr)-20:]), nil
}

// convertToInt converts various integer types into the required ABI integer format.
// Supports int8, int16, int32, int64, uint8, uint16, uint32, uint64, *big.Int, and string representations.
func convertToInt(ty eABI.Type, v interface{}) (interface{}, error) {
	switch val := v.(type) {
	case string:
		return parseIntFromString(ty, val)
	case int, int8, int16, int32, int64:
		return castIntSize(reflect.ValueOf(val).Int(), ty.Size), nil
	case uint, uint8, uint16, uint32, uint64:
		return castUintSize(reflect.ValueOf(val).Uint(), ty.Size), nil
	case *big.Int:
		return val, nil
	default:
		return nil, fmt.Errorf("unsupported type for integer conversion: %T", v)
	}
}

// parseIntFromString parses string representations of integers (decimal or hex) into appropriate numeric values.
func parseIntFromString(ty eABI.Type, s string) (interface{}, error) {
	if ty.T == eABI.IntTy && ty.Size <= 64 {
		tmp, err := strconv.ParseInt(s, 10, 64)
		if err != nil {
			return nil, fmt.Errorf("failed to parse int: %w", err)
		}
		return castIntSize(tmp, ty.Size), nil
	}

	if ty.T == eABI.UintTy && ty.Size <= 64 {
		tmp, err := strconv.ParseUint(s, 10, 64)
		if err != nil {
			return nil, fmt.Errorf("failed to parse uint: %w", err)
		}
		return castUintSize(tmp, ty.Size), nil
	}

	// Handle large integers (e.g., uint256)
	var intValue *big.Int
	if strings.HasPrefix(s, "0x") {
		intValue, _ = new(big.Int).SetString(s[2:], 16)
	} else {
		intValue, _ = new(big.Int).SetString(s, 10)
	}

	if intValue == nil {
		return nil, fmt.Errorf("failed to parse big.Int from: %s", s)
	}
	return intValue, nil
}

// castIntSize ensures signed integer types match the expected ABI size.
func castIntSize(value int64, size int) interface{} {
	switch size {
	case 8:
		return int8(value)
	case 16:
		return int16(value)
	case 32:
		return int32(value)
	case 64:
		return int64(value)
	default:
		return value
	}
}

// castUintSize ensures unsigned integer types match the expected ABI size.
func castUintSize(value uint64, size int) interface{} {
	switch size {
	case 8:
		return uint8(value)
	case 16:
		return uint16(value)
	case 32:
		return uint32(value)
	case 64:
		return uint64(value)
	default:
		return value
	}
}

// convertToBytes converts string input to a byte slice for ABI encoding.
// Supports both **hex** and **base64** encoding.
func convertToBytes(ty eABI.Type, v interface{}) (interface{}, error) {
	if data, ok := v.(string); ok {
		dataBytes, err := hex.DecodeString(data)
		if err != nil {
			dataBytes, err = base64.StdEncoding.DecodeString(data)
			if err != nil {
				return nil, err
			}
		}

		if ty.T == eABI.BytesTy || ty.Size == 0 {
			return dataBytes, nil
		}
		if len(dataBytes) != ty.Size {
			return nil, fmt.Errorf("invalid size: %d/%d", ty.Size, len(dataBytes))
		}
		switch ty.Size {
		case 1:
			value := [1]byte{}
			copy(value[:], dataBytes[:1])
			return value, nil
		case 2:
			value := [2]byte{}
			copy(value[:], dataBytes[:2])
			return value, nil
		case 8:
			value := [8]byte{}
			copy(value[:], dataBytes[:8])
			return value, nil
		case 16:
			value := [16]byte{}
			copy(value[:], dataBytes[:16])
			return value, nil
		case 32:
			value := [32]byte{}
			copy(value[:], dataBytes[:32])
			return value, nil
		}
	}
	return v, nil
}

// GetPaddedParam encodes input parameters into ABI-compliant byte arrays.
func GetPaddedParam(param []Param) ([]byte, error) {
	values := make([]interface{}, 0)
	arguments := eABI.Arguments{}

	for _, p := range param {
		if len(p) != 1 {
			return nil, fmt.Errorf("invalid param format: %+v", p)
		}

		for k, v := range p {
			// Determine the ABI type
			ty, err := eABI.NewType(k, "", nil)
			if err != nil {
				return nil, fmt.Errorf("invalid parameter type %s: %+v", k, err)
			}
			arguments = append(arguments, eABI.Argument{Name: "", Type: ty})

			// Handle arrays and slices
			if ty.T == eABI.SliceTy || ty.T == eABI.ArrayTy {
				switch ty.Elem.T {
				case eABI.AddressTy:
					tmp, ok := v.([]interface{})
					if !ok {
						return nil, fmt.Errorf("expected address array but got %+v", p)
					}

					if ty.T == eABI.ArrayTy && len(tmp) != ty.Size {
						return nil, fmt.Errorf("invalid array size, expected %d but got %d", ty.Size, len(tmp))
					}

					addresses := make([]common.Address, len(tmp))
					for i := range tmp {
						addr, err := convertToAddress(tmp[i])
						if err != nil {
							return nil, err
						}
						addresses[i] = addr
					}
					v = addresses

				case eABI.IntTy, eABI.UintTy:
					if ty.Elem.Size > 64 {
						tmp := make([]*big.Int, 0)
						tmpSlice, ok := v.([]string)
						if !ok {
							return nil, fmt.Errorf("expected array of uint but got %+v", p)
						}
						for i := range tmpSlice {
							val, err := parseBigInt(tmpSlice[i])
							if err != nil {
								return nil, err
							}
							tmp = append(tmp, val)
						}
						v = tmp
					}
				}
			}

			if ty.T == eABI.AddressTy {
				v, err = convertToAddress(v)
				if err != nil {
					return nil, err
				}
			}

			if (ty.T == eABI.IntTy || ty.T == eABI.UintTy) && reflect.TypeOf(v).Kind() == reflect.String {
				v, err = convertToInt(ty, v)
				if err != nil {
					return nil, err
				}
			}

			if ty.T == eABI.BytesTy || ty.T == eABI.FixedBytesTy {
				v, err = convertToBytes(ty, v)
				if err != nil {
					return nil, err
				}
			}

			values = append(values, v)
		}
	}

	return arguments.PackValues(values)
}

// parseBigInt parses a string into a *big.Int.
// Supports both decimal (e.g., "1000000") and hex formats (e.g., "0x1e8480").
func parseBigInt(s string) (*big.Int, error) {
	num := new(big.Int)

	if strings.HasPrefix(s, "0x") {
		if _, success := num.SetString(s[2:], 16); !success {
			return nil, errors.New("invalid hexadecimal number: " + s)
		}
	} else {
		if _, success := num.SetString(s, 10); !success {
			return nil, errors.New("invalid decimal number: " + s)
		}
	}

	return num, nil
}

// Pack encodes method signature and parameters into ABI-compliant bytes.
func Pack(method string, param []Param) ([]byte, error) {
	// Signature
	hasher := sha3.NewLegacyKeccak256()
	hasher.Write([]byte(method))
	signature := hasher.Sum(nil)[:4]

	pBytes, err := GetPaddedParam(param)
	if err != nil {
		return nil, err
	}
	signature = append(signature, pBytes...)
	return signature, nil
}
