package main

import (
	"fmt"
)

func main() {
	var A int
	var B int

	fmt.Scan(&A)
	fmt.Scan(&B)

	PROD := A * B

	fmt.Printf("PROD = %d\n", PROD)
}
