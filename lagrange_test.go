package polynomial

import (
	"math/big"
	"testing"
)

// Generates the given number (nhints) of points (x, y)
// q: modulo
func mkTestSet(p Poly, start *big.Int, nhints int, q *big.Int) Points {
	var pts Points = make([]Point, nhints)
	for i := 0; i < nhints; i++ {
		x := big.NewInt(int64(i))
		x.Add(x, start)
		pts[i] = Point{x, p.Eval(x, q)}
	}
	return pts
}

func TestMinusConst(t *testing.T) {
	cases := []struct {
		a   *big.Int
		ans Poly
	}{
		{
			big.NewInt(1),
			NewPolyInts(-1, 1),
		},
		{
			big.NewInt(-8),
			NewPolyInts(8, 1),
		},
		{
			big.NewInt(54321),
			NewPolyInts(-54321, 1),
		},
	}
	for _, c := range cases {
		res := xMinusConst(c.a)
		if res.Compare(&c.ans) != 0 {
			t.Errorf("(x-a) from %v != %v (your answer was %v)", c.a, c.ans, res)
		}
	}
}

func TestLagrange(t *testing.T) {
	cases := []struct {
		ps  Points
		m   *big.Int
		ans Poly
	}{
		{
			Points{
				Point{big.NewInt(1), big.NewInt(1)},
				Point{big.NewInt(2), big.NewInt(4)},
			},
			nil,
			NewPolyInts(0),
		},
		{
			Points{
				Point{big.NewInt(1), big.NewInt(1)},
				Point{big.NewInt(2), big.NewInt(4)},
				Point{big.NewInt(3), big.NewInt(9)},
			},
			big.NewInt(13),
			NewPolyInts(0, 0, 1),
		},
		{
			mkTestSet(NewPolyInts(43, 53, 45, 63, 43, 55, 75), big.NewInt(11), 8, big.NewInt(311)),
			big.NewInt(311),
			NewPolyInts(43, 53, 45, 63, 43, 55, 75),
		},
		{
			mkTestSet(NewPolyInts(43, 53, 45, 63, 43, 55, 75), big.NewInt(111), 8, big.NewInt(311)),
			big.NewInt(311),
			NewPolyInts(43, 53, 45, 63, 43, 55, 75),
		},
		{
			mkTestSet(NewPolyInts(43, 53, 45, 63, 43, 55, 75), big.NewInt(111), 10, big.NewInt(311)),
			big.NewInt(311),
			NewPolyInts(43, 53, 45, 63, 43, 55, 75),
		},
		{
			mkTestSet(NewPolyInts(1234561, 1234562, 1234563, 1234564, 1234565, 1234566, 1234567, 1234568, 1234569), big.NewInt(1234561), 11, big.NewInt(16769023)),
			big.NewInt(16769023),
			NewPolyInts(1234561, 1234562, 1234563, 1234564, 1234565, 1234566, 1234567, 1234568, 1234569),
		},
	}

	for _, c := range cases {
		res := c.ps.Lagrange(c.m)
		if res.Compare(&c.ans) != 0 {
			t.Errorf("Lagrange polynomial from %v [modulo: %v] != %v (your answer was %v)", c.ps, c.m, c.ans, res)
		}
	}
}
