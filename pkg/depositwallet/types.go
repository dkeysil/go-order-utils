// Package depositwallet provides EIP-712 signing primitives for Polymarket
// deposit wallets — ERC-1967 proxies whose isValidSignature follows ERC-7739.
//
// Two flows live here:
//
//   - Order signing for deposit wallets is handled in pkg/builder via
//     SignatureType POLY_1271; this package's exports are not used there.
//   - Wallet Batch signing — used at onboarding to authorize allowance
//     approvals — is the primary export of this package. Sign a Batch with
//     SignBatch and POST it to the relayer's WALLET-typed /submit endpoint.
package depositwallet

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

// Call is a single inner transaction inside a Batch. The wallet executes
// `target.call{value: value}(data)` once the batch signature is validated.
type Call struct {
	Target common.Address
	Value  *big.Int
	Data   []byte
}

// Batch is a set of Calls authorized by a single EIP-712 signature from the
// wallet's owner EOA. Submitted to the relayer as a WALLET transaction.
type Batch struct {
	Wallet   common.Address
	Nonce    *big.Int
	Deadline *big.Int
	Calls    []Call
}

// SignedBatch is a Batch plus its 65-byte ECDSA signature.
type SignedBatch struct {
	Batch

	Signature []byte
}
