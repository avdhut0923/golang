package main

import "fmt"

func modifyByReference(num *int) {
	*num = *num + 10
}

func main() {
	// var num int
	// num = 2

	// var ptr *int
	// ptr = &num

	// num := 3
	// ptr := &num

	// fmt.Println("pointer has value : ", ptr)
	// fmt.Println("NUM contain value : ", *ptr) //*ptr returns the value of box for which it stores pointer

	var pointer *int //default pointer value is nil
	if pointer == nil {
		fmt.Println("pointer is not assigned:", pointer)
	}

	value := 10
	modifyByReference(&value)
	fmt.Println("value modified is :", value)
}
