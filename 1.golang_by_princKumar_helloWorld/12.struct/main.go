package main

import "fmt"

type Person struct {
	firstName string
	LastName  string
	age       int
}
type contact struct {
	email string
	phone string
}
type address struct {
	house_no int
	city     string
	country  string
	khanda   string
}
type employee struct {
	person_detail  Person
	person_contact contact
	Person_address address
}

func main() {
	// var RAM Person
	// fmt.Println("Ram person : ", RAM)
	// RAM.firstName = "Ram"
	// RAM.LastName = "raghuvanshi"
	// RAM.age = 888
	// fmt.Println("Ram person : ", RAM)

	// method 2

	person1 := Person{
		firstName: "krishna",
		LastName:  "yadav",
		age:       125,
	}
	fmt.Println("person1 : ", person1)

	//method 3

	var person2 = new(Person)
	person2.firstName = "mahesh"
	person2.LastName = "OM"
	person2.age = 0

	fmt.Println("person2 is :", person2)

	var employee1 employee

	employee1.person_detail = Person{
		firstName: "ram",
		LastName:  "raghuvanshi",
		age:       125,
	}

	employee1.person_contact.email = "contact@gamil.com"
	employee1.person_contact.phone = "25235367676"

	employee1.Person_address = address{
		house_no: 99,
		city:     "ayodhya",
		country:  "bharat",
	}
	fmt.Println("employee1 :", employee1)
	employee1.Person_address.khanda = "bharatkhand"
	fmt.Println("employee1 :", employee1.Person_address.khanda)

}
