package main

import (
	"fmt"
	"sort"
)

type User struct {
	Name string
	Age  int
}

type UserSlice []User

/* implementasi interface sort Len(), Less(i, j), Swap(i, j) */
func (value UserSlice) Len() int {
	return len(value)
}

func (value UserSlice) Less(i, j int) bool {
	return value[i].Age < value[j].Age
}

func (value UserSlice) Swap(i, j int) {
	value[i], value[j] = value[j], value[i]
}

func main() {

	users := []User{
		{"Mochamad", 30},
		{"Rasyad", 20},
		{"Syafiq", 10},
		{"Riswansyah", 11},
	}

	fmt.Println(users)

	sort.Sort(UserSlice(users))

	fmt.Println(users)
}
