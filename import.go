package main

import (
	"belajar-golang-dasar/helper"
	"fmt"
)

/*
	- setiap nama variable / function yang di awali huruf besar
	bisa di export keluar
*/

func main() {
	helper.SayHello("rasyad")
	// helper.sayGoodBye("rasyad") // error ga bisa

	fmt.Println(helper.Application)
	// fmt.Println(helper.version) // // error ga bisa
}
