package main

import "fmt"

func main() {
	x := 7
	if x > 10 {
		fmt.Println("the number is greater than 10")
	} else if x > 5 {
		fmt.Println("the number is greater than 5 but less than 10")
	} else {
		fmt.Println("the number is less than 5")
	}

	y := 322
	p := 3
	if x > 5 && (y < 40 || p == 2) {
		fmt.Println("JAI SHREE RAM ! ")
	} else {
		fmt.Println("RAM KRISHNA HARI")
	}
}
