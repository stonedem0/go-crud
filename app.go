package main 

import (
	"fmt"
	"log"
	"net/http"

)

//handling GET and POST requests

func hello(w http.ResponseWriter, r *http.Request) {
    if r.URL.Path != "/" {
        http.Error(w, "404 not found.", http.StatusNotFound)
        return
	}
	switch r.Method {
		case "GET":     
		    fmt.Println("this is get")
			http.ServeFile(w, r, "home.html")
		case "POST": 
			fmt.Println("this is post")
			err := r.ParseForm()
			if err != nil {
				fmt.Fprintf(w, "Something wrong");
			}
		
			name := r.FormValue("username")
			fmt.Fprintf(w, "Hello, %s!", name)	
		
		default:
			fmt.Fprintf(w, "Sorry, only GET and POST methods are supported.")
		}
	}



func main() {

	http.HandleFunc("/", hello)
	fmt.Println("hi, noob! listening on 3000")

	if err := http.ListenAndServe(":3000", nil); err != nil {
		log.Fatal(err)
	}
}