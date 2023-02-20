package main

import "fmt"

func main() {
	var month = [12]string{
		"Januari",
		"Februari",
		"Maret",
		"April",
		"Mei",
		"Juni",
		"July",
		"Agustus",
		"September",
		"Oktober",
		"November",
		"Desember",
	}

	var slice1 = month[4:7]
	fmt.Println(slice1)
	fmt.Println(len(slice1))
	fmt.Println(cap(slice1))

	/*
		- ketika merubah array, slice yang reference terhadap array tersebut ikut berubah value nya
		- begitupun ketika merubah slice, array nya ikut berubah
	*/
	month[5] = "Ubah"
	fmt.Println(slice1)

	slice1[0] = "ubah oleh slice"
	fmt.Println(month)
	fmt.Println(slice1)

	var slice2 = month[10:]
	fmt.Println(slice2)

	/*
		- capacity sudah penuh, sehingga membuat array baru
		- sehingga array existing tidak ikut berubah
	*/
	var slice3 = append(slice2, "Bulan baru")
	fmt.Println(slice3)

	slice3[1] = "Bukan desember"
	fmt.Println(slice3)

	fmt.Println(slice2)
	fmt.Println(month)

	var slice4 = month[2:4]
	fmt.Println(slice4)

	var slice5 = append(slice4, "Tes masukin ke slice 5")
	fmt.Println(slice5)

	fmt.Println(month)

	/*
		make([]tipe_data, length, kapasitas)
	*/
	newSlice := make([]string, 2, 5)
	newSlice[0] = "Bang"
	newSlice[1] = "Messi"
	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	/* length dan capacity nya nyamain newSlice, biar sama persis dan data tidak terpotong */
	copySlice := make([]string, len(newSlice), cap(newSlice))
	copy(copySlice, newSlice)

	fmt.Println(copySlice)
}
