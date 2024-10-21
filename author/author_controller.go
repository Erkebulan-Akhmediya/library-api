package author

import (
	"encoding/json"
	"example.com/library/utils"
	"fmt"
	"net/http"
)

func Router() {
	http.HandleFunc("/author/all", Controller{}.getAll)
	http.HandleFunc("/author", Controller{}.post)
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
