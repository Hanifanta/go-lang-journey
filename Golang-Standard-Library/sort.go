package main

import (
	"fmt"
	"sort"
)

type User struct {
	Name string
	Age  int
}

type UsersSlice []User

func (user UsersSlice) Len() int {
	return len(user)
}

func (User UsersSlice) Less(i, j int) bool {
	return User[i].Age < User[j].Age
}

func (User UsersSlice) Swap(i, j int) {
	User[i], User[j] = User[j], User[i]
}

func main() {
	user := []User{
		{"John", 20},
		{"Jane", 30},
		{"Doe", 25},
		{"Foo", 10},
		{"Bar", 15},
	}

	sort.Sort(UsersSlice(user))

	fmt.Println(user)
}
