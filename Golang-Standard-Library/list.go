package main

import (
	"container/list"
	"fmt"
)

func main() {
	data := list.New()

	data.PushBack("Hanif")
	data.PushBack("Al")
	data.PushBack("Irsyad")

	head := data.Front()
	fmt.Println(head.Value)

	next := head.Next()
	fmt.Println(next.Value)

	next = next.Next()
	fmt.Println(next.Value)

	for e := data.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}
