package main

import (
	"fmt"
	"slices"
)

func main() {
	names := []string{"Hanif", "Al", "Irsyad"}
	values := []int{1, 2, 3, 4, 5}

	fmt.Println(slices.Min(values))
	fmt.Println(slices.Max(values))
	fmt.Println(slices.Contains(names, "Hanif"))
	fmt.Println(slices.Contains(names, "Al"))
	fmt.Println(slices.Contains(names, "Irsyad"))
	fmt.Println(slices.Index(names, "Hanif"))
}
