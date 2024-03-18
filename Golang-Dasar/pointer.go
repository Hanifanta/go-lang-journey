package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	// Pass by value
	// address1 := Address{"Subang", "Jawa Barat", "Indonesia"}
	// address2 := address1 // copy by value

	// address2.City = "Bandung"

	// fmt.Println(address1)
	// fmt.Println(address2)

	// Pass by reference
	address1 := Address{"Subang", "Jawa Barat", "Indonesia"}
	address2 := &address1 // copy by value

	address2.City = "Bandung"

	fmt.Println(address1)
	fmt.Println(address2)

}
