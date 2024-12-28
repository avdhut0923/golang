package main

import "fmt"

func add(a, b int) int {
	return a + b

}

func main() {
	fmt.Println("starting of the program :")

	defer fmt.Println("mid of the program :")
	defer fmt.Println("end of the program :")
	data := add(3, 4)
	fmt.Println("addition ka ans hai ", data)
}
