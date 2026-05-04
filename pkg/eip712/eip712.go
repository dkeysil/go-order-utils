package eip712

import (
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

// BuildEIP712DomainSeparator returns keccak256(encodeEIP712Domain{name,
// version, chainId, verifyingContract}) per EIP-712.
func BuildEIP712DomainSeparator(name, version common.Hash, chainID *big.Int, address common.Address) (common.Hash, error) {
	values := []any{
		eip712DomainHash,
		name,
		version,
		chainID,
		address,
	}

	encodedDomainSeparator, err := Encode(eip712Domain, values)
	if err != nil {
		return common.Hash{}, err
	}

	return crypto.Keccak256Hash(encodedDomainSeparator), nil
}

// BuildEIP712DomainSeparatorNoContract is the EIP-712 domain separator without
// a verifyingContract field; used by Polymarket's ClobAuth domain.
func BuildEIP712DomainSeparatorNoContract(name, version common.Hash, chainID *big.Int) (common.Hash, error) {
	values := []any{
		eip712DomainHashNoVerifyingContract,
		name,
		version,
		chainID,
	}

	encodedDomainSeparator, err := Encode(eip712DomainNoVerifyingContract, values)
	if err != nil {
		return common.Hash{}, err
	}

	return crypto.Keccak256Hash(encodedDomainSeparator), nil
}

// HashStruct returns keccak256(abiEncode(args, values)) — the inner
// hashStruct(s) of an EIP-712 typed message. The first entry of args/values is
// expected to be the type hash (bytes32). Nested structs and dynamic-bytes
// fields must be pre-hashed by the caller and passed as bytes32 placeholders.
func HashStruct(args []abi.Type, values []any) (common.Hash, error) {
	encoded, err := Encode(args, values)
	if err != nil {
		return common.Hash{}, err
	}
	return crypto.Keccak256Hash(encoded), nil
}

// HashFromStructHash returns the final EIP-712 digest given a precomputed
// struct hash: keccak256("\x19\x01" || domainSeparator || structHash).
func HashFromStructHash(domainSeparator, structHash common.Hash) common.Hash {
	rawData := fmt.Appendf(nil, "\x19\x01%s%s", string(domainSeparator[:]), string(structHash[:]))
	return crypto.Keccak256Hash(rawData)
}

// HashTypedDataV4 computes the final EIP-712 digest (0x1901 || domainSeparator
// || keccak256(abiEncode(args, values))).
func HashTypedDataV4(domainSeparator common.Hash, args []abi.Type, values []any) (common.Hash, error) {
	structHash, err := HashStruct(args, values)
	if err != nil {
		return common.Hash{}, err
	}
	return HashFromStructHash(domainSeparator, structHash), nil
}
