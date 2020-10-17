//This Go program will create a http server and provide some api calls
package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

//main function declaration
func main() {
	//init the mux routers
	r := mux.NewRouter()

	//Route handler / Enpoints
	r.HandleFunc("/api/books", getBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	r.HandleFunc("/api/books", createBooks).Methods("POST")
	r.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	r.HandleFunc("/api/books/{id}", deleteBooks).Methods("POST")

	//run the server
	log.Fatal(http.ListenAndServe(":8000", r))

	fmt.Println("Hello World!")
}
