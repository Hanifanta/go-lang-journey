package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("Hanif Al Irsyad", "Hanif"))
	fmt.Println(strings.Split("Hanif Al Irsyad", " "))
	fmt.Println(strings.ToLower("Hanif Al Irsyad"))
	fmt.Println(strings.ToUpper("Hanif Al Irsyad"))
	fmt.Println(strings.Trim("   Hanif Al Irsyad   ", " "))
	fmt.Println(strings.ReplaceAll("Hanif Al Irsyad", "Hanif", "Irsyad"))
}
