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

// HashTypedDataV4 computes the final EIP-712 digest (0x1901 || domainSeparator
// || keccak256(abiEncode(args, values))).
func HashTypedDataV4(domainSeparator common.Hash, args []abi.Type, values []any) (common.Hash, error) {
	encoded, err := Encode(args, values)
	if err != nil {
		return common.Hash{}, err
	}

	rawData := fmt.Appendf(nil, "\x19\x01%s%s", string(domainSeparator[:]), string(crypto.Keccak256Hash(encoded).Bytes()))
	return crypto.Keccak256Hash(rawData), nil
}
