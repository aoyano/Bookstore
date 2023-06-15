package microservice

import (
	"encoding/json"
	"log"
	"net/http"
)

type Book struct {
	ID       string  `json:"id"`
	Title    string  `json:"title"`
	Author   string  `json:"author"`
	Price    float64 `json:"price"`
	Quantity int     `json:"quantity"`
}

var books []Book

func main() {
	// Set up some initial data
	books = []Book{
		{ID: "1", Title: "Book 1", Author: "Author 1", Price: 10.99, Quantity: 5},
		{ID: "2", Title: "Book 2", Author: "Author 2", Price: 12.99, Quantity: 3},
		{ID: "3", Title: "Book 3", Author: "Author 3", Price: 9.99, Quantity: 8},
	}

	// Define the HTTP endpoints
	http.HandleFunc("/books", getBooksHandler)
	http.HandleFunc("/books/{id}", getBookHandler)

	// Start the server
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func getBooksHandler(w http.ResponseWriter, r *http.Request) {
	// Convert books slice to JSON
	jsonBytes, err := json.Marshal(books)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Set the Content-Type header and write the response
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonBytes)
}

func getBookHandler(w http.ResponseWriter, r *http.Request) {
	// Get the book ID from the URL path parameter
	id := r.URL.Query().Get("id")

	// Find the book with the given ID
	var book *Book
	for _, b := range books {
		if b.ID == id {
			book = &b
			break
		}
	}

	if book == nil {
		http.Error(w, "Book not found", http.StatusNotFound)
		return
	}

	// Convert book to JSON
	jsonBytes, err := json.Marshal(book)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Set the Content-Type header and write the response
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonBytes)
}
