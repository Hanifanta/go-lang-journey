package main

import "fmt"

type Filter func(string) string

func sayHello(name string, filter Filter) {
	filteredName := filter(name)
	fmt.Println("Hello", filteredName)
}

func spamFilter(name string) string {
	if name == "Anjing" {
		return "..."
	} else {
		return name
	}
}

func main() {
	sayHello("Hanif", spamFilter)

	// atau
	filter := spamFilter
	fmt.Println("Anjing", Filter)
}
