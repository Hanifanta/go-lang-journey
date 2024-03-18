package main

import "fmt"

func main() {
	// counter := 1

	// for counter <= 10 {
	// 	fmt.Println("Perulangan ke-", counter)
	// 	counter++
	// }

	for counter := 0; counter <= 10; counter++ {
		fmt.Println("Perulangan ke-", counter)
	}

	fmt.Println("Selesai")

	names := []string{"Hanif", "Al", "Irsyad"}

	// for i := 0; i < len(names); i++ {
	// 	fmt.Println(names[i])
	// }

	for i, name := range names {
		fmt.Println("Index", i, "=", name)
	}

	for _, name := range names {
		fmt.Println(name)
	}
}
