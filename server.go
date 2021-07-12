package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func runServer() {
	log.Println("Starting server...")
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}