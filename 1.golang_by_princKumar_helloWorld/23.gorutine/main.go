package main

import (
	"fmt"
	"time"
)

func sayNam() {
	fmt.Println("nam :jai shree ram")
	time.Sleep(1000 * time.Millisecond) // simuating some work
	fmt.Println("nam :ram krushna hari")
}
func sayJayhari() {
	fmt.Println("jai hanuman")
	time.Sleep(1100 * time.Millisecond)
	fmt.Println("jai jai shree ram")
}

func main() {
	fmt.Println("shree ganeshay namah")
	go sayNam()
	go sayJayhari()
	time.Sleep(1500 * time.Millisecond)

}
