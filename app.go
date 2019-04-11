package main 

import (
	"fmt"
	"log"
	"net/http"
	_ "github.com/mattn/go-sqlite3"
	
)

//handling GET and POST requests

func hello(db *Database) http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request) {
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
			
				username := r.FormValue("username")
				err = db.AddUser(username)
				if err != nil {
					log.Printf("Error adding user: %s", err)
					http.Error(w, "Failed to add user.", http.StatusInternalServerError)
					return
				}
			
				log.Printf("Added user %q", username)
				http.Redirect(w, r, "/", 302)
			default:
				fmt.Fprintf(w, "Sorry, only GET and POST methods are supported.")
		}
	}
}

func main() {
	
	db := &Database{}

	if err := db.Setup(); err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/", hello(db))
	
	fmt.Println("hi, noob! listening on 3000")

	if err := http.ListenAndServe(":3000", nil); err != nil {
		log.Fatal(err)
	}
}
