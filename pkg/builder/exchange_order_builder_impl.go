package builder

import (
	"crypto/ecdsa"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/polymarket/go-order-utils/pkg/eip712"
	"github.com/polymarket/go-order-utils/pkg/model"
	"github.com/polymarket/go-order-utils/pkg/signer"
	"github.com/polymarket/go-order-utils/pkg/utils"
)

type ExchangeOrderBuilderImpl struct {
	chainId            *big.Int
	saltGenerator      func() int64
	timestampGenerator func() int64
}

var _ ExchangeOrderBuilder = (*ExchangeOrderBuilderImpl)(nil)

// Option configures optional fields on an ExchangeOrderBuilderImpl at
// construction time. Use functional options instead of growing the constructor
// arg list for each new knob.
type Option func(*ExchangeOrderBuilderImpl)

// WithTimestampGenerator overrides the default unix-millisecond timestamp
// generator. Primarily useful for deterministic tests.
func WithTimestampGenerator(fn func() int64) Option {
	return func(b *ExchangeOrderBuilderImpl) {
		if fn != nil {
			b.timestampGenerator = fn
		}
	}
}

// NewExchangeOrderBuilderImpl builds a V2 CTF Exchange order builder.
//
// saltGenerator may be nil; it defaults to utils.GenerateRandomSalt.
// Use WithTimestampGenerator to override the default time.Now().UnixMilli().
func NewExchangeOrderBuilderImpl(chainId *big.Int, saltGenerator func() int64, opts ...Option) *ExchangeOrderBuilderImpl {
	if saltGenerator == nil {
		saltGenerator = utils.GenerateRandomSalt
	}
	b := &ExchangeOrderBuilderImpl{
		chainId:            chainId,
		saltGenerator:      saltGenerator,
		timestampGenerator: utils.GenerateTimestampMs,
	}
	for _, opt := range opts {
		opt(b)
	}
	return b
}

// BuildSignedOrder assembles a V2 order, hashes it under the EIP-712 domain
// for the given exchange, and attaches an ECDSA signature.
//
// The signature is verified locally before returning so that wiring mistakes
// (wrong signer address, wrong typehash, ...) surface immediately.
func (e *ExchangeOrderBuilderImpl) BuildSignedOrder(privateKey *ecdsa.PrivateKey, orderData *model.OrderData, contract model.VerifyingContract) (*model.SignedOrder, error) {
	order, err := e.BuildOrder(orderData)
	if err != nil {
		return nil, err
	}

	orderHash, err := e.BuildOrderHash(order, contract)
	if err != nil {
		return nil, err
	}

	signature, err := e.BuildOrderSignature(privateKey, orderHash)
	if err != nil {
		return nil, err
	}

	ok, err := signer.ValidateSignature(order.Signer, orderHash, signature)
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, fmt.Errorf("signature error")
	}

	return &model.SignedOrder{
		Order:     *order,
		Signature: signature,
	}, nil
}

// BuildOrder converts an OrderData into a fully-populated V2 Order, filling in
// Salt, Timestamp, Metadata, Builder, and defaulting Signer←Maker when the
// caller left Signer empty.
func (e *ExchangeOrderBuilderImpl) BuildOrder(orderData *model.OrderData) (*model.Order, error) {
	var signer common.Address
	if orderData.Signer == "" {
		signer = common.HexToAddress(orderData.Maker)
	} else {
		signer = common.HexToAddress(orderData.Signer)
	}

	var tokenId *big.Int
	var ok bool
	if tokenId, ok = new(big.Int).SetString(orderData.TokenId, 10); !ok {
		return nil, fmt.Errorf("can't parse TokenId: %s as valid *big.Int", orderData.TokenId)
	}

	var makerAmount *big.Int
	if makerAmount, ok = new(big.Int).SetString(orderData.MakerAmount, 10); !ok {
		return nil, fmt.Errorf("can't parse MakerAmount: %s as valid *big.Int", orderData.MakerAmount)
	}

	var takerAmount *big.Int
	if takerAmount, ok = new(big.Int).SetString(orderData.TakerAmount, 10); !ok {
		return nil, fmt.Errorf("can't parse TakerAmount: %s as valid *big.Int", orderData.TakerAmount)
	}

	var timestamp *big.Int
	if orderData.Timestamp == "" {
		timestamp = new(big.Int).SetInt64(e.timestampGenerator())
	} else if timestamp, ok = new(big.Int).SetString(orderData.Timestamp, 10); !ok {
		return nil, fmt.Errorf("can't parse Timestamp: %s as valid *big.Int", orderData.Timestamp)
	}

	metadata, err := parseBytes32(orderData.Metadata, "Metadata")
	if err != nil {
		return nil, err
	}

	builder, err := parseBytes32(orderData.Builder, "Builder")
	if err != nil {
		return nil, err
	}

	return &model.Order{
		Salt:          new(big.Int).SetInt64(e.saltGenerator()),
		Maker:         common.HexToAddress(orderData.Maker),
		Signer:        signer,
		TokenId:       tokenId,
		MakerAmount:   makerAmount,
		TakerAmount:   takerAmount,
		Side:          new(big.Int).SetInt64(int64(orderData.Side)),
		SignatureType: new(big.Int).SetInt64(int64(orderData.SignatureType)),
		Timestamp:     timestamp,
		Metadata:      metadata,
		Builder:       builder,
	}, nil
}

// BuildOrderHash returns the EIP-712 digest for the V2 Order under the
// requested verifying contract (CTF Exchange V2 or Neg-Risk CTF Exchange V2).
func (e *ExchangeOrderBuilderImpl) BuildOrderHash(order *model.Order, contract model.VerifyingContract) (model.OrderHash, error) {
	verifyingContract, err := utils.GetVerifyingContractAddress(e.chainId, contract)
	if err != nil {
		return model.OrderHash{}, err
	}

	domainSeparator, err := eip712.BuildEIP712DomainSeparator(_PROTOCOL_NAME, _PROTOCOL_VERSION, e.chainId, verifyingContract)
	if err != nil {
		return model.OrderHash{}, err
	}

	values := []interface{}{
		_ORDER_STRUCTURE_HASH,
		order.Salt,
		order.Maker,
		order.Signer,
		order.TokenId,
		order.MakerAmount,
		order.TakerAmount,
		uint8(order.Side.Uint64()),
		uint8(order.SignatureType.Uint64()),
		order.Timestamp,
		order.Metadata,
		order.Builder,
	}
	orderHash, err := eip712.HashTypedDataV4(domainSeparator, _ORDER_STRUCTURE, values)
	if err != nil {
		return model.OrderHash{}, err
	}

	return orderHash, nil
}

func (e *ExchangeOrderBuilderImpl) BuildOrderSignature(privateKey *ecdsa.PrivateKey, orderHash model.OrderHash) (model.OrderSignature, error) {
	return signer.Sign(privateKey, orderHash)
}

// parseBytes32 accepts an empty string (→ zero hash) or a 0x-prefixed hex
// string representing exactly 32 bytes. Returns an error for any other shape
// so malformed metadata/builder values fail fast rather than silently
// producing a wrong signature.
func parseBytes32(s, field string) (common.Hash, error) {
	if s == "" {
		return model.Bytes32Zero, nil
	}
	hex := s
	if len(hex) >= 2 && hex[:2] == "0x" {
		hex = hex[2:]
	}
	if len(hex) != 64 {
		return common.Hash{}, fmt.Errorf("invalid %s: expected 32-byte 0x-prefixed hex, got %q", field, s)
	}
	h := common.HexToHash(s)
	return h, nil
}
