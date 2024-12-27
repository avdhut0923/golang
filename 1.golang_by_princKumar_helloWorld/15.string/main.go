package main

import (
	"fmt"
	"strings"
)

func main() {
	data := "apple,orange,jackfruit,sweet lime"
	parts := strings.Split(data, ",")
	fmt.Println(parts)

	str := "one two threetwo  two five two six seven "
	cout := strings.Count(str, "two")
	fmt.Println("count:", cout) // gives 4

	strang := "    jai shree ram !    "
	trimmed := strings.TrimSpace(strang)
	fmt.Println("trimmed string :", trimmed)

	str1 := "ram"
	str2 := "krishna"
	str3 := "hari"
	result := strings.Join([]string{str1, str2, str3}, "_")
	fmt.Println("rsult:", result)

}
