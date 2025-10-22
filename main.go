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

func entryPoint(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "helloooo")
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", entryPoint)
	mux.HandleFunc("/hello/", helloHandler)
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, "OK")
	})

	log.Println("Server on :8080")
	log.Fatal(http.ListenAndServe(":8080", loggingMiddleware(mux)))
}
