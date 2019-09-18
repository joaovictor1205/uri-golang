package main

import (
	"fmt"
)

func main() {

	var CODE_PRODUCT1 int
	var NUMBER_PRODUCT1 int
	var PRICE_PRODUCT1 float64

	var CODE_PRODUCT2 int
	var NUMBER_PRODUCT2 int
	var PRICE_PRODUCT2 float64

	fmt.Scan(&CODE_PRODUCT1, &NUMBER_PRODUCT1, &PRICE_PRODUCT1)

	fmt.Scan(&CODE_PRODUCT2, &NUMBER_PRODUCT2, &PRICE_PRODUCT2)

	number1 := float64(NUMBER_PRODUCT1)
	number2 := float64(NUMBER_PRODUCT2)

	fmt.Printf("VALOR A PAGAR: R$ %.2f\n", (PRICE_PRODUCT1*number1)+(PRICE_PRODUCT2*number2))

}
