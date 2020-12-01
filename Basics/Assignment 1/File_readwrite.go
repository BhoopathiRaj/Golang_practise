package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	data, err := ioutil.ReadFile("new.txt")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(data))

	writeText()
}

func writeText() {
	message := []byte("hello write")
	err := ioutil.WriteFile("new.txt", message, 0644)
	if err != nil {
		fmt.Println(err)
	}

}