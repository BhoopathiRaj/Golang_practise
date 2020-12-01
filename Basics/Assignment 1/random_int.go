package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"time"
)

func main() {
	numThreads := runtime.NumCPU()
	runtime.GOMAXPROCS(numThreads)

	var numIntsToGenerate = 100000000
	var numIntsPerThread = numIntsToGenerate / numThreads

	ch := make(chan []int)

	fmt.Printf("Initiating single-threaded random number generation.\n")
	startSingleRun := time.Now()

	go makeRandomNumbers(numIntsToGenerate, ch)
	singleThreadIntSlice := <-ch
	elapsedSingleRun := time.Since(startSingleRun)
	fmt.Printf("Single-threaded run took %s\n", elapsedSingleRun)

	fmt.Printf("Initiating multi-threaded random number generation.\n")

	multiThreadIntSlice := make([][]int, numThreads)
	startMultiRun := time.Now()

	for i := 0; i < numThreads; i++ {
		go makeRandomNumbers(numIntsPerThread, ch)
	}
	for i := 0; i < numThreads; i++ {
		multiThreadIntSlice[i] = <-ch
	}
	elapsedMultiRun := time.Since(startMultiRun)
	fmt.Printf("Multi-threaded run took %s\n", elapsedMultiRun)
	fmt.Print(len(singleThreadIntSlice))
}

func makeRandomNumbers(numInts int, ch chan []int) {
	source := rand.NewSource(time.Now().UnixNano())
	generator := rand.New(source)
	result := make([]int, numInts)
	for i := 0; i < numInts; i++ {
		result[i] = generator.Intn(numInts * 100)
	}
	ch <- result
}
