package main

import "fmt"

func main() {

	/* Deklarasi sekaligus inisialisasi */
	person := map[string]string{
		"name":    "Messi",
		"address": "Argentina",
	}

	/* Menambah property */
	person["title"] = "Player"

	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["address"])

	/* Deklarasi saja*/
	book := make(map[string]string)
	book["title"] = "Belajar Golang"
	book["ups"] = "typo"

	fmt.Println(book)

	delete(book, "ups")
	fmt.Println(book)
}
