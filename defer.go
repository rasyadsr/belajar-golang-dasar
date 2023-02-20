package main

import "fmt"

func logging() {
	fmt.Println("Selesai Memanggil function")
}

func runApplication(value int) {
	/*
		- tidak peduli ada panic ataupun tidak, defer akan selalu di panggil
		- taruh di paling atas
	*/
	defer logging()
	fmt.Println("Run Application")
	result := 10 / value // namun jika dibagi 0 akan terjadi panic ( error ), sehingga kode di bawah nya tidak akan di eksekusi
	fmt.Println("Result", result)
	// logging()
}

func main() {
	runApplication(0)
}
