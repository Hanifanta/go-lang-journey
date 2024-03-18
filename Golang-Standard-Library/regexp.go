package main

import (
	"fmt"
	"regexp"
)

func main() {
	regex := regexp.MustCompile(`Ha([a-z])if`)

	fmt.Println(regex.MatchString("Hanif"))
	fmt.Println(regex.MatchString("Haiif"))
	fmt.Println(regex.MatchString("Hafif"))

	fmt.Println(regex.FindAllString("Hanif Hafif Haiif Hwquf", 10))
}
