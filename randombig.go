package polynomial

import (
	"crypto/rand"
	"math/big"
)

// Returns a big integer with the given nb bits
func RandomBigInt(nb int) *big.Int {
	b := make([]byte, nb)
	n, err := rand.Read(b)
	if err != nil || nb != n {
		panic("Failed to generate a seed.")
	}
	r := new(big.Int)
	r.SetBytes(b)
	return r
}
