package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	alamat1 := new(Address)
	alamat2 := alamat1

	alamat2.City = "Bandung"

	fmt.Println(alamat1)
	fmt.Println(alamat2)
}
