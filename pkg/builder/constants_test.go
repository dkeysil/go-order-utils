package builder

import (
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
)

// V2 ORDER_TYPEHASH pinned by the Solidity contract
// (ctf-exchange-v2/src/exchange/libraries/Structs.sol). If this changes, every
// signature this library produces is invalid — hence the pinned hex.
func TestOrderStructureHash_V2(t *testing.T) {
	expected := common.HexToHash("0xbb86318a2138f5fa8ae32fbe8e659f8fcf13cc6ae4014a707893055433818589")
	assert.Equal(t, expected.Hex(), _ORDER_STRUCTURE_HASH.Hex())
}
