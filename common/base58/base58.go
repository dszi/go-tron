//
// Copyright (C) 2024 dszi
//
// This file may be modified and distributed under the terms
// of the MIT license. See the LICENSE file for details.
// Repository: https://github.com/dszi/go-tron
//

package base58

import (
	"bytes"
	"crypto/sha256"
	"fmt"

	"github.com/btcsuite/btcutil/base58"
)

const (
	addressLength = 20
	prefixMainnet = 0x41
)

// Encode encodes the input bytes into a Base58 string using the Bitcoin alphabet.
func Encode(input []byte) string {
	return base58.Encode(input)
}

// EncodeCheck performs Base58Check encoding on the input bytes.
// It appends a 4-byte checksum (derived from a double SHA256 hash) to the input before encoding.
func EncodeCheck(input []byte) string {
	h0 := sha256.Sum256(input)
	h1 := sha256.Sum256(h0[:])
	dataWithChecksum := append(input, h1[:4]...)
	return Encode(dataWithChecksum)
}

// Decode decodes a Base58 encoded string into a byte slice.
// Returns an error if the decoding fails.
func Decode(input string) ([]byte, error) {
	decoded := base58.Decode(input)
	if len(decoded) == 0 {
		return nil, fmt.Errorf("decode error: empty result")
	}
	return decoded, nil
}

// DecodeCheck decodes a Base58Check encoded string and verifies its checksum.
// For a TRON address, the expected length is 1 (prefix) + 20 (address data) + 4 (checksum) = 25 bytes.
func DecodeCheck(input string) ([]byte, error) {
	decoded, err := Decode(input)
	if err != nil {
		return nil, err
	}

	if len(decoded) < 4 {
		return nil, fmt.Errorf("b58 check error: insufficient length")
	}

	expectedLen := addressLength + 4 + 1
	if len(decoded) != expectedLen {
		return nil, fmt.Errorf("invalid address length: %d", len(decoded))
	}

	if decoded[0] != prefixMainnet {
		return nil, fmt.Errorf("invalid prefix")
	}

	data := decoded[:len(decoded)-4]

	h0 := sha256.Sum256(data)
	h1 := sha256.Sum256(h0[:])
	checksum := decoded[len(decoded)-4:]

	if !bytes.Equal(h1[:4], checksum) {
		return nil, fmt.Errorf("b58 check error: checksum mismatch")
	}

	return data, nil
}
