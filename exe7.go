package main

import (
	"fmt"
	"net/http"
)

func RedirectorHandler(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path == "/legacy" {
		http.Redirect(w, r, "/v2", 301)
	}
	if r.URL.Path == "/v2" {
		fmt.Fprintln(w, "Welcome to version 2")
		return
	}
}
