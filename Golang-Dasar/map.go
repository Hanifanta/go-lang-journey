package main

import "fmt"

func main() {
	person := map[string]string{
		"names":   "Hanif Al Irsyad",
		"address": "DIY",
	}

	fmt.Println(person["names"])
	fmt.Println(person["address"])
	fmt.Println(person)

	// Function Map
	// len(map), map[key], map[key]=value, make(map[typekey]typevalue), delete(map,key)

	book := make(map[string]string)
	book["title"] = "Buku Golang"
	book["author"] = "Unknown"
	book["ups"] = "False"

	fmt.Println(book)

	delete(book, "ups")

	fmt.Println(book)
}
