package author

import (
	"encoding/json"
	"example.com/library/utils"
	"fmt"
	"net/http"
)

func Router() {
	http.HandleFunc("/author/all", getAll)
	http.HandleFunc("/author", post)
}

func getAll(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	rows, err := utils.ExecuteSql("select * from author")
	if err != nil {
		http.Error(w, fmt.Sprintf("Error executing query: %s", err.Error()), http.StatusInternalServerError)
		return
	}

	var authors []Author
	for rows.Next() {
		var author Author
		err = rows.Scan(&author.Id, &author.LastName, &author.FirstName)
		if err != nil {
			http.Error(w, fmt.Sprintf("Error scanning row: %s", err.Error()), http.StatusInternalServerError)
			return
		}
		authors = append(authors, author)
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(authors)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error encoding response: %s", err.Error()), http.StatusInternalServerError)
		return
	}
}

func post(w http.ResponseWriter, r *http.Request) {
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

	query := fmt.Sprintf("insert into author (last_name, first_name) values ('%s', '%s')", author.LastName, author.FirstName)
	_, err = utils.ExecuteSql(query)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error executing query: %s", err.Error()), http.StatusInternalServerError)
		return
	}
}
