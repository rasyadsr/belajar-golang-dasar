package main

import "fmt"

func main() {

	/*
		- child bisa mengakses data yang ada di parent, namun
		- parent tidak bisa mengakses data yang ada di child
	*/
	counter := 0

	increment := func() {
		fmt.Println("Increment")
		counter++
	}

	increment()
	increment()

	fmt.Println(counter)

}
