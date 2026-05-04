package builder

import (
	"encoding/binary"
	"encoding/hex"
	"math/big"
	"testing"

	"github.com/dkeysil/go-order-utils/pkg/depositwallet"
	"github.com/dkeysil/go-order-utils/pkg/eip712"
	"github.com/dkeysil/go-order-utils/pkg/model"
	"github.com/dkeysil/go-order-utils/pkg/signer"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/assert"
)

// EOA / Proxy / Safe goldens come from py-clob-client-v2==1.0.0 via
// scripts/gen_v2_vectors.py. POLY_1271 (ERC-7739) goldens come from
// @polymarket/clob-client-v2 at commit f7307a8.
var (
	chainID       = new(big.Int).SetInt64(80002)
	privateKey, _ = crypto.ToECDSA(common.Hex2Bytes("ac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80"))
	signerAddress = common.HexToAddress("0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266")

	// depositWalletAddr is a stand-in for a Polymarket deposit wallet contract
	// address used by POLY_1271 goldens. For these orders Maker == Signer ==
	// deposit wallet.
	depositWalletAddr = common.HexToAddress("0x000000000000000000000000000000000000dEaD")

	salt                 = int64(479249096354)
	timestampMs          = int64(1700000000000)
	nonZeroMetadataValue = common.HexToHash("0x1111111111111111111111111111111111111111111111111111111111111111")
	nonZeroBuilderValue  = common.HexToHash("0x2222222222222222222222222222222222222222222222222222222222222222")
)

func newFixedBuilder() *ExchangeOrderBuilderImpl {
	return NewExchangeOrderBuilderImpl(chainID, func() int64 { return salt }, func() int64 { return timestampMs })
}

func TestBuildOrder_DefaultsAndParsing(t *testing.T) {
	// Random salt + default timestamp: only assert fields are parsed and
	// defaults kick in.
	b := NewExchangeOrderBuilderImpl(chainID, nil, nil)

	order, err := b.BuildOrder(&model.OrderData{
		Maker:       signerAddress.Hex(),
		TokenId:     "1234",
		MakerAmount: "100000000",
		TakerAmount: "50000000",
		Side:        model.BUY,
	})
	assert.NoError(t, err)
	assert.NotNil(t, order)

	assert.Positive(t, order.Salt.Int64())
	assert.Equal(t, signerAddress, order.Maker)
	assert.Equal(t, signerAddress, order.Signer)
	assert.Equal(t, "1234", order.TokenId.String())
	assert.Equal(t, "100000000", order.MakerAmount.String())
	assert.Equal(t, "50000000", order.TakerAmount.String())
	assert.Equal(t, uint8(0), order.Side)
	assert.Equal(t, uint8(0), order.SignatureType)
	assert.Positive(t, order.Timestamp.Int64(), "default timestamp must be populated")
	assert.Equal(t, common.Hash{}, order.Metadata)
	assert.Equal(t, common.Hash{}, order.Builder)

	// Fixed salt + fixed timestamp: exact equality.
	b = newFixedBuilder()

	order, err = b.BuildOrder(&model.OrderData{
		Maker:       signerAddress.Hex(),
		TokenId:     "1234",
		MakerAmount: "100000000",
		TakerAmount: "50000000",
		Side:        model.BUY,
		Timestamp:   timestampMs,
	})
	assert.NoError(t, err)
	assert.Equal(t, salt, order.Salt.Int64())
	assert.Equal(t, timestampMs, order.Timestamp.Int64())
}

func TestBuildOrder_SignerOverrideAndMetadata(t *testing.T) {
	maker := common.HexToAddress("0xaFB8270A801862270FebB3763505b136491e557b")

	b := newFixedBuilder()

	order, err := b.BuildOrder(&model.OrderData{
		Maker:         maker.Hex(),
		Signer:        signerAddress.Hex(),
		TokenId:       "100",
		MakerAmount:   "50000000",
		TakerAmount:   "100000000",
		Side:          model.BUY,
		SignatureType: model.POLY_GNOSIS_SAFE,
		Timestamp:     timestampMs,
		Metadata:      nonZeroMetadataValue,
		Builder:       nonZeroBuilderValue,
	})
	assert.NoError(t, err)
	assert.Equal(t, maker, order.Maker)
	assert.Equal(t, signerAddress, order.Signer)
	assert.Equal(t, uint8(model.POLY_GNOSIS_SAFE), order.SignatureType)
	assert.Equal(t, nonZeroMetadataValue, order.Metadata)
	assert.Equal(t, nonZeroBuilderValue, order.Builder)
}

// TestV2Goldens asserts the Go builder reproduces py-clob-client-v2 digests
// and signatures byte-for-byte. If any of these drifts, on-chain orders will
// be rejected.
func TestV2Goldens(t *testing.T) {
	// hash is the EIP-712 order digest under the verifying-contract domain
	// (what BuildOrderHash returns). For POLY_1271 the operator validates
	// the wrapped signature, not this digest, so the field is left empty
	// and skipped.
	cases := []struct {
		name     string
		contract model.VerifyingContract
		sigType  model.SignatureType
		maker    string // empty → signerAddress (EOA path); set for POLY_1271
		signer   string // optional override
		metadata common.Hash
		builder  common.Hash
		hash     string
		sig      string
	}{
		{
			name:     "CTFExchange EOA",
			contract: model.CTFExchange,
			sigType:  model.EOA,
			hash:     "0x48d72286ccc2103a98c09668b7243830fa6f776091f1a6d96acfb7aa03f9928f",
			sig:      "00ebbea4a48aa0c49b9f2f8d118aae3cb04be3f12a54799b4be10194072a6d156d5a23a2be2926f115b6c2bcfb905f38f53b7ae348132742734997f62ae011b51c",
		},
		{
			name:     "NegRiskCTFExchange EOA",
			contract: model.NegRiskCTFExchange,
			sigType:  model.EOA,
			hash:     "0xd745a5e6ed75948fd0472b0815d6b013012c5179518915dcd7c4849bfa18ec5b",
			sig:      "0814c57e391daa07ddefcf17a4f9ffee36fa025c05eedd3f1d511c2c20e33da35f7a52fc38ea6ccd8a25c4d244faabe20f46778a367363398cac45555cc6b84b1b",
		},
		{
			name:     "CTFExchange POLY_GNOSIS_SAFE",
			contract: model.CTFExchange,
			sigType:  model.POLY_GNOSIS_SAFE,
			hash:     "0x24e5205f65d293ec044f06a91f10bda0a5922d4d49d94c198ef9055cf93f0054",
			sig:      "4ede1c23998c8d86a6cb3a4fd5fe94df824e4104f88dd5182905411ff6a568fa345d9ba51d3833ba76327e68beb06f4fb78ade58cc69e8c0859c89b548aad3131b",
		},
		{
			// POLY_1271 (Polymarket deposit wallet) signature is the
			// ERC-7739 composite. Maker and Signer are the wallet contract.
			// The EOA (privateKey) is what the wallet's isValidSignature
			// verifies internally. Golden generated by clob-client-v2@f7307a8.
			name:     "CTFExchange POLY_1271",
			contract: model.CTFExchange,
			sigType:  model.POLY_1271,
			maker:    depositWalletAddr.Hex(),
			signer:   depositWalletAddr.Hex(),
			sig:      "847acfdbf26728a4a084ba529cbf5c082f0b91cffd52e964785c08eb268ee5b120faae1e277da01f34f49e4593a036411b5f21ff61848e2356ec3571a9b0aa5b1ba440cbd865bc0c6243d7a8df9a8bf48a8827b0a4abbb61c30e96d305423af148a048f7941186908521f0b20e1527e197aead2546f4be1555880d18420deb2f994f726465722875696e743235362073616c742c61646472657373206d616b65722c61646472657373207369676e65722c75696e7432353620746f6b656e49642c75696e74323536206d616b6572416d6f756e742c75696e743235362074616b6572416d6f756e742c75696e743820736964652c75696e7438207369676e6174757265547970652c75696e743235362074696d657374616d702c62797465733332206d657461646174612c62797465733332206275696c6465722900ba",
		},
		{
			name:     "NegRiskCTFExchange POLY_1271",
			contract: model.NegRiskCTFExchange,
			sigType:  model.POLY_1271,
			maker:    depositWalletAddr.Hex(),
			signer:   depositWalletAddr.Hex(),
			sig:      "7e96a2d5890a3dd4cba3a28143eca6e9d7c6e754bb07bf96f281d8e31cbaf7e57b0594461dcb44d8ac433907ab04c0962a1816fc62d7675e31f5d60041d04af71b1468e39841f0a0e05d3762b235cb7acecabf8cd7b6d3a672965a5d415bf9378ea048f7941186908521f0b20e1527e197aead2546f4be1555880d18420deb2f994f726465722875696e743235362073616c742c61646472657373206d616b65722c61646472657373207369676e65722c75696e7432353620746f6b656e49642c75696e74323536206d616b6572416d6f756e742c75696e743235362074616b6572416d6f756e742c75696e743820736964652c75696e7438207369676e6174757265547970652c75696e743235362074696d657374616d702c62797465733332206d657461646174612c62797465733332206275696c6465722900ba",
		},
		{
			name:     "CTFExchange EOA non-zero metadata+builder",
			contract: model.CTFExchange,
			sigType:  model.EOA,
			metadata: nonZeroMetadataValue,
			builder:  nonZeroBuilderValue,
			hash:     "0x64b5ce14fbf56fa4cbbb41291a944029230cdaea196c9bbfe974c53ba14daf16",
			sig:      "de0dc5a17d30c677d17c8410798679ea1d634814b677c8541caaeda788830ee91ed4bfbc92313e70b3ebb43ad260fabcb8c2ff695f9d488625bc56182d443dbb1c",
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			b := newFixedBuilder()
			maker := tc.maker
			if maker == "" {
				maker = signerAddress.Hex()
			}
			data := &model.OrderData{
				Maker:         maker,
				Signer:        tc.signer,
				TokenId:       "1234",
				MakerAmount:   "100000000",
				TakerAmount:   "50000000",
				Side:          model.BUY,
				SignatureType: tc.sigType,
				Timestamp:     timestampMs,
				Metadata:      tc.metadata,
				Builder:       tc.builder,
			}

			order, err := b.BuildOrder(data)
			assert.NoError(t, err)

			if tc.hash != "" {
				hash, err := b.BuildOrderHash(order, tc.contract)
				assert.NoError(t, err)
				assert.Equal(t, tc.hash, hash.Hex())
			}

			signed, err := b.BuildSignedOrder(privateKey, data, tc.contract)
			assert.NoError(t, err)
			assert.NotNil(t, signed)
			assert.Equal(t, tc.sig, hex.EncodeToString(signed.Signature))
		})
	}
}

// TestBuildSignedOrder_POLY1271WireFormat decomposes the ERC-7739 composite
// signature and asserts each section matches the spec the deposit wallet's
// isValidSignature reads from:
//
//	innerSig(65) || appDomainSep(32) || contentsHash(32) || orderTypeString || uint16_BE(len)
func TestBuildSignedOrder_POLY1271WireFormat(t *testing.T) {
	b := newFixedBuilder()

	signed, err := b.BuildSignedOrder(privateKey, &model.OrderData{
		Maker:         depositWalletAddr.Hex(),
		Signer:        depositWalletAddr.Hex(),
		TokenId:       "1234",
		MakerAmount:   "100000000",
		TakerAmount:   "50000000",
		Side:          model.BUY,
		SignatureType: model.POLY_1271,
		Timestamp:     timestampMs,
	}, model.CTFExchange)
	assert.NoError(t, err)
	assert.NotNil(t, signed)
	assert.Equal(t, depositWalletAddr, signed.Signer)

	const (
		innerSigLen     = 65
		domainSepLen    = common.HashLength
		contentsHashLen = common.HashLength
		lenSuffixLen    = 2
	)
	sig := signed.Signature
	typeStrLen := len(orderTypeString)
	assert.Len(t, sig, innerSigLen+domainSepLen+contentsHashLen+typeStrLen+lenSuffixLen, "composite signature length mismatch")

	domainSepEnd := innerSigLen + domainSepLen
	contentsHashEnd := domainSepEnd + contentsHashLen
	typeBytesEnd := contentsHashEnd + typeStrLen

	innerSig := sig[:innerSigLen]
	appDomainSep := sig[innerSigLen:domainSepEnd]
	contentsHash := sig[domainSepEnd:contentsHashEnd]
	typeBytes := sig[contentsHashEnd:typeBytesEnd]
	lenSuffix := sig[typeBytesEnd:]

	expectedDomainSep, err := eip712.BuildEIP712DomainSeparator(
		protocolName, protocolVersion, chainID,
		common.HexToAddress("0xE111180000d2663C0091e4f400237545B87B996B"), // CTFExchange amoy
	)
	assert.NoError(t, err)
	assert.Equal(t, expectedDomainSep[:], appDomainSep, "appDomainSep section")

	expectedContentsHash, err := hashOrderStruct(&signed.Order)
	assert.NoError(t, err)
	assert.Equal(t, expectedContentsHash[:], contentsHash, "contentsHash section")

	assert.Equal(t, orderTypeString, typeBytes, "orderTypeString section")
	assert.Equal(t, binary.BigEndian.AppendUint16(nil, uint16(typeStrLen)), lenSuffix, "uint16 BE length suffix")

	// Recompute the digest the EOA signed over and verify innerSig recovers
	// to the EOA derived from privateKey.
	contentsHashHash := common.BytesToHash(contentsHash)
	outerStruct, err := eip712.HashStruct(typedDataSignArgs, []any{
		typedDataSignTypeHash,
		contentsHashHash,
		depositwallet.DomainName,
		depositwallet.DomainVersion,
		chainID,
		depositWalletAddr,
		common.Hash{},
	})
	assert.NoError(t, err)
	digest := eip712.HashFromStructHash(common.BytesToHash(appDomainSep), outerStruct)

	ok, err := signer.ValidateSignature(crypto.PubkeyToAddress(privateKey.PublicKey), digest, innerSig)
	assert.NoError(t, err)
	assert.True(t, ok, "innerSig must ecrecover to the EOA")
}

func TestBuildOrderHash_UnsupportedChain(t *testing.T) {
	b := NewExchangeOrderBuilderImpl(big.NewInt(1), nil, nil)
	order, err := b.BuildOrder(&model.OrderData{
		Maker:       signerAddress.Hex(),
		TokenId:     "1",
		MakerAmount: "1",
		TakerAmount: "1",
		Side:        model.BUY,
		Timestamp:   timestampMs,
	})
	assert.NoError(t, err)

	_, err = b.BuildOrderHash(order, model.CTFExchange)
	assert.Error(t, err)
}
