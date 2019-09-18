package main

import (
	"fmt"
)

func main() {
	var A float64
	var B float64
	var MEDIA float64

	fmt.Scan(&A)
	fmt.Scan(&B)

	MEDIA = ((A * 3.5) + (B * 7.5)) / (3.5 + 7.5)

	fmt.Printf("MEDIA = %.5f\n", MEDIA)
}
