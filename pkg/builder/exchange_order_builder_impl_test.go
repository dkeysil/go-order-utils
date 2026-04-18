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

// Golden vectors are produced by scripts/gen_v2_vectors.py (py-clob-client-v2).
var (
	chainID       = new(big.Int).SetInt64(80002)
	privateKey, _ = crypto.ToECDSA(common.Hex2Bytes("ac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80"))
	signerAddress = common.HexToAddress("0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266")

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
	cases := []struct {
		name     string
		contract model.VerifyingContract
		sigType  model.SignatureType
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
			// POLY_1271 Signer is a smart contract whose isValidSignature()
			// authorizes the provided privateKey's EOA. The signature bytes
			// still come from the EOA — the local ecrecover check is skipped.
			name:     "CTFExchange POLY_1271",
			contract: model.CTFExchange,
			sigType:  model.POLY_1271,
			hash:     "0x476135297e4fc5246f359be914c1f083dd27b71081bdce54d453f6317895139f",
			sig:      "a54cdd9888dcbb3ea9715158ac08b27042be945a28cac645f28b92e23f9b4f9a16ac91aec9d85a5c9ecda794f0f3728cf42acc8d0f967a2e6671c1d1df2317e51b",
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
			data := &model.OrderData{
				Maker:         signerAddress.Hex(),
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

			hash, err := b.BuildOrderHash(order, tc.contract)
			assert.NoError(t, err)
			assert.Equal(t, tc.hash, hash.Hex())

			signed, err := b.BuildSignedOrder(privateKey, data, tc.contract)
			assert.NoError(t, err)
			assert.NotNil(t, signed)
			assert.Equal(t, tc.sig, hex.EncodeToString(signed.Signature))
		})
	}
}

// TestBuildSignedOrder_POLY1271SkipsEcrecover exercises the contract-wallet
// path: Signer is a smart-contract address that the EOA privateKey is
// authorized to sign for. The local ecrecover check cannot match the contract
// address, so the method must return the signature without erroring. This
// case fails under a V1-style unconditional ValidateSignature.
func TestBuildSignedOrder_POLY1271SkipsEcrecover(t *testing.T) {
	b := newFixedBuilder()
	safeLikeContract := common.HexToAddress("0x000000000000000000000000000000000000dEaD")

	signed, err := b.BuildSignedOrder(privateKey, &model.OrderData{
		Maker:         safeLikeContract.Hex(),
		Signer:        safeLikeContract.Hex(),
		TokenId:       "1234",
		MakerAmount:   "100000000",
		TakerAmount:   "50000000",
		Side:          model.BUY,
		SignatureType: model.POLY_1271,
		Timestamp:     timestampMs,
	}, model.CTFExchange)
	assert.NoError(t, err)
	assert.NotNil(t, signed)
	assert.Len(t, signed.Signature, 65)
	assert.Equal(t, safeLikeContract, signed.Signer)
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
