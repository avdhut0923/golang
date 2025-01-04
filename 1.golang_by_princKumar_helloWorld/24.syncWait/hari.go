package main

import (
	"fmt"
	"sync"
)

func worker(i int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("shree ganeshay namah %d \n", i)
	// some task happens
	fmt.Printf("worked %d \n", i)
}

func main() {
	// start 3 worker go routine

	var wg sync.WaitGroup
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go worker(i, &wg)
		fmt.Println("jai mahakal")

	}
	wg.Wait()
}
