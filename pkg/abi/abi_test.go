//
// Copyright (C) 2024 dszi
//
// This file may be modified and distributed under the terms
// of the MIT license. See the LICENSE file for details.
// Repository: https://github.com/dszi/go-tron
//

package abi

import (
	"encoding/hex"
	"fmt"
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestABIParamBasicTypes verifies encoding for basic data types like string, uint8, and uint256.
func TestABIParamBasicTypes(t *testing.T) {
	ss, _ := new(big.Int).SetString("500000000000000000000", 10)

	b, err := GetPaddedParam([]Param{
		{"string": "TRX Test Token"},
		{"string": "TRX"},
		{"uint8": uint8(10)},
		{"uint256": ss},
	})
	require.Nil(t, err)
	assert.Len(t, b, 256, fmt.Sprintf("Wrong length %d/%d", len(b), 256))

	b, err = GetPaddedParam([]Param{
		{"string": "TRX Test Token"},
		{"string": "TRX"},
		{"uint8": "10"},
		{"uint256": ss.String()},
	})
	require.Nil(t, err)
	assert.Len(t, b, 256, fmt.Sprintf("Wrong length %d/%d", len(b), 256))
}

// TestABIParamAddressArray checks correct encoding of an address array.
func TestABIParamAddressArray(t *testing.T) {
	param, err := LoadFromJSON(`
	[
		{"address[2]":["TRGhNNfnmgLegT4zHNjEqDSADjgmnHvubJ", "TRGhNNfnmgLegT4zHNjEqDSADjgmnHvubJ"]}
	]
	`)
	require.Nil(t, err)

	b, err := GetPaddedParam(param)
	require.Nil(t, err)

	assert.Len(t, b, 64, fmt.Sprintf("Wrong length %d/%d", len(b), 64))
	assert.Equal(t, "000000000000000000000000a7d8a35b260395c14aa456297662092ba3b76fc0000000000000000000000000a7d8a35b260395c14aa456297662092ba3b76fc0", hex.EncodeToString(b))
}

// TestABIParamUint256Array tests encoding of an array of uint256 values.
func TestABIParamUint256Array(t *testing.T) {
	b, err := GetPaddedParam([]Param{
		{"uint256[2]": []string{"100000000000000000000", "200000000000000000000"}},
	})
	require.Nil(t, err)

	assert.Len(t, b, 64, fmt.Sprintf("Wrong length %d/%d", len(b), 64))
	assert.Equal(t, "0000000000000000000000000000000000000000000000056bc75e2d6310000000000000000000000000000000000000000000000000000ad78ebc5ac6200000", hex.EncodeToString(b))
}

// TestABIParamBytes32 validates encoding for a fixed-length bytes32 type.
func TestABIParamBytes32(t *testing.T) {
	param, err := LoadFromJSON(`
	[
		{"bytes32": "0001020001020001020001020001020001020001020001020001020001020001"}
	]
	`)
	require.Nil(t, err)

	b, err := GetPaddedParam(param)
	require.Nil(t, err)

	assert.Len(t, b, 32, fmt.Sprintf("Wrong length %d/%d", len(b), 32))
	assert.Equal(t, "0001020001020001020001020001020001020001020001020001020001020001", hex.EncodeToString(b))
}

// TestABIParamHexUint256 tests the handling of both decimal and hexadecimal uint256 values.
func TestABIParamHexUint256(t *testing.T) {
	b, err := GetPaddedParam([]Param{
		{"uint256": "123456"},
		{"uint256": "0x1E240"},
	})
	require.Nil(t, err)

	assert.Len(t, b, 64, fmt.Sprintf("Wrong length %d/%d", len(b), 64))
	assert.Equal(t, "000000000000000000000000000000000000000000000000000000000001e240000000000000000000000000000000000000000000000000000000000001e240", hex.EncodeToString(b))
}
