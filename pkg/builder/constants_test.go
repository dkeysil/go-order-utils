package builder

import (
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
)

// TestOrderStructureHashV2 pins the V2 ORDER_TYPEHASH to the Solidity constant
// in ctf-exchange-v2 (Structs.sol). If this drifts, every signature this
// library produces is invalid.
func TestOrderStructureHashV2(t *testing.T) {
	expected := common.HexToHash("0xbb86318a2138f5fa8ae32fbe8e659f8fcf13cc6ae4014a707893055433818589")
	assert.Equal(t, expected.Hex(), orderStructureHash.Hex())
}
