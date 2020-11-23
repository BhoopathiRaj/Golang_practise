package main

import (
	"fmt"
	"time"
)

func main() {
	today := time.Now()

	switch today.Day() {
	case 5:
		fmt.Println("Today is 5th. Clean your house.")
	case 10:
		fmt.Println("Today is 10th. Buy some snacks.")
	case 15:
		fmt.Println("Today is 15th. Take oil bath.")
	case 23:
		fmt.Println("Today is 23th. Buy Medicine.")
	case 31:
		fmt.Println("Party tonight.")
	default:
		fmt.Println("No information available for that day.")
	}
}
