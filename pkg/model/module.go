// Package model contains the order data types and enums exchanged with the
// V2 CTF Exchange order builder.
package model

// VerifyingContract identifies which V2 exchange the order is bound to.
type VerifyingContract = int

const (
	// CTFExchange is the standard CTF Exchange V2 contract.
	CTFExchange VerifyingContract = iota
	// NegRiskCTFExchange is the negative-risk CTF Exchange V2 contract.
	NegRiskCTFExchange
)
