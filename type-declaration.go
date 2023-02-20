package main

import "fmt"

func main() {

	type NoKtp string
	type IsMarried bool

	var noKtpOrang NoKtp = "1234567890"
	fmt.Println(noKtpOrang)

	var statusMarried IsMarried = true
	fmt.Println(statusMarried)
}