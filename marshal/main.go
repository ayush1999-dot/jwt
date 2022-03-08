package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type person struct {
	Name string
}

func main() {
	p1 := person{Name: "ayush"}
	p2 := person{Name: "shetty"}

	TP := []person{p1, p2}
	mmarshal(TP)

}

func mmarshal(TP []person) {
	bs, err := json.Marshal(TP)
	if err != nil {
		log.Panic(err)
	}
	fmt.Println(string(bs))
	unmarshall(bs)
}

func unmarshall(bs []byte) {
	var tp []person
	err := json.Unmarshal(bs, &tp)
	if err != nil {
		log.Panic(err)
	}
	fmt.Println(tp)
}
