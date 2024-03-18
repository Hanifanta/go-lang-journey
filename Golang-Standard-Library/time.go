package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println(now.Local())

	utc := time.Date(2023, 12, 29, 0, 0, 0, 0, time.UTC)
	fmt.Println(utc.Local())

	formatter := "2006-01-02 15:04:05"
	value := "2023-12-29 00:00:00"
	valueTime, err := time.Parse(formatter, value)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(valueTime)
	}
}
