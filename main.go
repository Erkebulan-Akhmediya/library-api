package main

import (
	"example.com/library/author"
	"example.com/library/book"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	author.Router(r)
	book.Router(r)
	http.Handle("/", r)
	err := http.ListenAndServe(":2100", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
