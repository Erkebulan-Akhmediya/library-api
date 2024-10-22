package book

import (
	"encoding/json"
	"example.com/library/utils"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

type Controller struct {
	service Service
}

func Router(router *mux.Router) {
	r := router.PathPrefix("/book").Subrouter()
	r.HandleFunc("", Controller{}.post)
	r.HandleFunc("/all", Controller{}.getAll)
}

func (c Controller) post(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	var book Book
	err := utils.ParseRequest(r, &book)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error parsing request: %s", err.Error()), http.StatusBadRequest)
		return
	}

	err = c.service.post(book)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error posting book: %s", err.Error()), http.StatusInternalServerError)
		return
	}
}

func (c Controller) getAll(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	books, err := c.service.getAll()
	if err != nil {
		http.Error(w, fmt.Sprintf("Error getting books: %s", err.Error()), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(books)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error encoding books: %s", err.Error()), http.StatusInternalServerError)
		return
	}
}
