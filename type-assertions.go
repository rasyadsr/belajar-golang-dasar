package main

import "fmt"

func random() interface{} {
	return "Ups"
}

func main() {
	var result interface{} = random()

	// cara ketika, kita sudah tau type data dari return value nya
	var resultString string = result.(string) // jangan salah konversi, jika salah akan terjadi panic
	fmt.Println(resultString)

	switch value := result.(type) {
	case string:
		fmt.Println("String", value)
	case int:
		fmt.Println("Integer", value)
	default:
		fmt.Println("Unknown Type")
	}
}
