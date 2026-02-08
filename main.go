package main

import (
	//"errors"
	"fmt"
	//"io"
	//"log"
	"net/http"
	//"os"
)

func  getHealth(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte("Service is running"))
}

func getHello(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusBadRequest)
	w.Write([]byte("Error: name parameter is required"))
	return
	}

	fmt.Fprintf(w, "Hello, %s!\n", name)
}
func getRoot(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	fmt.Fprintf(w, "Welcome to the API!\n")
}

func main() {
	
	http.HandleFunc("/health", getHealth)
	http.HandleFunc("/hello", getHello)

	fmt.Println("Server starting on http://localhost:8080/")
	http.ListenAndServe(":8080", nil)
}
