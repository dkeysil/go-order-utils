package builder

import (
	"crypto/ecdsa"
	"encoding/binary"
	"fmt"
	"math/big"

	"github.com/dkeysil/go-order-utils/pkg/depositwallet"
	"github.com/dkeysil/go-order-utils/pkg/eip712"
	"github.com/dkeysil/go-order-utils/pkg/model"
	"github.com/dkeysil/go-order-utils/pkg/signer"
	"github.com/dkeysil/go-order-utils/pkg/utils"
	"github.com/ethereum/go-ethereum/common"
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
	for contract, version := range protocolVersions {
		addr, err := utils.GetVerifyingContractAddress(chainID, contract)
		if err != nil {
			continue
		}
		sep, err := eip712.BuildEIP712DomainSeparator(protocolName, version, chainID, addr)
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

// BuildSignedOrder assembles an Order and attaches its signature.
//
// For EOA / POLY_PROXY / POLY_GNOSIS_SAFE the signature is the 65-byte ECDSA
// over the EIP-712 order digest, locally verified via ecrecover.
//
// For POLY_1271 the signature is an ERC-7739 composite for Polymarket deposit
// wallets:
//
//	innerSig(65) || appDomainSep(32) || contentsHash(32) || orderTypeString || uint16_BE(len)
//
// innerSig is the EOA's signature over the EIP-712 digest of a TypedDataSign
// struct under the CTFExchange app domain; the wallet's own EIP-712 domain
// (DepositWallet/1, verifyingContract = order.Signer) is encoded inside that
// struct. The wallet validates this on-chain via isValidSignature, so no
// local ecrecover-vs-Signer check is performed.
func (e *ExchangeOrderBuilderImpl) BuildSignedOrder(privateKey *ecdsa.PrivateKey, orderData *model.OrderData, contract model.VerifyingContract) (*model.SignedOrder, error) {
	order, err := e.BuildOrder(orderData)
	if err != nil {
		return nil, err
	}

	if order.SignatureType == uint8(model.POLY_1271) {
		signature, err := e.buildPoly1271Signature(privateKey, order, contract)
		if err != nil {
			return nil, err
		}
		return &model.SignedOrder{Order: *order, Signature: signature}, nil
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

	structHash, err := hashOrderStruct(order)
	if err != nil {
		return model.OrderHash{}, err
	}
	return eip712.HashFromStructHash(domainSeparator, structHash), nil
}

// BuildOrderSignature signs orderHash with privateKey and returns the 65-byte
// ECDSA signature.
func (e *ExchangeOrderBuilderImpl) BuildOrderSignature(privateKey *ecdsa.PrivateKey, orderHash model.OrderHash) (model.OrderSignature, error) {
	return signer.Sign(privateKey, orderHash)
}

// hashOrderStruct returns hashStruct(Order) — the EIP-712 inner struct hash,
// reused as `contentsHash` in the ERC-7739 wrap for POLY_1271 orders.
func hashOrderStruct(order *model.Order) (common.Hash, error) {
	return eip712.HashStruct(orderStructure, []any{
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
	})
}

// buildPoly1271Signature produces the ERC-7739 composite signature accepted
// by Polymarket deposit wallets' isValidSignature.
func (e *ExchangeOrderBuilderImpl) buildPoly1271Signature(privateKey *ecdsa.PrivateKey, order *model.Order, contract model.VerifyingContract) (model.OrderSignature, error) {
	appDomainSep, ok := e.domainSeparators[contract]
	if !ok {
		return nil, fmt.Errorf("unsupported verifying contract %d for chain %s", contract, e.chainID)
	}

	contentsHash, err := hashOrderStruct(order)
	if err != nil {
		return nil, err
	}

	outerStructHash, err := eip712.HashStruct(typedDataSignArgs, []any{
		typedDataSignTypeHash,
		contentsHash,
		depositwallet.DomainName,
		depositwallet.DomainVersion,
		e.chainID,
		order.Signer, // verifyingContract = the deposit wallet itself
		common.Hash{},
	})
	if err != nil {
		return nil, err
	}

	digest := eip712.HashFromStructHash(appDomainSep, outerStructHash)
	innerSig, err := signer.Sign(privateKey, digest)
	if err != nil {
		return nil, err
	}

	out := make([]byte, 0, len(innerSig)+common.HashLength*2+len(orderTypeString)+2)
	out = append(out, innerSig...)
	out = append(out, appDomainSep[:]...)
	out = append(out, contentsHash[:]...)
	out = append(out, orderTypeString...)
	out = binary.BigEndian.AppendUint16(out, uint16(len(orderTypeString)))
	return out, nil
}
