package main

import "fmt"

func getGoodBye(name string) string {
	return "Good Bye " + name
}

func main() {
	/* variable sayGoodBye berisi function getGoodBye*/
	sayGoodBye := getGoodBye

	result := sayGoodBye("Messi")
	fmt.Println(result)

}
