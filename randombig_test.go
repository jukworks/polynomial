package polynomial

import "testing"

func TestRandomBigInt(t *testing.T) {
	for i := 50; i < 100; i++ {
		p := RandomBigInt(i)
		if p.BitLen() < (i-1)*8 || p.BitLen() > i*8 {
			t.Errorf("A request for a random big integer with %v bits returns %v (%v bits)", i, p, p.BitLen())
		}
	}
}
