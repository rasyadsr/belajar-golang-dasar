package main

import (
	"belajar-golang-dasar/database"
	/* _, adalah blank identifier kalau gamau di pake gunakan _ aja*/
	_ "belajar-golang-dasar/helper"
	"fmt"
)

func main() {
	result := database.GetDatabase()
	fmt.Println(result)
}
