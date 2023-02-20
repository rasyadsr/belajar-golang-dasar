package main

import "fmt"

func main() {
	// jumlah_data_yang_bisa_ditampung = length
	// var namaArray [jumlah_data_yang_bisa_ditampung]tipe_ata
	var names [2]string
	names[0] = "Bang"
	names[1] = "Messi"
	// names[2] = "Artinya apa" // error

	fmt.Println(names)
	fmt.Println(names[0])
	fmt.Println(names[1])

	/* Deklarasi sekaligus inisialisasi */
	var values = [3]int{
		90,
		80,
		70,
	}
	
	fmt.Println(values)
	fmt.Println(values[0])
	fmt.Println(values[1])
	fmt.Println(values[2])

	/* function di array */
	// jumlah_data_yang_bisa_ditampung = length
	/* len(), untuk mendapatkan jumlah_data_yang_bisa_ditampung */
	fmt.Println(len(names)) 
	fmt.Println(len(values))

	/* array[index] = value untuk merubah array */
	names[0] = "Pak"
	fmt.Println(names)

	/* array[index] untuk mengambil value dari array berdasarkan index */
	fmt.Println(names[1])
}