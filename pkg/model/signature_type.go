package model

// SignatureType identifies the wallet signing mode used to produce an order
// signature. Values and their semantics mirror SignatureTypeV2 in
// py-clob-client-v2 and @polymarket/clob-client-v2.
type SignatureType = int

const (
	// EOA is an ECDSA EIP712 signature signed directly by an externally-owned
	// account.
	EOA SignatureType = iota

	// POLY_PROXY is an EIP712 signature signed by an EOA that owns a Polymarket
	// proxy wallet.
	POLY_PROXY

	// POLY_GNOSIS_SAFE is an EIP712 signature signed by an EOA that owns a
	// Polymarket Gnosis safe.
	POLY_GNOSIS_SAFE

	// POLY_1271 is an EIP1271 signature signed by a smart contract wallet or
	// vault. Added in V2.
	POLY_1271
)
