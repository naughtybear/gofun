package main

import (
	"fmt"
)

type Product struct {
	name  string
	price float64
}

func (product Product) String() string {
	return fmt.Sprintf("%s (%.2f)", product.name, product.price)
}

func main() {
	products := []*Product{{"Spanner", 3.99}, {"Wrench", 2.48}, {"Screwdriver", 1.99}}
	fmt.Println(products)

	for _, product := range products {
		product.price *= 2
	}
	fmt.Println(products)

	// modify slice
	s := []string{"A", "B", "C", "D", "E", "F", "G"}
	t := []string{"K", "L", "M", "N"}
	u := []string{"m", "n", "o", "p", "q", "r"}

	s = append(s, "h", "i", "j")
	s = append(s, t...)
	s = append(s, u[2:5]...)
	b := []byte{'U', 'V'}
	letters := "WXY"
	b = append(b, letters...) // append a string
	fmt.Printf("%v\n%s\n", s, b)

	fmt.Println()

	// copy demo
	fmt.Println(products)
	prodCopy := make([]*Product, 2)
	copy(prodCopy, products)
	fmt.Println(prodCopy)
	prodCopy[0].price = 0.99
	fmt.Println(prodCopy)
	fmt.Println(products) // shallow copy
}
