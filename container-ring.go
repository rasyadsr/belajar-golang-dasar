package main

import (
	"container/ring"
	"fmt"
	"strconv"
)

func main() {
	var data *ring.Ring = ring.New(5) // membuat ring berisi 5 value

	for i := 0; i < data.Len(); i++ {
		data.Value = "Data " + strconv.FormatInt(int64(i), 10)
		data = data.Next() // biar datanya ke-assign ke data selanjutnya
	}

	fmt.Println(*data)

	/*
		- untuk melakukan iterasi gunakan function DO,
		- tidak bisa menggunakan for loop, karena ring akan selalu berputar ( tidak memiliki akhir )
	*/
	data.Do(func(value any) {
		fmt.Println(value)
	})
}
