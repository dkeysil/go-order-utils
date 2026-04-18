package builder

import (
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/polymarket/go-order-utils/pkg/eip712"
)

// EIP-712 domain for CTF Exchange V2. Domain version is "2"; name stays
// "Polymarket CTF Exchange" (same string as V1).
var (
	_PROTOCOL_NAME    = crypto.Keccak256Hash([]byte("Polymarket CTF Exchange"))
	_PROTOCOL_VERSION = crypto.Keccak256Hash([]byte("2"))
)

// _ORDER_STRUCTURE is the ABI tuple used to encode the V2 Order struct for
// EIP-712 hashing. Order and count must exactly match _ORDER_STRUCTURE_HASH
// below and the Solidity Order struct in ctf-exchange-v2.
var (
	_ORDER_STRUCTURE = []abi.Type{
		eip712.Bytes32, // typehash
		eip712.Uint256, // salt
		eip712.Address, // maker
		eip712.Address, // signer
		eip712.Uint256, // tokenId
		eip712.Uint256, // makerAmount
		eip712.Uint256, // takerAmount
		eip712.Uint8,   // side
		eip712.Uint8,   // signatureType
		eip712.Uint256, // timestamp
		eip712.Bytes32, // metadata
		eip712.Bytes32, // builder
	}
)

// _ORDER_STRUCTURE_HASH must equal 0xbb86318a2138f5fa8ae32fbe8e659f8fcf13cc6ae4014a707893055433818589
// (the ORDER_TYPEHASH constant in ctf-exchange-v2 Structs.sol). See the unit
// test in constants_test.go.
var (
	_ORDER_STRUCTURE_HASH = crypto.Keccak256Hash(
		[]byte("Order(uint256 salt,address maker,address signer,uint256 tokenId,uint256 makerAmount,uint256 takerAmount,uint8 side,uint8 signatureType,uint256 timestamp,bytes32 metadata,bytes32 builder)"),
	)
)
