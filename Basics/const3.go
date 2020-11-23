package main

import "fmt"

const (
	PRODUCT  = "Mobile"
	QUANTITY = 50
	PRICE    = 50.50
	STOCK    = true
)

const a int16 = 8

func main() {
	const PRODUCT = "Cellphone"
	var b int16 = 42
	fmt.Printf("%v %T \n", QUANTITY+b+a, QUANTITY+b+a)
	fmt.Println(PRICE)
	fmt.Println(PRODUCT)
	fmt.Println(STOCK)
}
