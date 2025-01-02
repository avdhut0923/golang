package main

import (
	"fmt"
	"io"
	"net/http"
)

type todo struct {
	UserId    int    `json:"UserId"`
	Id        int    `JSON : "ID"`
	Title     string `JSON : "Title"`
	Colpleted bool   `JSON : "Colpleted"`
}

func main() {
	fmt.Println("learning about crud operation in golang")
	res, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err != nil {
		fmt.Println("error while GET url", err)
		return
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		fmt.Println("error while fetching data", res.Status)
		return
	}

	data, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("error while reading body", err)
		return
	}
	fmt.Println("data :", string(data))
}
