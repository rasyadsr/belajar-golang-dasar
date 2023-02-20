package main

import "fmt"

func newMap(name string) map[string]string {
	if name == "" {
		return nil
	} else {
		return map[string]string{
			"name": name,
		}
	}
}

func main() {
	var person map[string]string = newMap("Tes")
	if person != nil {
		fmt.Println(person)
	} else {
		fmt.Println("Data Kosong!")
	}

	var other map[string]string = newMap("")
	if other != nil {
		fmt.Println(other)
	} else {
		fmt.Println("Data Kosong!")
	}
}
