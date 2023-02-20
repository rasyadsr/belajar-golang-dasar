package main

import "fmt"

func getFullName() (string, string) {
	return "Bang", "Messi"
}

func main() {
	firstName, lastName := getFullName()
	fmt.Println(firstName)
	fmt.Println(lastName)

	// ignore jika tidak di butuhkan menggunakan _
	_, last := getFullName()
	fmt.Println(last)
}
