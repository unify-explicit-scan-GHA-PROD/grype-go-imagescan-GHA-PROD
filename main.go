package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello from Go on Docker with Alpine!")
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Starting server on :8080...")
	http.ListenAndServe(":8080", nil)
}
