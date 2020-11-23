package main

import (
	"fmt"
)

func main() {
	x := 10

	if x >= 50 && x <= 100 {
		fmt.Println("Germany")
	} else if x <= 50 {
		fmt.Println("Japan")
	} else {
		fmt.Println("Canada")
	}
}
