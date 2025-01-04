package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"
)

type todo struct {
	UserId    int    `json:"UserId"`
	Id        int    `json : "Id"`
	Title     string `json : "Title"`
	Completed bool   `json : "Completed"`
}

func performpostreq() {
	todo := Todo{
		"userId":    1,
		"id":        1,
		"title":     "Jai shree ram",
		"completed": true,
	}

	// convert the Todo struct to json
	jsonData, err := json.Marshal(todo)
	if err !=nil {
		fmt.Println("error marsheling : ",err)
		return 
	}
	// convert json data to string data 
	 jsonString := string(jsonData)

	//  convert string data to string 
	jsonReader := string.NewReader(jsonString)

	myurl := "https://jsonplaceholder.typicode.com/todos"

	// send post request 
	res,err !=http.post(myurl,"application/json",jsonReader)
	if err != nil{
		fmt.Println("error sending request : ",err)
		return
	}

	defer res.Body.Close()

	data , _ := ioutil.ReadAll(res.Body)
	fmt.Println("response : ",string(data))
}

func main() {
	// fmt.Println("learning about crud operation in golang")
	// res, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	// if err != nil {
	// 	fmt.Println("error while GET url", err)
	// 	return
	// }
	// defer res.Body.Close()

	// if res.StatusCode != http.StatusOK {
	// 	fmt.Println("error while fetching data", res.Status)
	// 	return
	// }

	// // 	data, err := io.ReadAll(res.Body)
	// // 	if err != nil {
	// // 		fmt.Println("error while reading body", err)
	// // 		return
	// // 	}
	// // 	fmt.Println("data :", string(data))

	// var todo Todo
	// err = json.NewDecoder(res.Body).Decode(&todo)
	// if err != nil {
	// 	fmt.Println("Error decoding : ", err)
	// 	return
	// }
	// fmt.Println("Todo: ", todo)
	performpostreq()

}
