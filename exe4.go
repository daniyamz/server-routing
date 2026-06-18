package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func CalculateHandler(w http.ResponseWriter, r *http.Request) {
	op := r.URL.Query().Get("op")
	a := r.URL.Query().Get("a")
	b := r.URL.Query().Get("b")

	a_int, err := strconv.Atoi(a)
	if err != nil {
		http.Error(w, "400 Bad Request", http.StatusBadRequest)
		return
	}
	b_int, err := strconv.Atoi(b)
	if err != nil {
		http.Error(w, "400 Bad Request", http.StatusBadRequest)
		return
	}
	result := 0
	switch op {
	case "add":
		result = a_int + b_int
		fmt.Fprintln(w, result)
	case "subtract":
		result = a_int - b_int
		fmt.Fprintln(w, result)
	case "multiply":
		result = a_int * b_int
		fmt.Fprintln(w, result)
	}
}
