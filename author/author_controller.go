package author

import (
	"encoding/json"
	"example.com/library/utils"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func Router(router *mux.Router) {
	r := router.PathPrefix("/author").Subrouter()
	r.HandleFunc("", Controller{}.post)
	r.HandleFunc("/all", Controller{}.getAll)
	r.HandleFunc("/{id}", Controller{}.authorId)
}

type Controller struct {
	service Service
}

func (c Controller) getAll(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	authors, err := c.service.getAll()
	if err != nil {
		http.Error(w, fmt.Sprintf("Error getting authors: %s", err.Error()), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(authors)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error encoding response: %s", err.Error()), http.StatusInternalServerError)
		return
	}
}

func (c Controller) post(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	var author Author
	err := utils.ParseRequest(r, &author)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error parsing request: %s", err.Error()), http.StatusBadRequest)
		return
	}

	err = c.service.post(author)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error posting author: %s", err.Error()), http.StatusInternalServerError)
		return
	}
}

func (c Controller) authorId(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, fmt.Sprintf("Error parsing author id: %s", err.Error()), http.StatusBadRequest)
		return
	}

	if r.Method == "PUT" {
		var author Author
		err := utils.ParseRequest(r, &author)
		if err != nil {
			http.Error(w, fmt.Sprintf("Error parsing request: %s", err.Error()), http.StatusBadRequest)
			return
		}
		err = c.service.update(id, author)
		if err != nil {
			http.Error(w, fmt.Sprintf("Error updating author: %s", err.Error()), http.StatusInternalServerError)
			return
		}
	} else if r.Method == "DELETE" {
		err := c.service.delete(id)
		if err != nil {
			http.Error(w, fmt.Sprintf("Error deleting author: %s", err.Error()), http.StatusInternalServerError)
			return
		}
	} else if r.Method == "GET" {
		author, err := c.service.GetById(id)
		if err != nil {
			http.Error(w, fmt.Sprintf("Error getting author: %s", err.Error()), http.StatusNotFound)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		err = json.NewEncoder(w).Encode(author)
		if err != nil {
			http.Error(w, fmt.Sprintf("Error encoding response: %s", err.Error()), http.StatusInternalServerError)
			return
		}
	} else {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}
}
