package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("hey,what's your name?")
	// var name string
	// fmt.Scan(&name)
	// fmt.Println("hello mr.", name)

	reader := bufio.NewReader(os.Stdin)
	name, _ := reader.ReadString('\n') //delimitor
	fmt.Println("hello Mr.", name)
}
