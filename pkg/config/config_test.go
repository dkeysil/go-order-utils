package config

import (
	"bytes"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
)

func TestGetContracts(t *testing.T) {
	var (
		amoy = &Contracts{
			Exchange:        common.HexToAddress("0xE111180000d2663C0091e4f400237545B87B996B"),
			NegRiskExchange: common.HexToAddress("0xe2222d279d744050d28e00520010520000310F59"),
			NegRiskAdapter:  common.HexToAddress("0xd91E80cF2E7be2e162c6513ceD06f1dD0dA35296"),
			Collateral:      common.HexToAddress("0xC011a7E12a19f7B1f670d46F03B03f3342E82DFB"),
			Conditional:     common.HexToAddress("0x69308FB512518e39F9b16112fA8d994F4e2Bf8bB"),
		}

		matic = &Contracts{
			Exchange:        common.HexToAddress("0xE111180000d2663C0091e4f400237545B87B996B"),
			NegRiskExchange: common.HexToAddress("0xe2222d279d744050d28e00520010520000310F59"),
			NegRiskAdapter:  common.HexToAddress("0xd91E80cF2E7be2e162c6513ceD06f1dD0dA35296"),
			Collateral:      common.HexToAddress("0xC011a7E12a19f7B1f670d46F03B03f3342E82DFB"),
			Conditional:     common.HexToAddress("0x4D97DCd97eC945f40cF65F87097ACe5EA0476045"),
		}
	)

	c, err := GetContracts(80002)
	assert.NotNil(t, c)
	assert.Nil(t, err)
	assert.True(t, bytes.Equal(c.Exchange[:], amoy.Exchange[:]))
	assert.True(t, bytes.Equal(c.NegRiskExchange[:], amoy.NegRiskExchange[:]))
	assert.True(t, bytes.Equal(c.NegRiskAdapter[:], amoy.NegRiskAdapter[:]))
	assert.True(t, bytes.Equal(c.Collateral[:], amoy.Collateral[:]))
	assert.True(t, bytes.Equal(c.Conditional[:], amoy.Conditional[:]))

	c, err = GetContracts(137)
	assert.NotNil(t, c)
	assert.Nil(t, err)
	assert.True(t, bytes.Equal(c.Exchange[:], matic.Exchange[:]))
	assert.True(t, bytes.Equal(c.NegRiskExchange[:], matic.NegRiskExchange[:]))
	assert.True(t, bytes.Equal(c.NegRiskAdapter[:], matic.NegRiskAdapter[:]))
	assert.True(t, bytes.Equal(c.Collateral[:], matic.Collateral[:]))
	assert.True(t, bytes.Equal(c.Conditional[:], matic.Conditional[:]))

	c, err = GetContracts(100000)
	assert.Nil(t, c)
	assert.NotNil(t, err)
}
