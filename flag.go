package main

import (
	"flag"
	"fmt"
)

func main() {
	/*
		variable" ini berisi pointer string
	*/
	var host *string = flag.String("host", "default-value", "Put your database host")
	var username *string = flag.String("username", "default-value", "Put your database username")
	var password *string = flag.String("password", "default-value", "Put your database password")

	flag.Parse() // perintah ini wajib di panggil

	// hanya akan mendapatkan nilai pointer nya
	fmt.Println("Host :", host)
	fmt.Println("User :", username)
	fmt.Println("Password :", password)

	// go run flag.go -host=localhost -username=root -password=root

	// gunakan * di awal nama varibel untuk mendapatkan nilai value nya
	fmt.Println("Host :", *host)
	fmt.Println("User :", *username)
	fmt.Println("Password :", *password)
}
