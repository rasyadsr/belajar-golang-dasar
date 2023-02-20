package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("Mochamad Rasyad", "Rasyad"))
	fmt.Println(strings.Contains("Mochamad Rasyad", "Budi"))

	fmt.Println(strings.Split("Mochamad Rasyad", " "))

	fmt.Println(strings.ToLower("Mochamad Rasyad"))
	fmt.Println(strings.ToUpper("Mochamad Rasyad"))
	fmt.Println(strings.ToTitle("Mochamad Rasyad"))

	fmt.Println(strings.Trim("    Mochamad Rasyad         ", " "))

	fmt.Println(strings.ReplaceAll("Eko Eko Eko Eka Eki Rasyad", "Eko", "Ubah"))
}
