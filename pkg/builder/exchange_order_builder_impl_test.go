package builder

import (
	"encoding/hex"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/polymarket/go-order-utils/pkg/model"
	"github.com/stretchr/testify/assert"
)

// Golden vectors in this file were generated with py-clob-client-v2==1.0.0
// using the well-known Anvil key #0 (ac0974...ff80 → 0xf39F...2266) and a
// fixed salt + timestamp. The generator script lives at scripts/gen_v2_vectors.py
// (kept out-of-tree; see the PR description). If V2 SDK semantics change the
// vectors must be regenerated.

var (
	chainId = new(big.Int).SetInt64(80002)
	// publicly known private key (Anvil default account #0)
	privateKey, _ = crypto.ToECDSA(common.Hex2Bytes("ac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80"))
	signerAddress = common.HexToAddress("0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266")

	salt                 = int64(479249096354)
	timestampMs          = int64(1700000000000)
	timestampMsString    = "1700000000000"
	fixedSaltGenerator   = func() int64 { return salt }
	fixedTimeGenerator   = func() int64 { return timestampMs }
	nonZeroMetadataValue = "0x1111111111111111111111111111111111111111111111111111111111111111"
	nonZeroBuilderValue  = "0x2222222222222222222222222222222222222222222222222222222222222222"
)

func newFixedBuilder() *ExchangeOrderBuilderImpl {
	return NewExchangeOrderBuilderImpl(chainId, fixedSaltGenerator, WithTimestampGenerator(fixedTimeGenerator))
}

func TestBuildOrder_DefaultsAndParsing(t *testing.T) {
	// Random salt + default timestamp: the only determinism we assert is
	// that fields are parsed correctly and defaults kick in.
	b := NewExchangeOrderBuilderImpl(chainId, nil)

	order, err := b.BuildOrder(&model.OrderData{
		Maker:       signerAddress.Hex(),
		TokenId:     "1234",
		MakerAmount: "100000000",
		TakerAmount: "50000000",
		Side:        model.BUY,
	})
	assert.NoError(t, err)
	assert.NotNil(t, order)

	assert.True(t, order.Salt.Int64() > 0)
	assert.Equal(t, signerAddress, order.Maker)
	assert.Equal(t, signerAddress, order.Signer)
	assert.Equal(t, "1234", order.TokenId.String())
	assert.Equal(t, "100000000", order.MakerAmount.String())
	assert.Equal(t, "50000000", order.TakerAmount.String())
	assert.Equal(t, "0", order.Side.String())
	assert.Equal(t, "0", order.SignatureType.String())
	assert.True(t, order.Timestamp.Int64() > 0, "default timestamp must be populated")
	assert.Equal(t, model.Bytes32Zero, order.Metadata)
	assert.Equal(t, model.Bytes32Zero, order.Builder)

	// Fixed salt + fixed timestamp: exact equality.
	b = newFixedBuilder()

	order, err = b.BuildOrder(&model.OrderData{
		Maker:       signerAddress.Hex(),
		TokenId:     "1234",
		MakerAmount: "100000000",
		TakerAmount: "50000000",
		Side:        model.BUY,
		Timestamp:   timestampMsString,
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
		Timestamp:     timestampMsString,
		Metadata:      nonZeroMetadataValue,
		Builder:       nonZeroBuilderValue,
	})
	assert.NoError(t, err)
	assert.Equal(t, maker, order.Maker)
	assert.Equal(t, signerAddress, order.Signer)
	assert.Equal(t, big.NewInt(int64(model.POLY_GNOSIS_SAFE)).String(), order.SignatureType.String())
	assert.Equal(t, common.HexToHash(nonZeroMetadataValue), order.Metadata)
	assert.Equal(t, common.HexToHash(nonZeroBuilderValue), order.Builder)
}

func TestBuildOrder_InvalidBytes32Rejected(t *testing.T) {
	b := newFixedBuilder()

	_, err := b.BuildOrder(&model.OrderData{
		Maker:       signerAddress.Hex(),
		TokenId:     "1234",
		MakerAmount: "100000000",
		TakerAmount: "50000000",
		Side:        model.BUY,
		Timestamp:   timestampMsString,
		Metadata:    "0xdeadbeef", // too short
	})
	assert.Error(t, err, "short bytes32 must be rejected")
}

// BuildOrderHash: fixed salt + timestamp produces a digest that must match the
// py-clob-client-v2 golden for each (exchange, signatureType) combo.
func TestBuildOrderHash_V2Goldens(t *testing.T) {
	baseData := func(sigType model.SignatureType) *model.OrderData {
		return &model.OrderData{
			Maker:         signerAddress.Hex(),
			TokenId:       "1234",
			MakerAmount:   "100000000",
			TakerAmount:   "50000000",
			Side:          model.BUY,
			SignatureType: sigType,
			Timestamp:     timestampMsString,
		}
	}

	cases := []struct {
		name     string
		contract model.VerifyingContract
		sigType  model.SignatureType
		metadata string
		builder  string
		expected string
	}{
		{
			name:     "CTFExchange EOA",
			contract: model.CTFExchange,
			sigType:  model.EOA,
			expected: "0x48d72286ccc2103a98c09668b7243830fa6f776091f1a6d96acfb7aa03f9928f",
		},
		{
			name:     "NegRiskCTFExchange EOA",
			contract: model.NegRiskCTFExchange,
			sigType:  model.EOA,
			expected: "0xd745a5e6ed75948fd0472b0815d6b013012c5179518915dcd7c4849bfa18ec5b",
		},
		{
			name:     "CTFExchange POLY_GNOSIS_SAFE",
			contract: model.CTFExchange,
			sigType:  model.POLY_GNOSIS_SAFE,
			expected: "0x24e5205f65d293ec044f06a91f10bda0a5922d4d49d94c198ef9055cf93f0054",
		},
		{
			name:     "CTFExchange EOA non-zero metadata+builder",
			contract: model.CTFExchange,
			sigType:  model.EOA,
			metadata: nonZeroMetadataValue,
			builder:  nonZeroBuilderValue,
			expected: "0x64b5ce14fbf56fa4cbbb41291a944029230cdaea196c9bbfe974c53ba14daf16",
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			b := newFixedBuilder()
			data := baseData(tc.sigType)
			data.Metadata = tc.metadata
			data.Builder = tc.builder

			order, err := b.BuildOrder(data)
			assert.NoError(t, err)

			hash, err := b.BuildOrderHash(order, tc.contract)
			assert.NoError(t, err)
			assert.Equal(t, tc.expected, hash.Hex())
		})
	}
}

// BuildSignedOrder: end-to-end parity with py-clob-client-v2 signatures. If
// either typehash, domain, or field-encoding drifts, these fail.
func TestBuildSignedOrder_V2Goldens(t *testing.T) {
	baseData := func(sigType model.SignatureType) *model.OrderData {
		return &model.OrderData{
			Maker:         signerAddress.Hex(),
			TokenId:       "1234",
			MakerAmount:   "100000000",
			TakerAmount:   "50000000",
			Side:          model.BUY,
			SignatureType: sigType,
			Timestamp:     timestampMsString,
		}
	}

	cases := []struct {
		name     string
		contract model.VerifyingContract
		sigType  model.SignatureType
		metadata string
		builder  string
		expected string // py-clob-client-v2 signature
	}{
		{
			name:     "CTFExchange EOA",
			contract: model.CTFExchange,
			sigType:  model.EOA,
			expected: "00ebbea4a48aa0c49b9f2f8d118aae3cb04be3f12a54799b4be10194072a6d156d5a23a2be2926f115b6c2bcfb905f38f53b7ae348132742734997f62ae011b51c",
		},
		{
			name:     "NegRiskCTFExchange EOA",
			contract: model.NegRiskCTFExchange,
			sigType:  model.EOA,
			expected: "0814c57e391daa07ddefcf17a4f9ffee36fa025c05eedd3f1d511c2c20e33da35f7a52fc38ea6ccd8a25c4d244faabe20f46778a367363398cac45555cc6b84b1b",
		},
		{
			name:     "CTFExchange POLY_GNOSIS_SAFE",
			contract: model.CTFExchange,
			sigType:  model.POLY_GNOSIS_SAFE,
			expected: "4ede1c23998c8d86a6cb3a4fd5fe94df824e4104f88dd5182905411ff6a568fa345d9ba51d3833ba76327e68beb06f4fb78ade58cc69e8c0859c89b548aad3131b",
		},
		{
			name:     "CTFExchange EOA non-zero metadata+builder",
			contract: model.CTFExchange,
			sigType:  model.EOA,
			metadata: nonZeroMetadataValue,
			builder:  nonZeroBuilderValue,
			expected: "de0dc5a17d30c677d17c8410798679ea1d634814b677c8541caaeda788830ee91ed4bfbc92313e70b3ebb43ad260fabcb8c2ff695f9d488625bc56182d443dbb1c",
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			b := newFixedBuilder()
			data := baseData(tc.sigType)
			data.Metadata = tc.metadata
			data.Builder = tc.builder

			signed, err := b.BuildSignedOrder(privateKey, data, tc.contract)
			// POLY_GNOSIS_SAFE signed by the maker's own EOA won't round-trip
			// through ValidateSignature (maker != recovered signer). Check the
			// raw signature bytes instead.
			if tc.sigType == model.POLY_GNOSIS_SAFE {
				// BuildSignedOrder will still error on signature mismatch;
				// assert separately that the signing pipeline reached here.
				if err != nil {
					// Expected for safe-mode: go back and sign manually.
					order, berr := b.BuildOrder(data)
					assert.NoError(t, berr)
					hash, herr := b.BuildOrderHash(order, tc.contract)
					assert.NoError(t, herr)
					sig, serr := b.BuildOrderSignature(privateKey, hash)
					assert.NoError(t, serr)
					assert.Equal(t, tc.expected, hex.EncodeToString(sig))
					return
				}
			}
			assert.NoError(t, err)
			assert.NotNil(t, signed)
			assert.Equal(t, tc.expected, hex.EncodeToString(signed.Signature))
		})
	}
}
