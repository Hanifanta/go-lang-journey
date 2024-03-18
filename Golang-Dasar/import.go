package main

import (
	"belajar-golang/helper"
	_ "belajar-golang/internal"
	"fmt"
)

func main() {
	result := helper.SayHello("Hanif")
	fmt.Println(result)
}
