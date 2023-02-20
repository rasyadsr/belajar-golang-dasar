package main

import (
	"fmt"
	"regexp"
)

func main() {
	var regex *regexp.Regexp = regexp.MustCompile("e([a-z])o")

	fmt.Println(regex.MatchString("eko"))
	fmt.Println(regex.MatchString("edo"))
	fmt.Println(regex.MatchString("eDo"))

	var search []string = regex.FindAllString("eko eka edo eto eyo eki", 2) // jika -1, itu akan mengambil semua data yang match
	fmt.Println(search)
}
