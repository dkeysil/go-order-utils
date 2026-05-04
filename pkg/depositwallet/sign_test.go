package depositwallet

import (
	"encoding/hex"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/assert"
)

// Goldens cross-checked against @polymarket/clob-client-v2's signTypedData
// over the Batch type. Inputs match the fixtures in pkg/builder so they can
// be cross-referenced.
var (
	testPrivateKey, _ = crypto.ToECDSA(common.Hex2Bytes("ac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80"))
	testWallet        = common.HexToAddress("0x000000000000000000000000000000000000dEaD")
	testChainID       = big.NewInt(80002)

	usdcAddr = common.HexToAddress("0x2791Bca1f2de4661ED88A30C99A7a9449Aa84174")
	ctfAddr  = common.HexToAddress("0x4D97DCd97eC945f40cF65F87097ACe5EA0476045")
)

// approveData is `IERC20.approve(exchange, type(uint256).max)` calldata.
func approveData() []byte {
	return common.FromHex("0x095ea7b3000000000000000000000000e111180000d2663c0091e4f400237545b87b996bffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff")
}

// setApprovalForAllData is `IERC1155.setApprovalForAll(exchange, true)` calldata.
func setApprovalForAllData() []byte {
	return common.FromHex("0xa22cb465000000000000000000000000e111180000d2663c0091e4f400237545b87b996b0000000000000000000000000000000000000000000000000000000000000001")
}

func TestSignBatch_USDCAndCTFApprovals(t *testing.T) {
	batch := Batch{
		Wallet:   testWallet,
		Nonce:    big.NewInt(0),
		Deadline: big.NewInt(0),
		Calls: []Call{
			{Target: usdcAddr, Value: big.NewInt(0), Data: approveData()},
			{Target: ctfAddr, Value: big.NewInt(0), Data: setApprovalForAllData()},
		},
	}

	signed, err := SignBatch(testPrivateKey, testChainID, batch)
	assert.NoError(t, err)
	assert.NotNil(t, signed)

	wantSig := "1a1b8ba12375acf5bc052cf00695566dd6f9d267e5c3a31cffeac5c6fe7433a53f332bf4e97acaea1990cdf4bb5e1d02d8dca224c1df9a337f2ce5ecb85656f71c"
	assert.Equal(t, wantSig, hex.EncodeToString(signed.Signature))
}

func TestSignBatch_Empty(t *testing.T) {
	batch := Batch{
		Wallet:   testWallet,
		Nonce:    big.NewInt(0),
		Deadline: big.NewInt(0),
		Calls:    nil,
	}
	signed, err := SignBatch(testPrivateKey, testChainID, batch)
	assert.NoError(t, err)
	wantSig := "9358a00f9f5174a12f890d320f6be587b2ef4eb68158bae17e0542fc1ba548db0b2d5bff72817dba9ffe74f174615a102d5475d8b7dd7e1c5b054d403bfd7b531c"
	assert.Equal(t, wantSig, hex.EncodeToString(signed.Signature))
}

func TestSignBatch_OneCallEmptyData(t *testing.T) {
	batch := Batch{
		Wallet:   testWallet,
		Nonce:    big.NewInt(1),
		Deadline: big.NewInt(9999999999),
		Calls: []Call{
			{Target: usdcAddr, Value: big.NewInt(0), Data: nil},
		},
	}
	signed, err := SignBatch(testPrivateKey, testChainID, batch)
	assert.NoError(t, err)
	wantSig := "73b63e82df4148a2615a8175cfe80591439a3f7fa1d6168f8d169f73071b0e93309c155c3993b51ef1c85493bf2d770eb391d4b61b54c1a6a8b5291a8af1e50a1c"
	assert.Equal(t, wantSig, hex.EncodeToString(signed.Signature))
}

// HashCalls of an empty slice must be keccak256("") — matches Solidity's
// abi.encode of an empty dynamic array.
func TestHashCalls_Empty(t *testing.T) {
	got, err := HashCalls(nil)
	assert.NoError(t, err)
	assert.Equal(t, crypto.Keccak256Hash(nil), got)
}
