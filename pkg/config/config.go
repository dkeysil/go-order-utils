package config

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"
)

// Contracts bundles the Polymarket protocol addresses for a single chain.
//
// Exchange and NegRiskExchange point at the V2 CTF Exchange contracts (both
// addresses are chain-agnostic — Polygon and Amoy share the same deployed
// addresses per the V2 SDK config).
type Contracts struct {
	Exchange        common.Address
	NegRiskExchange common.Address
	NegRiskAdapter  common.Address
	Collateral      common.Address
	Conditional     common.Address
}

var (
	// V2 exchange addresses are identical on mainnet (137) and Amoy (80002),
	// mirroring @polymarket/clob-client-v2 config.
	_EXCHANGE_V2          = common.HexToAddress("0xE111180000d2663C0091e4f400237545B87B996B")
	_NEG_RISK_EXCHANGE_V2 = common.HexToAddress("0xe2222d279d744050d28e00520010520000310F59")
	_NEG_RISK_ADAPTER     = common.HexToAddress("0xd91E80cF2E7be2e162c6513ceD06f1dD0dA35296")

	// Collateral is pUSD in V2, replacing USDC.e. Value mirrors
	// @polymarket/clob-client-v2 config (same on both chains).
	_COLLATERAL = common.HexToAddress("0xC011a7E12a19f7B1f670d46F03B03f3342E82DFB")

	_AMOY_CONTRACTS = &Contracts{
		Exchange:        _EXCHANGE_V2,
		NegRiskExchange: _NEG_RISK_EXCHANGE_V2,
		NegRiskAdapter:  _NEG_RISK_ADAPTER,
		Collateral:      _COLLATERAL,
		Conditional:     common.HexToAddress("0x69308FB512518e39F9b16112fA8d994F4e2Bf8bB"),
	}

	_MATIC_CONTRACTS = &Contracts{
		Exchange:        _EXCHANGE_V2,
		NegRiskExchange: _NEG_RISK_EXCHANGE_V2,
		NegRiskAdapter:  _NEG_RISK_ADAPTER,
		Collateral:      _COLLATERAL,
		Conditional:     common.HexToAddress("0x4D97DCd97eC945f40cF65F87097ACe5EA0476045"),
	}
)

func GetContracts(chainId int64) (*Contracts, error) {
	switch chainId {
	case 137:
		return _MATIC_CONTRACTS, nil
	case 80002:
		return _AMOY_CONTRACTS, nil
	default:
		return nil, fmt.Errorf("invalid chain id")
	}
}
