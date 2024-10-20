package utils

import (
	"database/sql"
	"encoding/json"
	_ "github.com/lib/pq"
	"io/ioutil"
	"net/http"
)

func ExecuteSql(query string) (*sql.Rows, error) {
	connStr := "user=postgres password=123456 database=postgres port=2101 sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	defer db.Close()

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	return rows, nil
}

func ParseRequest(r *http.Request, res interface{}) error {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}

	defer r.Body.Close()

	err = json.Unmarshal(body, res)
	if err != nil {
		return err
	}

	return nil
}
