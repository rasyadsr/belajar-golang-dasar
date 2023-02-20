package main

import "fmt"

func main() {
	var (
		nilai32 int32 = 10000000
		nilai64 int64 = int64(nilai32)

		/*
			terjadi integer overflow, karena variable nilai32 terlalu besar angkanya
			untuk int8, jika berlebih maka hasil konversi berisi nilai min dari tipe data tersebut
			min : -128
			max : 128
		*/
		nilai8 int8 = int8(nilai32)
	)

	fmt.Println(nilai32)
	fmt.Println(nilai64)
	fmt.Println(nilai8)

	var name string = "Rasyad"
	var r byte = name[0]
	fmt.Println("r byte :", r)
	var rToString string = string(r)
	fmt.Println("rToString string :", rToString)
}
