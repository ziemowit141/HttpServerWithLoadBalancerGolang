package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Person struct {
	ID int `json:"ID"`
	Name string `json:"Name"`
	EmailAddress string `json:"EmailAddress"`
	Phone string `json:"Phone"`
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func addUserHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		reqBody, _ := ioutil.ReadAll(r.Body)
		var person Person
		json.Unmarshal(reqBody, &person)
		fmt.Println(person.ID, person.Name, person.EmailAddress, person.Phone)
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