// Package builder assembles, hashes, and signs V2 CTF Exchange orders.
package builder

import (
	"crypto/ecdsa"

	"github.com/dkeysil/go-order-utils/pkg/model"
)

// ExchangeOrderBuilder turns OrderData into signed V2 CTF Exchange orders.
//
//go:generate mockery --name ExchangeOrderBuilder
type ExchangeOrderBuilder interface {
	// BuildSignedOrder assembles an Order, computes the EIP-712 digest, and
	// returns the order plus its ECDSA signature.
	BuildSignedOrder(privateKey *ecdsa.PrivateKey, orderData *model.OrderData, contract model.VerifyingContract) (*model.SignedOrder, error)

	// BuildOrder populates an Order from OrderData, filling defaults (salt,
	// timestamp, signer ← maker) as needed.
	BuildOrder(orderData *model.OrderData) (*model.Order, error)

	// BuildOrderHash returns the EIP-712 digest of an Order under the given
	// verifying contract.
	BuildOrderHash(order *model.Order, contract model.VerifyingContract) (model.OrderHash, error)

	// BuildOrderSignature produces the 65-byte ECDSA signature over orderHash.
	BuildOrderSignature(privateKey *ecdsa.PrivateKey, orderHash model.OrderHash) (model.OrderSignature, error)
}
