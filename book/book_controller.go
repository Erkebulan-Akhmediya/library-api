package book

import (
	"encoding/json"
	"example.com/library/utils"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type Controller struct {
	service Service
}

func Router(router *mux.Router) {
	r := router.PathPrefix("/book").Subrouter()
	r.HandleFunc("", Controller{}.post)
	r.HandleFunc("/all", Controller{}.getAll)
	r.HandleFunc("/{id}", Controller{}.bookId)
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

func (c Controller) bookId(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, fmt.Sprintf("Error parsing book id: %s", err.Error()), http.StatusBadRequest)
		return
	}

	if r.Method == "GET" {
		book, err := c.service.getById(id)
		if err != nil {
			http.Error(w, fmt.Sprintf("Error getting book: %s", err.Error()), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		err = json.NewEncoder(w).Encode(book)
		if err != nil {
			http.Error(w, fmt.Sprintf("Error encoding book: %s", err.Error()), http.StatusInternalServerError)
			return
		}
	} else if r.Method == "PUT" {
		var book Book
		err := utils.ParseRequest(r, &book)
		if err != nil {
			http.Error(w, fmt.Sprintf("Error parsing request: %s", err.Error()), http.StatusBadRequest)
			return
		}

		err = c.service.update(id, book)
		if err != nil {
			http.Error(w, fmt.Sprintf("Error updating book: %s", err.Error()), http.StatusInternalServerError)
			return
		}
	} else if r.Method == "DELETE" {
		err := c.service.delete(id)
		if err != nil {
			http.Error(w, fmt.Sprintf("Error deleting book: %s", err.Error()), http.StatusInternalServerError)
			return
		}
	} else {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}
}
