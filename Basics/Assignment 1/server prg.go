package main

import (
	"fmt"
	"net/http"
	"sync"
)

func main() {
	
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello world")
	})

	http.HandleFunc("/new", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello New world")
	})

	http.ListenAndServe(":9000", nil)
}
