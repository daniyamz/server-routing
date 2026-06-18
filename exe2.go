package main

import (
	"fmt"
	"net/http"
)

func NameHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
	name := "Guest!"

	str := r.URL.Query().Get("name")
	if str != "" {
		name = str
	}
	fmt.Fprintf(w, "Hello, %v!", name)
}
