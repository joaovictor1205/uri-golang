package main

import (
	"fmt"
	"math"
)

func main() {
	var R float64
	pi := 3.14159

	fmt.Scan(&R)

	A := pi * math.Pow(R, 2)

	fmt.Printf("A=%.4f\n", A)
}
