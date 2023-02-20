package main

import (
	"fmt"
)

func main() {

	name := "Budi"

	switch name {
	case "Budi":
		fmt.Println("Hello Budi")
	case "Joko":
		fmt.Println("Hello Joko")
	default:
		fmt.Println("Hello")
	}

	/* short statement */
	switch length := len(name); length > 5 {
	case true:
		fmt.Println("Nama terlalu panjang")
	case false:
		fmt.Println("Nama sudah benar")
	}

	/* switch tanpa kondisi */
	otherLength := len(name)
	switch {
	case otherLength > 5:
		fmt.Println("Nama lumayan panjang")
	case otherLength > 10:
		fmt.Println("Nama terlalu panjang")
	default:
		fmt.Println("Nama sudah benar")
	}

}
