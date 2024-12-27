package main

import (
	"fmt"
	"strconv"
)

func main() {
	var num int = 32
	// fmt.Println("number is ", num)
	// fmt.Printf("type of num is %T\n", num)

	// var data float64 = float64(num)
	// data = data + 1.23
	// fmt.Println("number is ", data)
	// fmt.Printf("type of data is %T\n", data)

	num = 123
	str := strconv.Itoa(num)
	fmt.Println("str is ", str)
	fmt.Printf("type of str is %T\n", str)

	number_str := "123"
	number_int, _ := strconv.Atoi(number_str)
	fmt.Println("number_int is ", number_int)
	fmt.Printf("type of number_int is %T\n", number_int)

	num_string := "3.14"
	num_float, _ := strconv.ParseFloat(num_string, 64)
	fmt.Println(" num_float is ", num_float)
	fmt.Printf("type of num_float is %T\n", num_float)

}
