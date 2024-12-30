package main

import (
	"fmt"
	"os"
)

func main() {
	// file, err := os.Create("ram.txt")
	// if err != nil {
	// 	fmt.Println("error while creating a file", err)
	// 	return
	// }
	// defer file.Close()

	// content := "jai shree ram bolo jai shree ram"
	// byte, erro := io.WriteString(file, content+"\n")
	// fmt.Println("byte size of given file is ", byte)
	// if erro != nil {
	// 	fmt.Println("error while writing in the file:", erro)
	// 	return
	// }
	// fmt.Println("succesfully created and written a file")

	//CREATE buffer to read file content
	// buffer := make([]byte, 1024)

	// // read the file content
	// for {
	// 	n, err := file.Read(buffer)
	// 	if err == io.EOF {
	// 		break
	// 	}
	// 	if err != nil {
	// 		fmt.Println("error while reading file", err)
	// 		return
	// 	}

	// 	fmt.Println(string(buffer[:n]))

	// }
	// fmt.Println("jai shree ram")

	//read the entire file into a byte slice
	content, err := os.ReadFile("ram.txt")
	if err != nil {
		fmt.Println("error while reading this file", err)
		return
	}
	fmt.Println(string(content))
}
