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
	case 23:
		fmt.Println("Today is 23th. Buy Medicine.")
		fallthrough
	case 25:
		fmt.Println(" Visit a doctor.")
		fallthrough
	case 31:
		fmt.Println("Party tonight.")
		break
		fmt.Println("take rest")
	default:
		fmt.Println("No information available for that day.")
	}
}
