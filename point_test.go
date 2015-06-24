package polynomial

import (
	"fmt"
	"math/big"
)

// Printing example for Point data sturcture
func ExamplePrintPoint() {
	p := Point{big.NewInt(1), big.NewInt(2)}
	fmt.Println(p)
	q := Point{big.NewInt(1234567890), big.NewInt(987654321)}
	fmt.Println(q)
	// Output:
	// (1, 2)
	// (1234567890, 987654321)
}

// Printing example for Points data structure
func ExamplePrintPoints() {
	ps := Points{
		Point{big.NewInt(1), big.NewInt(2)},
		Point{big.NewInt(12345), big.NewInt(54321)},
	}
	fmt.Println(ps)
	// Output:
	// Point #1 (1, 2)
	// Point #2 (12345, 54321)
}
