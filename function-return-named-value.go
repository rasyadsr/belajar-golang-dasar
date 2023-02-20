package main

import "fmt"

func getFullName2() (firstName string, middleName string, lastName string) {
	firstName = "Bang"
	middleName = "Messi"
	lastName = "Suii"

	// cukup return saja, karena kita sudah deklarasikan pada return named nya
	return
}

func main() {
	firstName, middleName, lastName := getFullName2()
	fmt.Println(firstName)
	fmt.Println(middleName)
	fmt.Println(lastName)
}
