package main

import "fmt"

// func divide(a, b float32) (float32, error) {
// 	if b == 0 {
// 		return 0, fmt.Errorf("denominator can not be 0")
// 	}
// 	return a / b, nil
// }

func divide(a, b float32) (float32, string) {
	if b == 0 {
		return 0, "denominator can not be 0"
	}
	return a / b, "nil"
}

func main() {

	data, _ := divide(10, 0) //blank identifier to ignore the second value
	fmt.Println("division of given no.s is ", data)
}
