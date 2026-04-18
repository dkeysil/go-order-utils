package utils

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/polymarket/go-order-utils/pkg/model"
	"github.com/stretchr/testify/assert"
)

func TestGetVerifyingContractAddress(t *testing.T) {
	// V2 exchanges share the same addresses on mainnet and Amoy.
	exchangeV2 := common.HexToAddress("0xE111180000d2663C0091e4f400237545B87B996B")
	negRiskExchangeV2 := common.HexToAddress("0xe2222d279d744050d28e00520010520000310F59")

	// amoy
	contract, err := GetVerifyingContractAddress(big.NewInt(80002), model.CTFExchange)
	assert.NoError(t, err)
	assert.Equal(t, exchangeV2.Hex(), contract.Hex())

	contract, err = GetVerifyingContractAddress(big.NewInt(80002), model.NegRiskCTFExchange)
	assert.NoError(t, err)
	assert.Equal(t, negRiskExchangeV2.Hex(), contract.Hex())

	// polygon mainnet
	contract, err = GetVerifyingContractAddress(big.NewInt(137), model.CTFExchange)
	assert.NoError(t, err)
	assert.Equal(t, exchangeV2.Hex(), contract.Hex())

	contract, err = GetVerifyingContractAddress(big.NewInt(137), model.NegRiskCTFExchange)
	assert.NoError(t, err)
	assert.Equal(t, negRiskExchangeV2.Hex(), contract.Hex())

	// wrong network
	_, err = GetVerifyingContractAddress(big.NewInt(1), model.CTFExchange)
	assert.Error(t, err)

	_, err = GetVerifyingContractAddress(big.NewInt(1), model.NegRiskCTFExchange)
	assert.Error(t, err)
}
