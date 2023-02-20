package main

import (
	"fmt"
	"strings"
)

/*
- callback func(string) string
- namaFunction func(type data parameter) return value
- jika lebihd dari satu parameter, pisahkan dengan koma
- func(string, number, bool)
*/
func sayHelloWithFilter(name string, callback func(string) string) {
	nameFiltered := callback(name)
	fmt.Println("Hello", nameFiltered)
}

func spamFilter(name string) string {
	if name == "Anjing" {
		return "..."
	} else {
		return name
	}
}

func upperToFilter(name string) string {
	return strings.ToUpper(name)
}

/* function type declaration */
type Filter func(string) string

func sayHelloWithFilterUsingTypeDeclaration(name string, callback Filter) {
	nameFiltered := callback(name)
	fmt.Println(nameFiltered)
}

func main() {
	sayHelloWithFilter("Messi", spamFilter)
	sayHelloWithFilter("Anjing", spamFilter)

	sayHelloWithFilter("kecil", upperToFilter)
}
