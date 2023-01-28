package main

import (
	"fmt"
	"net/http"
	"zkfmapf123/go-server/utils"
)

/*
	/ 		index.html
	/hello	func
	/form	form.html
*/

const PORT = 8080
func main(){
	fileServer := http.FileServer(http.Dir("./static"))	
	router(fileServer)

	fmt.Printf("Starting server at port %d", PORT)
	if err := http.ListenAndServe(fmt.Sprintf(":%d",PORT),nil); err !=nil {
		panic(err)
	}
}

func router(fServer http.Handler){

	http.Handle("/", fServer)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)
}

// hello
func helloHandler(w http.ResponseWriter, r *http.Request) {
	if err := utils.ResponseError((r.URL.Path != "/hello"), w, "404 not found", http.StatusNotFound); err != nil {
		return
	}

	if err := utils.ResponseError((r.Method != "GET"), w, "method is not supported", http.StatusNotFound); err != nil {
		return 
	}

	fmt.Fprintf(w, "hello!")	
}

// form
func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err : %v", err)
	}
	fmt.Fprintf(w, "POST request successful")

	name,password := r.FormValue("name"), r.FormValue("password")
	fmt.Fprintf(w, "Name = %s\n", name)
	fmt.Fprintf(w, "Password = %s\n", password)

}