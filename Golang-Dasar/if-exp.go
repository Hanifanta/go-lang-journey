package main

import "fmt"

func main() {
	name := "Hanif Al Irsyad"

	if name == "Hanif" {
		fmt.Println("Hello Hanif!")
	} else if name == "Hanif Al" {
		fmt.Println("Hello Hanif Al!")
	} else if name == "Hanif Al Irsyad" {
		fmt.Println("Hello Hanif Al Irsyad!")
	} else {
		fmt.Println("Maaf!")
	}

	if length := len(name); length > 5 {
		fmt.Println("Nama terlalu panjang")
	} else {
		fmt.Println("Nama sudah benar")
	}
}
