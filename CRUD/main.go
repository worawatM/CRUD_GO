package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Book struct {
	ID     int     `json: "id"`
	Title  string  `json: "title"`
	Author *Author `json:	"author"`
}
type Author struct {
	Name     string `json: "name"`
	Lastname string `json:	"lastname"`
}

var books []Book

func getAllBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

func createBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var book Book

	json.NewDecoder(r.Body).Decode(&book)
	books = append(books, book)
	json.NewEncoder(w).Encode(book)
}

func getBooks(w http.ResponseWriter, r *http.Request) {
	// var idParam string = mux.Vars(r)["id"]
	idParam := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idParam)
	id = id - 1
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte("ID could not be convert to Integer"))
		return
	}
	if id > len(books) {
		w.WriteHeader(404)
		w.Write([]byte("Don't have This Book in Data"))
		return
	}
	book := books[id]
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(book)

}

func updateBooks(w http.ResponseWriter, r *http.Request) {
	// var idParam string = mux.Vars(r)["id"]
	idParam := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idParam)
	id = id - 1
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte("ID could not be convert to Integer"))
		return
	}
	if id > len(books) {
		w.WriteHeader(404)
		w.Write([]byte("Don't have This Book in Data"))
		return
	}
	// Path update new data
	var upBook Book
	json.NewDecoder(r.Body).Decode(&upBook)
	books[id] = upBook
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(upBook)

}

// Have Some Bug in Function Delete
func deleteBooks(w http.ResponseWriter, r *http.Request) {
	// var idParam string = mux.Vars(r)["id"]
	idParam := mux.Vars(r)
	id, err := strconv.Atoi(idParam["id"])
	if err != nil {
		http.Error(w, "Invalid User ID", http.StatusBadRequest)
		return
	}
	for i, book := range books {
		if book.ID == id {
			books = append(books[:i], books[i+1:]...)
			w.WriteHeader(http.StatusOK)
			return
		}
	}
	http.Error(w, "User not Found", http.StatusNotFound)
}
func main() {
	r := mux.NewRouter()
	books = append(books, Book{ID: 1, Title: "Book One", Author: &Author{Name: "Worawat", Lastname: "Mongkholbut"}})
	books = append(books, Book{ID: 2, Title: "My Fail Love Story", Author: &Author{Name: "Worawat", Lastname: "Mongkholbut"}})

	//Get All Book
	r.HandleFunc("/api/books", getAllBooks).Methods("GET")

	// Get Book By id
	r.HandleFunc("/api/books/{id}", getBooks).Methods("GET")

	// Create
	r.HandleFunc("/api/books", createBooks).Methods("POST")

	// Update
	r.HandleFunc("/api/books/{id}", updateBooks).Methods("PUT")

	// Delete
	r.HandleFunc("/api/books/{id}", deleteBooks).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", r))
}
