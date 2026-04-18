package model

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

type OrderSignature = []byte

type OrderHash = common.Hash

// Zero-value bytes32 used for unset Metadata / Builder fields.
var Bytes32Zero = common.Hash{}

// OrderData is the caller-facing input to the V2 CTF Exchange order builder.
//
// Fields removed relative to V1: Taker, Expiration (no longer signed), Nonce, FeeRateBps.
// Fields added: Timestamp (unix milliseconds), Metadata (bytes32), Builder (bytes32).
type OrderData struct {
	// Maker of the order, i.e the source of funds for the order
	Maker string

	// Token Id of the CTF ERC1155 asset to be bought or sold.
	// If BUY, this is the tokenId of the asset to be bought, i.e the makerAssetId
	// If SELL, this is the tokenId of the asset to be sold, i.e the takerAssetId
	TokenId string

	// Maker amount, i.e the max amount of tokens to be sold
	MakerAmount string

	// Taker amount, i.e the minimum amount of tokens to be received
	TakerAmount string

	// The side of the order, BUY or SELL
	Side Side

	// Signer of the order. Optional, if it is not present the signer is the maker of the order.
	Signer string

	// Signature type used by the Order. Default value 'EOA'.
	SignatureType SignatureType

	// Order creation time, unix milliseconds as a base-10 string.
	// Optional; defaults to time.Now().UnixMilli() when empty.
	// Replaces V1's Nonce for uniqueness.
	Timestamp string

	// Arbitrary bytes32 metadata attached to the order (0x-prefixed hex).
	// Optional; defaults to the zero hash when empty.
	Metadata string

	// Builder code (bytes32, 0x-prefixed hex) used for builder attribution.
	// Optional; defaults to the zero hash when empty.
	Builder string
}

// Order is the canonical V2 signed struct hashed into the EIP-712 digest.
//
// Field order here must match pkg/builder.constants.go `_ORDER_STRUCTURE` /
// `_ORDER_STRUCTURE_HASH` or signatures will be rejected onchain.
type Order struct {
	// Unique salt to ensure entropy
	Salt *big.Int

	// Maker of the order, i.e the source of funds for the order
	Maker common.Address

	// Signer of the order
	Signer common.Address

	// Token Id of the CTF ERC1155 asset to be bought or sold.
	TokenId *big.Int

	// Maker amount, i.e the max amount of tokens to be sold
	MakerAmount *big.Int

	// Taker amount, i.e the minimum amount of tokens to be received
	TakerAmount *big.Int

	// The side of the order, BUY or SELL
	Side *big.Int

	// Signature type used by the Order
	SignatureType *big.Int

	// Unix timestamp in milliseconds at which the order was created
	Timestamp *big.Int

	// Arbitrary bytes32 metadata
	Metadata common.Hash

	// Builder code (bytes32)
	Builder common.Hash
}

type SignedOrder struct {
	Order

	// The order signature
	Signature OrderSignature
}
