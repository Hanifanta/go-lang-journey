package main

import "fmt"

func main() {
	name := "Hanif"

	switch name {
	case "Hanif":
		fmt.Println("Hai Hanif!")
	case "Hanif Al":
		fmt.Println("Hai Hanif Al!")
	default:
		fmt.Println("Hai boleh kenalan?")
	}

	switch length := len(name); length > 5 {
	case true:
		fmt.Println("Nama terlalu panjang")
	case false:
		fmt.Println("Nama sudah benar")
	}

	length := len(name)
	switch {
	case length > 10:
		fmt.Println("Nama terlalu panjang")
	case length > 5:
		fmt.Println("Dikit lagi benar")
	default:
		fmt.Println("Nama sudah benar")
	}
}
