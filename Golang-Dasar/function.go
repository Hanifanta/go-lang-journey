package main

import "fmt"

// Function
// func sayHello() {
// 	fmt.Println("Hello")
// }

// Function Parameter
func sayHelloTo(firstName string, lastName string) {
	fmt.Println("Hello", firstName, lastName)
}

// Function Return Value
func getHello(name string) string {
	helo := "Hello " + name
	return helo
}

func main() {
	sayHelloTo("Hanif", "Al Irsyad")
	sayHelloTo("Budi", "Joko")

	result := getHello("Hanif")
	fmt.Println(result)
}
