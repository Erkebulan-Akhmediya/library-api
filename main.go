package main

import (
	"example.com/library/author"
	"fmt"
	"net/http"
)

func main() {
	author.Router()
	err := http.ListenAndServe(":2100", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
