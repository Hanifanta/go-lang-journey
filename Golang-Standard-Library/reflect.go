package main

import (
	"fmt"
	"reflect"
)

type Sample struct {
	Name string `required:"true" max:"10"`
}

type Person struct {
	Name string `required:"true" max:"10"`
}

func readField(value any) {
	valueType := reflect.TypeOf(value)
	fmt.Println("Type Name", valueType.Name())
	for i := 0; i < valueType.NumField(); i++ {
		valueField := valueType.Field(i)
		fmt.Println(valueField.Name, "with Type", valueField.Type)
		fmt.Println(valueField.Tag.Get("required"))
		fmt.Println(valueField.Tag.Get("max"))
	}
}

func IsValid(value any) (result bool) {
	result = true
	t := reflect.TypeOf(value)
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		if field.Tag.Get("required") == "true" {
			data := reflect.ValueOf(value).Field(i).Interface()
			result = data != ""
			if result == false {
				return result
			}
		}
	}
	return result
}

func main() {
	readField(Sample{"John Doe"})
	readField(Person{"John Doe"})

	person := Person{"Hanif"}
	fmt.Println("Valid", IsValid(person))
}
