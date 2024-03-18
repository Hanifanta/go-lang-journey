package main

import (
	"encoding/csv"
	"os"
)

func main() {
	writer := csv.NewWriter(os.Stdout)

	_ = writer.Write([]string{"Hanif", "Al", "Irsyad"})
	_ = writer.Write([]string{"Hanif", "Al", "Irsyad"})
	_ = writer.Write([]string{"Hanif", "Al", "Irsyad"})

	writer.Flush()

}
