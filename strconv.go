package main

import (
	"fmt"
	"strconv"
)

func main() {
	boolean, err := strconv.ParseBool("true")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(boolean)
	}

	number, err := strconv.ParseInt("100000", 10, 64)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(number)
	}
}
