package main

import "fmt"

func main() {
	const A = "New"
	var B = "World"

	// Concat strings.
	var helloWorld = A + " " + B
	helloWorld += "!"
	fmt.Println(helloWorld)

	// Compare strings.
	fmt.Println(A == "New")
	fmt.Println(A == B)

}