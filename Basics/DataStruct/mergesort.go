package main

import (
	"fmt"
	"sort"
)

func divideList(list []int, cut int) (tmp, s []int) {
	tmp = append(list[cut:])
	s = append(list[:cut])
	return tmp, s
}

func sortList(list []int, c chan []int) {
	sort.Ints(list)
	c <- list
}

func main() {
	// define the size of the number we want to sort
	size := 5

	// put the number list into a slice
	list := []int{-9, -11, 12, 13, 9}
	fmt.Println("Your digits input : ", list)

	// divide the number list into 4 different slices
	cut := size / 4
	list, s1 := divideList(list, cut)
	cut = len(list) / 3
	list, s2 := divideList(list, cut)
	cut = len(list) / 2
	list, s3 := divideList(list, cut)
	s4 := list

	// sort in different go routines
	c := make(chan []int, 4)
	go sortList(s1, c)
	go sortList(s2, c)
	go sortList(s3, c)
	go sortList(s4, c)

	sortedS1 := <-c
	sortedS2 := <-c
	sortedS3 := <-c
	sortedS4 := <-c

	fmt.Println(sortedS1)
	fmt.Println(sortedS2)
	fmt.Println(sortedS3)
	fmt.Println(sortedS4)
	fmt.Println("-----------------------")

	// merge the 4 sorted slices
	sortedList1 := append(sortedS1, sortedS2...)
	fmt.Println(sortedList1)
	sortedList2 := append(sortedS3, sortedS4...)
	fmt.Println(sortedList2)
	finalList := append(sortedList1, sortedList2...)
	fmt.Println(finalList)

	// we need to sort it again
	sort.Ints(finalList)
	fmt.Println("Here your digits sorted: ", finalList)
}