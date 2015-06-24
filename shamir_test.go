package polynomial

import (
	"math/big"
	"testing"
)

func TestGenRandomShares(t *testing.T) {
	q := big.NewInt(179424691)

	for i := 0; i < 10; i++ {
		ps, p := GenRandomShares(10, 7, q)
		for j := len(ps); j >= 7; j-- {
			var testps Points = ps[:j]
			recoveredPoly := testps.Lagrange(q)
			if p[0].Cmp(recoveredPoly[0]) != 0 {
				t.Errorf("Recovering Shamir's secret sharing (with %v points) fails: the secret was %v (your answer was %v)", len(testps), p[0], recoveredPoly[0])
			}
		}
		for j := 6; j >= 3; j-- {
			var testps Points = ps[:j]
			recoveredPoly := testps.Lagrange(q)
			if p[0].Cmp(recoveredPoly[0]) == 0 {
				t.Errorf("Recovering Shamir's secret sharing (with %v points) should fail but the answer was correct: the secret was %v (your answer was %v)", len(testps), p[0], recoveredPoly[0])
			}
		}
	}
}
