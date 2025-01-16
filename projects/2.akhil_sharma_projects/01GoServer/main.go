package main

import (
	"fmt"
	"log"
	"net/http"
)

func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err : %v", err)
		return
	}
	fmt.Fprintf(w, "Post request succesful \n")
	name := r.FormValue("name")
	address := r.FormValue("address")
	fmt.Fprintf(w, "Name = %s\n", name)
	fmt.Fprintf(w, "Address = %s\n", address)
}

func namoHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/namo" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "method is not supported	", http.StatusNotFound)
		return
	}
	fmt.Fprintf(w, "namo vasudevaya !")
}

func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer) //  / WILL serve index.html file
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/namo", namoHandler)

	fmt.Printf("starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil { // important line
		log.Fatal(err)
	}
}
