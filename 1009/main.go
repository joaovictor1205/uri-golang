package main

import (
	"fmt"
)

func main() {

	var NAME string
	var SALARY float64
	var SALE float64

	fmt.Scan(&NAME)
	fmt.Scan(&SALARY)
	fmt.Scan(&SALE)

	TOTAL := SALARY + (SALE * 0.15)

	fmt.Printf("TOTAL = R$ %.2f\n", TOTAL)
}
