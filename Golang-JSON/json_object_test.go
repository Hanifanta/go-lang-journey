package Golang_JSON

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Address struct {
	Street     string
	Country    string
	PostalCode string
}

type Customer struct {
	FirstName  string
	MiddleName string
	LastName   string
	Age        int
	Married    bool
	Hobbies    []string
	Addresses  []Address
}

func TestJSONObj(t *testing.T) {
	customer := Customer{
		FirstName:  "Hanif",
		MiddleName: "Al",
		LastName:   "Irsyad",
		Age:        23,
		Married:    false,
		Hobbies:    []string{"Coding", "Gaming", "Reading"},
		Addresses: []Address{
			{
				Street:     "Jl. Raya Cipadung",
				Country:    "Indonesia",
				PostalCode: "99999",
			},
			{
				Street:     "Jl. Raya Cibiru",
				Country:    "Indonesia",
				PostalCode: "88888",
			},
		},
	}

	bytes, _ := json.Marshal(customer)
	fmt.Println(string(bytes))
}
