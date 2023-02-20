package main

import "fmt"

func Ups(i int) interface{} {
	if i == 1 {
		return i
	} else if i == 2 {
		return true
	} else {
		return "Ups"
	}
}

func main() {
	/*
		Walauun kita tahu ketika mengirim 1, akan mereturn 1 tersebut yang bernilai int,
		namun kita tidak bisa langsung deklarasikan tipe data nya dengan int,
		kita harus menggunakan interface{}
		var data int = Ups(1)
	*/
	var data interface{} = Ups(1)
	fmt.Println(data)
}
