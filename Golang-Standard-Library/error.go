package main

import (
	"errors"
	"fmt"
)

var (
	ValidationError = errors.New("validation error")
	NotFoundError   = errors.New("not found")
)

func GetById(id string) error {
	if id == "" {
		return ValidationError
	}
	if id != "Hanif" {
		return NotFoundError
	}
	return nil
}

func main() {
	err := GetById("Hanif")
	if err != nil {
		if errors.Is(err, ValidationError) {
			fmt.Println("validation error")
		} else if errors.Is(err, NotFoundError) {
			fmt.Println("not found error")
		} else {
			fmt.Println("Unknown error")
		}
	}
}
