package main

import "fmt"

func main() {
	var name string = "messi"

	if name == "messi" {
		fmt.Println("Hello messi")
	} else if name == "dodo" {
		fmt.Println("Hello dodo")
	} else {
		fmt.Println("Hello")
	}

	// if short statement
	if length := len(name); length > 5 {
		fmt.Println("Nama sudah benar")
	} else {
		fmt.Println("Nama terlalu panjang!")
	}
}
