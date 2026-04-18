// Package signer wraps go-ethereum ECDSA signing with the Ethereum 27/28
// recovery-byte convention used by Polymarket order signatures.
package signer

import (
	"bytes"
	"crypto/ecdsa"
	"errors"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

// ErrInvalidSignatureLen is returned by ValidateSignature when the signature
// is not exactly 65 bytes long.
var ErrInvalidSignatureLen = errors.New("invalid signature length")

// Sign produces a 65-byte ECDSA signature over hashedData with V adjusted to
// the Ethereum 27/28 convention.
func Sign(privateKey *ecdsa.PrivateKey, hashedData common.Hash) ([]byte, error) {
	sign, err := crypto.Sign(hashedData.Bytes(), privateKey)
	if err != nil {
		return nil, err
	}
	sign[64] += 27
	return sign, err
}

// ValidateSignature recovers the signer address from signature/hashedData and
// compares it to the expected signer. Accepts both Ethereum-style (27/28) and
// raw (0/1) recovery bytes.
func ValidateSignature(signer common.Address, hashedData common.Hash, signature []byte) (bool, error) {
	sigCopy := make([]byte, len(signature))
	copy(sigCopy, signature)

	if len(sigCopy) != 65 {
		return false, ErrInvalidSignatureLen
	}

	if sigCopy[64] != 0 && sigCopy[64] != 1 { // in case of ledger signing v might already be 0 or 1
		sigCopy[64] -= 27 // Transform V from 27/28 to 0/1 according to the yellow paper
	}

	sigPublicKey, err := crypto.Ecrecover(hashedData.Bytes(), sigCopy)
	if err != nil {
		return false, err
	}

	recoveredPublicKey, err := crypto.UnmarshalPubkey(sigPublicKey)
	if err != nil {
		return false, err
	}

	recoveredAddress := crypto.PubkeyToAddress(*recoveredPublicKey)
	return bytes.Equal(signer.Bytes(), recoveredAddress.Bytes()), nil
}
