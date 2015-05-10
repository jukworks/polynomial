package polynomial

import (
	"crypto/rand"
	"math/big"
)

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

func GenShares(n, k int, q *big.Int) (ps Points, p Poly) {
	if q.ProbablyPrime(100) == false {
		panic("The modulo should be a prime.")
	}
	// k차 다항식을 만들고, modulo q를 이용. n개의 points를 돌려주면 됨.
	size := q.BitLen()/8 + 1
	p = make([]*big.Int, k)
	for i := 0; i < k; i++ {
		coeff := RandomBigInt(size)
		coeff.Mod(coeff, q)
		p[i] = coeff
	}
	ps = make([]Point, n)
	for i := 0; i < n; i++ {
		r := RandomBigInt(size)
		r.Mod(r, q)
		var t Point
		t.x = r
		t.y = p.Eval(r, q)
		ps[i] = t
	}
	return
}
