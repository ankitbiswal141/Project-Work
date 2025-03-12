package main

import (
	"fmt"
	"net/http"
	"log"
)

func helloHandler(res http.ResponseWriter,req *http.Request){
	if req.URL.Path != "/hello"{
		http.Error(res,"404 not found", http.StatusNotFound)
		return
	}

	if req.Method != "GET"{
		http.Error(res,"Method not supported",http.StatusNotFound)
		return
	}
}

func formHandler(res http.ResponseWriter, req *http.Request){
	if err != req.ParseForm(); err != nil {
		fmt.Fprintf(res,"Parse Form error: %v",err )
		return
	}
	fmt.Println("POST request successfull.")

	name:=req.FormValue("name")
	age:= req.FormValue("age")

	fmt.Fprintf(res,"name= %s\n",name)
	fmt.Fprintf(res,"age= %d\n", age)

	

}

func main(){
	fileServer:= http.FileServer(http.Dir("./static"))
	http.Handle("/",fileServer)


}