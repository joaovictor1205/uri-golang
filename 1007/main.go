package main

import (
	"fmt"
)

func main() {

	var A int
	var B int
	var C int
	var D int
	var DIFERENCA int

	fmt.Scan(&A)
	fmt.Scan(&B)
	fmt.Scan(&C)
	fmt.Scan(&D)

	DIFERENCA = (A * B) - (C * D)

	fmt.Printf("DIFERENCA = %d\n", DIFERENCA)
}
