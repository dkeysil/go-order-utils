package builder

import (
	"github.com/dkeysil/go-order-utils/pkg/eip712"
	"github.com/dkeysil/go-order-utils/pkg/model"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

var (
	protocolName = crypto.Keccak256Hash([]byte("Polymarket CTF Exchange"))

	// All exchanges share the domain name; only the domain version differs.
	protocolVersions = map[model.VerifyingContract]common.Hash{
		model.CTFExchange:        crypto.Keccak256Hash([]byte("2")),
		model.NegRiskCTFExchange: crypto.Keccak256Hash([]byte("2")),
		model.CTFExchangeV3:      crypto.Keccak256Hash([]byte("3")),
	}
)

var (
	orderStructure = []abi.Type{
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

	// orderTypeString is the EIP-712 encodeType output for the Order struct.
	// Embedded verbatim into ERC-7739 wrapped POLY_1271 signatures, so its
	// byte sequence must remain stable.
	orderTypeString = []byte("Order(uint256 salt,address maker,address signer,uint256 tokenId,uint256 makerAmount,uint256 takerAmount,uint8 side,uint8 signatureType,uint256 timestamp,bytes32 metadata,bytes32 builder)")

	orderStructureHash = crypto.Keccak256Hash(orderTypeString)
)

// ERC-7739 wrap for POLY_1271 deposit-wallet orders.
//
// On-wire signature layout:
//
//	innerSig(65) || appDomainSep(32) || contentsHash(32) || orderTypeString || uint16_BE(len)
//
// where innerSig is over the EIP-712 digest of the TypedDataSign struct under
// the CTFExchange app domain, and TypedDataSign carries the wallet's own
// EIP-712 domain (DepositWallet/1, verifyingContract = order.Signer) inside.
var (
	// typedDataSignArgs matches:
	//   bytes32 typeHash, bytes32 contents, bytes32 nameHash,
	//   bytes32 versionHash, uint256 chainId, address verifyingContract,
	//   bytes32 salt
	typedDataSignArgs = []abi.Type{
		eip712.Bytes32,
		eip712.Bytes32,
		eip712.Bytes32,
		eip712.Bytes32,
		eip712.Uint256,
		eip712.Address,
		eip712.Bytes32,
	}

	// EIP-712 nested-type encoding: outer struct followed by referenced
	// struct types in alphabetical order. Only Order is referenced.
	typedDataSignTypeHash = crypto.Keccak256Hash(append(
		[]byte("TypedDataSign(Order contents,string name,string version,uint256 chainId,address verifyingContract,bytes32 salt)"),
		orderTypeString...,
	))
)
