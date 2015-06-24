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
		strs[i] = fmt.Sprintf("Point #%v %v", i+1, p)
	}
	return strings.Join(strs, "\n")
}
