package main

import "fmt"

func main() {
	studentGrade := make(map[string]int)

	studentGrade["ram"] = 99
	studentGrade["krishna"] = 98
	studentGrade["hari"] = 97
	studentGrade["radha"] = 100

	// fmt.Println("marks of krishna is ", studentGrade["krishna"])
	// fmt.Println("marks of radha is ", studentGrade["radha"])
	// studentGrade["hari"] = 101
	// fmt.Println("marks of hari is ", studentGrade["hari"])

	// delete(studentGrade, "ram")
	// fmt.Println("marks of ram is ", studentGrade["ram"])

	// grades, exists := studentGrade["dev"]
	// fmt.Println("grades of dev : ", grades)
	// fmt.Println("dev exists : ", exists)

	// Grades, Exists := studentGrade["ram"]
	// fmt.Println("grades of ram : ", Grades)
	// fmt.Println("ram exists : ", Exists)

	// for index, values := range studentGrade {
	// 	fmt.Printf("index is   %s   and value is %d\n", index, values)
	// }

	person := map[string]int{
		"brahma":     109879,
		"vishnu":     342772,
		"mahesh":     325775,
		"adi_shakti": 883000729,
	}
	for index, values := range person {
		fmt.Printf("index is %s ,and value is %d\n", index, values)
	}

}
