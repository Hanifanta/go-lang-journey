package main

import "fmt"

type validationError struct {
	Message string
}

func (v *validationError) Error() string {
	return v.Message
}

type notFoundError struct {
	Message string
}

func (n *notFoundError) Error() string {
	return n.Message
}

func SaveData(id string, data any) error {
	if id == "" {
		return &validationError{"Validation Error"}
	}
	if id != "Hanif" {
		return &notFoundError{"Not Found Error"}
	}

	return nil
}

func main() {
	err := SaveData("Budi", nil)
	if err != nil {
		// if validationError, ok := err.(*validationError); ok {
		// 	fmt.Println(validationError.Error())
		// } else if notFoundError, ok := err.(*notFoundError); ok {
		// 	fmt.Println(notFoundError.Error())
		// } else {
		// 	fmt.Println(err.Error())
		// }

		switch FinalError := err.(type) {
		case *validationError:
			fmt.Println(FinalError.Error())
		case *notFoundError:
			fmt.Println(FinalError.Error())
		default:
			fmt.Println(FinalError.Error())
		}

	} else {
		fmt.Println("Success")
	}

}
