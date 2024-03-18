package main

import (
	"flag"
	"fmt"
)

func main() {
	username := flag.String("username", "root", "Database username")
	password := flag.String("password", "root", "Database password")
	host := flag.String("host", "localhost", "Database host")
	port := flag.Int("port", 3306, "Database port")

	flag.Parse()

	fmt.Println(*username, *password, *host, *port)

}
