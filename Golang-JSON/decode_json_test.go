package Golang_JSON

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestDecodeJSON(t *testing.T) {
	jsonString := `{"FirstName":"Hanif","MiddleName":"Al","LastName":"Irsyad","Age":23,"Married":false,"Hobbies":["Coding","Gaming","Reading"],"Addresses":[{"Street":"Jl. Raya Cipadung","Country":"Indonesia","PostalCode":"99999"},{"Street":"Jl. Raya Cibiru","Country":"Indonesia","PostalCode":"88888"}]}`
	jsonBytes := []byte(jsonString)

	customer := &Customer{}

	err := json.Unmarshal(jsonBytes, customer)
	if err != nil {
		panic(err)
	}

	fmt.Println(customer)
	fmt.Println(customer.FirstName)
	fmt.Println(customer.MiddleName)
	fmt.Println(customer.LastName)
	fmt.Println(customer.Age)
	fmt.Println(customer.Married)
	fmt.Println(customer.Hobbies)
	fmt.Println(customer.Addresses)
	fmt.Println(customer.Addresses[0])
	fmt.Println(customer.Addresses[1])
}

func TestOnlyJSONArray(t *testing.T) {
	jsonString := `[{"Street":"Jl. Raya Cipadung","Country":"Indonesia","PostalCode":"99999"},{"Street":"Jl. Raya Cibiru","Country":"Indonesia","PostalCode":"88888"}]`
	jsonBytes := []byte(jsonString)

	addresses := &[]Address{}

	err := json.Unmarshal(jsonBytes, addresses)
	if err != nil {
		panic(err)
	}

	fmt.Println(addresses)
}
