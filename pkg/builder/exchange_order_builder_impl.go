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

// ExchangeOrderBuilderImpl is the default ExchangeOrderBuilder implementation.
type ExchangeOrderBuilderImpl struct {
	chainID            *big.Int
	saltGenerator      func() int64
	timestampGenerator func() int64
	domainSeparators   map[model.VerifyingContract]common.Hash
}

var _ ExchangeOrderBuilder = (*ExchangeOrderBuilderImpl)(nil)

// NewExchangeOrderBuilderImpl builds a V2 CTF Exchange order builder.
//
// saltGenerator / timestampGenerator may be nil; they default to
// utils.GenerateRandomSalt / utils.GenerateTimestampMs respectively.
//
// The domain separators for both supported verifying contracts are
// precomputed at construction time, so unsupported chain IDs surface their
// error on the first signing call rather than at every call.
func NewExchangeOrderBuilderImpl(chainID *big.Int, saltGenerator func() int64, timestampGenerator func() int64) *ExchangeOrderBuilderImpl {
	if saltGenerator == nil {
		saltGenerator = utils.GenerateRandomSalt
	}
	if timestampGenerator == nil {
		timestampGenerator = utils.GenerateTimestampMs
	}

	domainSeparators := map[model.VerifyingContract]common.Hash{}
	for _, contract := range []model.VerifyingContract{model.CTFExchange, model.NegRiskCTFExchange} {
		addr, err := utils.GetVerifyingContractAddress(chainID, contract)
		if err != nil {
			continue
		}
		sep, err := eip712.BuildEIP712DomainSeparator(protocolName, protocolVersion, chainID, addr)
		if err != nil {
			continue
		}
		domainSeparators[contract] = sep
	}

	return &ExchangeOrderBuilderImpl{
		chainID:            chainID,
		saltGenerator:      saltGenerator,
		timestampGenerator: timestampGenerator,
		domainSeparators:   domainSeparators,
	}
}

// BuildSignedOrder assembles an Order, hashes it, and attaches an ECDSA signature.
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

// BuildOrder converts an OrderData into a fully-populated V2 Order, defaulting
// Salt, Timestamp, Signer (← Maker), and the zero-valued Metadata/Builder.
func (e *ExchangeOrderBuilderImpl) BuildOrder(orderData *model.OrderData) (*model.Order, error) {
	maker := common.HexToAddress(orderData.Maker)
	signerAddr := maker
	if orderData.Signer != "" {
		signerAddr = common.HexToAddress(orderData.Signer)
	}

	tokenID, ok := new(big.Int).SetString(orderData.TokenId, 10)
	if !ok {
		return nil, fmt.Errorf("can't parse TokenId: %s as valid *big.Int", orderData.TokenId)
	}

	makerAmount, ok := new(big.Int).SetString(orderData.MakerAmount, 10)
	if !ok {
		return nil, fmt.Errorf("can't parse MakerAmount: %s as valid *big.Int", orderData.MakerAmount)
	}

	takerAmount, ok := new(big.Int).SetString(orderData.TakerAmount, 10)
	if !ok {
		return nil, fmt.Errorf("can't parse TakerAmount: %s as valid *big.Int", orderData.TakerAmount)
	}

	ts := orderData.Timestamp
	if ts == 0 {
		ts = e.timestampGenerator()
	}

	return &model.Order{
		Salt:          new(big.Int).SetInt64(e.saltGenerator()),
		Maker:         maker,
		Signer:        signerAddr,
		TokenId:       tokenID,
		MakerAmount:   makerAmount,
		TakerAmount:   takerAmount,
		Side:          uint8(orderData.Side),
		SignatureType: uint8(orderData.SignatureType),
		Timestamp:     new(big.Int).SetInt64(ts),
		Metadata:      orderData.Metadata,
		Builder:       orderData.Builder,
	}, nil
}

// BuildOrderHash returns the EIP-712 digest for order under the given
// verifying contract.
func (e *ExchangeOrderBuilderImpl) BuildOrderHash(order *model.Order, contract model.VerifyingContract) (model.OrderHash, error) {
	domainSeparator, ok := e.domainSeparators[contract]
	if !ok {
		return model.OrderHash{}, fmt.Errorf("unsupported verifying contract %d for chain %s", contract, e.chainID)
	}

	values := []any{
		orderStructureHash,
		order.Salt,
		order.Maker,
		order.Signer,
		order.TokenId,
		order.MakerAmount,
		order.TakerAmount,
		order.Side,
		order.SignatureType,
		order.Timestamp,
		order.Metadata,
		order.Builder,
	}
	return eip712.HashTypedDataV4(domainSeparator, orderStructure, values)
}

// BuildOrderSignature signs orderHash with privateKey and returns the 65-byte
// ECDSA signature.
func (e *ExchangeOrderBuilderImpl) BuildOrderSignature(privateKey *ecdsa.PrivateKey, orderHash model.OrderHash) (model.OrderSignature, error) {
	return signer.Sign(privateKey, orderHash)
}
