package main

import "fmt"

func main() {
	var ujian = 80
	var absensi = 70

	var lulusUjian = ujian >= 80
	var lulusAbsensi = absensi >= 80
	fmt.Println(lulusUjian)
	fmt.Println(lulusAbsensi)

	var lulus = lulusAbsensi && lulusUjian
	fmt.Println(lulus)

	/* Namun biasanya langsung seperti ini */
	fmt.Println(ujian >= 80 && absensi >= 80)
}