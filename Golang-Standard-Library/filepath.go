package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	fmt.Println(filepath.Dir("/home/username/file.txt"))
	fmt.Println(filepath.Base("/home/username/file.txt"))
	fmt.Println(filepath.IsAbs("/home/username/file.txt"))
	fmt.Println(filepath.IsLocal("/home/username/file.txt"))
	fmt.Println(filepath.Ext("home/username/file.txt"))
	fmt.Println(filepath.Join("/home", "username", "file.txt"))

}
