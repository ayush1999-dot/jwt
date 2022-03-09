package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type person struct {
	First string
}

func main() {

	http.HandleFunc("/bar", bar)
	http.HandleFunc("/foo", foo)
	http.ListenAndServe(":8080", nil)
}
func bar(wr http.ResponseWriter, r *http.Request) {

	p := person{First: "ayush"}
	json.NewEncoder(wr).Encode(&p)
}
func foo(wr http.ResponseWriter, r *http.Request) {
	var p person
	json.NewDecoder(r.Body).Decode(&p)
	fmt.Println(p)
}
