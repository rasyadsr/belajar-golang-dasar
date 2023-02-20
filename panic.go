package main

import "fmt"

func endApp() {

	/*
		- untuk menangkap panic, letakan recover di function yang akan di jalankan setelah nya,
		- contoh nya adalah defer function
	*/
	err := recover()
	if err != nil {
		fmt.Println("Error : ", err)
	}

	fmt.Println("Aplikasi selesai")
}

func runApp(error bool) {
	defer endApp()

	if error {
		panic("Aplikasi error") // seperti throw new Exception
	}

	fmt.Println("Aplikasi berjalan")
}

func main() {
	// runApp(false)
	runApp(true)
}
