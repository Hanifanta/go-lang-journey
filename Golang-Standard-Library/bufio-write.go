package main

import (
	"bufio"
	"os"
)

func main() {
	writer := bufio.NewWriter(os.Stdout)
	_, _ = writer.WriteString("Hello World!\n")
	writer.Flush()

}
