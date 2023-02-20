package main

import "fmt"

/* variadic function */
func sumAll(numbers ...int) int {
	total := 0
	for _, v := range numbers {
		total += v
	}
	return total
}

func main() {
	total := sumAll(10, 10, 10)
	fmt.Println(total)

	slice := []int{10, 10, 20}
	total = sumAll(slice...)
	fmt.Println(total)
}
