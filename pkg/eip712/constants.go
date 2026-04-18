// Package eip712 provides primitives for computing EIP-712 typed-data digests
// compatible with the Polymarket protocol.
package eip712

import (
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/crypto"
)

var (
	eip712Domain = []abi.Type{
		Bytes32, // typehash
		Bytes32, // name
		Bytes32, // version
		Uint256, // chainId
		Address, // verifyingContract
	}

	eip712DomainHash = crypto.Keccak256Hash(
		[]byte("EIP712Domain(string name,string version,uint256 chainId,address verifyingContract)"),
	)

	eip712DomainNoVerifyingContract = []abi.Type{
		Bytes32, // typehash
		Bytes32, // name
		Bytes32, // version
		Uint256, // chainId
	}

	eip712DomainHashNoVerifyingContract = crypto.Keccak256Hash(
		[]byte("EIP712Domain(string name,string version,uint256 chainId)"),
	)
)
