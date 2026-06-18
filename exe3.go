package main

import (
	"fmt"
	"io"
	"net/http"
)

func CountHander(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		fmt.Fprintln(w, "Send a POST request with text to count words")
		return
	}
	if r.Method == http.MethodPost {
		count, err := io.ReadAll(r.Body)

		if err != nil {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
		fmt.Fprintf(w, "Count %v", len(count))
	}
}
