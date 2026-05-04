package depositwallet

import (
	"crypto/ecdsa"
	"math/big"

	"github.com/dkeysil/go-order-utils/pkg/eip712"
	"github.com/dkeysil/go-order-utils/pkg/signer"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

// HashCall returns hashStruct(Call) — keccak256(callTypeHash || abi.encode(
// target, value, keccak256(data))).
func HashCall(c Call) (common.Hash, error) {
	dataHash := crypto.Keccak256Hash(c.Data)
	return eip712.HashStruct(callArgs, []any{
		callTypeHash,
		c.Target,
		c.Value,
		dataHash,
	})
}

// HashCalls returns the Call[] field encoding: keccak256(concat(HashCall(c)))
// over each call. An empty slice yields keccak256("") per EIP-712.
func HashCalls(calls []Call) (common.Hash, error) {
	concat := make([]byte, 0, len(calls)*common.HashLength)
	for i := range calls {
		h, err := HashCall(calls[i])
		if err != nil {
			return common.Hash{}, err
		}
		concat = append(concat, h[:]...)
	}
	return crypto.Keccak256Hash(concat), nil
}

// HashBatch returns hashStruct(Batch).
func HashBatch(b Batch) (common.Hash, error) {
	callsHash, err := HashCalls(b.Calls)
	if err != nil {
		return common.Hash{}, err
	}
	return eip712.HashStruct(batchArgs, []any{
		batchTypeHash,
		b.Wallet,
		b.Nonce,
		b.Deadline,
		callsHash,
	})
}

// BatchDigest returns the final EIP-712 digest of a Batch under the
// DepositWallet domain. The domain separator is built on demand so this
// works for any chain or wallet address without a pre-registered list.
func BatchDigest(chainID *big.Int, b Batch) (common.Hash, error) {
	domainSep, err := eip712.BuildEIP712DomainSeparator(DomainName, DomainVersion, chainID, b.Wallet)
	if err != nil {
		return common.Hash{}, err
	}
	structHash, err := HashBatch(b)
	if err != nil {
		return common.Hash{}, err
	}
	return eip712.HashFromStructHash(domainSep, structHash), nil
}

// SignBatch produces a 65-byte ECDSA signature over the batch digest. The
// caller is responsible for passing the EOA that owns the wallet.
func SignBatch(privateKey *ecdsa.PrivateKey, chainID *big.Int, b Batch) (*SignedBatch, error) {
	digest, err := BatchDigest(chainID, b)
	if err != nil {
		return nil, err
	}
	sig, err := signer.Sign(privateKey, digest)
	if err != nil {
		return nil, err
	}
	return &SignedBatch{Batch: b, Signature: sig}, nil
}
