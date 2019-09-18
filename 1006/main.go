package main

import (
	"fmt"
)

func main() {
	var A float64
	var B float64
	var C float64
	var MEDIA float64

	fmt.Scan(&A)
	fmt.Scan(&B)
	fmt.Scan(&C)

	MEDIA = ((A * 2) + (B * 3) + (C * 5)) / (2 + 3 + 5)

	fmt.Printf("MEDIA = %.1f\n", MEDIA)
}
