package main

import (
	"fmt"
	"github.com/gorilla/schema"
)

type Person struct {
	Names []string
	Phone string
}

func main() {
	values := map[string][]string{
		"Names": {"John", "Smith"},
		"Phone": {"999-999-999"},
	}
	person := new(Person)
	decoder := schema.NewDecoder()
	decoder.Decode(person, values)

	fmt.Println(person)
}
