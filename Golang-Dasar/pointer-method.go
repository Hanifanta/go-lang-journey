package main

import "fmt"

type Man struct {
	Name string
}

func (man *Man) Married() {
	man.Name = "Mr. " + man.Name
}

func main() {
	hanif := Man{"Hanif"}
	hanif.Married()

	fmt.Printf(hanif.Name)

}
