package main

import "fmt"

func main() {
	var names [3]string

	names[0] = "Hanif"
	names[1] = "Al"
	names[2] = "Irsyad"

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

	var values = [...]int{
		90,
		20,
		54,
		140,
		141,
	}
	fmt.Println(values)
	fmt.Println(values[0])
	fmt.Println(values[1])
	fmt.Println(values[2])

	// Function Array
	fmt.Println(len(values))
}
