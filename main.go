package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimPrefix(r.URL.Path, "/hello/")
	if path == "" {
		fmt.Fprintln(w, "Hello, Stranger")
		return
	}
	fmt.Fprintf(w, "Hello, %s!\n", path)
}

func main() {
	http.HandleFunc("/hello/", helloHandler)
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, "OK")
	})

	log.Println("Server on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
