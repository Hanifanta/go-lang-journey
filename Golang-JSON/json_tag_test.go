package Golang_JSON

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Product struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	ImageURL string `json:"image_url"`
}

func TestJSONTag(t *testing.T) {
	product := Product{
		Id:       "P001",
		Name:     "Product 1",
		ImageURL: "https://example.com/product1.jpg",
	}

	bytes, _ := json.Marshal(product)
	fmt.Println(string(bytes))
}

func TestJSONTagDecode(t *testing.T) {
	jsonString := `{"id":"P001","name":"Product 1","image_url":"https://example.com/product1.jpg"}`
	jsonBytes := []byte(jsonString)

	product := &Product{}

	json.Unmarshal(jsonBytes, product)
	fmt.Println(product)
}
