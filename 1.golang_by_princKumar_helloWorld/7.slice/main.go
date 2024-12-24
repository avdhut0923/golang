package main

import "fmt"

func main() {
	// numbers := []int{1, 2, 3, 4, 5}
	// numbers = append(numbers, 3, 4, 5, 6, 7, 8, 23)
	// fmt.Println("numbers : ", numbers)
	// fmt.Printf("numbers has datatypes %T\n", numbers)
	// fmt.Println("length:", len(numbers))

	// name := []string{}
	// fmt.Println("name : ", name)

	// numbers := make([]int, 3, 5)

	// numbers = append(numbers, 23)
	// numbers = append(numbers, 2)
	// numbers = append(numbers, 24)
	// numbers = append(numbers, 243)

	// fmt.Println("slice", numbers)
	// fmt.Println("length", len(numbers))
	// fmt.Println("capacity:", cap(numbers))

	stock := make([]int, 0)

	fmt.Println("slice", stock)
	fmt.Println("length", len(stock))
	fmt.Println("capacity:", cap(stock))

}
