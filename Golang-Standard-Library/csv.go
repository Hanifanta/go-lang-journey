package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"strings"
)

func main() {
	csvString := "hanif,al,irsyad\n" +
				"hanif,al,irsyad\n" +
				"hanif,al,irsyad\n" +
				"hanif,al,irsyad"

	reader := csv.NewReader(strings.NewReader(csvString))

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}

		fmt.Println(record)
	}
}
