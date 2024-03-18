package main

import "fmt"

func getFullName() (firstName, middleName, lastName string) {
	firstName = "Hanif"
	middleName = "Al"
	lastName = "Irsyad"

	return firstName, middleName, lastName
}

func main() {
	a, b, c := getFullName()
	fmt.Println(a, b, c)
}
