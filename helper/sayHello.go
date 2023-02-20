package helper

import "fmt"

var Application = "belajar-golang-dasar"
var version = "v1.0.0"

func SayHello(name string) {
	fmt.Println("Hello", name)
}

func sayGoodBye(name string) {
	fmt.Println("Good Bye", name)
}
