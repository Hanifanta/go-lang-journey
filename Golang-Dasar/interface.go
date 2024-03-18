package main

import "fmt"

type HasName interface {
	GetName() string
}

func SayHello(value HasName) {
	fmt.Println("Hello", value.GetName())
}

type Person struct {
	Name string
}

type Animal struct {
	Name string
}

func (person Person) GetName() string {
	return person.Name
}

func (animal Animal) GetName() string {
	return animal.Name
}

func main() {
	person := Person{Name: "Hanif"}
	SayHello(person)

	animal := Animal{Name: "Cat"}
	SayHello(animal)

}
