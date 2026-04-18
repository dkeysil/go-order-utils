package utils

import (
	"crypto/rand"
	"math"
	"math/big"
	"time"
)

// GenerateRandomSalt returns a non-negative 32-bit-bounded salt, matching the
// JS reference implementation (`Math.random() * 2**32`).
func GenerateRandomSalt() int64 {
	maxInt := math.Pow(2, 32)
	nBig, _ := rand.Int(rand.Reader, big.NewInt(int64(maxInt)))
	return nBig.Int64()
}

// GenerateTimestampMs returns the current unix time in milliseconds. It is
// the default timestamp source used by the V2 order builder and replaces
// V1's `Nonce` for order uniqueness.
func GenerateTimestampMs() int64 {
	return time.Now().UnixMilli()
}
