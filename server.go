package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func addUserHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		reqBody, _ := ioutil.ReadAll(r.Body)
		fmt.Println(string(reqBody))
	default:
		fmt.Fprintf(w, "Sorry, only POST method is supported.")
	}
}


func runServer() {
	log.Println("Starting server...")
	http.HandleFunc("/", handler)
	http.HandleFunc("/addUser", addUserHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}