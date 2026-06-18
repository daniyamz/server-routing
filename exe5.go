package main

import (
	"fmt"
	"net/http"
)

func AgentHander(w http.ResponseWriter, r *http.Request) {
	Agnet := r.Header.Get("User-Agent")
	fmt.Fprintln(w, Agnet)
}
