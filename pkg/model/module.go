// Package model contains the order data types and enums exchanged with the
// CTF Exchange order builder.
package model

// VerifyingContract identifies which exchange the order is bound to.
type VerifyingContract = int

const (
	// CTFExchange is the standard CTF Exchange V2 contract.
	CTFExchange VerifyingContract = iota
	// NegRiskCTFExchange is the negative-risk CTF Exchange V2 contract.
	NegRiskCTFExchange
	// CTFExchangeV3 is the combos (RFQ) exchange contract, EIP-712 domain version "3".
	CTFExchangeV3
)
