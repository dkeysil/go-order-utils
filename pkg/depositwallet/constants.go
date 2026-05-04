package depositwallet

import (
	"github.com/dkeysil/go-order-utils/pkg/eip712"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/crypto"
)

// DomainName is the keccak256 of the DepositWallet EIP-712 domain name.
// Exported so pkg/builder's ERC-7739 wrap for POLY_1271 orders can embed the
// same wallet domain without recomputing.
var DomainName = crypto.Keccak256Hash([]byte("DepositWallet"))

// DomainVersion is the keccak256 of the DepositWallet EIP-712 domain version.
var DomainVersion = crypto.Keccak256Hash([]byte("1"))

var (
	// EIP-712 nested-type encoding: outer struct followed by referenced
	// struct types in alphabetical order. Only Call is referenced.
	batchTypeHash = crypto.Keccak256Hash([]byte(
		"Batch(address wallet,uint256 nonce,uint256 deadline,Call[] calls)" +
			"Call(address target,uint256 value,bytes data)",
	))

	callTypeHash = crypto.Keccak256Hash([]byte(
		"Call(address target,uint256 value,bytes data)",
	))

	batchArgs = []abi.Type{
		eip712.Bytes32, // typehash
		eip712.Address, // wallet
		eip712.Uint256, // nonce
		eip712.Uint256, // deadline
		eip712.Bytes32, // keccak256(concat(callHashes))
	}

	callArgs = []abi.Type{
		eip712.Bytes32, // typehash
		eip712.Address, // target
		eip712.Uint256, // value
		eip712.Bytes32, // keccak256(data)
	}
)
