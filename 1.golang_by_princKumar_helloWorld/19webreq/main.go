package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	fmt.Println("learing web services")
	res, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err != nil {
		fmt.Println("error while GET responsce", err)
		return
	}
	defer res.Body.Close()
	fmt.Printf("type of response : %T \n ", res)
	// fmt.Println("response : ", res)
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("erro reading responce ", err)
		return
	}
	fmt.Println("responce :", string(data))

}
