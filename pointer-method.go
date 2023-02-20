package main

import "fmt"

type Man struct {
	Name string
}

func (man *Man) Married() {
	man.Name = "Mr. " + man.Name
	// fmt.Println("Di method", man.Name)
}

func main() {
	rasyad := Man{"Rasyad"}
	fmt.Println(rasyad)
	rasyad.Married()
	fmt.Println(rasyad)
}
