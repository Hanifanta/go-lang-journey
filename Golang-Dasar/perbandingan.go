package main

import "fmt"

func main() {
	// Operasi Perbandingan
	// >, <, >=, <=, ==, !=
	var (
		nama1 = "Hanif"
		nama2 = "Hanif"
	)

	var result bool = nama1 == nama2

	fmt.Println(result)

	var result2 bool = nama1 != nama2

	fmt.Println(result2)

}
