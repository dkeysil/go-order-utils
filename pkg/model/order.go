package model

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

type OrderSignature = []byte

type OrderHash = common.Hash

// OrderData is the caller-facing input to the V2 CTF Exchange order builder.
type OrderData struct {
	// Maker of the order, i.e the source of funds for the order
	Maker string

	// Token Id of the CTF ERC1155 asset to be bought or sold.
	TokenId string

	// Maker amount, i.e the max amount of tokens to be sold
	MakerAmount string

	// Taker amount, i.e the minimum amount of tokens to be received
	TakerAmount string

	// The side of the order, BUY or SELL
	Side Side

	// Signer of the order. Optional, if it is not present the signer is the maker of the order.
	Signer string

	// Signature type used by the Order.
	SignatureType SignatureType

	// Order creation time, unix milliseconds. Zero means "use the builder's
	// timestamp generator" (by default time.Now().UnixMilli()).
	Timestamp int64

	// Arbitrary bytes32 metadata attached to the order. Zero value is allowed.
	Metadata common.Hash

	// Builder code (bytes32) used for builder attribution. Zero value is allowed.
	Builder common.Hash
}

// Order is the canonical V2 signed struct hashed into the EIP-712 digest.
//
// Field order here must match builder._ORDER_STRUCTURE.
type Order struct {
	Salt          *big.Int
	Maker         common.Address
	Signer        common.Address
	TokenId       *big.Int
	MakerAmount   *big.Int
	TakerAmount   *big.Int
	Side          uint8
	SignatureType uint8
	Timestamp     *big.Int
	Metadata      common.Hash
	Builder       common.Hash
}

type SignedOrder struct {
	Order

	Signature OrderSignature
}
