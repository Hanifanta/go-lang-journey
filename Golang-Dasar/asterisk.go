package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	address1 := Address{"Subang", "Jawa Barat", "Indonesia"}
	address2 := &address1

	address2.City = "Bandung"

	fmt.Println(address1)
	fmt.Println(address2)

	// address2 = &Address{"Malang", "Jawa Timur", "Indonesia"}
	*address2 = Address{"Surabaya", "Jawa Timur", "Indonesia"}
	fmt.Println(address1)
	fmt.Println(address2)
}
