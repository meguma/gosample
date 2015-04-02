package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Person2 struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Email   string `json:"-"`
	Age     int    `json:"age"`
	Address string `json:"address,omitempty"`
	memo    string
}

func main() {
	person2 := &Person2{
		ID:      1,
		Name:    "Gopher",
		Email:   "gopher@example.org",
		Age:     5,
		Address: "",
		memo:    "golang lover",
	}

	b, err := json.Marshal(person2)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(b))
}
