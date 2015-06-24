package polynomial

import "math/big"

// Generates P = (x-a) for the given a
func xMinusConst(a *big.Int) Poly {
	b := new(big.Int)
	b.Neg(a)
	return Poly{b, big.NewInt(1)}
}

// If m is not given (i.e. nil), return P = 0
// This library only handles polynomials with BigInteger coefficients
func (ps Points) Lagrange(m *big.Int) (lag Poly) {
	if m == nil {
		return NewPolyInts(0)
	}
	lag = NewPolyInts(0) // lag will store the sum of Polynomial L_{x}s (L1, L2, L3, ...)
	n := len(ps)         // # of hints
	for i, con, deno := 0, new(big.Int), new(big.Int); i < n; i++ {
		lx := NewPolyInts(1)
		con.Set(ps[i].y)
		for j := 0; j < n; j++ { // calculate L_{x}
			if i == j {
				continue
			}
			lx = lx.Mul(xMinusConst(ps[j].x), m) // * (x-n)
			deno.Sub(ps[i].x, ps[j].x)
			deno.Mod(deno, m)
			deno.ModInverse(deno, m)
			con.Mul(con, deno)
			con.Mod(con, m) // * y * denominator
		}
		for k := 0; k <= lx.GetDegree(); k++ { // all coefficients * evaluated constant
			lx[k].Mul(lx[k], con)
			lx[k].Mod(lx[k], m)
		}
		lx.trim()
		// fmt.Printf("Lt[%v] = %v (Constant part: %v)\n", i, lx, con)
		lag = lag.Add(lx, m)
		// fmt.Println("+ =", lag)
	}
	lag.trim()
	return
}
