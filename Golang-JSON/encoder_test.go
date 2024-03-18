package Golang_JSON

import (
	"encoding/json"
	"os"
	"testing"
)

func TestEncoder(t *testing.T) {
	writer, _ := os.Create("customer.json")

	encoder := json.NewEncoder(writer)

	customer := Customer{
		FirstName:  "Hanif",
		MiddleName: "Al",
		LastName:   "Irsyad",
	}

	encoder.Encode(customer)
}
