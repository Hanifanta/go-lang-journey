package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

func main() {
	input := strings.NewReader("Hello,\n Reader!")

	reader := bufio.NewReader(input)

	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}

		fmt.Println(string(line))
	}
}
