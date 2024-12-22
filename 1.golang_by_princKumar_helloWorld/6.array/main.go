package main

import "fmt"

func main() {
	fmt.Println("we are learnig array in golang")

	// var names [5]string

	// names[0] = "ram"
	// names[1] = "krishna"
	// fmt.Println("names are ", names)

	// numbers := [9]int{1, 2, 3, 56}
	// fmt.Println("numbers are ", numbers)

	// fmt.Println("length of numbers array is ", len(numbers))

	// fmt.Println("element at first index is ", names[1])

	var names [5]string
	names[0] = "4"
	names[3] = "4"
	names[2] = "2"

	fmt.Println("names array contains ", names)

	fmt.Printf("names array contains %q ", names)

}
