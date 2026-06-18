package main

import (
	"fmt"
	"net/http"
)

func PingHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	if r.URL.Path != "/ping" {
		http.Error(w, "path not allowed", http.StatusNotFound)
		return
	}
	fmt.Fprintf(w, "pong")
}
func main() {
	http.HandleFunc("/ping", PingHandler)
	http.HandleFunc("/hello", NameHandler)
	http.HandleFunc("/count", CountHander)
	http.HandleFunc("/calculate", CalculateHandler)
	http.HandleFunc("/agent", AgentHander)
	http.HandleFunc("/dashboard", DashboardHandler)
	http.HandleFunc("/legacy", RedirectorHandler)
	http.HandleFunc("/v2", RedirectorHandler)
	http.ListenAndServe(":8080", nil)
}
