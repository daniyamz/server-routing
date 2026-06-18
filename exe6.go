package main

import (
	"fmt"
	"net/http"
)

func DashboardHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/dashboard" {
		http.Error(w, "404", http.StatusNotFound)
	}
	API := r.Header.Get("X-API-Key")
	handcodeded := "secret123"

	if API != handcodeded {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	} else {
		fmt.Fprintln(w, "Welcome")
	}
}
