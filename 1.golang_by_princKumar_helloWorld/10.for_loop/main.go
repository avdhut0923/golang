package main

import (
	"fmt"
)

func main() {
	// for i := 0; i < 8; i++ {
	// 	fmt.Println("numbers are :", i)
	// }
	// counter := 0
	// for {
	// 	fmt.Println("jai shree ram!")
	// 	counter++
	// 	if counter == 3 {
	// 		break
	// 	}
	// }

	// numbers := []int{2, 3, 5, 8, 4, 6}
	// for index, values := range numbers {
	// 	fmt.Printf("numbers are %d, values are %d\n", index, values)
	// }

	naam := "jai shree ram"
	for index, values := range naam {
		fmt.Printf("index is %d, charter is %c\n", index, values)
	}
}
