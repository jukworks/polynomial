package polynomial

import (
	"fmt"
	"math/big"
	"strings"
)

// Point type represents a coordinate (x, y) where x and y are big integers
type Point struct {
	x, y *big.Int
}

// Points type represents a set of Point type
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
