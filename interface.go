package main

import "fmt"

// cara membuat kontrak interface
type HasName interface {
	GetName() string
}

// sebuah function yang menerima tipe HasName
func SayHelloLagi(hasName HasName) {
	fmt.Println("Hello", hasName.GetName())
}

type Person struct {
	Name string
}

// person ini, sudah implement interface HasName
// implement function harus sama persis dengan interface
func (person Person) GetName() string {
	return person.Name
}

func main() {
	var rasyad Person
	rasyad.Name = "Rasyad"

	SayHelloLagi(rasyad)
}
