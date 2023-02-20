package main

import "fmt"

func main() {

	/* for loop ( mirip while ) */
	nilaiAwal := 1
	for nilaiAwal <= 10 {
		fmt.Println("Perulangan ke-", nilaiAwal)
		nilaiAwal++
	}

	/* for loop pada umumnya*/
	for i := 1; i <= 10; i++ {
		fmt.Println("Perulangan ke-", i)
	}

	/* for range ( mirip foreach )*/
	days := []string{
		"Senin",
		"Selasa",
		"Rabu",
		"Kamis",
		"Jumat",
	}
	for i, v := range days {
		fmt.Println("Index -", i, " Value :", v)
	}
}
