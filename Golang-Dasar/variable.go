package main

import "fmt"

func main() {
	// Variable dengan deklarasi tipe data
	// var name string

	// name = "Hanif Al Irsyad"
	// fmt.Println(name)

	// Variable tanpa deklarasi tipe data menggunakan "VAR"
	// var name = "Hanif Al Irsyad"
	// fmt.Println(name)

	// Variable menggunakan := tanpa var dan hanya boleh dilakukan pada variable pertama
	// name := "Hanif Al Irsyad"
	// fmt.Println(name)

	// name = "Hanif Al Irsyad tanpai :="
	// fmt.Println(name)

	// Multiple Variable
	var (
		firstName = "Hanif"
		lastName  = "Al Irsyad"
	)

	fmt.Println(firstName)
	fmt.Println(lastName)
}
