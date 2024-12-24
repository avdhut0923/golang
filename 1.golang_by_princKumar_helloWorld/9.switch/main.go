package main

import "fmt"

func main() {
	// bhagvan := "mahakal"

	// switch bhagvan {
	// case "ram", "datta guru", "parshuram":
	// 	fmt.Println("vishnu")
	// case "bhole baba", "shiv ji", "mahakal":
	// 	fmt.Println("mahesh")
	// case "sopandev":
	// 	fmt.Print("bramha")
	// case "muktabai":
	// 	fmt.Println("maya")
	// }

	temp := 33
	switch {
	case temp < 0:
		fmt.Println("freezing")
	case temp < 10:
		fmt.Println("cold")
	case temp < 20:
		fmt.Println("normal")
	case temp > 21 && temp < 50:
		fmt.Println("hot")
	default:
		fmt.Println("you are not on earth")
	}

}
