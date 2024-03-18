package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

// Struct Method
func (customer Customer) sayHello(name string) {
	fmt.Println("Hello", name, "My Name is", customer.Name)
}

func main() {
	var hanif Customer
	hanif.Name = "Hanif Al Irsyad"
	hanif.Address = "Indonesia"
	hanif.Age = 23

	fmt.Println(hanif)
	fmt.Println(hanif.Name)
	fmt.Println(hanif.Address)
	fmt.Println(hanif.Age)

	// Struct Literal
	joko := Customer{
		Name:    "Joko",
		Address: "Indonesia",
		Age:     23,
	}
	fmt.Println(joko)

	budi := Customer{"Budi", "Indonesia", 23}
	fmt.Println(budi)

	// Struct Method
	budi.sayHello("Hanif")
}
