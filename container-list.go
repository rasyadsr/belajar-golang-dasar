package main

import (
	"container/list"
	"fmt"
)

func main() {
	data := list.New()

	data.PushBack("Rasyad")
	data.PushBack("Syafiq")
	data.PushBack("Riswansyah")
	data.PushFront("Mochamad")

	// data.Front().Prev().Prev()

	fmt.Println(data.Front().Value)
	fmt.Println(data.Back().Value)

	// depan -> belakang
	for element := data.Front(); element != nil; element = element.Next() {
		fmt.Println(element.Value)
	}

	// belakang -> depan
	for element := data.Back(); element != nil; element = element.Prev() {
		fmt.Println(element.Value)
	}
}
