package main

import "fmt"

func main() {
	/* Mendeklarasikan variable dengan tipe data */
	var name string

	name = "bang"
	fmt.Println(name)

	// name = true // error

	name = "siapa"
	fmt.Println(name)

	/*
		Menginisialisasi langsung variable,
		sehingga kita tidak perlu mendeklarasikan tipe data variable tersebut
	*/
	var ismun = "bahlul"
	fmt.Println(ismun)

	/*
		Tidak wajib menggunakan keyword var,
		ketika deklarasi variable menggunakan :=
	*/
	ngaran := "abdi"
	fmt.Println(ngaran)

	/*
		Deklarasi multiple variable
	*/
	var (
		firstname = "Mochamad"
		lastname  = "Rasyad"
	)
	fmt.Println(firstname, " ", lastname)
}
