package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args
	// fmt.Println("Argument :", args)

	/*
		- Data parameter asli dimulai dari index ke- 1
		- index ke 0 nya berisi path lokasi file ini di ekseskusi
	*/

	if len(args) == 1 {
		fmt.Println("No Parameter!")
	} else {
		for _, v := range args {
			fmt.Println(v)
		}
	}

	name, err := os.Hostname()
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Hostname :", name)
	}

	GOPATH := os.Getenv("GOPATH")
	fmt.Println("GOPATH :", GOPATH)
}
