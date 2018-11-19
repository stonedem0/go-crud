package main 

import (
	"fmt"
	"log"
	"net/http"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"

)

func hello(w http.ResponseWriter, r *http.Request) {
    if r.URL.Path != "/" {
        http.Error(w, "404 not found.", http.StatusNotFound)
        return
	}
	switch r.Method {
		case "GET":     
		    fmt.Println("this is get")
			http.ServeFile(w, r, "random.html")
		case "POST": 
			fmt.Println("this is post")
			err := r.ParseForm()
			if err != nil {
				fmt.Fprintf(w, "no idea what but something wrong");
			}
		
			name := r.FormValue("username")
			fmt.Fprintf(w, "Hello, %s!", name)	
		
		default:
			fmt.Fprintf(w, "Sorry, only GET and POST methods are supported.")
		}
	}



func main() {
	db, err := sql.Open("mysql",
		"user:password@tcp(127.0.0.1:3306)/hello")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	http.HandleFunc("/", hello)
	fmt.Println("hi, noob! listening on 3000")

	if err := http.ListenAndServe(":3000", nil); err != nil {
		log.Fatal(err)
	}
// 	db, err := sql.Open("mysql", "theUser:thePassword@/theDbName")
//     if err != nil {
//       panic(err)
//  }
} 