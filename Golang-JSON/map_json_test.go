package Golang_JSON

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestMap(t *testing.T) {
	jsonString := `{"id":"P001","name":"Product 1","image_url":"https://example.com/product1.jpg"}`
	jsonBytes := []byte(jsonString)

	var result map[string]interface{}

	json.Unmarshal(jsonBytes, &result)

	fmt.Println(result)
	fmt.Println(result["id"])
	fmt.Println(result["name"])
	fmt.Println(result["image_url"])
}

func TestMapEncode(t *testing.T) {
	product := map[string]interface{}{
		"id":        "P001",
		"name":      "Product 1",
		"image_url": "https://example.com/product1.jpg",
	}

	bytes, _ := json.Marshal(product)
	fmt.Println(string(bytes))
}
