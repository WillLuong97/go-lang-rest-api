//This Go program will create a http server and provide some api calls
package main

//import new packages into the array
import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

//Book struct (Model): (similar to class in Python and Java)
type Book struct {
	ID     string  `json:"id"`
	Isbn   string  `json:"isbn"`
	Title  string  `json:"title"`
	Author *Author `json:"author"`
}

//Init books variablea as a slice Book struct
var books []Book

//Author struct
type Author struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

//Get all books:
func getBooks(w http.ResponseWriter, r *http.Request) {
	//setting the header with a content type
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

//Get single book:
func getBook(w http.ResponseWriter, r *http.Request) {
	//setting the header with a content type
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) // get the params
	//Loop through the books and find with id

	//range is used to loop a range or slice of data
	for _, item := range books {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}

	}
	//output the list of books that we found from the id
	json.NewEncoder(w).Encode(&Book{})
}

//create a new book
func createBooks(w http.ResponseWriter, r *http.Request) {
	//setting the header with a content type
	w.Header().Set("Content-Type", "application/json")
	var book Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	book.ID = strconv.Itoa(rand.Intn(10000000)) //Mock ID - not safe
	books = append(books, book)
	//output the book found from ID
	json.NewEncoder(w).Encode(book)

}

//Update the book
func updateBook(w http.ResponseWriter, r *http.Request) {
	//setting the header with a content type
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	//find the book to update based on the ID
	for index, item := range books {
		if item.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)
			var book Book
			_ = json.NewDecoder(r.Body).Decode(&book)
			book.ID = params["id"] //set the updated book ID to the id of the parameter
			books = append(books, book)
			//output the book found from ID
			json.NewEncoder(w).Encode(book)
			return
		}
	}
	//output the book found from ID
	json.NewEncoder(w).Encode(books)

}

//Delete the book
func deleteBooks(w http.ResponseWriter, r *http.Request) {
	//setting the header with a content type
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	//find the book to update based on the ID
	for index, item := range books {
		if item.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)
			break
		}
	}
	//output the book found from ID
	json.NewEncoder(w).Encode(books)

}

//main function declaration
func main() {
	//init the mux routers
	r := mux.NewRouter()

	//Mock data - @todo - implement database
	books = append(books, Book{ID: "1", Isbn: "44873", Title: "Book One", Author: &Author{Firstname: "John", Lastname: "Doe"}})
	books = append(books, Book{ID: "2", Isbn: "545466", Title: "Book Two", Author: &Author{Firstname: "Steve", Lastname: "Smith"}})
	books = append(books, Book{ID: "3", Isbn: "56454", Title: "Book Three", Author: &Author{Firstname: "David", Lastname: "Doe"}})
	books = append(books, Book{ID: "4", Isbn: "515126", Title: "Book Four", Author: &Author{Firstname: "Linda", Lastname: "Doe"}})

	//Route handler(Enpoints for our API)
	r.HandleFunc("/api/books", getBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	r.HandleFunc("/api/books", createBooks).Methods("POST")
	r.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	r.HandleFunc("/api/books/{id}", deleteBooks).Methods("DELETE")

	//run the server
	log.Fatal(http.ListenAndServe(":8000", r))

	fmt.Println("Hello World!")
}
