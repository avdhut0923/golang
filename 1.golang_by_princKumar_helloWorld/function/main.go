package main

import "fmt"

func simpleFunc() {
	fmt.Println("jai jai shree ram")
}

// func add(a, b int) int {
// 	return a + b
// }

// func add(a, b int) (result int) {
// 	result = a + b
// 	return
// }

func mul(a, b int) int {
	return a * b
}

func main() {
	// fmt.Println("jai shree ram")
	// simpleFunc()
	// addi := add(3, 55)
	// fmt.Println("addition ka answer hai", addi)
	data := mul(3, 4)
	fmt.Println("ans of multiplication is ", data)

}
