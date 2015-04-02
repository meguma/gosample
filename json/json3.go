package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Person3 struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Email   string `json:"-"`
	Age     int    `json:"age"`
	Address string `json:"address,omitempty"`
	memo    string
}

func main() {
	var person3 Person3

	b := []byte(`{"id":1,"name":"Gopher","age":5}`)
	err := json.Unmarshal(b, &person3)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(person3) // {1 Gopher  5  }
}
