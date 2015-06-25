package polynomial

import (
	"crypto/rand"
	"math/big"
)

// Returns a big integer with the given nb bytes
func RandomBigInt(nb int) *big.Int {
	b := make([]byte, nb)
	rand.Read(b)
	r := new(big.Int)
	r.SetBytes(b)
	return r
}
