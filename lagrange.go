package polynomial

import (
	"fmt"
	"math/big"
	"strings"
)

type Point struct {
	x, y *big.Int
}

type Points []Point

func (p Point) String() string {
	return fmt.Sprintf("(%v, %v)", p.x, p.y)
}

func (ps Points) String() string {
	strs := make([]string, len(ps))
	for i, p := range ps {
		strs[i] = fmt.Sprintf("%v", p)
	}
	return strings.Join(strs, ",")
}

/*
	a를 주면 (x-a) 다항식을 반환
*/
func xMinusConst(a *big.Int) Poly {
	b := new(big.Int)
	b.Neg(a)
	return Poly{b, big.NewInt(1)}
}

/*
	라그랑주 다항식 기법은 modulo 연산을 하지 않고도 사용할 수 있지만,
	Big Integer를 이용한 다항식만 취급하기 때문에 분수 처리를 할 수 없고,
	따라서 나머지 연산을 통해 모두 정수화 시켜서 계산한다.
	따라서 파라미터 m이 주어지지 않으면 P = 0 다항식을 반환하도록 했다.
*/
func (ps Points) Lagrange(m *big.Int) (lag Poly) {
	// fmt.Println("--------------------------------")
	if m == nil {
		return NewPolyInts(0)
	}
	lag = NewPolyInts(0) // Lx 다항식(L1, L2, L3, ...)들을 모두 더하기 위한 변수
	n := len(ps)         // 주어진 점(힌트)의 개수
	for i := 0; i < n; i++ {
		lx := NewPolyInts(1)
		con := new(big.Int)
		con.Set(ps[i].y)
		deno := new(big.Int)
		for j := 0; j < n; j++ { // Lx를 계산하는 루프
			if i == j {
				continue
			}
			lx = lx.Mul(xMinusConst(ps[j].x), m) // (x-n) 형태를 곱해주고,
			deno.Sub(ps[i].x, ps[j].x)
			deno.Mod(deno, m)
			deno.ModInverse(deno, m)
			con.Mul(con, deno)
			con.Mod(con, m) // y값과 분모값을 모두 곱해준다
		}
		// fmt.Println(lx, "*", con)
		for k := 0; k < len(lx); k++ { // 계산된 상수를 각 coefficient에 곱해준다.
			lx[k].Mul(lx[k], con)
			if m != nil {
				lx[k].Mod(lx[k], m)
			}
		}
		lx.trim()
		// fmt.Printf("Lt[%v] = %v (Constant part: %v)\n", i, lx, con)
		lag = lag.Add(lx, m)
		// fmt.Println("+ =", lag)
	}
	lag.trim()
	return
}
