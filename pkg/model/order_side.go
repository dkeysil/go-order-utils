package model

// Side is the direction of an order: BUY or SELL.
type Side = int

const (
	// BUY signals an order buying the outcome token.
	BUY Side = iota
	// SELL signals an order selling the outcome token.
	SELL
)
