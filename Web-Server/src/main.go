package main

import (
	"fmt"
	"log"
	"net/http"
)

func homeHandle(res http.ResponseWriter, req *http.Request) {

	if req.URL.Path != "/home" {
		http.Error(res, "404 not found!", http.StatusNotFound)
		return
	}

	if req.Method != "GET" {
		http.Error(res, "Method not found", http.StatusNotFound)
		return
	}

	fmt.Println("Home function called")
}

func formHandle(res http.ResponseWriter, req *http.Request) {

	if err := req.ParseForm(); err != nil {
		fmt.Fprintf(res, "Parse form error: %v", err)
		return
	}
	fmt.Fprintf(res, "POST request successfull.")

	name := req.FormValue("name")
	age := req.FormValue("age")

	fmt.Fprintf(res, "Name=%v\n", name)
	fmt.Fprintf(res, "Age=%v\n", age)
}

func main() {

	fileServer := http.FileServer(http.Dir("../static"))
	http.Handle("/", fileServer)

	http.HandleFunc("/home", homeHandle)
	http.HandleFunc("/form", formHandle)

	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal(err)
	}

}
