package main

import (
	"fmt"
	"net/url"
)

func main() {
	fmt.Println("SHREE GANESHAY NAMA:")
	myurl := "https://www.merriam-webster.com/thesaurus/remedy?nearby-entries"
	fmt.Printf("Type of Url: %T \n", myurl)

	parsedURl, err := url.Parse(myurl)
	if err != nil {
		fmt.Println("cant parse Url", err)
		return
	}
	fmt.Printf("Type of Url: %T \n", parsedURl)

	fmt.Println("Scheme : ", parsedURl.Scheme)
	fmt.Println("Host : ", parsedURl.Host)
	fmt.Println("path : ", parsedURl.Path)
	fmt.Println("RawQuery : ", parsedURl.RawQuery)

	// modify url components
	parsedURl.Path = "/newPath"
	parsedURl.RawQuery = "jaishreeram/bolo.jaishreeram"

	newUrl := parsedURl.String()

	fmt.Println("new URL : ", newUrl)

}
