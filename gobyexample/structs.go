package main

import "fmt"

type person struct {
	name string
	age  int
}

func changePerson(p person) {
	p.age = 10
	p.name = "lisi"
}

func changePersonPtr(p *person) {
	p.age = 10
	p.name = "lisi"
}

func main() {
	fmt.Println(person{"Bob", 3})
	fmt.Println(person{name: "Fred"})
	fmt.Println(&person{name: "Ann", age: 10})

	s := person{"Sean", 50}
	fmt.Println(s.name)

	sp := &s
	fmt.Println(sp.age) // automatically dereferenced

	sp.age = 29
	fmt.Println(sp)

	fmt.Println(s)
	changePerson(s)
	fmt.Println(sp)

	fmt.Println(sp)
	changePersonPtr(sp)
	fmt.Println(sp)
}
