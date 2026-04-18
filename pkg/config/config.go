// Package config pins the deployed Polymarket protocol addresses keyed by
// chain id.
package config

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"
)

// Contracts bundles the Polymarket V2 protocol addresses for a single chain.
type Contracts struct {
	Exchange        common.Address
	NegRiskExchange common.Address
	NegRiskAdapter  common.Address
	Collateral      common.Address
	Conditional     common.Address
}

var (
	exchangeV2        = common.HexToAddress("0xE111180000d2663C0091e4f400237545B87B996B")
	negRiskExchangeV2 = common.HexToAddress("0xe2222d279d744050d28e00520010520000310F59")
	negRiskAdapter    = common.HexToAddress("0xd91E80cF2E7be2e162c6513ceD06f1dD0dA35296")
	collateral        = common.HexToAddress("0xC011a7E12a19f7B1f670d46F03B03f3342E82DFB")

	amoyContracts = &Contracts{
		Exchange:        exchangeV2,
		NegRiskExchange: negRiskExchangeV2,
		NegRiskAdapter:  negRiskAdapter,
		Collateral:      collateral,
		Conditional:     common.HexToAddress("0x69308FB512518e39F9b16112fA8d994F4e2Bf8bB"),
	}

	maticContracts = &Contracts{
		Exchange:        exchangeV2,
		NegRiskExchange: negRiskExchangeV2,
		NegRiskAdapter:  negRiskAdapter,
		Collateral:      collateral,
		Conditional:     common.HexToAddress("0x4D97DCd97eC945f40cF65F87097ACe5EA0476045"),
	}
)

// GetContracts returns the Polymarket V2 contract set for the given chain id.
// Supported chains: 137 (Polygon mainnet), 80002 (Amoy).
func GetContracts(chainID int64) (*Contracts, error) {
	switch chainID {
	case 137:
		return maticContracts, nil
	case 80002:
		return amoyContracts, nil
	default:
		return nil, fmt.Errorf("invalid chain id")
	}
}
