package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	type Person struct {
		Name string
		Age  int
		Car  []string
	}

	p1 := &Person{Name: "Rahul", Age: 30, Car: []string{"Ford", "Tata"}}
	data, _ := json.Marshal(p1)
	fmt.Println(string(data))

}
